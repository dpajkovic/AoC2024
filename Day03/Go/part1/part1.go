// Copyright (c) Miloš Rackov 2024
// Distributed under the Boost Software License, Version 1.0.
// (See accompanying file LICENSE or copy at
// https://www.boost.org/LICENSE_1_0.txt)
package part1

import (
	"regexp"
	"strconv"
)

func Part1(input []string) int {
	s := 0
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	for _, i := range input {
		m := re.FindAllStringSubmatch(i, -1)
		for _, v := range m {
			a, _ := strconv.Atoi(v[1])
			b, _ := strconv.Atoi(v[2])
			s += a * b
		}
	}

	return s
}
