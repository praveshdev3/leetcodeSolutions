func minimumSum(num int) int {
    var nums []int
    for num > 0{
        n := num%10
        nums = append(nums,n)
        num = num/10
    }
    sort.Ints(nums)
    fmt.Println(nums)
    num1 := nums[0]*10 + nums[2]
    num2 := nums[1]*10 + nums[3]
    return num1 + num2
}
