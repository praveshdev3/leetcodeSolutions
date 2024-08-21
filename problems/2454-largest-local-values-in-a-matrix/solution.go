func largestLocal(grid [][]int) [][]int {
    n := len(grid)
    resp := make([][]int,n-2)

    for i := 0; i < n-2; i++{
        resp[i] = make([]int, n-2)
    }

    for i:=0; i<n-2; i++{
        for j:=0; j<n-2; j++{
            resp[i][j] = findMax(grid,i,j)
        }
    }

    return resp
}

func findMax(grid [][]int, i, j int) int {
    maxElement := 0
    for k:=i; k<i+3; k++{
        for m:=j; m<j+3; m++{
            if maxElement < grid[k][m]{
                maxElement = grid[k][m]
            }
        }
    }
    return maxElement
}



