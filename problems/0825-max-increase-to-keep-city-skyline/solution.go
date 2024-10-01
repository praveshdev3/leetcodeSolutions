func maxIncreaseKeepingSkyline(grid [][]int) int {
    maxRow := make([]int,len(grid))
    maxCol := make([]int,len(grid))

    for i := range len(grid) {
        for j,v := range grid[i] {
            if maxRow[i] < grid[j][i] {
                maxRow[i] = grid[j][i]
            }
            if maxCol[i] < v {
                maxCol[i] = v
            }
        }
    }
    sum := 0
      for i := range len(grid) {
        for j,v := range grid[i] {
            sum += min(maxRow[i],maxCol[j]) - v
        }
    }
    return sum 
}
