func isValid(s string) bool {
    if len(s)%2 != 0{
        return false
    }

    stack := make([]rune,0)

    for _,c := range s{
        if c == '{' || c == '(' || c == '['{
            stack = append(stack,c)
        } else if len(stack) > 0{
            switch c{
                case ')':
                    if stack[len(stack)-1] != '('{
                        return false
                    } else{
                        stack = stack[:len(stack)-1]
                    }
                case '}':
                    if stack[len(stack)-1] != '{'{
                        return false
                    } else{
                        stack = stack[:len(stack)-1]
                    }
                case ']':
                    if stack[len(stack)-1] != '['{
                        return false
                    } else{
                        stack = stack[:len(stack)-1]
                    }
            }
        } else {
            return false
        }
    }
    if len(stack) == 0{
        return true
    }
    return false
}
