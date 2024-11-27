func hIndex(citations []int) int {
    newArr := make([]int, len(citations)+1)
    for _,num := range citations{
        if num >= len(citations){
            newArr[len(citations)] = newArr[len(citations)] + 1
        } else {
            newArr[num] = newArr[num] + 1
        }
    }
fmt.Println(newArr)
    citation := 0
    for i:=len(newArr)-1;i>=0;i--{
        citation += newArr[i]
        if i <= citation{
            return i
        }
    }
    return 0
}
