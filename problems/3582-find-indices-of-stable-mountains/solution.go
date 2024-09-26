func stableMountains(height []int, threshold int) []int {
    var indices []int
    for index,_ := range height{
        if index > 0 && height[index-1] > threshold {
            indices = append(indices,index)
        }
    }
    return indices
}
