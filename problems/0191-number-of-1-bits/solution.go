func hammingWeight(n int) int {
    numOfOne := 0
    for i:=0; i<32; i++{
        if n & 1 == 1{
           numOfOne++ 
        }
        n >>= 1
    }
    return numOfOne
}
