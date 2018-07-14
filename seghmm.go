package gcws


type HMMSeg struct {
}

func (seg *HMMSeg) Cut(sentence string, cutAll, HMM bool) []string {
	if cutAll {
		return seg.CutAll([]rune(sentence))
	} else if HMM {
		return seg.CutDAGWithHMM([]rune(sentence))
	} else {
		return seg.CutDGA([]rune(sentence))
	}
}

func (seg *HMMSeg) CutDGA(sentence []rune) (res []string) {
	DAG := Get_DAG(sentence)
	router := Calc(sentence, DAG)
	buf := ""
	for x := 0; x < len(sentence); {
		y := int(router[x][1] + 1)
		lWord := sentence[x:y]
		if len(lWord) == 1 {
			buf += string(lWord)
			x = y
		} else {
			if len(buf) != 0 {
				res = append(res, buf)
				buf = ""
			}
			res = append(res, string(lWord))
			x = y
		}

	}
	if len(buf) != 0 {
		res = append(res, buf)
		buf = ""
	}
	return
}

func (seg *HMMSeg) CutAll(sentence []rune) (res []string) {
	DAG := Get_DAG(sentence)
	oldJ := -1
	for k, L := range DAG {
		if len(L) == 1 && k > oldJ {
			word := string(sentence[k : L[0]+1])
			res = append(res, word)
			oldJ = L[0]
		} else {
			for _, j := range L {
				if j > k {
					word := string(sentence[k : j+1])
					res = append(res, word)
					oldJ = j
				}
			}
		}
	}
	return
}

func (seg *HMMSeg) CutDAGWithHMM(senetence []rune) (res []string) {
	return
}