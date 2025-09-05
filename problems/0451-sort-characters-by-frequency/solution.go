func frequencySort(s string) string {

    arr := make([]int,255)

    for i:=0; i<len(s); i++{
        arr[s[i]]++
    }

    b := []byte(s)
    sort.Slice(b, func(i , j int) bool {
        if arr[b[i]] == arr[b[j]]{
            return b[i] > b[j]
        }
        return arr[b[i]] > arr[b[j]]
    })

    return string(b)
}
