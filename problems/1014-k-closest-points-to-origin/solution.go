func kClosest(points [][]int, k int) [][]int {
    for i:=1; i<len(points); i++{
        if i>=1{
        distance1 := points[i-1][0] * points[i-1][0] + points[i-1][1] * points[i-1][1]
        distance2 := points[i][0] * points[i][0] + points[i][1] * points[i][1]
        if distance1 > distance2{
            points[i], points[i-1] = points[i-1], points[i]
            i-=2
            }
        }
    }
    return points[:k]
}
