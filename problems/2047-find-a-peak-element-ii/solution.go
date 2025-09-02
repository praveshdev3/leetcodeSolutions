import "fmt"
func findPeakGrid(mat [][]int) []int {
	cols := len(mat[0])

    low := 0
    high := cols - 1

    for low <= high{
        mid := (low + high)/2
        maxRowIdx := maxInColumn(mid, mat)
        if (mid == 0 || mat[maxRowIdx][mid] > mat[maxRowIdx][mid-1]) &&
        (mid == cols-1 || mat[maxRowIdx][mid] > mat[maxRowIdx][mid+1]){
            return []int{maxRowIdx, mid}
        }else if mid > 0 && mat[maxRowIdx][mid] < mat[maxRowIdx][mid-1]{
            high = mid - 1
        }else{
            low = mid + 1
        }
    }

	return []int{-1, -1}
}

func maxInColumn(colIdx int, mat [][]int) int {
    maxRowIdx := -1
    max := -1
    for i := 0; i < len(mat); i++{
        if mat[i][colIdx] > max{
            maxRowIdx = i   
            max = mat[i][colIdx]
        }
    }
    return maxRowIdx
}
