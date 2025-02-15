func search(nums []int, target int) int {
    low,high := 0,len(nums)-1
    for low <= high{
        mid := (low+high)/2

        if nums[mid] == target{
            return mid
        }

        if nums[low] <= nums[mid]{
            if target < nums[mid] && target >= nums[low] {
                high = mid - 1
            } else{
                low = mid + 1
            }
        } else {
            if target > nums[mid] && target <= nums[high] {
                low = mid + 1
            } else{
                high = mid - 1
            }
        }
    }
    return -1
}

    // for index,num := range nums {
    //     if target == num{
    //         return index
    //     }
    // }
    // return -1
