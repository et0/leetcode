package lh2402

import "container/heap"

type Item struct {
	room int
	end  int
}

type EndMeeting []*Item

type IntHeap []int

func NewIntHeap() *IntHeap {
	h := &IntHeap{}
	heap.Init(h)

	return h
}

// Len возвращает длину кучи
func (h IntHeap) Len() int { return len(h) }

// Less определяет порядок элементов (min-heap: h[i] < h[j])
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap меняет местами элементы
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push добавляет элемент в кучу
func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Pop извлекает минимальный элемент (корень кучи)
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]  // Берём последний элемент
	*h = old[:n-1] // Уменьшаем длину кучи
	return x
}

func NewEndMeeting() *EndMeeting {
	h := &EndMeeting{}
	heap.Init(h)

	return h
}

// Len возвращает длину кучи
func (h EndMeeting) Len() int { return len(h) }

// Less определяет порядок элементов (min-heap: h[i] < h[j])
func (h EndMeeting) Less(i, j int) bool {
	// если время завершения равное, то берём меньшую комнату
	if h[i].end == h[j].end {
		// fmt.Println("LESS", h[i], h[j])
		return h[i].room < h[j].room
	}

	return h[i].end < h[j].end
}

// Swap меняет местами элементы
func (h EndMeeting) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push добавляет элемент в кучу
func (h *EndMeeting) Push(x any) {
	*h = append(*h, x.(*Item))
}

// Pop извлекает минимальный элемент (корень кучи)
func (h *EndMeeting) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]  // Берём последний элемент
	*h = old[:n-1] // Уменьшаем длину кучи
	return x
}
