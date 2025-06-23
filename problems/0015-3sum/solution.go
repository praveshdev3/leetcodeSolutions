func threeSum(nums []int) [][]int {
    var result [][]int
    sort.Ints(nums)

    for i:=0; i<len(nums)-2; i++{
        if i!=0 && nums[i]==nums[i-1]{
            continue
        }

        target := 0-nums[i]
        j:=i+1
        k:=len(nums)-1

        for j<k{
            if nums[j] + nums[k] == target{
                result = append(result,[]int{nums[i],nums[j],nums[k]})
                j++
                k--
                for j<k && nums[j] == nums[j-1]{
                    j++
                }
                for j<k && nums[k] == nums[k+1]{
                    k--
                }
            }else if nums[j] + nums[k] > target{
                k--
            }else{
                j++
            }
        }
    }
    return result
}
