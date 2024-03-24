## Para pruebas en local

Para poder levantar el proyecto necesita crear **.aws/credentials**
y agregar las sigueintes valores:
```
[stori]
aws_access_key_id= key
aws_secret_access_key= secret
```
Posteriormente entrar en el archivo .env y rellenar las variables de entorno.

## ejemplo de curl

cargar file en s3 para la prueba con el siguiente servicio

```
curl --location 'localhost:8080/loadFile' \
--header 'Content-Type: application/ecmascript' \
--form 'file=@"/C:/Users/max12/OneDrive/Documentos/proyectogo/transactionData.csv"' \
--form 'data="{
    \"path\": \"archivo.csv\"
}"'
```

Leer archivo del s3 y mandar reporte por correo

```
curl --location 'localhost:8080/sendMail' \
--header 'Content-Type: application/json' \
--data-raw '{
    "mail":"agrega correo",
    "path":"archivo.csv"
}'
```
Ejecutar el docker-compose para levantar base de datos,
crear schema con:
```
CREATE DATABASE nombre_del_schema;
```
ejecutar main-test.go que esta
ubidado en ***/cmd/lambda*** para levantar aplicacion


## variables de entorno
```
db_hostname =localhost
db_port     =3306
db_password = 
db_name   = stori
db_user     = root
smtp_hostname=smtp.gmail.com
smtp_port=587
smtp_password=
smtp_mail=
aws_region=us-east-1
bucket_name=
```

### Generar archivo para lambda
```
//paso 1
make build
//paso 2
make zip
```
