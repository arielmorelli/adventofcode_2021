package giant_squid

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "regexp"
)

func convertLineToNumbers(raw_str, separator string) []int {
    var final []int
    trimmed := strings.TrimSpace(raw_str)
    for _, value := range strings.Split(string(trimmed), separator) {
        converted, _ := strconv.Atoi(value)
        final = append(final, converted)
    }
    return final
}

func convertBingoBoard(l0, l1, l2, l3, l4 string) *BingoBoard {
    var convertedLines [][]int = make([][]int, 5)
    spaces := regexp.MustCompile(`\s+`)

    convertedLines[0] = convertLineToNumbers(spaces.ReplaceAllString(l0, " "), " ")
    convertedLines[1] = convertLineToNumbers(spaces.ReplaceAllString(l1, " "), " ")
    convertedLines[2] = convertLineToNumbers(spaces.ReplaceAllString(l2, " "), " ")
    convertedLines[3] = convertLineToNumbers(spaces.ReplaceAllString(l3, " "), " ")
    convertedLines[4] = convertLineToNumbers(spaces.ReplaceAllString(l4, " "), " ")
    return NewBingoBoard(convertedLines)
}

func GiantSquid(input []string) int {
    var bingoNumbers []int = convertLineToNumbers(input[0], ",")
    var bingoBoards []*BingoBoard
    var board *BingoBoard
    var finished bool

    for i := 1; i < len(input) ; i++ {
        if input[i] == "" { continue }
        bingoBoards = append(bingoBoards, convertBingoBoard(input[i],
                                            input[i+1],
                                            input[i+2],
                                            input[i+3],
                                            input[i+4],
                                          ))
        i += 4
    }

    for _, number := range bingoNumbers {
        for _, board = range bingoBoards {
            finished = (*board).Bingo(number)
            if finished {
                return (*board).GetAllUnmarked()*number
            }
        }
    }
    return 0
}

func GiantSquidLastToWin(input []string) int {
    var bingoNumbers []int = convertLineToNumbers(input[0], ",")
    var bingoBoards, unfinishedBoards []*BingoBoard
    var board *BingoBoard
    var finished bool

    for i := 1; i < len(input) ; i++ {
        if input[i] == "" { continue }
        bingoBoards = append(bingoBoards, convertBingoBoard(input[i],
                                            input[i+1],
                                            input[i+2],
                                            input[i+3],
                                            input[i+4],
                                          ))
        i += 4
    }

    unfinishedBoards = bingoBoards
    for _, number := range bingoNumbers {
        bingoBoards = unfinishedBoards
        unfinishedBoards = []*BingoBoard{}
        for _, board = range bingoBoards {
            finished = (*board).Bingo(number)
            if finished {
                if len(bingoBoards) == 1 {
                    return (*board).GetAllUnmarked()*number
                }
            } else {
                unfinishedBoards = append(unfinishedBoards, board)
            }
        }
    }
    return 0
}

func Run() {
    data, err := ioutil.ReadFile("inputs/giant_squid")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }

    lines := strings.Split(string(data), "\n")

    fmt.Printf("4 - Part 1: %d\n", GiantSquid(lines))
    fmt.Printf("4 - Part 2: %d\n", GiantSquidLastToWin(lines))
}
