func canConstruct(ransomNote string, magazine string) bool {
    arr := make([]int,26)
    for _,r := range ransomNote{
        i := r-'a'
        arr[i] = arr[i] + 1
    }
    for _,r := range magazine{
        i := r-'a'
        arr[i] = arr[i] - 1
    }
    for _,i := range arr{
        if i > 0{
            return false
        }
    }
    return true
}
