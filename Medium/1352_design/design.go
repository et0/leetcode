package lm1352_design // https://leetcode.com/problems/product-of-the-last-k-numbers/description/

import "fmt"

type ProductOfNumbers struct {
	last []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		last: []int{1},
	}
}

func (prod *ProductOfNumbers) Add(num int) {
	if num == 0 {
		prod.last = []int{1}
	} else {
		prod.last = append(prod.last, prod.last[len(prod.last)-1]*num)
	}
}

func (prod *ProductOfNumbers) GetProduct(k int) int {
	size := len(prod.last)

	if k >= len(prod.last) {
		return 0
	}

	return prod.last[size-1] / prod.last[size-1-k]
}

func Wrapper(todo []string, data []int) {
	obj := Constructor()

	for i := 1; i < len(todo); i++ {
		switch todo[i] {
		case "add":
			obj.Add(data[i])
		case "getProduct":
			fmt.Println(obj.GetProduct(data[i]))
		}
	}

	fmt.Println(obj)
}
