func reverse(str string) (result string) { 
    for _, v := range str { 
        result = string(v) + result 
    } 
    return
} 
func stringToBinary(x int) string {
    if x==0 {
        return "0"
    }
    
    res := ""
    for x > 0 {
        if x%2==0 {
            res += "0"
        } else {
            res += "1"
        }
        x/=2
    }
    ans := reverse(res)
    return ans
}
func convertDateToBinary(date string) string {
    var yr string = date[0:4];
    var mn string = date[5:7];
    var dy string = date[8:10];
    
    year,_ := (strconv.Atoi(yr))
    month,_ := (strconv.Atoi(mn))
    day,_ := (strconv.Atoi(dy))
    yearConverted := stringToBinary(year)
    monthConverted := stringToBinary(month)
    dayConverted:= stringToBinary(day)
    return yearConverted+"-"+monthConverted+"-"+dayConverted;
    
}
