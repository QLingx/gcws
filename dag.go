package gcws

import (
	"math"
)

func Get_DAG(sentence []rune) map[int][]int {
	DAG := make(map[int][]int)
	length := len(sentence)
	for k := 0; k < length; k++ {
		tmplist := make([]int, 0)
		i := k
		frag := string(sentence[k])
		_, ok := FREQ[frag]
		for ; i < length && ok; {
			if int(FREQ[frag]) != 0 {
				tmplist = append(tmplist, i)
			}
			i += 1
			if i == length {
				frag = string(sentence[k:i])
			} else {
				frag = string(sentence[k : i+1])
			}
			_, ok = FREQ[frag]
		}
		if len(tmplist) == 0 {
			tmplist = append(tmplist, k)
		}
		DAG[k] = tmplist
	}
	return DAG
}

func Calc(sentence []rune, DAG map[int][]int) (route map[int][2]float64) {
	length := len(sentence)
	route = make(map[int][2]float64)
	route[length] = [2]float64{0, 0}
	logTotal := math.Log(TOTAL)
	for idx := length - 1; idx > -1; idx-- {
		tmp := make([][2]float64, 0)
		for _, x := range DAG[idx] {
			v, ok := FREQ[string(sentence[idx:x+1])]
			if !ok {
				v = 1
			}
			tuple := [2]float64{math.Log(v) - logTotal + route[x+1][0], float64(x)}
			tmp = append(tmp, tuple)
		}
		route[idx] = max(tmp)
	}
	return
}

func max(arr [][2]float64) [2]float64 {
	res := arr[0]
	for _, v := range arr {
		if res[0] < v[0] {
			res = v
		}
	}
	return res
}
