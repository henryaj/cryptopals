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

var _ = Describe("RepeatingKeyXOR", func() {
	It("applies each byte of the key sequentially", func() {
		input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
		expected:= HexStringToBytes(`0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272`+
			`a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`)

		result:=RepeatingKeyXOR([]byte(input), []byte("ICE"))

		Expect(result).To(Equal(expected))
	})
})

func slicesToString(input [][]byte) (result []string) {
	for _, slice := range input {
		result = append(result, string(slice))
	}
	return result
}