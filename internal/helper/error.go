package helper

import (
	"fmt"
	"net/http"
)

func ApiBadRequest(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func PrintlnError(err error, msj string) {
	if err != nil {
		fmt.Println(msj, err)
	}

}

func ErrorPanic(err error) {
	if err != nil {
		fmt.Printf("fail %v", err)
		panic(err)
	}
}
