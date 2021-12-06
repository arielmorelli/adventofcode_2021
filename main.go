package main

import (
    "fmt"
    sonarSweep "github.com/arielmorelli/adventofcode_2021/src/1_sonar_sweep"
    dive "github.com/arielmorelli/adventofcode_2021/src/2_dive"
    binaryDiagnostic "github.com/arielmorelli/adventofcode_2021/src/3_binary_diagnostic"
    giantSquid "github.com/arielmorelli/adventofcode_2021/src/4_giant_squid"

)

func main() {
    fmt.Println("Outputs:")
    sonarSweep.Run()
    fmt.Println("--------------------")
    dive.Run()
    fmt.Println("--------------------")
    binaryDiagnostic.Run()
    fmt.Println("--------------------")
    giantSquid.Run()
}
