package config

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gocarina/gocsv"
	"log"
	"mime/multipart"
	"os"
	"sync"
)

type Aws struct {
	Bucket  string
	session *session.Session
}

const (
	REGION = "sa-east-1"
	BUCKET = "bucket_name"
)

var (
	s3Once     sync.Once
	s3Instance *Aws
)

func NewAwsClient() *Aws {
	s3Once.Do(func() {
		s3Instance = createClient()
	})
	return s3Instance
}

func createClient() *Aws {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(REGION),
		Credentials: credentials.NewSharedCredentials("", "stori"),
	})

	if err != nil {
		fmt.Println("Error al crear la sesión de AWS:", err)
		return nil
	}
	return &Aws{
		session: sess,
		Bucket:  os.Getenv(BUCKET),
	}
}

func (t Aws) GetFile(key string, transactionFileRecord interface{}) error {
	s := s3.New(t.session)
	resp, err := s.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(t.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Println("Error al obtener el objeto desde S3:", err)
		return err
	}
	defer resp.Body.Close()

	if err := gocsv.Unmarshal(resp.Body, transactionFileRecord); err != nil {
		log.Fatalf("Error al deserializar el archivo CSV: %v", err)
		return err
	}
	return nil
}

func (t Aws) PutObjet(filename string, file multipart.File) {
	uploader := s3manager.NewUploader(t.session)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(t.Bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		fmt.Printf("fail put s3 %v", err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, t.Bucket)
}
