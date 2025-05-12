func isPalindrome(x int) bool {
   reversedNum := 0
   num := x

   for x != 0{
    reversedNum = reversedNum*10 + x%10
    x = x/10
   }

   if reversedNum != num || num < 0{
    return false
   }

    return true
}
