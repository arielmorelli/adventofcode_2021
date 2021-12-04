package main

import (
    "fmt"
    sonar_sweep "github.com/arielmorelli/adventofcode_2021/src/1_sonar_sweep"
    dive "github.com/arielmorelli/adventofcode_2021/src/2_dive"
    binary_diagnostic "github.com/arielmorelli/adventofcode_2021/src/3_binary_diagnostic"
)

func main() {
    fmt.Println("Outputs:")
    sonar_sweep.Run()
    dive.Run()
    binary_diagnostic.Run()
}
