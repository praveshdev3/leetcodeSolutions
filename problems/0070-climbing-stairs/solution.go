func climbStairs(n int) int {
    res := 0

    secondNext,next := 0,0
    for i:=1; i<=n; i++{
        if i==1{
            res=1
        } else if i==2{
            res=2
        }else{
            res=next+secondNext
        }
        next = secondNext
        secondNext = res
    }
    return res
}
