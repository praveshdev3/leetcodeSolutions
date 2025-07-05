func isPalindrome(x int) bool {
    reversedNum := 0
    n := x

    if n < 0{
        return false
    }

    for x != 0{
        num := x%10
        x = x/10
        reversedNum = reversedNum*10 + num
    }

    if reversedNum != n{
        return false
    }

    return true
}

