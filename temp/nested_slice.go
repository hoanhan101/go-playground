package main

import (
	"fmt"
	"strconv"
	"strings"
)

type config struct {
	n        int
	endnames [][]string
	ends     []string
}

func main() {
	cfg := &config{}

	cfg.n = 5
	cfg.endnames = make([][]string, cfg.n)

	fmt.Println(cfg.endnames)

	for i := 0; i < cfg.n; i++ {
		cfg.start1(i)
	}

	fmt.Println(cfg.endnames)
	fmt.Println(cfg.ends)
}

func (cfg *config) start1(i int) {

	cfg.endnames[i] = make([]string, cfg.n)
	for j := 0; j < cfg.n; j++ {
		s := []string{strconv.Itoa(i), strconv.Itoa(j)}
		cfg.endnames[i][j] = strings.Join(s, "")
	}

	cfg.ends = make([]string, cfg.n)
	for j := 0; j < cfg.n; j++ {
		cfg.ends[j] = cfg.endnames[i][j]
	}
}
