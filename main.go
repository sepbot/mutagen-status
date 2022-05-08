package main

import (
	"context"
	"fmt"
	"github.com/mutagen-io/mutagen/cmd/mutagen/daemon"
	"github.com/mutagen-io/mutagen/pkg/selection"
	"github.com/mutagen-io/mutagen/pkg/service/synchronization"
	"os"
	"strings"
)

func run() (string, bool, error) {
	gc, err := daemon.Connect(true, false)
	if err != nil {
		return "", false, fmt.Errorf("failed to connect to daemon: %v\n", err)
	}

	defer func() {
		if err := gc.Close(); err != nil {
			fmt.Printf("failed to close daemon connection: %v\n", err)
		}
	}()

	sync := synchronization.NewSynchronizationClient(gc)

	resp, err := sync.List(
		context.Background(),
		&synchronization.ListRequest{
			Selection: &selection.Selection{All: true},
		},
	)
	if err != nil {
		return "", false, fmt.Errorf("failed to get list of syncs: %v\n", err)
	}

	if err := resp.EnsureValid(); err != nil {
		return "", false, fmt.Errorf("list response was not valid: %v\n", err)
	}

	states := resp.GetSessionStates()

	output := make([]string, len(states))
	for i, state := range states {
		conflicts := len(state.GetConflicts())
		healthy := conflicts == 0 && isHealthy(state.GetStatus())

		if healthy {
			output[i] = "🟢"
		} else {
			output[i] = fmt.Sprintf("🔴(%v)", conflicts)
		}
	}

	return strings.Join(output, " "), true, nil
}

func main() {
	out, healthy, err := run()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Print(out)
	if healthy {
		os.Exit(0)
	} else {
		os.Exit(2)
	}
}
