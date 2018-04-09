package encode_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/henryaj/cryptopals/encode"
)

var _ = Describe("FixedXOR", func() {
	It("XORs two equal-lenth slices of bytes", func() {
		input1 := HexStringToBytes("1c0111001f010100061a024b53535009181c")
		input2 := HexStringToBytes("686974207468652062756c6c277320657965")

		expected := HexStringToBytes("746865206b696420646f6e277420706c6179")

		Expect(FixedXOR(input1, input2)).To(Equal(expected))
	})
})

var _ = Describe("SingleByteXOR", func() {
	It("returns all possible outputs from XORing a string against a single byte", func() {
		input := HexStringToBytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

		results := SingleByteXOR(input)
		Expect(slicesToString(results)).To(ContainElement("Cooking MC's like a pound of bacon"))
	})
})

func slicesToString(input [][]byte) (result []string) {
	for _, slice := range input {
		result = append(result, string(slice))
	}
	return result
}