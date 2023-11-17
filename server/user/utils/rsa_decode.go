package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

func RsaDecode(encryptedData string) (string, error) {
	encryptedDecodeBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	block, _ := pem.Decode([]byte(KeyForPwd.PrivateKey))
	priKey, parseErr := x509.ParsePKCS8PrivateKey(block.Bytes)
	if parseErr != nil {
		fmt.Println(parseErr.Error())
		return "", errors.New("解析私钥失败")
	}
	originalData, encryptErr := rsa.DecryptPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), encryptedDecodeBytes)
	if encryptErr != nil {
		fmt.Println(encryptErr.Error())
	}
	return string(originalData), encryptErr
}
