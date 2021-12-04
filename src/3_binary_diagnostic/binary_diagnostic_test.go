package binary_diagnostic

import "testing"
import "fmt"

func RunTestCaseToDec(t *testing.T, entries []int, expected int) {
    increases := ToDec(entries)
    if increases != expected {
        t.Errorf("Input: '%v'. Wanted %d got %d", entries, expected, increases)
    }
}

func TestToDec(t *testing.T) {
     var input []int = make([]int, 0)
     RunTestCaseToDec(t, input, 0)

     input = []int{0}
     RunTestCaseToDec(t, input, 0)

     input = []int{1}
     RunTestCaseToDec(t, input, 1)

     input = []int{1, 0}
     RunTestCaseToDec(t, input, 2)

     input = []int{0, 1}
     RunTestCaseToDec(t, input, 1)

     input = []int{1, 1, 1}
     RunTestCaseToDec(t, input, 7)

     input = []int{1, 0, 0}
     RunTestCaseToDec(t, input, 4)

     input = []int{1, 0, 1}
     RunTestCaseToDec(t, input, 5)

     input = []int{0, 0, 0}
     RunTestCaseToDec(t, input, 0)
}

func RunTestCaseGetOpposite(t *testing.T, entries []int, expected []int) {
    result := GetOpposite(entries)
    if len(expected) != len(result) {
        t.Errorf("'%v' different len of %v", result, expected)
        return
    }
    for i := 0; i < len(result); i++ {
        if expected[i] != result[i] {
            t.Errorf("Input: '%v'. Wanted %v got %v", entries, expected, result)
            return
        }
    }
}

func TestGetOpposite(t *testing.T) {
     var input, expected []int

     input = []int{0}
     expected = []int{1}
     RunTestCaseGetOpposite(t, input, expected)

     input = []int{1}
     expected = []int{0}
     RunTestCaseGetOpposite(t, input, expected)

     input = []int{1, 0}
     expected = []int{0, 1}
     RunTestCaseGetOpposite(t, input, expected)

     input = []int{0, 1}
     expected = []int{1, 0}
     RunTestCaseGetOpposite(t, input, expected)

     input = []int{1, 1, 1}
     expected = []int{0, 0, 0}
     RunTestCaseGetOpposite(t, input, expected)

     input = []int{1, 0, 0}
     expected = []int{0, 1, 1}
     RunTestCaseGetOpposite(t, input, expected)

     input = []int{1, 0, 1}
     expected = []int{0, 1, 0}
     RunTestCaseGetOpposite(t, input, expected)

     input = []int{0, 0, 0}
     expected = []int{1, 1, 1}
     RunTestCaseGetOpposite(t, input, expected)
}

func RunTestCase(t *testing.T, entries []string, expected []int) {
    result := BinaryDiagnostic(entries)
    if len(expected) != len(result) {
        t.Errorf("'%v' different len of %v", result, expected)
        return
    }
    for i := 0; i < len(result); i++ {
        if expected[i] != result[i] {
            t.Errorf("Input: '%v'. Wanted %v got %v", entries, expected, result)
            return
        }
    }
}

func TestBinaryDiagnostic(t *testing.T){
    var input []string
    var expected []int
    RunTestCase(t, input, expected)

    input = []string{}
    expected = []int{}
    RunTestCase(t, input, expected)

    input = []string{
        "1",
    }
    expected = []int{1}
    RunTestCase(t, input, expected)

    input = []string{
        "10",
    }
    expected = []int{1, 0}
    RunTestCase(t, input, expected)

    input = []string{
        "101",
    }
    expected = []int{1, 0, 1}
    RunTestCase(t, input, expected)

    input = []string{
        "101",
        "101",
    }
    expected = []int{1, 0, 1}
    RunTestCase(t, input, expected)

    input = []string{
        "000",
        "111",
    }
    expected = []int{1, 1, 1}
    RunTestCase(t, input, expected)

    input = []string{
        "000",
        "111",
        "000",
    }
    expected = []int{0, 0, 0}
    RunTestCase(t, input, expected)

    input = []string{
        "000",
        "000",
        "000",
    }
    expected = []int{0, 0, 0}
    RunTestCase(t, input, expected)

    input = []string{
        "111",
        "111",
        "111",
    }
    expected = []int{1, 1, 1}
    RunTestCase(t, input, expected)

    input = []string{
        "00100",
        "11110",
        "10110",
        "10111",
        "10101",
        "01111",
        "00111",
        "11100",
        "10000",
        "11001",
        "00010",
        "01010",
    }
    expected = []int{1, 0, 1, 1, 0}
    RunTestCase(t, input, expected)
}

