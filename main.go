package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Ошибка генерации:", err)
		return
	}

	privFile, _ := os.Create("private.pem")
	defer privFile.Close()
	pem.Encode(privFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	pubASN1, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	pubFile, _ := os.Create("public.pem")
	defer pubFile.Close()
	pem.Encode(pubFile, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	})

	fmt.Println("Keys created: private.pem and public.pem")
}
