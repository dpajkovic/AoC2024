//	Copyright (c) Miloš Rackov 2024
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import "fmt"

func Part2(input []string) int {

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

	s := 0

	for i := range l {
		s += count(l[i], r) * l[i]
	}

	return s
}

func count(x int, s []int) int {
	c := 0
	for _, v := range s {
		if v == x {
			c++
		}
	}
	return c
}
