package gcws

type Seg interface {
	Cut(string)
}

var WordList []string

func init() {

}

/*
最大匹配算法:
	根据词典,采用贪心策略进行分词
*/
type MaxMatchSeg struct {
}

func MaxLength(list map[string]float64) int {
	maxLength := 0
	for k, _ := range FREQ {
		if len([]rune(k)) > maxLength {
			maxLength = len([]rune(k))
		}
	}
	return maxLength
}

func Find(wlist map[string]float64, seg string) bool {
	_, ok := wlist[seg]
	return ok
}

func (seg *MaxMatchSeg) Cut(sentence []rune) []string {
	return seg.rightMatch(sentence)
}

func (seg *MaxMatchSeg) leftMatch(sentence []rune) []string {
	maxlength := MaxLength(FREQ)
	var words []string
	for cursor := 0; cursor < len(sentence); {
		for l := maxlength; l >= 1; l-- {
			if (cursor + l) <= len(sentence) {
				seg := sentence[cursor : cursor+l]
				if Find(FREQ, string(seg)) || len(seg) == 1 {
					words = append(words, string(seg))
					cursor += l
					break
				}
			}
		}
	}
	return words
}

func (seg *MaxMatchSeg) rightMatch(sentence []rune) []string {
	maxlength := MaxLength(FREQ)
	var words []string
	for cursor := len(sentence); cursor > 0; {
		for l := maxlength; l >= 1; l-- {
			if (cursor - l) >= 0 {
				seg := sentence[cursor-l : cursor]
				if Find(FREQ, string(seg)) || len(seg) == 1 {
					words = append(words, string(seg))
					cursor -= l
					break
				}
			}
		}
	}
	return words
}






