func reformatDate(date string) string {
    dates := strings.Split(date," ")

    dat := getDate(dates[0])
    if len(getDate(dates[0])) == 1{
        dat = fmt.Sprintf("0%s",getDate(dates[0]))
    }
    return fmt.Sprintf("%s-%s-%s",dates[2],getMonth(dates[1]),dat)
}

func getMonth(month string) string{
    switch month{
        case "Jan":
            return "01"
        case "Feb":
            return "02"
        case "Mar":
            return "03"
        case "Apr":
            return "04"
        case "May":
            return "05"
        case "Jun":
            return "06"
        case "Jul":
            return "07"
        case "Aug":
            return "08"
        case "Sep":
            return "09"
        case "Oct":
            return "10"
        case "Nov":
            return "11"
        case "Dec":
            return "12"
        default:
            return ""
    }
}

func getDate(date string) string {
    d := []byte(date)
    d = d[:len(d)-2]
    return string(d)
}
