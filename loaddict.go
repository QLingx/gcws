package gcws

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var (
	FREQ  = make(map[string]float64)
	TOTAL float64
)

func init() {
	FREQ, TOTAL = Gen_Pfdict("./dict.txt")
}

func Gen_Pfdict(fileName string) (map[string]float64, float64) {
	var total = 0.0
	var lfreq = make(map[string]float64)
	fd, err := os.OpenFile(fileName, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
	buf := bufio.NewReader(fd)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}
		//解析line
		linearr := strings.Split(line, " ")
		word := linearr[0]
		freq, err := strconv.ParseFloat(linearr[1], 64)
		if err != nil {
			freq = 0
		}
		lfreq[word] = freq
		total += freq
		for i := 0; i < len(word); i++ {
			wfrag := string(word[0:i+1])
			if _, ok := lfreq[wfrag]; !ok {
				lfreq[wfrag] = 0
			}
		}
	}
	fd.Close()
	return lfreq, total
}
