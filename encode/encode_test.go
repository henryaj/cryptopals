package encode_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"

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
		expected := HexStringToBytes(`0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272` +
			`a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`)

		result := RepeatingKeyXOR([]byte(input), []byte("ICE"))

		Expect(result).To(Equal(expected))
	})
})

var _ = Describe("GetNumSetBits", func() {
	It("calculates the number of set bits in a byte", func() {
		// 00100001 (65), 2 set bits
		// 00000011 (3), 2 set bits

		Expect(GetNumSetBits(65)).To(Equal(2))
		Expect(GetNumSetBits(3)).To(Equal(2))
	})
})

var _ = Describe("HammingDistance", func() {
	It("calculates the Hamming distance between two strings", func() {
		Expect(HammingDistance([]byte("this is a test"), []byte("wokka wokka!!!"))).To(Equal(37))
	})
})

var _ = Describe("GuessKeySize", func() {
	It("returns the top 3 most likely key sizes", func() {
		ciphertext, err := ioutil.ReadFile("repeating_key_xor.txt")
		Expect(err).NotTo(HaveOccurred())

		result := GuessKeySize(ciphertext)

		Expect(result).To(Equal([]int{15, 3, 38}))
	})
})

var _ = Describe("BreakRepeatingKeyXOR", func() {
	It("foos", func() {
		ciphertext, err := ioutil.ReadFile("repeating_key_xor.txt")
		Expect(err).NotTo(HaveOccurred())

		BreakRepeatingKeyXOR(ciphertext, 15)
	})
})

func slicesToString(input [][]byte) (result []string) {
	for _, slice := range input {
		result = append(result, string(slice))
	}
	return result
}
