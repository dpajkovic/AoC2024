//	Copyright (c) Miloš Rackov 2024
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part1

import (
	"fmt"
	"sort"
)

func Part1(input []string) int {

	l, r := []int{}, []int{}

	for _, s := range input {
		// split s into l and r on space
		li, ri := 0, 0
		_, err := fmt.Sscanf(s, "%d %d", &li, &ri)
		if err != nil {
			panic(err)
		}
		l = append(l, li)
		r = append(r, ri)
	}
	// sort l and r
	sort.Ints(l)
	sort.Ints(r)
	d := 0
	for i := range l {
		d += abs(l[i] - r[i])
	}
	return d
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
