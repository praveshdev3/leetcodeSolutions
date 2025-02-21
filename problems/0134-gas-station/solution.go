func canCompleteCircuit(gas []int, cost []int) int {
    globalFuelLeft, fuelLeft, start := 0,0,0
    for index := range gas{
        globalFuelLeft += gas[index] - cost[index]
        fuelLeft += gas[index] - cost[index]
        if fuelLeft < 0{
            fuelLeft = 0
            start = index + 1
        } 
    }

    if globalFuelLeft < 0{
        return -1
    }

    return start
}
