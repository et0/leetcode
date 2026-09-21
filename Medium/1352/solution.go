package main // https://leetcode.com/problems/product-of-the-last-k-numbers/description/

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
