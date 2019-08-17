package commands

import (
	"fmt"
	"testing"
)

func TestFormatReplicasNumber(t *testing.T) {
	var tests = []struct {
		replicas int32
		want     string
	}{
		{0, "there are 0 replicas"},
		{1, "there is 1 replica"},
		{2, "there are 2 replicas"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.replicas), func(t *testing.T) {
			got := formatReplicasNumber(tt.replicas)
			if got != tt.want {
				t.Fatalf("got %s, expecting %s", got, tt.want)
			}
		})
	}
}
