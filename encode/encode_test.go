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

		expected := HexStringToBytes("746865206b696420646f6e277420706c617")

		Expect(FixedXOR(input1, input2)).To(Equal(expected))
	})
})
