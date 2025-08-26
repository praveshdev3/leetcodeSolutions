func search(nums []int, target int) int {
    low := 0
    high := len(nums)-1
    pos := -1
    for low<=high{
        mid := (low+high)/2
        if nums[mid] == target{
            pos = mid
            break
        }else if nums[low]<=nums[mid]{
            if target < nums[mid] && target >= nums[low]{
                high = mid - 1
            }else{
                low = mid + 1
            }
        }else{
            if target > nums[mid] && target <= nums[high]{
                low = mid + 1
            }else{
                high = mid - 1
            }
        }
    }
    return pos
}