func RunTestCaseForRatings(t *testing.T, entries []string, expected1, expected2 []int) {
    result1, result2 := BinaryDiagnosticForRatings(entries)
    if len(expected1) != len(result1) {
        t.Errorf("'%v' different len of %v", result1, expected1)
        return
    }
    if len(expected2) != len(result2) {
        t.Errorf("'%v' different len of %v", result2, expected2)
        return
    }
}

func TestBinaryDiagnosticForRatings(t *testing.T){
    var input []string
    var expected1, expected2 []int
    RunTestCaseForRatings(t, input, expected1, expected2)

    input = []string{
        "101",
        "000",
    }
    expected1 = []int{1, 0, 1}
    expected2 = []int{0, 0, 0}
    RunTestCaseForRatings(t, input, expected1, expected2)

    input = []string{
        "101",
        "011",
        "000",
    }
    expected1 = []int{0, 1, 1}
    expected2 = []int{1, 0, 1}
    RunTestCaseForRatings(t, input, expected1, expected2)

    input = []string{
        "00100",
        "11110",
        "10110",
        "10111",
        "10101",
        "01111",
        "00111",
        "11100",
        "10000",
        "11001",
        "00010",
        "01010",
    }
    expected1 = []int{1, 0, 1, 1, 1}
    expected2 = []int{0, 1, 0, 1, 0}
    RunTestCaseForRatings(t, input, expected1, expected2)

    fmt.Println(ToDec(expected1))
    fmt.Println(ToDec(expected2))
}

func RunTestCaseGet0And1(t *testing.T, entries [][]int, index, expected1, expected2 int) {
    zeros, ones := Get0and1FromListOfList(entries, index)
    if zeros != expected1 || ones != expected2 {
        t.Errorf("'%v' wrong return: zeros: %v ones: %v", entries, zeros, ones)
        return
    }
}

func TestGet0and1FromListOfList(t *testing.T) {
    var input [][]int
    var zeros, ones int

    input = [][]int{
        {1, 1, 1},
        {0, 0, 0},
    }
    zeros = 1
    ones = 1
    RunTestCaseGet0And1(t, input, 1, zeros, ones)

    input = [][]int{
        {1, 1, 1},
        {0, 0, 1},
        {0, 0, 0},
        {0, 0, 0},
    }
    zeros = 2
    ones = 2
    RunTestCaseGet0And1(t, input, 2, zeros, ones)
}

func RunTestCaseFilterUsingIndex(t *testing.T, input [][]int, index, target int, expected [][]int) {
    result := FilterListUsingIndex(input, index, target)
    if len(expected) != len(expected) {
        t.Errorf("Different sizes: expected: %v got %v", len(expected), len(result))
        return
    }
    for i := 0; i < len(result); i++ {
        for j := 0; j < len(result[i]); j++ {
            if expected[i][j] != result[i][j] {
                t.Errorf("'%v' wrong return: expected: %v got %v", input, expected, result)
                return
            }
        }
    }
}

func TestFilterUsingIndex(t *testing.T) {
    var input, expected [][]int
    var index, target int

    target = 1

    input = [][]int{
        {1, 0, 1},
        {0, 1, 1},
    }

    index = 0
    expected = [][]int{
        {1, 0, 1},
    }
    RunTestCaseFilterUsingIndex(t, input, index, target, expected)

    index = 1
    expected = [][]int{
        {0, 1, 1},
    }
    RunTestCaseFilterUsingIndex(t, input, index, target, expected)

    index = 2
    expected = [][]int{
        {1, 0, 1},
        {0, 1, 1},
    }
    RunTestCaseFilterUsingIndex(t, input, index, target, expected)
}
