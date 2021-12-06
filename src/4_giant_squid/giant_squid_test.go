package giant_squid

import "testing"


func TestConvertBingoNumber(t *testing.T) {
    var input string
    var converted, expected []int

    input = "1"
    expected = []int{1}
    converted = convertLineToNumbers(input, ",")
    for i, value := range expected {
        if converted[i] != value {
            t.Errorf("Expected %v got %v", expected, converted)
        }
    }

    input = "1,2,5,6"
    expected = []int{1, 2, 5, 6}
    converted = convertLineToNumbers(input, ",")
    for i, value := range expected {
        if converted[i] != value {
            t.Errorf("Expected %v got %v", expected, converted)
        }
    }

    input = "1 2 5 6"
    expected = []int{1, 2, 5, 6}
    converted = convertLineToNumbers(input, " ")
    for i, value := range expected {
        if converted[i] != value {
            t.Errorf("Expected %v got %v", expected, converted)
        }
    }

    input = " 1 2 5 6 "
    expected = []int{1, 2, 5, 6}
    converted = convertLineToNumbers(input, " ")
    for i, value := range expected {
        if converted[i] != value {
            t.Errorf("Expected %v got %v", expected, converted)
        }
    }
}

func TestConvertBingoBoard(t *testing.T) {
    input0 := "63  5 10 69 57"
    input1 := "64 88 27 40 76"
    input2 := "59 20 58 90  6"
    input3 := "74 32 72 16 26"
    input4 := "50 17  7 93 94"
    expected := [][]int{
        {63,  5, 10, 69, 57},
        {64, 88, 27, 40, 76},
        {59, 20, 58, 90,  6},
        {74, 32, 72, 16, 26},
        {50, 17,  7, 93, 94},
    }
    newBoard := convertBingoBoard(input0, input1, input2, input3, input4)
    for i := 0; i < len(expected); i++ {
        for j := 0; j < len(expected[i]); j++ {
            if (*newBoard).Lines[i][j].Value != expected[i][j] {
                t.Errorf("%v is different from %v", expected, (*newBoard).Lines)
            }
            if (*newBoard).Lines[i][j].Marked {
                t.Errorf("%v inniting with marked value", expected)
            }
        }
    }
}

func TestGiantSquid(t *testing.T) {
    var input []string
    var expected, result int

    input = []string{
        "1,2,3,4,5",
        "",
        " 1  2  3  4  5",
        " 6  7  8  9 10",
        "11 12 13 14 15",
        "16 17 18 19 20",
        "21 22 23 24 25",
    }
    expected = (6+7+8+9+10+11+12+13+14+15+16+17+18+19+20+21+22+23+24+25)*5
    result = GiantSquid(input)
    if expected != result {
        t.Errorf("Expected %v got %v", expected, result)
    }

    input = []string{
        "1,2,3,4,5",
        "",
        " 1  2  3  4  5",
        " 6  7  8  9 10",
        "11 12 13 14 15",
        "16 17 18 19 20",
        "21 22 23 24 25",
        "",
        " 6  7  8  9 10",
        "11 12 13 14 15",
        "16 17 18 19 20",
        "21 22 23 24 25",
        "26 27 28 29 30",
        "",
    }
    expected = (6+7+8+9+10+11+12+13+14+15+16+17+18+19+20+21+22+23+24+25)*5
    result = GiantSquid(input)
    if expected != result {
        t.Errorf("Expected %v got %v", expected, result)
    }

    input = []string{
        "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
        "",
        "22 13 17 11  0",
        " 8  2 23  4 24",
        "21  9 14 16  7",
        " 6 10  3 18  5",
        " 1 12 20 15 19",
        "",
        " 3 15  0  2 22",
        " 9 18 13 17  5",
        "19  8  7 25 23",
        "20 11 10 24  4",
        "14 21 16 12  6",
        "",
        "14 21 17 24  4",
        "10 16 15  9 19",
        "18  8 23 26 20",
        "22 11 13  6  5",
        " 2  0 12  3  7",
    }
    expected = 4512
    result = GiantSquid(input)
    if expected != result {
        t.Errorf("Expected %v got %v", expected, result)
    }
}

func TestGiantSquidLastToWin(t *testing.T) {
    var input []string
    var expected, result int

    input = []string{
        "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
        "",
        "22 13 17 11  0",
        " 8  2 23  4 24",
        "21  9 14 16  7",
        " 6 10  3 18  5",
        " 1 12 20 15 19",
        "",
        " 3 15  0  2 22",
        " 9 18 13 17  5",
        "19  8  7 25 23",
        "20 11 10 24  4",
        "14 21 16 12  6",
        "",
        "14 21 17 24  4",
        "10 16 15  9 19",
        "18  8 23 26 20",
        "22 11 13  6  5",
        " 2  0 12  3  7",
    }
    expected = 1924
    result = GiantSquidLastToWin(input)
    if expected != result {
        t.Errorf("Expected %v got %v", expected, result)
    }
}

