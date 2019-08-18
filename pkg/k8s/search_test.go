package k8s

import (
	"testing"
)

func TestNormalizeName(t *testing.T) {
	var tests = []struct {
		userName string
		want     string
	}{
		{"nginx-ingress-controller", "nginxingresscontroller"},
		{"cert manager", "certmanager"},
		{"App manager", "appmanager"},
	}

	for _, tt := range tests {
		t.Run(tt.userName, func(t *testing.T) {
			got := normalizeName(tt.userName)
			if got != tt.want {
				t.Fatalf("got %s, expecting %s", got, tt.want)
			}
		})
	}
}
