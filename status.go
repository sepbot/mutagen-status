package main

import (
	"github.com/mutagen-io/mutagen/pkg/synchronization"
)

type Status int

const (
	StatusHealthy = iota
	StatusConnecting
	StatusInProgress
	StatusNotHealthy
)

func isHealthy(status synchronization.Status) Status {
	switch status {
	case synchronization.Status_Disconnected:
		return StatusNotHealthy
	case synchronization.Status_HaltedOnRootEmptied:
		return StatusNotHealthy
	case synchronization.Status_HaltedOnRootDeletion:
		return StatusNotHealthy
	case synchronization.Status_HaltedOnRootTypeChange:
		return StatusNotHealthy
	case synchronization.Status_ConnectingAlpha:
		return StatusConnecting
	case synchronization.Status_ConnectingBeta:
		return StatusConnecting
	case synchronization.Status_Watching:
		return StatusHealthy
	case synchronization.Status_Scanning:
		return StatusInProgress
	case synchronization.Status_WaitingForRescan:
		return StatusConnecting
	case synchronization.Status_Reconciling:
		return StatusInProgress
	case synchronization.Status_StagingAlpha:
		return StatusInProgress
	case synchronization.Status_StagingBeta:
		return StatusInProgress
	case synchronization.Status_Transitioning:
		return StatusInProgress
	case synchronization.Status_Saving:
		return StatusInProgress
	default:
		return StatusNotHealthy
	}
}
