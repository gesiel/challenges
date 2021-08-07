package main

func bestRhombicAreaFrame(matrix [][]int, radius int) []int {
	result := []int{-1, -1, -1}
	resultMap := make(map[int]bool)

	lines := len(matrix)
	columns := len(matrix[0])

	for l := radius-1; l <= lines-radius; l++ {
		for c := radius-1; c <= columns-radius; c++ {
			sum := matrix[l][c]
			for r := 1; r < radius; r++ {
				sum += matrix[l+r][c] + matrix[l-r][c] + matrix[l][c+r] + matrix[l][c-r]

				for r2 := 1; r >= 2 && r2 < r; r2++ {
					sum += matrix[l+r-r2][c+r2] + matrix[l+r-r2][c-r2] + matrix[l-r+r2][c+r2] + matrix[l-r+r2][c-r2]
				}
			}

			if !resultMap[sum] {
				resultMap[sum] = true

				if sum > result[0] {
					result[2] = result[1]
					result[1] = result[0]
					result[0] = sum
				} else if sum > result[1] {
					result[2] = result[1]
					result[1] = sum
				} else if sum > result[2] {
					result[2] = sum
				}
			}
		}
	}

	var r []int
	for _,v := range result {
		if v > -1 {
			r = append(r, v)
		}
	}
	return r
}

/*

Let's define a rhombic area of size k as the set of all cells that are less than k steps away (up, down, left, or right)
from a given center cell in a matrix.

In the image below, the black cell is the center of the rhombic area and every cell contains a number - the distance to
the center, which is called radius. Cells with the same radius are colored in the same color.

Given a rectangular matrix of integers matrix, and an integer radius, your task is the following:

For each of the possible center cells of matrix, find the sum of all elements on the border of the surrounding rhombic a
rea of size radius. Note: only consider the rhombic areas whose elements all fit within the bounds of matrix. From among
these sums, find the 3 highest distinct values. Return an array containing the 3 highest distinct sums, in order of
decreasing value. If there are fewer than 3 distinct sums, return a shorter array containing all the distinct sums in
decreasing order.

For
  [[4, 1, 4, 5, 3],
   [6, 2, 7, 2, 1],
   [10, 0, 4, 2, 7]]
and
	radius = 2,
the output should be bestRhombicAreaFrame(matrix, radius) = [15, 14, 12].

=====

[input] array.array.integer matrix

A rectangular matrix of integers.

Guaranteed constraints:
1 ≤ matrix.length ≤ 100,
1 ≤ matrix[i].length ≤ 100,
0 ≤ matrix[i][j] ≤ 100.

[input] integer radius

An integer representing the radius of the rhombic areas.

Guaranteed constraints:
1 ≤ 2·radius - 1 ≤ min(matrix.length, matrix[i].length).

[output] array.integer

*/
