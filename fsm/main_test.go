package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/looplab/fsm"
)

func Test_fsm(t *testing.T) {
	consumerState := fsm.NewFSM(
		"stopped",
		fsm.Events{
			{Name: "start", Src: []string{"stopped"}, Dst: "started"},
			{Name: "pause", Src: []string{"started"}, Dst: "paused"},
		},
		fsm.Callbacks{},
	)

	maintenanceState := fsm.NewFSM(
		"normal",
		fsm.Events{
			{Name: "normal", Src: []string{"maintenance"}, Dst: "normal"},
			{Name: "maintenance", Src: []string{"normal"}, Dst: "maintenance"},
		},
		fsm.Callbacks{
			"normal":      func(_ context.Context, e *fsm.Event) { fmt.Println("back to normal", e) },
			"maintenance": func(_ context.Context, e *fsm.Event) { fmt.Println("maintenance mode", e) },
		},
	)

	err := consumerState.Event(context.Background(), "start")
	if err != nil {
		fmt.Println(err)
	}
	err = maintenanceState.Event(context.Background(), "maintenance")
	if err != nil {
		fmt.Println(err)
	}
	err = consumerState.Event(context.Background(), "pause")
	if err != nil {
		fmt.Println(err)
	}
}
