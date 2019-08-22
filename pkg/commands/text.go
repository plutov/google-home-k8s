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

// a. 1 node
// b. 2 nodes
func formatClusterSize(size int) string {
	if size == 1 {
		return "there is 1 node"
	}

	return fmt.Sprintf("there are %d nodes", size)
}
