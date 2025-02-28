func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)

    for i := 0; i < numCourses; i++{
        graph[i] = make([]int, 0)
    }

    for _,prerequisite := range prerequisites{
        graph[prerequisite[0]] = append(graph[prerequisite[0]],prerequisite[1])
    }

    visited := make(map[int]int, numCourses)
    for i := 0; i < numCourses; i++{
        if hasCycle(graph, visited, i){
            return false
        }
    }

    return true
}

func hasCycle(graph [][]int, visited map[int]int, num int) bool {
    n,ok := visited[num];
    if ok{
        if n == 1{
            return true
        } else if n == -1{
            return false
        }
    }
    visited[num] = 1
    for _,neighbour := range graph[num]{
        if hasCycle(graph, visited, neighbour){
            return true
        }
    }

    visited[num] = -1
    return false
}
