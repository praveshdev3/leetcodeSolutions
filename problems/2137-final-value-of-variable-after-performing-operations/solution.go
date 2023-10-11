func finalValueAfterOperations(operations []string) int {
    x := 0
    for _,operation :=range operations{
        if operation == "X++" || operation == "++X"{
            x = x+1
        } else {
            x = x-1
        }
    }
    return x
}
