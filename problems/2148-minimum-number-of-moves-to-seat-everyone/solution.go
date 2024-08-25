func minMovesToSeat(seats []int, students []int) int {
    result := 0
    sort.Ints(seats)
    sort.Ints(students)
    for i:=0; i<len(students); i++ {
        result += abs(students[i] - seats[i])
    }
    return result
}

func abs(a int) int {
    if a<0 {
        a=-a
    }
    return a
}
