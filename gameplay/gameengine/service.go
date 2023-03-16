// Copyright (C) 2021 - 2023 Creditor Corp. Group.
// See LICENSE for copying information.

package gameengine

import (
	"sort"

	"github.com/zeebo/errs"
)

// ErrGameEngine indicates that there was an error in the service.
var ErrGameEngine = errs.Class("game engine service error")

// Service is handling clubs related logic.
//
// architecture: Service
type Service struct {
}

// NewService is a constructor for game engine service.
func NewService() *Service {
	return &Service{}
}

const (
	minPlace = 0
	maxPlace = 83
)

// GetPlayersMoves get all player possible moves.
func (service *Service) GetPlayersMoves(playerPlace int) ([]int, error) {
	top := []int{77, 70, 63, 56, 49, 42, 35, 28, 21, 14, 7, 0}
	bottom := []int{6, 13, 20, 27, 34, 41, 48, 55, 62, 69, 76, 83}
	exceptions := []int{1, 5, 78, 82}

	if playerPlace < 0 || playerPlace > 83 {
		return []int{}, ErrGameEngine.New("player place can not be more 83 or les than 0, player place is %d", playerPlace)
	}
	var stepInWidth []int

	switch {
	case contains(top, playerPlace):
		stepInWidth = append(stepInWidth, playerPlace, playerPlace+1, playerPlace+2)

	case contains(bottom, playerPlace):
		stepInWidth = append(stepInWidth, playerPlace-2, playerPlace-1, playerPlace)

	case contains(exceptions, playerPlace):
		stepInWidth = append(stepInWidth, playerPlace-1, playerPlace, playerPlace+1)

	case playerPlace == 8:
		stepInWidth = append(stepInWidth, playerPlace-1, playerPlace, playerPlace+1, playerPlace+2)

	case playerPlace == 12:
		stepInWidth = append(stepInWidth, playerPlace-2, playerPlace-1, playerPlace, playerPlace+1)

	default:
		stepInWidth = append(stepInWidth, playerPlace-2, playerPlace-1, playerPlace, playerPlace+1, playerPlace+2)
	}

	var moves []int

	for _, w := range stepInWidth {
		min := w - 14
		max := w + 14
		moves = append(moves, min, min+7, max-7, max, w)
	}

	sort.Ints(moves)
	moves = removeMin(moves, minPlace)
	moves = removeMax(moves, maxPlace)

	return moves, nil
}

func removeMin(l []int, min int) []int {
	for i, v := range l {
		if v >= min {
			return l[i:]
		}
	}
	return l
}
func removeMax(l []int, max int) []int {
	for i, v := range l {
		if v > max {
			return l[:i]
		}
	}
	return l
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
