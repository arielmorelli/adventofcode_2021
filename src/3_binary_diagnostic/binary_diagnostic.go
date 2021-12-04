package binary_diagnostic

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func ToDec(target_value []int) int {
    var dec, exp int = 0, 1

    for i := len(target_value)-1; i >= 0; i-- {
        dec += target_value[i]*exp
        exp = exp<<1
    }

    return dec
}

func GetOpposite(target_value []int) []int {
    var opposite []int = make([]int, len(target_value))
    for i := 0; i < len(target_value); i++ {
        opposite[i] += (1-target_value[i])
    }

    return opposite
}

func Get0and1FromListOfList(target_value [][]int, index int) (int, int){
    var count0, count1 int
    for i := 0; i < len(target_value); i++ {
        if target_value[i][index] == 0 {
            count0++
        } else {
            count1++
        }
    }

    return count0, count1
}

func FilterListUsingIndex(target_list [][]int, index, target_value int) [][]int {
    var final_list [][]int

    if len(target_list) == 1 {
        return target_list
    }

    for i :=0; i < len(target_list); i++ {
        if target_list[i][index] == target_value {
            final_list = append(final_list, target_list[i])
        }
    }

    return final_list
}

func BinaryDiagnosticForRatings(input []string) ([]int, []int) {
    var converted_list, list_o2, list_co2 [][]int
    var converted, index int

    if len(input) == 0 {
        return []int{}, []int{}
    }

    for i := 0; i < len(input); i++ {
        if input[i] == "" {
            break
        }
        converted_list = append(converted_list, make([]int, len(input[i])))

        for j := 0; j < len(input[i]); j++ {
            converted, _ = strconv.Atoi(string(input[i][j]))
            converted_list[i][j] = converted
        }
    }

    list_o2, list_co2 = converted_list, converted_list
    index = 0
    for len(list_o2) > 1 {
        zeros, ones := Get0and1FromListOfList(list_o2, index)
        if zeros > ones {
            list_o2 = FilterListUsingIndex(list_o2, index, 0)
        } else if zeros < ones {
            list_o2 = FilterListUsingIndex(list_o2, index, 1)
        } else {
            list_o2 = FilterListUsingIndex(list_o2, index, 1)
        }
        index++
    }

    index = 0
    for len(list_co2) > 1 {
        zeros, ones := Get0and1FromListOfList(list_co2, index)
        if zeros > ones {
            list_co2 = FilterListUsingIndex(list_co2, index, 1)
        } else if zeros < ones {
            list_co2 = FilterListUsingIndex(list_co2, index, 0)
        } else {
            list_co2 = FilterListUsingIndex(list_co2, index, 0)
        }
        index++
    }

    return list_o2[0], list_co2[0]
}

func BinaryDiagnostic(input []string) []int {
    var final, count []int
    var converted, threshold int = 0, 0

    if len(input) == 0 {
        return []int{}
    }

    count = make([]int, len(input[0]))

    for i := 0; i < len(input); i++ {
        if input[i] == "" {
            break
        }

        for j := 0; j < len(input[i]); j++ {
            converted, _ = strconv.Atoi(string(input[i][j]))
            count[j] += converted
        }
        threshold += (i+1)%2
    }

    final = make([]int, len(count))

    for i := 0; i < len(count); i++ {
        if count[i] >= threshold {
            final[i] = 1
        } else {
            final[i] = 0
        }
    }

    return final
}


func Run() {
    data, err := ioutil.ReadFile("inputs/binary_diagnostic")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }

    lines := strings.Split(string(data), "\n")

    most_common_value := BinaryDiagnostic(lines)
    opposite := GetOpposite(most_common_value)

    fmt.Printf("3 - Part 1: %d\n", ToDec(most_common_value)*ToDec(opposite))

    o2_list, co2_list := BinaryDiagnosticForRatings(lines)
    fmt.Printf("3 - Part 2: %d\n", (ToDec(o2_list)* ToDec(co2_list)))
}
