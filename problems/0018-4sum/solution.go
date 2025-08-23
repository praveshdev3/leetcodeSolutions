func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    return kSum(nums, target, 0, 4)
}

func kSum(nums []int, target, start, k int) [][]int {
    res := [][]int{}

    if start == len(nums){
        return res
    }

    averageVal := target/k

    if nums[start] > averageVal || averageVal > nums[len(nums)-1]{
        return res
    }

    if k==2{
        return twoSum(nums, target, start)
    }

    for i:=start; i<len(nums);i++{
        if i==start || nums[i-1]!=nums[i]{
            for _,subset := range kSum(nums, target-nums[i], i+1, k-1){
                res = append(res, append([]int{nums[i]},subset...))
            }
        }
    }

    return res
}

func twoSum(nums []int, target, start int) [][]int{
    res := [][]int{}
    l := start
    h := len(nums)-1
    for l < h{
        sum := nums[l] + nums[h]
        if sum < target || (l > start && nums[l] == nums[l-1]){
            l++
        }else if sum > target || (h < len(nums)-1 && nums[h] == nums[h+1]){
            h--
        }else{
            res = append(res, []int{nums[l],nums[h]})
            l++
            h--
        } 
    }
    return res
}
