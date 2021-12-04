package sonar_sweep

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func SonarSweep(input []string) int {
    var converted int
    var err error
    var increases int = 0

    if len(input) <= 1 {
        return increases
    }

    converted_input := make([]int, len(input))
    for index, value := range(input) {
        if value == "" {
            break
        }
        converted, err = strconv.Atoi(value)
        if err != nil {
            panic(fmt.Sprintf("Error converting %v. %v", value, err))
        }
        converted_input[index] = converted
    }

    for i := 0; i < len(converted_input)-1; i++ {
        if converted_input[i] <= converted_input[i+1] {
            increases += 1
        }
    }
    return increases
}

func Run() {

    data, err := ioutil.ReadFile("inputs/sonar_sweep")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }

    lines := strings.Split(string(data), "\n")

    fmt.Printf("1. %d\n", SonarSweep(lines))
}
