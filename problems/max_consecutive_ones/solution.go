func findMaxConsecutiveOnes(nums []int) int {
  max := 0
    con := 0
    for _, num := range nums {
        if num == 1 {
            con++
        } else {
            if con > max {
                max = con
            }
            con = 0
        }
    }
    if con > max {
        max = con
    }
    return max}
    