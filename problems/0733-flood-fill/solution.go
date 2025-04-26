func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    colour := image[sr][sc]
    if colour != color{
        dfs(image,sr,sc,colour,color)
    }
    return image
}

func dfs(image [][]int, r, c, color, newColor int) {
	if image[r][c] == color {
		image[r][c] = newColor
		if r >= 1 {
			dfs(image, r-1, c, color, newColor)
		}
		if c >= 1 {
			dfs(image, r, c-1, color, newColor)
		}
		if r+1 < len(image) {
			dfs(image, r+1, c, color, newColor)
		}
		if c+1 < len(image[0]) {
			dfs(image, r, c+1, color, newColor)
		}
	}
}
