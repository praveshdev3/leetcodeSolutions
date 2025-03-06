func maxTwoEvents(events [][]int) int {
    sort.Slice(events, func(i,j int) bool {
        return events[i][1] < events[j][1]
    })

    n := len(events)
    maxArr := make([]int,n)
    maxArr[0] = events[0][2]

    for i:=1; i<n; i++{
        maxArr[i] = max(events[i][2],maxArr[i-1])
    }

    maxSum := 0

    for i:=0; i<n; i++{
        start := events[i][0]
        value := events[i][2]

        idx := binarySearch(events,start-1)
        currentSum := value
        if idx != -1{
            currentSum += maxArr[idx]
        }

        maxSum = max(maxSum,currentSum)
    }
    return maxSum
}

func binarySearch(events [][]int, targetEnd int)int{
    start, end, res := 0, len(events)-1, -1
    for start <= end{
        mid := start + (end-start)/2
        if events[mid][1] <= targetEnd{
            res=mid
            start=mid+1
        }else{
            end = mid-1
        }
    }
    return res
}

func max(a,b int) int{
    if a>b {
        return a
    }
    return b
}
