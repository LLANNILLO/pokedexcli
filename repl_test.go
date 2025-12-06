package main

import (
	"testing"
	"github.com/google/go-cmp/cmp"
)

func TestCleanInput (t *testing.T) {

	tests := map[string]struct{
		input string
		output []string
	}{
		"hello world testing": {input: " hello world ", output:[]string{"hello","world"}},
		"pokemon testing": {input:"Charmander Bulbasur PIKACHU", output:[]string{"charmander","bulbasur","pikachu"}},
	}

	for test, tc := range tests{
		t.Run(test, func(t *testing.T) {
			got := cleanInput(tc.input)
			diff := cmp.Diff(tc.output, got)

			if len(got) != len(tc.output) {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}

			for i := range got {
				word := got[i]
				outputWord := tc.output[i]

				if word != outputWord {
					t.Fatalf("expected: %#v, got: %#v", outputWord, got)
				}
			}
		})
	}

}
