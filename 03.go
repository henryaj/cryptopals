// package foo

package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"strings"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Result struct {
	output string
	score  int
}

type Results []Result

func (rs Results) Len() int {
	return len(rs)
}
func (rs Results) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}
func (rs Results) Less(i, j int) bool {
	return rs[i].score < rs[j].score
}

func main() {
	input := decodeHex(os.Args[1])
	chars := make([]byte, 256)
	for i := 0; i < 256; i++ {
		chars[i] = byte(i)
	}

	var results Results

	for _, char := range chars {
		result := make([]byte, len(input))
		for i, b := range input {
			result[i] = b ^ char
		}

		results = append(results, Result{output: string(result)})
	}

	for _, result := range results {
		var score int
		for _, letter := range letters {
			count := strings.Count(result.output, string(letter))
			score = score + count
		}
		result.score = score
	}

	sort.Sort(results)
	fmt.Println(results[0:25])
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
