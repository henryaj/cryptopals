package encode

import (
	"encoding/base64"
	"encoding/hex"
)

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func HammingDistance(input1, input2 string) int {
	var count int
	xor := FixedXOR([]byte(input1), []byte(input2))
	for _, byte := range xor {
		count += GetNumSetBits(int(byte))
	}

	return count
}

func GetNumSetBits(input int) int {
	var count int

	for input > 0 {
		count += input & 1
		input >>= 1
	}

	return count
}

func HexStringToBytes(input string) []byte {
	b, err := hex.DecodeString(input)
	if err != nil {
		panic("unable to decode hex")
	}

	return b
}

func HexBytesToBase64(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func HexBytesToString(hexBytes []byte) string {
	return hex.EncodeToString(hexBytes)
}

func FixedXOR(input, xor []byte) []byte {
	dest := make([]byte, len(input))

	for i := range input {
		b := input[i] ^ xor[i]
		dest[i] = b
	}

	return dest
}

func SingleByteXOR(input []byte) [][]byte {
	var result [][]byte

	for _, xorTarget := range chars {
		xorTargetByte := byte(xorTarget)
		dest := make([]byte, len(input))
		for i := range input {
			dest[i] = input[i] ^ xorTargetByte
		}

		result = append(result, dest)
	}

	return result
}

func RepeatingKeyXOR(input, key []byte) []byte {
	result := make([]byte, len(input))
	var keyPointer int

	for i := range input {
		result[i] = input[i] ^ key[keyPointer]
		if keyPointer < (len(key) - 1) {
			keyPointer++
		} else {
			keyPointer = 0
		}
	}

	return result
}
