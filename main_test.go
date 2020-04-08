package main

import "testing"

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

func Test_FullTests(t *testing.T) {
	cases := map[string]string{
		"110": "0",
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