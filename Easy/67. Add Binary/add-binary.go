package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	sizeA, sizeB := len(a), len(b)
	sizeR := 0
	if sizeA > sizeB {
		sizeR = sizeA + 1
	} else {
		sizeR = sizeB + 1
	}

	r := make([]int, sizeR, sizeR)
	fmt.Println(r)
	sizeA--
	sizeB--
	sizeR--
	perenos := false
	for sizeA >= 0 && sizeB >= 0 {
		fmt.Println(string(a[sizeA]), string(b[sizeB]))
		if a[sizeA] == '0' && b[sizeB] == '0' {
			if perenos {
				r[sizeR] = 1
			} else {
				r[sizeR] = 0
			}
			perenos = false
		} else if (a[sizeA] == '1' && b[sizeB] == '0') || (a[sizeA] == '0' && b[sizeB] == '1') {
			if perenos {
				r[sizeR] = 0
			} else {
				r[sizeR] = 1
				perenos = false
			}
		} else {
			r[sizeR] = 0
			if perenos {
				r[sizeR] = 1
			} else {
				r[sizeR] = 0
			}
			perenos = true
		}

		sizeA--
		sizeB--
		sizeR--
	}

	if sizeA >= 0 {
		for sizeA >= 0 {
			fmt.Println(string(a[sizeA]))
			if a[sizeA] == '0' {
				if perenos {
					r[sizeR] = 1
				} else {
					r[sizeR] = 0
				}
				perenos = false
			} else {
				if perenos {
					r[sizeR] = 0
				} else {
					r[sizeR] = 1
					perenos = false
				}
			}

			sizeA--
			sizeR--
		}
	}

	if sizeB >= 0 {
		for sizeB >= 0 {
			fmt.Println(string(b[sizeB]))
			if b[sizeB] == '0' {
				if perenos {
					r[sizeR] = 1
				} else {
					r[sizeR] = 0
				}
				perenos = false
			} else {
				if perenos {
					r[sizeR] = 0
				} else {
					r[sizeR] = 1
					perenos = false
				}
			}

			sizeB--
			sizeR--
		}
	}

	if perenos {
		r[0] = 1
	}

	s := ""
	for i := len(r) - 1; i >= 0; i-- {
		if i == 0 && r[i] == 0 {
			continue
		}
		s = strconv.Itoa(r[i]) + s
	}

	fmt.Println(s)

	return ""
}

func main() {
	fmt.Println(addBinary("10", "101"))
}
