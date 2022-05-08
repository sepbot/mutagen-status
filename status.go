package main

import (
	"github.com/mutagen-io/mutagen/pkg/synchronization"
)

func isHealthy(status synchronization.Status) bool {
	return status != synchronization.Status_Disconnected &&
		status != synchronization.Status_HaltedOnRootEmptied &&
		status != synchronization.Status_HaltedOnRootDeletion &&
		status != synchronization.Status_HaltedOnRootTypeChange &&
		status != synchronization.Status_ConnectingAlpha &&
		status != synchronization.Status_ConnectingBeta
}
