package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Case struct {
	CaseNum  int
	Credit   int
	Items    int
	Prices   Ints
	SolveFor Prices
	Spent    int
	Purchase Ints
}

type Price struct {
	price, index int
}

type Prices []*Price

//go:generate goast write impl github.com/jamesgarfield/sliceops
type Ints []int
type Cases []*Case

//go:generate goast write impl pipeline.go
type CaseChan chan *Case

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	in := bufio.NewScanner(file)
	ans := solve(parse(in))
	for _, a := range ans {
		fmt.Println(a)
	}
}

func solve(c []*Case) (ans []string) {
	var cases Cases = c
	done := make(chan bool)

	var pipe CaseChan = make(chan *Case)

	w := runtime.NumCPU()

	go func() {
		for _, sc := range cases {
			pipe <- sc
		}
	}()

	var solved Cases = pipe.Fan(done, w, solveCase).Collect(done, cases.Len())

	solved.Sort(func(a, b *Case) bool { return a.CaseNum < b.CaseNum })

	for _, a := range solved {
		ans = append(ans, formatAnswer(a.CaseNum, a.Purchase))
	}

	return
}

func formatAnswer(n int, ans []int) string {
	return fmt.Sprintf("Case #%d: %d %d", n, ans[0], ans[1])
}

func solveCase(c *Case) *Case {
	c.setup()
	subcases := c.divide()

	done := make(chan bool)

	var pipe CaseChan = make(chan *Case)
	w := runtime.NumCPU()

	go func() {
		for _, sc := range subcases {
			pipe <- sc
		}
	}()

	answers := pipe.Fan(done, w, solveSubCase).Filter(done, func(c *Case) bool { return c != nil }).Collect(done, 1)
	ans := answers[0]
	return ans
}

func solveSubCase(c *Case) *Case {
	for i := 0; i < c.SolveFor.Len(); i++ {
		price := c.SolveFor[i]
		newSpent := price.price + c.Spent
		//Value is too large, move to next value
		if newSpent > c.Credit {
			continue
		}

		c.Spent = newSpent
		c.Purchase = append(c.Purchase, price.index)

		//Failed to find answer in 2 items
		if c.Purchase.Len() == 2 && newSpent < c.Credit {
			return nil
		}
	}

	if c.Spent < c.Credit {
		return nil
	}

	c.Purchase.Sort(func(a, b int) bool { return a < b })
	return c
}

func (c *Case) setup() {
	for idx, p := range c.Prices {
		c.SolveFor = append(c.SolveFor, &Price{p, idx + 1})
	}
	c.SolveFor = c.SolveFor.Where(func(p *Price) bool { return p.price < c.Credit })
	c.SolveFor.Sort(func(a, b *Price) bool { return b.price < a.price })
}

func (c *Case) divide() (result Cases) {
	for i := 0; i < c.SolveFor.Len(); i++ {
		newCase := &Case{
			CaseNum:  c.CaseNum,
			Credit:   c.Credit,
			Items:    c.Items,
			Prices:   c.Prices,
			SolveFor: c.SolveFor[i:],
		}
		result = append(result, newCase)
	}
	return
}

func (s Ints) String() string {
	strs := []string{}
	for _, i := range s {
		strs = append(strs, strconv.Itoa(i))
	}
	return strings.Join(strs, ", ")
}

func (s Prices) String() string {
	strs := []string{}
	for _, i := range s {
		strs = append(strs, strconv.Itoa(i.price))
	}
	return strings.Join(strs, ", ")
}
