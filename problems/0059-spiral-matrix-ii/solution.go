func generateMatrix(n int) [][]int {
    matrix := make([][]int, n)
    
    for i := range matrix{
        matrix[i] = make([]int, n)
    }

    if n == 0{
        return matrix
    }

    left := 0
    right := n-1
    top := 0
    bottom := n-1
    num := 1

    for left <= right && top <= bottom {
        for i := left; i <= right; i++{
            matrix[top][i] = num
            num++
        } 
        top++

        for i := top; i <= bottom; i++{
            matrix[i][right] = num
            num++
        } 
        right--

        for i := right; i >= left; i--{
            matrix[bottom][i] = num
            num++
        } 
        bottom--

        for i := bottom; i >= top; i--{
            matrix[i][left] = num
            num++
        } 
        left++
    }
    return matrix
}
