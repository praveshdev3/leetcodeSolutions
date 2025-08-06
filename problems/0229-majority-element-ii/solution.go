func majorityElement(nums []int) []int {
    result := make([]int,0)
    
    counter1, counter2, count1, count2 := 0,1,0,0

    for _,num := range nums{
        if num == counter1{
            count1++
        }else if num == counter2{
            count2++
        }else if count1 == 0{
            counter1 = num
            count1 = 1
        }else if count2 == 0{
            counter2 = num
            count2 = 1
        }else{
            count1--
            count2--
        }
    }

    count1,count2 = 0,0
    for _,num := range nums{
        if num == counter1{
            count1++
        }else if num == counter2{
            count2++
        }
    }

    if count1 > len(nums)/3{
        result = append(result,counter1)
    }
    if count2 > len(nums)/3{
        result = append(result,counter2)
    }

    return result
}
