package lh2402 // https://leetcode.com/problems/meeting-rooms-iii/description/

import (
	"container/heap"
	"sort"
)

func mostBooked(n int, meetings [][]int) int {
	// Информация о кол-ве встреч по комнатам
	rooms := make([]int, n)

	// Список свободных комнат
	freeRooms := NewIntHeap()
	for i := range n {
		heap.Push(freeRooms, i)
	}

	// Список, когда заканчивается встреча
	endMeeting := NewEndMeeting()

	// Сортируем встречи по времени их начала
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	for _, m := range meetings {
		// Освобождаем комнаты, если есть занятые и время встречи уже прошло
		for endMeeting.Len() > 0 && (*endMeeting)[0].end <= m[0] {
			r := heap.Pop(endMeeting).(*Item)
			heap.Push(freeRooms, r.room)
		}

		// если есть свободные комнаты
		if freeRooms.Len() > 0 {
			// извлекаем номер свободной комнаты
			room := heap.Pop(freeRooms).(int)

			// увеличиваем счётчик проведенных встреч в комнатах
			rooms[room]++

			// добавляем в кучу время освобождения и номер комнаты
			heap.Push(endMeeting, &Item{room, m[1]})

			continue
		}

		// извлекаем время и комнату, которая освободится в ближайщее время
		item := heap.Pop(endMeeting).(*Item)

		// увеличиваем счётчик проведенных встреч в комнатах
		rooms[item.room]++

		// добавляем время освобождения
		if item.end > m[0] {
			// встреча отложена, до начала первой освобождающейся комнаты
			heap.Push(endMeeting, &Item{item.room, item.end + (m[1] - m[0])})
		} else {
			// Одна из комнат освободилась раньше, чем должна начаться встреча
			heap.Push(endMeeting, &Item{item.room, m[1]})
		}
	}

	min := 0
	for i := 1; i < n; i++ {
		if rooms[i] > rooms[min] {
			min = i
		}
	}

	return min
}

func Wrapper(n int, meetings [][]int) int {
	return mostBooked(n, meetings)
}
