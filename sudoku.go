package main

import (
    "fmt"
    "sync"
)

type SudokuPuzzle [][]int

var sudoku SudokuPuzzle
var wg sync.WaitGroup
var valid []int

func main() {
    sudoku = SudokuPuzzle{
        []int{6,2,4,5,3,9,1,8,7},
        []int{5,1,9,7,2,8,6,3,4},
        []int{8,3,7,6,1,4,2,9,5},
        []int{1,4,3,8,6,5,7,2,9},
        []int{9,5,8,2,4,7,3,6,1},
        []int{7,6,2,3,9,1,4,5,8},
        []int{3,7,1,9,5,6,8,4,2},
        []int{4,9,6,1,8,2,5,7,3},
        []int{2,8,5,4,7,3,9,1,6},
    }

    valid = make([]int, 0)
    wg.Add(27)

    //cols
    for i := 1; i <= 9; i++ {
        go validate(1, i, 9, i)
    }

    //rows
    for i := 1; i <= 9; i++ {
        go validate(i, 1, i, 9)
    }

    //3x3 subgrids
    for i := 0; i < 3; i++ {
        go validate((3*i)+1, 1, 3+(i*3), 3)
        go validate((3*i)+1, 4, 3+(i*3), 6)
        go validate((3*i)+1, 7, 3+(i*3), 9)
    }

    wg.Wait()

    tmp := 0
    for _,i := range valid {
        tmp += i
    }
    if tmp == 27 {
        fmt.Println("Sudoku solution is valid")
    } else {
        fmt.Println("Sudoku solution is not valid")
    }
}

func validate(x int, y int, w int, h int) {
    defer wg.Done()

    one_to_nine := make([]int, 9)
    for i := x-1; i < w; i++ {
        for j := y-1; j < h; j++ {
            v := sudoku[i][j] - 1
            if one_to_nine[v] != 1 {
                one_to_nine[v] = 1
            }
        }
    }

    tmp := 0
    for _, i := range one_to_nine {
        tmp += i;
    }

    if tmp == 9 {
        valid = append(valid, 1)
    }
}
