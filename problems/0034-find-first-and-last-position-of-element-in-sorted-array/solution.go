func searchRange(nums []int, target int) []int {
   return []int{firstIndex(nums,target),lastIndex(nums,target)}
}

func firstIndex(nums []int, target int) int{
    lo := 0
    hi := len(nums)-1
    pos := -1
    for lo <= hi{
        mid := (lo+hi)/2
        if nums[mid] >= target{
            if nums[mid] == target{
                pos = mid
            }
            hi = mid-1
        }else{
            lo = mid+1
        }
    }
    return pos
}

func lastIndex(nums []int, target int) int{
    lo := 0
    hi := len(nums)-1
    pos := -1
    for lo <= hi{
        mid := (lo+hi)/2
        if nums[mid] > target{
            hi = mid-1
        }else{
            if nums[mid] == target{
                pos = mid
            }
            lo = mid+1
        }
    }
    return pos
}
