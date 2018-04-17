package encode

import (
	"fmt"
	"sort"
)

type guess struct {
	keyLength int
	score     int
}

type guesses []guess

func (g guesses) Len() int {
	return len(g)
}
func (g guesses) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func (g guesses) Less(i, j int) bool {
	return g[i].score < g[j].score
}

func GuessKeySize(input []byte) []int {
	var maxKeyLength = 40
	var results guesses

	for keysize := 2; keysize <= maxKeyLength; keysize++ {
		one := input[0:keysize]
		two := input[keysize : keysize*2]

		score := HammingDistance(one, two) / keysize
		results = append(results, guess{keyLength: keysize, score: score})
	}

	sort.Sort(results)
	return []int{results[0].keyLength, results[1].keyLength, results[2].keyLength}
}

func BreakRepeatingKeyXOR(input []byte, keyLength int) []byte {
	var blocks [][]byte
	numBlocks := len(input) / keyLength

	for i := 0; i < numBlocks; i++ {
		blocks = append(blocks, input[i:keyLength*i])
	}

	fmt.Println(blocks)

	return []byte{}
}
