func minEatingSpeed(piles []int, h int) int {
    maxSpeed := 0
    for _,pile := range piles{
        if pile>maxSpeed{
            maxSpeed = pile
        }
    }
    low := 1
    high := maxSpeed
    ans := 0
    for low <= high{
        mid := (low+high)/2
        if hoursNeeded(piles, mid) > h{
            low = mid+1
        }else{
            ans = mid
            high = mid-1
        }
    }

    return ans
}

func hoursNeeded(piles []int, k int) int{
    hours := 0
    for _,pile := range piles{
        hr := pile/k
        hours += hr
        if pile%k != 0{
            hours += 1
        }
    }
    return hours
}
