func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    noOfEmployees := 0
    for _ , h := range hours {
        if h >= target {
            noOfEmployees++
        }
    }
    return noOfEmployees
}
