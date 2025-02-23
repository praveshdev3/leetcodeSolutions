func numIslands(grid [][]byte) int {
    count := 0

    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[i]); j++{
            if grid[i][j] == '1'{
                count++
                dfs(grid, i, j)
            }
        }
    }

    return count
}

func dfs(grid [][]byte, r,c int){
    if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[r]) || grid[r][c] != '1'{
        return
    }

    grid[r][c] = '2'
    dfs(grid, r+1, c)
    dfs(grid, r-1, c)
    dfs(grid, r, c+1)
    dfs(grid, r, c-1)
}
