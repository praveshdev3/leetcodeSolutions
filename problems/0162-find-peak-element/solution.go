func findPeakElement(nums []int) int {
    low := 0
    high := len(nums)-1
    for low<=high{
        mid := (low+high)/2
        if mid > 0 && nums[mid] < nums[mid-1]{
            high = mid-1
        }else if mid < len(nums)-1 && nums[mid] < nums[mid+1]{
            low=mid+1
        }else{
            return mid
        }
    }
    return -1
}
