package main

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_ValidInput(t *testing.T) {
	cases := map[string]bool{
		"0": true,
		"1": true,
		"2": false,
		"x": false,
	}

	for expectedInput, expectedOutput := range cases {
		actualOutput := validInput(expectedInput)
		if actualOutput != expectedOutput {
			t.Errorf("Expected case %s to return %t, got %t", expectedInput, expectedOutput, actualOutput)
		}
	}
}

func Test_Transition(t *testing.T) {
	type inputCase struct {
		CurrentState string
		Input        string
	}

	cases := map[inputCase]string{
		inputCase{CurrentState: "S0", Input: "0"}: "S0",
		inputCase{CurrentState: "S0", Input: "1"}: "S1",
		inputCase{CurrentState: "S1", Input: "0"}: "S2",
		inputCase{CurrentState: "S1", Input: "1"}: "S0",
		inputCase{CurrentState: "S2", Input: "0"}: "S1",
		inputCase{CurrentState: "S2", Input: "1"}: "S2",
	}

	for expectedInput, expectedOutput := range cases {
		actualOutput := transition(expectedInput.CurrentState, expectedInput.Input)
		if actualOutput != expectedOutput {
			t.Errorf("Expected case %+v to return %s, got %s", expectedInput, expectedOutput, actualOutput)
		}
	}
}

func Test_InterpretFinalState(t *testing.T) {
	type interpretOutput struct {
		Output string
		Final  bool
	}
	cases := map[string]interpretOutput{
		"S0":       interpretOutput{Output: "0", Final: true},
		"S1":       interpretOutput{Output: "1", Final: true},
		"S2":       interpretOutput{Output: "2", Final: true},
		"Snotreal": interpretOutput{Output: "invalid input", Final: false},
		"":         interpretOutput{Output: "invalid input", Final: false},
	}

	for expectedInput, expectedOutput := range cases {
		actualOutput, final := interpretFinalState(expectedInput)
		if final != expectedOutput.Final {
			t.Errorf("Expected case %s to have finality %t, got %t", expectedInput, expectedOutput.Final, final)
		}
		if actualOutput != expectedOutput.Output {
			t.Errorf("Expected case %s to return %s, got %s", expectedInput, expectedOutput.Output, actualOutput)
		}
	}
}

func Test_Fsa(t *testing.T) {
	cases := map[string]string{
		"110":  "0",
		"1010": "1",
	}

	for expectedInput, expectedOutput := range cases {
		actualOutput, err := fsa(expectedInput)
		if err != nil {
			t.Error(err)
		}
		if actualOutput != expectedOutput {
			t.Errorf("Expected case %s to return %s, got %s", expectedInput, expectedOutput, actualOutput)
		}
	}
}

func Test_FsaBig(t *testing.T) {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		testnum := rand.Int63()
		inputBinary := strconv.FormatInt(testnum, 2)
		expectedOutput := strconv.FormatInt(testnum%3, 10)

		actualOutput, err := fsa(inputBinary)
		if err != nil {
			t.Error(err)
		}
		if actualOutput != expectedOutput {
			t.Errorf("Expected input %s to return %s, got %s", string(testnum), expectedOutput, actualOutput)
		}
	}
}
