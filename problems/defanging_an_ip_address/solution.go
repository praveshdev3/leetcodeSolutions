import "bytes"

func defangIPaddr(address string) string {
    var strBuffer bytes.Buffer
    for _,c := range address{
        if c == '.'{
            strBuffer.WriteString("[.]")
        } else {
            strBuffer.WriteRune(c)
        }
    }
    return strBuffer.String()
}