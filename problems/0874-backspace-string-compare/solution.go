func backspaceCompare(s string, t string) bool {
    srr,k := processString(s)
    trr,p := processString(t)

    if k != p{
        return false
    }


    for i := 0; i<k;i++{
        if srr[i] != trr[i]{
            return false
        }
    }
    return true
}

func processString(s string) ([]rune,int) {
    arr := make([]rune,len(s))
    k := 0
    for _, c := range s{
        if c != '#'{
            arr[k] = c
            k++
        } else if k > 0 {
            k--
        }
    }
    for i:=k;i<len(s);i++{
        arr[i]=0
    }
    return arr,k
}
