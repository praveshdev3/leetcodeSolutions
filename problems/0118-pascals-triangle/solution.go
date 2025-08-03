func generate(numRows int) [][]int {
    if numRows <= 0{
        return [][]int{}
    }

    if numRows == 1{
        return [][]int{[]int{1}}
    }

    res := make([][]int, numRows)
    res[0] = []int{1}

    for i:=1; i<numRows; i++{
        arr := make([]int,i+1)
        for j:=0; j<len(arr); j++{
            if j==0 || j==len(arr)-1{
                arr[j] = 1
            }else{
                arr[j] = res[i-1][j] + res[i-1][j-1]
            }
        }
        res[i] = arr
    }

    return res
}
