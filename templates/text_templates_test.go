package templates

import "testing"

func Test_displayPerson(t *testing.T) {
	tests := []struct {
		name string
		person Person
		output string
	}{
		{"Siva Bio", Person{ Name: "Siva", Job: "A Software Developer"}, "Name: Siva, Job: A Software Developer"},
		{"Prasad Bio", Person{ Name: "Prasad", Job: "A Lawyer"}, "Name: Prasad, Job: A Lawyer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := displayPerson(tt.person); got != tt.output {
				t.Errorf("displayPage() = %v, want %v", got, tt.output)
			}
		})
	}
}