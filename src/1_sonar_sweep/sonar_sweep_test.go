package sonar_sweep

import "testing"


func RunTestCase(t *testing.T, entries []string, expected int) {
    increases := SonarSweep(entries)
    if increases != expected {
        t.Errorf("Input: '%v'. Wanted %d got %d", entries, expected, increases)
    }
}

func TestAll(t *testing.T){
    RunTestCase(t, []string{}, 0)
    RunTestCase(t, []string{"1"}, 0)
    RunTestCase(t, []string{"1", "2"}, 1)
    RunTestCase(t, []string{"2", "1"}, 0)
    RunTestCase(t, []string{"1", "2", "3"}, 2)
    RunTestCase(t, []string{"1", "3", "2"}, 1)
    RunTestCase(t, []string{"3", "2", "1"}, 0)
}
