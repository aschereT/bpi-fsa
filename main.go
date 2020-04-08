package main

import (
	"fmt"
	"os"
	"strings"
)

const initialState = "S0"

var inputAlphabet = map[string]bool{
	"0": true,
	"1": true,
}

var stateTransitionTable = map[string]map[string]string{
	"0": map[string]string{
		"S0": "S0",
		"S1": "S2",
		"S2": "S1",
	},
	"1": map[string]string{
		"S0": "S1",
		"S1": "S0",
		"S2": "S2",
	},
}

var finalStates = map[string]string{
	"S0": "0",
	"S1": "1",
	"S2": "2",
}

func validInput(input string) bool {
	_, ex := inputAlphabet[input]
	return ex
}

func transition(currentState string, input string) (finalState string) {
	return stateTransitionTable[input][currentState]
}

func interpretFinalState(state string) (output string, final bool) {
	output, final = finalStates[state]
	if !final {
		output = "invalid input"
	}
	return
}

func fsa(inputString string) (string, error) {
	strings := strings.Split(inputString, "")

	state := initialState
	for i := 0; i < len(strings); i++ {
		currentInput := strings[i]
		if validInput(currentInput) {
			state = transition(state, currentInput)
		} else {
			return "", fmt.Errorf("Invalid input %s at index %d", currentInput, i)
		}
	}

	output, final := interpretFinalState(state)
	if !final {
		return output, fmt.Errorf("Last state is not final state")
	}
	return output, nil
}

func main() {
	inputs := os.Args[1:]

	if len(inputs) == 0 {
		fmt.Println("Usage: bpi-fsa input1 input2 ...")
		fmt.Println("Example: bpi-fsa 110 1010")
	}

	for _, input := range inputs {
		output, err := fsa(input)
		if err != nil {
			fmt.Println(input, ":", output, "#", err)
		} else {
			fmt.Println(input, ":", output)
		}
	}
}
