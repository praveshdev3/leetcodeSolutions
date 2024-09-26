func countPoints(points [][]int, queries [][]int) []int {
    var result []int
    for _,query := range queries {
        x := float64(query[0])
        y := float64(query[1])
        radius := float64(query[2])
        count := 0
        for _,point := range points{
            if (float64(point[0]) - x)*(float64(point[0]) - x) + (float64(point[1]) - y)*(float64(point[1]) - y) <= radius * radius {
                count++
            }
        }
        result = append(result, count)
    }
    return result
}   
