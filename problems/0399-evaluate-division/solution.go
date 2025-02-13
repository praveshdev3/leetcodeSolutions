func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    graph := make(map[string]map[string]float64)
    seen := make(map[string]bool)

    for i:=0; i<len(equations); i++{
        a,b := equations[i][0], equations[i][1]
        if _,ok := graph[a]; !ok{
            graph[a] = make(map[string]float64)
        }
        graph[a][b] = values[i]
        if _,ok := graph[b]; !ok{
            graph[b] = make(map[string]float64)
        }
        graph[b][a] = 1/values[i]
    }

    var divide func(string,string) float64
    divide = func(a,b string) float64{
        if a==b{
            return 1
        }
        seen[a] = true
        for key,value:=range graph[a]{
            if seen[key]{
                continue
            }
            res := divide(key,b)
            if res > 0{
                return res*value
            }
        }
        return -1.0
    }

    var isExists func(string,string) bool
    isExists = func(a,b string) bool{
        if _,ok := graph[a]; !ok {
            return false
        }
        if _,ok := graph[b]; !ok {
            return false
        }
        return true
    }

    ans := make([]float64,0)
    for i:=0; i<len(queries); i++{
        a,b := queries[i][0], queries[i][1]
        if !isExists(a,b){
            ans = append(ans,-1.00)
            continue
        }
        seen = make(map[string]bool)
        ans = append(ans,divide(a,b))
    }

    return ans
}
