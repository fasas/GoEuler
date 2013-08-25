package euler

import (

)

/**
* Project Euler #18 :
* 
* By starting at the top of the triangle below and moving
* to adjacent numbers on the row below, Find the maximum total
* from top to bottom of the triangle.
*
* NOTE: the triangle is actually equalateral.
* See the appropriate project euler page. It's only
* shown different for code formatting purposes.
*/

var triangle = [][]int{
	{75},
	{95, 64,},
	{17, 47, 82,},
	{18, 35, 87, 10},
	{20, 4, 82, 47, 65},
	{19, 1, 23, 75, 3, 34},
	{88, 2, 77, 73, 7, 63, 67},
	{99, 65, 4, 28, 6, 16, 70, 92,},
	{41, 41, 26, 56, 83, 40, 80, 70, 33},
	{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23}}

func Euler18() int {
	triangleHeight := len(triangle)
	
	for i := triangleHeight-2; i >=0; i-- {
		traceTriangle(i)
	}
	return triangle[0][0]
}

func traceTriangle(row int) {
	cols := len(triangle[row])-1
	for col := 0; col <= cols; col++ {
		mergeTriangle(row, col)
	}
}

func mergeTriangle(row int, col int) {
	left := triangle[row+1][col]
	right := triangle[row+1][col+1]
	
	if left > right {
		triangle[row][col] += left
	} else {
		triangle[row][col] += right
	}
}
