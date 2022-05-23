package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func splitThousands(in, separator string) string {
	if separator == "" {
		return in
	}
	sep := []byte(separator)
	numOfDigits := len(in)
	if in[0] == '-' {
		numOfDigits--
	}
	numOfCommas := (numOfDigits - 1) / 3
	out := make([]byte, len(in)+numOfCommas*len(sep))
	if in[0] == '-' {
		in, out[0] = in[1:], '-'
	}
	k := 0
	for i, j := len(in)-1, len(out)-1; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			k = 0
			j -= len(sep)
			copy(out[j:], sep)
		}
	}
}

func main() {
	var sep string
	flag.StringVar(&sep, "sep", " ", "thousands separator")
	var maxExp int64
	flag.Int64Var(&maxExp, "max", 64, "maximum exponent")
	flag.Parse()

	if maxExp < 0 {
		fmt.Println("exponent must not be negative")
		os.Exit(1)
	}

	max := new(big.Int).Exp(big.NewInt(2), big.NewInt(maxExp), nil)
	var ss []string
	for i := big.NewInt(1); i.Cmp(max) < 1; i.Lsh(i, 1) {
		ss = append(ss, splitThousands(i.String(), sep))
	}

	width := len(ss[len(ss)-1])
	expWidth := len(strconv.FormatInt(maxExp, 10))
	for i, s := range ss {
		fmt.Printf("%*d: %*s\n", expWidth, i, width, s)
	}
}
