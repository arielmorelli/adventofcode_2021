package dive

import "testing"

func RunTestCase(t *testing.T, entries []string, expected int) {
    increases := Dive(entries)
    if increases != expected {
        t.Errorf("Input: '%v'. Wanted %d got %d", entries, expected, increases)
    }
}

func RunTestCaseWithAim(t *testing.T, entries []string, expected int) {
    increases := DiveWithAim(entries)
    if increases != expected {
        t.Errorf("Input: '%v'. Wanted %d got %d", entries, expected, increases)
    }
}

func TestDive(t *testing.T){
    var input []string
    RunTestCase(t, []string{}, 0)
    RunTestCase(t, []string{"forward 10"}, 0)
    RunTestCase(t, []string{"up 10"}, 0)
    RunTestCase(t, []string{"down 10"}, 0)

    input = []string{
        "forward 3", "down 5",
    }
    RunTestCase(t, input, 15)

    input = []string{
        "forward 3", "up 5",
    }
    RunTestCase(t, input, -15)

    input = []string{
        "forward 5", "down 5", "up 5",
    }
    RunTestCase(t, input, 0)

    input = []string{
        "down 5", "up 3",
    }
    RunTestCase(t, input, 0)

    input = []string{
        "forward 5", "down 5", "forward 8",
        "up 3", "down 8", "forward 2",
    }
    RunTestCase(t, input, 150)
}

func TestDiveWithAim(t *testing.T){
    var input []string
    RunTestCaseWithAim(t, []string{}, 0)
    RunTestCaseWithAim(t, []string{"forward 10"}, 0)
    RunTestCaseWithAim(t, []string{"up 10"}, 0)
    RunTestCaseWithAim(t, []string{"down 10"}, 0)

    input = []string{
        "forward 3", "down 5",
    }
    RunTestCaseWithAim(t, input, 0)

    input = []string{
        "forward 3", "up 5",
    }
    RunTestCaseWithAim(t, input, 0)

    input = []string{
        "forward 5", "down 5", "up 5",
    }
    RunTestCaseWithAim(t, input, 0)

    input = []string{
        "down 5", "up 3",
    }
    RunTestCaseWithAim(t, input, 0)

    input = []string{
        "forward 5", "down 5", "forward 5",
    }
    RunTestCaseWithAim(t, input, 10*5*5)

    input = []string{
        "forward 5", "down 5", "forward 8",
        "up 3", "down 8", "forward 2",
    }
    RunTestCaseWithAim(t, input, 900)
}
