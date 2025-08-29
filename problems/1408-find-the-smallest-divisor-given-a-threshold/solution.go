func smallestDivisor(nums []int, threshold int) int {
    high := 0
    for _,num := range nums{
        if num > high{
            high = num
        }
    }
    low := 1
    ans := -1
    for low <= high{
        mid := (low+high)/2
        if !isValid(nums,threshold,mid){
            low = mid+1
        }else{
            ans=mid
            high = mid-1
        }
    }
    return ans
}

func isValid(nums []int, threshold, mid int) bool {
    sum := 0
    for _,num := range nums{
        n := num/mid
        sum += n
        if num%mid != 0{
            sum += 1
        }
    }
    if sum <= threshold{
        return true
    }
    return false
}
