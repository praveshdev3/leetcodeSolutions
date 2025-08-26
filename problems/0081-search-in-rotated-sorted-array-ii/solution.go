func search(nums []int, target int) bool {
    low := 0
    high := len(nums)-1
    for low <= high{
        mid := (low+high)/2
        if nums[mid] == target{
            return true
        }
        if nums[low] == nums[mid]{
            low++
            continue
        }
        if nums[low] <= nums[mid]{
            if target < nums[mid] && target >= nums[low]{
                high = mid - 1
            }else{
                low = mid+1
            }
        }else{
            if target > nums[mid] && target <= nums[high]{
                low = low + 1
            }else{
                high = high-1
            }
        }
    }
    return false
}
