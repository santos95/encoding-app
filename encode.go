package main

import (
	b64 "encoding/base64"
	"fmt"
)

func EncodeString(text string) string {

	fmt.Println("Encoding string ...")
	encStr := b64.StdEncoding.EncodeToString([]byte(text))

	return encStr
}
