package embedded

import (
	_ "embed"
	"math/rand/v2"
	"regexp"

	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
	"github.com/SladeThe/word-of-wisdom/internal/server/repositories"
)

var (
	//go:embed data/word_of_wisdom.txt
	wordOfWisdomData string

	wordOfWisdomOffsets = func() [][2]int {
		separatorOffsets := regexp.MustCompile("\r\n|\r|\n").FindAllStringIndex(wordOfWisdomData, -1)
		if len(separatorOffsets) == 0 {
			return nil
		}

		offsets := make([][2]int, len(separatorOffsets))
		offsets[0] = [2]int{0, separatorOffsets[0][0]}

		for i, separatorOffset := range separatorOffsets[1:] {
			offsets[i+1] = [2]int{separatorOffsets[i][1], separatorOffset[0]}
		}

		return offsets
	}()
)

type WordOfWisdom struct {
}

var _ repositories.WordOfWisdom = (*WordOfWisdom)(nil)

func NewWordOfWisdom() WordOfWisdom {
	return WordOfWisdom{}
}

func (w WordOfWisdom) OneRandom() (entities.WordOfWisdom, error) {
	if len(wordOfWisdomOffsets) == 0 {
		return entities.WordOfWisdom{}, repositories.ErrNotFound
	}

	offset := wordOfWisdomOffsets[rand.IntN(len(wordOfWisdomOffsets))]
	return entities.WordOfWisdom{Text: wordOfWisdomData[offset[0]:offset[1]]}, nil
}
