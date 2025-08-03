func spiralOrder(matrix [][]int) []int {
    left := 0
    right := len(matrix[0])-1
    top := 0
    bottom := len(matrix)-1

    spiral := make([]int, 0)


    for left <= right || top <= bottom{
        for i:=left; i<=right; i++{
            spiral = append(spiral, matrix[top][i])
        }
        top++
        if top > bottom{
            break
        }
        for i:=top; i<=bottom; i++{
            spiral = append(spiral, matrix[i][right])
        }
        right--
        if left > right{
            break
        }
        for i:=right; i>=left; i--{
            spiral = append(spiral, matrix[bottom][i])
        }
        bottom--
        if top > bottom{
            break
        }
        for i:=bottom; i>=top; i--{
            spiral = append(spiral, matrix[i][left])
        }
        left++
        if left > right{
            break
        }
    }

    return spiral
}
