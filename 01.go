package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	data := os.Args[1]
	hex := decodeHex(data)

	fmt.Println(encode64(hex))
}

func decodeHex(hexStr string) []byte {
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		panic("unable to decode hex")
	}

	return b
}

func encode64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}
