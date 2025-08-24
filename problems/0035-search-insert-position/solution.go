func searchInsert(nums []int, target int) int {
    pos := -1
    low := 0
    high := len(nums)-1
    for low<=high{
        mid := (low+high)/2
        if nums[mid] == target{
            pos = mid-1
            break
        }else if nums[mid] > target{
            high = mid-1
        }else{
            pos = mid
            low = mid+1
        }
    }
    return pos+1
}
