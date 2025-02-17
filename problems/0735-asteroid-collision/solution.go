func asteroidCollision(asteroids []int) []int {
    result := make([]int,0)

    for i:=0; i<len(asteroids); i++{
        asteroid := asteroids[i]

        if len(result) == 0 || asteroid > 0 || result[len(result)-1] < 0{
            result = append(result,asteroid)
        } else{
            top := result[len(result)-1]

            if -asteroid == top{
                result = result[:len(result)-1]
            } else if -asteroid > top{
                result = result[:len(result)-1]
                i--
            }
        }
    }

    return result
}
