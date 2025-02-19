type item struct{
    n int
    bytes []byte
}

func decodeString(s string) string {
    num := 0
    st := []item{{1,[]byte{}}}

    for i := range s {
        switch {
            case s[i] == '0':
                num *= 10
            case s[i] > '0' && s[i] <= '9':
                num = num * 10 + int(s[i] - '0')
            case s[i] == '[':
                st = append(st, item{num, []byte{}})
                num = 0
            case s[i] == ']':
                tmp := st[len(st)-1]
                st = st[:len(st)-1]
                for j:=0; j<tmp.n; j++{
                    st[len(st)-1].bytes = append(st[len(st)-1].bytes, tmp.bytes...)
                }
            default:
                st[len(st)-1].bytes = append(st[len(st)-1].bytes, s[i])
        }
    }

    return string(st[0].bytes)
}
