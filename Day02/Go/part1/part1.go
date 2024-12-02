//	Copyright (c) Miloš Rackov 2024
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part1

import (
	"strconv"
	"strings"
)

func Part1(input []string) int {
	s := 0

	for _, l := range input {
		t := strings.Split(l, " ")
		r := []int{}
		for _, v := range t {
			i, _ := strconv.Atoi(v)
			r = append(r, i)
		}

		if isSafe(r) {
			s++
		}
	}

	return s
}

func isSafe(a []int) bool {
	var isInOrder func(a, b int) bool
	if isAsc(a[0], a[1]) {
		isInOrder = isAsc
	} else {
		isInOrder = isDesc
	}
	for i := 0; i < len(a)-1; i++ {
		if !isInLimit(a[i], a[i+1]) || !isInOrder(a[i], a[i+1]) {
			return false
		}
	}
	return true
}

func isInLimit(a, b int) bool {
	return abs(a-b) <= 3
}

func isAsc(a, b int) bool {
	return a < b
}

func isDesc(a, b int) bool {
	return a > b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
