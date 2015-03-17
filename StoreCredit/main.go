package main

type Case struct {
	Credit   int
	Items    int
	Prices   []int
	Purchase []int
}

//go:generate goast write impl pipeline.go
type CaseChan chan *Case

func main() {

}
