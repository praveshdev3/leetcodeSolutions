func splitArray(nums []int, k int) int {
    max := 0
    sum := 0

    for _,num := range nums{
        if num > max{
            max = num
        }
        sum += num
    }
    
    left := max
    right := sum 

    for left < right{
        mid := (left + right) / 2
        if canSplit(nums,k,mid) {
            right = mid
        }else{
            left = mid + 1
        }
    }
    return left
}

func canSplit(nums []int, k int, maxSum int) bool{
    count := 1
    sum := 0
    for _, num := range nums{
        sum += num
        if sum > maxSum{
            count++
            sum = num
        }
        if count > k{
            return false
        }
    }
    return true
}
