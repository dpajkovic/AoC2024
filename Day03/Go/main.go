//	Copyright (c) Miloš Rackov 2024
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package main

import (
	"AoC2024D03/part1"
	"AoC2024D03/part2"
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := scanner.Text()
		input = append(input, l)
	}

	p1 := part1.Part1(input)
	fmt.Printf("Result for part 1: %d\n", p1)
	p2 := part2.Part2(input)
	fmt.Printf("Result for part 2: %d\n", p2)
}
