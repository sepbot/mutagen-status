package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mutagen-io/mutagen/cmd/external"
	"github.com/mutagen-io/mutagen/cmd/mutagen/daemon"
	"github.com/mutagen-io/mutagen/pkg/selection"
	"github.com/mutagen-io/mutagen/pkg/service/synchronization"
	"os"
	"strings"
)

func init() {
	external.UsePathBasedLookupForDaemonStart = true
}

func run() (string, int, error) {
	gc, err := daemon.Connect(true, false)
	if err != nil {
		return "", 0, fmt.Errorf("failed to connect to daemon: %v\n", err)
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
		return "", 0, fmt.Errorf("failed to get list of syncs: %v\n", err)
	}

	if err := resp.EnsureValid(); err != nil {
		return "", 0, fmt.Errorf("list response was not valid: %v\n", err)
	}

	states := resp.GetSessionStates()

	exitCode := 0
	output := make([]string, len(states))
	for i, state := range states {
		health, conflicts := getHealth(state)

		// conflicts will override
		if conflicts > 0 {
			health = StatusNotHealthy
		}

		switch health {
		case StatusConnecting:
			output[i] = "\U0001F7E1"
			exitCode += 1
		case StatusInProgress:
			output[i] = "\U0001F535"
		case StatusHealthy:
			output[i] = "\U0001F7E2"
		case StatusNotHealthy:
			fallthrough
		default:
			exitCode += 1
			output[i] = "\U0001F534"
		}

		if conflicts > 0 {
			output[i] = output[i] + fmt.Sprintf("%v", conflicts)
		}
	}

	return strings.Join(output, " "), exitCode, nil
}

func main() {
	var quiet bool
	var template string

	flag.BoolVar(&quiet, "quiet", false, "don't produce any output if healthy")
	flag.StringVar(
		&template, "template", "%v",
		"template the output using %v where the regular output should go eg: mutagen:(%v)",
	)

	flag.Parse()

	out, exitCode, err := run()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(127)
	}

	if exitCode != 0 || !quiet {
		fmt.Print(strings.ReplaceAll(template, "%v", out))
	}

	os.Exit(exitCode)
}
