package pascaltriangle

import "fmt"

func Generate(numRows int) [][]int {
    triangle := make([][]int, numRows, numRows)
    for i := 0; i < numRows; i++ {
        if i == 0 {
            triangle[i] = []int{ 1 }
        } else {
            triangle[i] = make([]int, i + 1, i + 1)
            for j := 0; j < len(triangle[i]); j++ {
                if j == 0 {
                    triangle[i][j] = 1
                } else if j == i {
                    triangle[i][j] = 1
                } else {
                    triangle[i][j] = triangle[i - 1][j - 1] + 
                    triangle[i - 1][j]
                }
            }
        }
    }

    fmt.Printf("%v", triangle)

    return triangle
}
