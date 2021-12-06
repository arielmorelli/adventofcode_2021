package giant_squid

import "testing"

func TestCreateBingoBoard(t *testing.T) {
    var input [][]int
    var newBoard *BingoBoard
     
    input = [][]int{
        {63,  5, 10, 69, 57},
        {64, 88, 27, 40, 76},
        {59, 20, 58, 90,  6},
        {74, 32, 72, 16, 26},
        {50, 17,  7, 93, 94},
    }
    newBoard = NewBingoBoard(input)
    for i := 0; i < len(input); i++ {
        for j := 0; j < len(input[i]); j++ {
            if (*newBoard).Lines[i][j].Value != input[i][j] {
                t.Errorf("%v is different from %v", input, (*newBoard).Lines)
            }
            if (*newBoard).Lines[i][j].Marked {
                t.Errorf("%v inniting with marked value", input)
            }
        }
    }
    for i := 0; i < len((*newBoard).remainingPerCols); i++ {
        if (*newBoard).remainingPerCols[i] != 5 {
            t.Error("Init remaining cols wrong")
        }
        if (*newBoard).remainingPerLines[i] != 5 {
            t.Error("Init remaining cols wrong")
        }
    }
}

func TestBingo(t *testing.T) {
    var bingo bool
    var newBoard *BingoBoard = NewBingoBoard([][]int{
        {01, 02, 03, 04, 05},
        {06, 07,  8,  9, 10},
        {11, 12, 13, 14, 15},
        {16, 17, 18, 19, 20},
        {21, 22, 23, 24, 25},
    })

    (*newBoard).Bingo(8)
    if ! (*newBoard).Lines[1][2].Marked {
        t.Errorf("Mark is not working for %v", 8)
    }
    if (*newBoard).remainingPerLines[1] == 0 {
        t.Errorf("Remaining lines is not working for %v", 8)
    }

    // test for col
    bingo = (*newBoard).Bingo(4)
    if bingo {
        t.Errorf("Wrong bingo time for %v", 4)
    }
    bingo = (*newBoard).Bingo(9)
    if bingo {
        t.Errorf("Wrong bingo time for %v", 9)
    }
    bingo = (*newBoard).Bingo(14)
    if bingo {
        t.Errorf("Wrong bingo time for %v", 14)
    }
    bingo = (*newBoard).Bingo(19)
    if bingo {
        t.Errorf("Wrong bingo time for %v", 19)
    }
    bingo = (*newBoard).Bingo(24)
    if (*newBoard).remainingPerCols[3] != 0 {
        t.Errorf("Remaining col is not working for col %v", 3)
    }
    if ! bingo {
        t.Errorf("Wrong bingo time for %v should be true", 24)
    }

    // test for line
    bingo = (*newBoard).Bingo(11)
    if bingo {
        t.Errorf("Wrong bingo time for %v", 11)
    }
    bingo = (*newBoard).Bingo(12)
    if bingo {
        t.Errorf("Wrong bingo time for %v", 12)
    }
    bingo = (*newBoard).Bingo(13)
    if bingo {
        t.Errorf("Wrong bingo time for %v", 13)
    }
    bingo = (*newBoard).Bingo(15)
    if (*newBoard).remainingPerLines[2] != 0 {
        t.Errorf("Remaining lines is not working for line %v", 2)
    }
    if ! bingo {
        t.Errorf("Wrong bingo time for %v should be true", 24)
    }
}

func TestMultiplyMarkedVsNonMarked(t *testing.T) {
    var result, expected int

    var newBoard *BingoBoard = NewBingoBoard([][]int{
        {01, 02, 03, 04, 05},
        {06, 07,  8,  9, 10},
        {11, 12, 13, 14, 15},
        {16, 17, 18, 19, 20},
        {21, 22, 23, 24, 25},
    })
    _ = (*newBoard).Bingo(1)
    _ = (*newBoard).Bingo(6)
    _ = (*newBoard).Bingo(11)
    _ = (*newBoard).Bingo(16)
    _ = (*newBoard).Bingo(21)
    _ = (*newBoard).Bingo(3)
    _ = (*newBoard).Bingo(8)
    _ = (*newBoard).Bingo(13)
    _ = (*newBoard).Bingo(18)
    _ = (*newBoard).Bingo(23)
    result = (*newBoard).GetAllUnmarked()
    expected = (2+7+12+17+22+4+9+14+19+24+5+10+15+20+25) 
    if result != expected {
        t.Errorf("Multiply is not working expected %v got %v", expected, result)
    }
}
