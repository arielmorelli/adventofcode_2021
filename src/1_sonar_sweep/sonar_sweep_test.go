package sonar_sweep

import "testing"


func RunTestCase(t *testing.T, entries []int, expected int) {
    increases := SonarSweep(entries)
    if increases != expected {
        t.Errorf("Input: '%v'. Wanted %d go %d", entries, expected, increases)
    }
}

func TestAll(t *testing.T){
    RunTestCase(t, []int{}, 0)
    RunTestCase(t, []int{1}, 0)
    RunTestCase(t, []int{1, 2}, 1)
    RunTestCase(t, []int{2, 1}, 0)
    RunTestCase(t, []int{1, 2, 3}, 2)
    RunTestCase(t, []int{1, 3, 2}, 1)
    RunTestCase(t, []int{3, 2, 1}, 0)
}
