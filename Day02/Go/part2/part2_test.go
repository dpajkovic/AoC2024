//	Copyright (c) Miloš Rackov 2024
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package part2

import (
	"bufio"
	"os"
	"testing"
)

func TestPart2(t *testing.T) {
	file1, err := os.Open("../testinput.txt")
	if err != nil {
		panic(err)
	}

	var testInput []string

	scanner := bufio.NewScanner(file1)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := scanner.Text()
		testInput = append(testInput, l)
	}

	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test input",
			args: args{input: testInput},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.input); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
