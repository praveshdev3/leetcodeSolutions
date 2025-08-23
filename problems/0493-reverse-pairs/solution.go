func reversePairs(nums []int) int {
    reversePairs := 0
    reversePairs, nums = countReversePairs(nums, reversePairs)
    return reversePairs
}

func countReversePairs(nums []int, reversePairs int)(int,[]int){
    if len(nums) <= 1{
        return reversePairs, nums
    }

    mid := len(nums)/2

    reversePairs1, result1 := countReversePairs(nums[:mid], reversePairs)
    reversePairs2, result2 := countReversePairs(nums[mid:], reversePairs)
    l,h,count := 0,0,0
    result := make([]int,0)
    for l<len(result1) && h<len(result2){
        if result1[l] > 2 * result2[h]{
            count += len(result1) - l
            h++
        }else{
            l++
        }
    }

    l,h = 0,0
    for l<len(result1) && h<len(result2){
        if result1[l] > result2[h]{
            result = append(result,result2[h])
            h++
        }else{
            result = append(result,result1[l])
            l++
        }
    }

    result = append(result, result1[l:]...)
    result = append(result, result2[h:]...)
    return reversePairs1 + reversePairs2 + count, result
}
