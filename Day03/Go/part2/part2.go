// Copyright (c) Miloš Rackov 2024
// Distributed under the Boost Software License, Version 1.0.
// (See accompanying file LICENSE or copy at
// https://www.boost.org/LICENSE_1_0.txt)
package part2

import (
	"regexp"
	"strconv"
)

func Part2(input []string) int {
	s := 0
	rb := regexp.MustCompile(`do\(\)(.*?)don't\(\)`)
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	i := "do()" + input[0] + "don't()"
	mb := rb.FindAllStringSubmatch(i, -1)

	newInput := ""
	for _, v := range mb {
		newInput += v[1]
	}

	m := re.FindAllStringSubmatch(newInput, -1)
	for _, v := range m {
		a, _ := strconv.Atoi(v[1])
		b, _ := strconv.Atoi(v[2])
		s += a * b
	}

	return s
}
