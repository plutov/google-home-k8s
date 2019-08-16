package commands

import "fmt"

// a. there are 3 replicas
// b. there is 1 replica
func formatReplicasNumber(replicasCount int32) string {
	if replicasCount == 1 {
		return "there is 1 replica"
	}

	return fmt.Sprintf("there are %d replicas", replicasCount)
}
