package lm2402 // https://leetcode.com/problems/maximum-matching-of-players-with-trainers/description/

import (
	"sort"
)

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)

	out := 0
	sizeP, sizeT := len(players), len(trainers)
	for i, j := 0, 0; i < sizeP && j < sizeT; {
		if players[i] > trainers[j] {
			j++

			continue
		}

		out++

		i++
		j++
	}

	return out
}

func Wrapper(players []int, trainers []int) int {
	return matchPlayersAndTrainers(players, trainers)
}
