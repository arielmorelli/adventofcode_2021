package giant_squid

type bingoEntry struct {
    Value int
    Marked bool
}

type BingoBoard struct {
    Lines [][]bingoEntry
    remainingPerCols, remainingPerLines []int
}

func NewBingoBoard(entries [][]int) *BingoBoard {
    var newBingoBoard *BingoBoard = new(BingoBoard)

    (*newBingoBoard).Lines = make([][]bingoEntry, 5)
    for i := 0; i < len(entries); i++ {
        (*newBingoBoard).Lines[i] = make([]bingoEntry, 5)
        for j := 0; j < len(entries); j++ {
            (*newBingoBoard).Lines[i][j].Value = entries[i][j]
            (*newBingoBoard).Lines[i][j].Marked = false
        }
    }
    (*newBingoBoard).remainingPerCols = []int{5, 5, 5, 5, 5}
    (*newBingoBoard).remainingPerLines = []int{5, 5, 5, 5, 5}

    return newBingoBoard
}

func (b BingoBoard) Bingo(value int) bool {
    for i := 0; i < len(b.Lines); i++ {
        for j := 0; j < len(b.Lines[i]); j++ {
            if b.Lines[i][j].Value == value {
                b.Lines[i][j].Marked = true
                b.remainingPerLines[i]--
                b.remainingPerCols[j]--

                if b.remainingPerLines[i] == 0 ||  b.remainingPerCols[j] == 0 {
                    return true
                }
                return false
            }
        }
    }
    return false
}

func (b BingoBoard) GetAllUnmarked() int {
    var unmarked int
    for i := 0; i < len(b.Lines); i++ {
        for j := 0; j < len(b.Lines[i]); j++ {
            if ! b.Lines[i][j].Marked {
                unmarked += b.Lines[i][j].Value
            }        
        }
    }
    return unmarked
}
