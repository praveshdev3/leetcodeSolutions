func searchMatrix(matrix [][]int, target int) bool {
    rows := len(matrix)
    cols := len(matrix[0])

    top := 0
    bottom := rows-1
    rowToSearch := -1

    for top <= bottom{
        mid := (top + bottom)/2
        if matrix[mid][0] <=target && target <= matrix[mid][cols-1]{
            rowToSearch = mid
            break
        }else if matrix[mid][0] > target{
            bottom = mid-1
        }else{
            top = mid+1
        }
    }

    if rowToSearch == -1{
        return false
    }

    left := 0
    right := cols-1
    for left <= right{
        mid := (left+right)/2
        if matrix[rowToSearch][mid] == target{
            return true
        }else if matrix[rowToSearch][mid] > target{
            right = mid - 1
        }else {
            left = mid + 1
        }
    }

    return false
}
