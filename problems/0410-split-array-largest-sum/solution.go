func splitArray(nums []int, k int) int {
    high := 0
    low := 0

    for _,num := range nums{
        if num > low{
            low = num
        }
        high += num
    }

    for low <= high{
        mid := (low+high)/2
        p := isValid(nums,mid)
        if p > k{
            low = mid+1
        }else{
            high = mid-1
        }
    }

    return low
}

func isValid(nums []int, mid int)int{
    sum := 0
    count := 0
    for _,num := range nums{
        if sum + num <= mid{
            sum += num
        }else{
            count++
            sum = num
        }
    }
    return count+1
}
