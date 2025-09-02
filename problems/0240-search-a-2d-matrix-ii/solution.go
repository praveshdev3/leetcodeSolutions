func searchMatrix(matrix [][]int, target int) bool {
    rows := len(matrix)
    cols := len(matrix[0])
    row := 0
    col := cols-1

    for row < rows && col >= 0{
        if target > matrix[row][col]{
            row++
        }else if target < matrix[row][col]{
            col--
        }else{
            return true
        }
    }

    return false
}
