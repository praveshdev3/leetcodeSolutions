func shipWithinDays(weights []int, days int) int {
    low := 0
    high := 0
    for _,weight := range weights{
        if weight > low{
           low = weight
        }
        high += weight
    }
    ans := 0
    for low <= high{
        mid := (low+high)/2
        if !isValidDays(weights,mid,days){
            low = mid+1
        }else{
            ans = mid
            high = mid-1
        }
    }
    return ans
}

func isValidDays(weights []int, mid,days int) bool {
    dayCount := 0
    sum := 0
    for _,weight := range weights{
        if sum+weight <= mid{
            sum += weight
        }else{
            dayCount += 1
            sum = weight
        }
    }
    if sum > 0{
        dayCount += 1
    }
    if dayCount <= days{
        return true
    }
    return false
}
