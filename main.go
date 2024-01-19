package main

import (
	"fmt"

	"github.com/AshirwadPradhan/tracksslcerts/sslcertchecker"
)

func main() {

	domains := []string{"google.com", "ashiprad.com", "example.com"}

	for _, domain := range domains {
		sslInfo := sslcertchecker.NewSSLCertInfo(domain)
		sslInfo.Validate()
		fmt.Printf("Info: %+v\n", sslInfo)
	}

}
