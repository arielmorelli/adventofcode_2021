package sonar_sweep

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func SonarSweep(input []int) int {
    increases := 0
    if len(input) <= 1 {
        return increases
    }
    for i := 0; i < len(input)-1; i++ {
        if input[i] <= input[i+1] {
            increases += 1
        }
    }
    return increases
}

func Run() {
    var converted int

    data, err := ioutil.ReadFile("inputs/sonar_sweep")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }

    lines := strings.Split(string(data), "\n")
    input := make([]int, len(lines))
    for index, value := range(lines) {
        if value == "" {
            break
        }
        converted, err = strconv.Atoi(value)
        if err != nil {
            fmt.Printf("Error converting %v. %v\n", value, err)
            return
        }
        input[index] = converted
    }

    fmt.Printf("1. %d\n", SonarSweep(input))
}
