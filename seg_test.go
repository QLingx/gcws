package gcws

import (
	"fmt"
	"testing"
)

func init() {
}

func Test_MaxMatch(t *testing.T) {
	seg := new(MaxMatchSeg)
	fmt.Println(seg.Cut([]rune("阳光灿烂的日子")))
	fmt.Println(seg.Cut([]rune("武汉市长江大桥")))
	hmmseg := new(HMMSeg)
	fmt.Println(hmmseg.CutDGA([]rune("阳光灿烂的日子")))
	fmt.Println(hmmseg.CutDGA([]rune("武汉市长江大桥")))
	fmt.Println(hmmseg.CutAll([]rune("阳光灿烂的日子")))
	fmt.Println(hmmseg.CutAll([]rune("武汉市长江大桥")))
}
