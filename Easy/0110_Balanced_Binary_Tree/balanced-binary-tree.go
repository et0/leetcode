package balancedbinarytree

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	// Cтруктура с деревом (указатель на корень) и длинной его поддеревьев
	type depth struct {
		root  *TreeNode
		left  float64 // глубина левого поддерева
		right float64 // глубина правого поддерева
	}

	// стек с корнями
	stack := make([]depth, 0, 5000)

	// флаг, нужен для метки когда поднимаемся обратно вверх по дереву
	flag := true

	for {
		if root == nil {
			break
		}

		// Если флаг положительный - это новая ветка. Если отрицательный - старая ветка
		if flag {
			stack = append(stack, depth{root, 0, 0})
		} else {
			flag = true
		}

		if root.Left != nil {
			root = root.Left
			root.Val = 0                        // изменяем значение, если 0 - ушли влево
			stack[len(stack)-1].left++          // увеличиваем счетчик переходов влево, т.к на следующей итерации уходим влево
			stack[len(stack)-1].root.Left = nil // обнуляем переход влево, что бы не зацикливаться

			continue
		}

		if root.Right != nil {
			root = root.Right
			root.Val = 1                         // изменяем значение, если 1 - ушли вправо
			stack[len(stack)-1].right++          // увеличиваем счетчик переходов вправо, т.к на следующей итерации уходим вправо
			stack[len(stack)-1].root.Right = nil // обнуляем переход вправо, что бы не зацикливаться

			continue
		}

		// текущий элемент не имеет веток, сравниваем длину поддеревьев
		if math.Abs(stack[len(stack)-1].left-stack[len(stack)-1].right) > 1.0 {
			return false
		}

		// запоминаем элемент и кол-во веток
		val, left, right := stack[len(stack)-1].root.Val, stack[len(stack)-1].left, stack[len(stack)-1].right
		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			if val == 0 { // если элемент "левый", то прибавляем к родителю слева длину ветки
				stack[len(stack)-1].left += math.Max(left, right)
			} else { // если элемент "правый", то прибавляем к родителю справа длину ветки
				stack[len(stack)-1].right += math.Max(left, right)
			}

			root = stack[len(stack)-1].root

			// ставим флаг, что бы не создавался новый элемент в стеке, т.к. мы поднимаемся наверх по дереву
			flag = false

			continue
		}

		break
	}

	return true
}

func Wrapper(root *TreeNode) bool {
	return isBalanced(root)
}
