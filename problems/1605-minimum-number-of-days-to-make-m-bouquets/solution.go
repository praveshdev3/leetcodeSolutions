func minDays(bloomDay []int, m int, k int) int {
    if m*k > len(bloomDay){
        return -1
    }

    maxDays := 0
    minDays := bloomDay[0]
    for _,day := range bloomDay{
        if day > maxDays{
            maxDays = day
        }
        if day < minDays{
            minDays = day
        }
    }

    low := minDays
    high := maxDays

    for low <= high{
        mid := (low+high)/2
        if !isValidDay(bloomDay, m, k, mid){
            low = mid+1
        }else{
            high=mid-1
        }
    }
    return low
}

func isValidDay(bloomDay []int, m,k, day int) bool{
    count := 0
    maxBouquets := 0
    for _,d := range bloomDay{
        if day >= d{
            count++
        }else{
            maxBouquets += count/k
            count = 0
        }
    }
    maxBouquets += count/k
    if maxBouquets >= m{
        return true
    }
    return false
}
