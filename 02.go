package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	hex := decodeHex(os.Args[1])
	xorTarget := decodeHex(os.Args[2])

	dest := make([]byte, len(hex))

	for i, _ := range hex {
		b := hex[i] ^ xorTarget[i]
		dest[i] = b
	}

	fmt.Println(encodeHex(dest))
}

func decodeHex(hexStr string) []byte {
	b, err := hex.DecodeString(hexStr)
	if err != nil {
		panic("unable to decode hex")
	}

	return b
}

func encodeHex(hexBytes []byte) string {
	return hex.EncodeToString(hexBytes)

}

func encode64(bytes []byte) string {
	return base64.StdEncoding.EncodeToString(bytes)
}
