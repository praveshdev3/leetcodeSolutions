func isIsomorphic(s string, t string) bool {

    if len(s) != len(t){
        return false
    }

    map1 := make([]int, 128)
    map2 := make([]int, 128)

    for i:=0; i<=len(s)-1; i++{

        sCh := s[i]
        tCh := t[i]

        if map1[sCh] == 0 && map2[tCh] == 0{
            map1[sCh] = int(tCh)
            map2[tCh] = int(sCh)
        }else if map1[sCh] != int(tCh) || map2[tCh] != int(sCh){
            return false
        }
    } 

    return true
}
