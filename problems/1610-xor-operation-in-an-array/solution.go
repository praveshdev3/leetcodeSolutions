func xorOperation(n int, start int) int {
    xorOfElems := 0
    for i:=0; i<n; i++{
        num := start + 2 * i
        xorOfElems = xorOfElems ^ num
    }
    return xorOfElems
}
