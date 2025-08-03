func setZeroes(matrix [][]int)  {
    firstRowZero, firstColumnZero := false,false

     for i:=0; i<len(matrix); i++{
        if matrix[i][0] == 0{
            firstColumnZero = true
            break
        }
    }

    for i:=0; i<len(matrix[0]); i++{
        if matrix[0][i] == 0{
            firstRowZero = true
            break
        }
    }

    for i:=0; i<len(matrix); i++{
        for j:=0; j<len(matrix[i]); j++{
            if matrix[i][j] == 0{
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }

    for i:=1; i<len(matrix); i++{
        if matrix[i][0] == 0{
            for j:=1; j<len(matrix[i]); j++{
                matrix[i][j] = 0
            }
        }
    }

    for i:=1; i<len(matrix[0]); i++{
        if matrix[0][i] == 0{
            for j:=1; j<len(matrix); j++{
                matrix[j][i] = 0
            }
        }
    }

    if firstRowZero{
        for j:=0; j<len(matrix[0]); j++{
            matrix[0][j] = 0
        }
    }

    if firstColumnZero{
        for j:=0; j<len(matrix); j++{
            matrix[j][0] = 0
        }
    }
}
