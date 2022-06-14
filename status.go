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

func getHealth(state *synchronization.State) (Status, int) {
	conflicts := len(state.GetConflicts())
	if conflicts > 0 {
		return StatusNotHealthy, conflicts
	}

	switch state.GetStatus() {
	case synchronization.Status_Disconnected:
		return StatusNotHealthy, conflicts
	case synchronization.Status_HaltedOnRootEmptied:
		return StatusNotHealthy, conflicts
	case synchronization.Status_HaltedOnRootDeletion:
		return StatusNotHealthy, conflicts
	case synchronization.Status_HaltedOnRootTypeChange:
		return StatusNotHealthy, conflicts
	case synchronization.Status_ConnectingAlpha:
		return StatusConnecting, conflicts
	case synchronization.Status_ConnectingBeta:
		return StatusConnecting, conflicts
	case synchronization.Status_Watching:
		return StatusHealthy, conflicts
	case synchronization.Status_Scanning:
		return StatusInProgress, conflicts
	case synchronization.Status_WaitingForRescan:
		return StatusConnecting, conflicts
	case synchronization.Status_Reconciling:
		return StatusInProgress, conflicts
	case synchronization.Status_StagingAlpha:
		return StatusInProgress, conflicts
	case synchronization.Status_StagingBeta:
		return StatusInProgress, conflicts
	case synchronization.Status_Transitioning:
		return StatusInProgress, conflicts
	case synchronization.Status_Saving:
		return StatusInProgress, conflicts
	default:
		return StatusNotHealthy, conflicts
	}
}
