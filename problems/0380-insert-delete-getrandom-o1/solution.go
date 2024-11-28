type RandomizedSet struct {
    m map[int]int
    l []int
}


func Constructor() RandomizedSet {
    m := make(map[int]int)
    l := make([]int, 0)
    return RandomizedSet{m,l}
}


func (this *RandomizedSet) Insert(val int) bool {
    if _,ok := this.m[val]; ok {
        return false
    }
    this.m[val] = len(this.l)
    this.l = append(this.l, val)
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _,ok := this.m[val]; !ok {
        return false
    }
    n := len(this.l)
    i := this.m[val]
    this.l[i], this.l[n-1] = this.l[n-1], this.l[i]
    this.m[this.l[i]]=i
    this.l = this.l[:n-1]
    delete(this.m,val)
    return true
}


func (this *RandomizedSet) GetRandom() int {
    i := rand.Intn(len(this.l))
    return this.l[i]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
