package k8s

import (
	"testing"
)

func TestValidateResourceType(t *testing.T) {
	var tests = []struct {
		userType string
		wantType string
	}{
		{"unknown type", ""},
		{"replicaset", ResourceTypeReplicaSet},
		{"rs", ResourceTypeReplicaSet},
		{" Deployment", ResourceTypeDeployment},
	}

	c := &Client{}
	for _, tt := range tests {
		t.Run(tt.userType, func(t *testing.T) {
			got, _ := c.ValidateResourceType(tt.userType)
			if got != tt.wantType {
				t.Fatalf("got %s, expecting %s", got, tt.wantType)
			}
		})
	}
}
