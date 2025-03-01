func maximumSwap(num int) int {
    numStr := strconv.Itoa(num)
    max := -1
    maxIndex := -1
    swapi := -1
    swapj := -1

    for i:=len(numStr)-1;i>=0;i--{
        if int(numStr[i] - '0') > max {
            max = int(numStr[i] - '0')
            maxIndex = i
        }
        if int(numStr[i] - '0') < max{
            swapi = i
            swapj = maxIndex
        }
    }

    if swapi > -1 && swapj > -1{
        s := make([]byte,0)
        for i,_ := range numStr{
            if i == swapj{
                s = append(s,numStr[swapi])
            } else if i == swapi{
                s = append(s,numStr[swapj])
            } else{
                s = append(s,numStr[i])
            }
        }
        numStr = string(s)
    }

    num,_ = strconv.Atoi(numStr)
    return num
}
