package gcws

import (
	"testing"
	"fmt"
)

func TestGet_DAG(t *testing.T) {
	str := []rune("阳光灿烂的日子")
	dag := Get_DAG(str)
	fmt.Println(dag)
	router := Calc(str, dag)
	fmt.Println(router)
}
