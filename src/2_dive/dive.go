package dive

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

func Dive(input []string) int {
    var horizontal, depth, value int = 0, 0, 0
    var cmd string

    for i := 0; i < len(input); i++ {
        if input[i] == "" {
            break
        }
        raw_cmd := strings.Split(input[i], " ")
        cmd = raw_cmd[0]
        value, _ = strconv.Atoi(raw_cmd[1])

        if cmd == "forward" {
            horizontal += value
        } else if cmd == "down" {
            depth += value
        } else {
            depth -= value
        }
    }
    return horizontal*depth
}

func DiveWithAim(input []string) int {
    var horizontal, depth, aim, value int = 0, 0, 0, 0
    var cmd string

    for i := 0; i < len(input); i++ {
        if input[i] == "" {
            break
        }
        raw_cmd := strings.Split(input[i], " ")
        cmd = raw_cmd[0]
        value, _ = strconv.Atoi(raw_cmd[1])

        if cmd == "forward" {
            horizontal += value
            depth += value*aim
        } else if cmd == "down" {
            aim += value
        } else {
            aim -= value
        }
    }
    return horizontal*depth
}

func Run() {
    data, err := ioutil.ReadFile("inputs/dive")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }

    lines := strings.Split(string(data), "\n")

    fmt.Printf("2. Part 1: %d\n", Dive(lines))
    fmt.Printf("2. Part 2: %d\n", DiveWithAim(lines))
}
