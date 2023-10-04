func distMoney(money int, children int) int {
    money = money - children
    if money < 0 {
        return -1
    }
    eightChildren :=  money / 7
    restMoney := money % 7
    if eightChildren == children && restMoney == 0 {
        return eightChildren
    } else if restMoney == 3 && eightChildren == children - 1{
        return eightChildren - 1
    } else if eightChildren > children - 1{
        return children - 1
    }
    return eightChildren
}