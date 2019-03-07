package main

import (
	"fmt"

	"go.i3wm.org/i3"
)

func main() {
	recv := i3.Subscribe(i3.WindowEventType, i3.WorkspaceEventType, i3.OutputEventType,
		i3.ModeEventType, i3.BarconfigUpdateEventType, i3.BindingEventType, i3.ShutdownEventType,
		i3.TickEventType)

	for recv.Next() {
		ev := recv.Event() //.(*i3.WindowEvent)

		eventType := "eventType"
		var change string

		switch ev.(type) {
		case *i3.WindowEvent:
			eventType = "window"
			change = ev.(*i3.WindowEvent).Change
		case *i3.TickEvent:
			eventType = "tick"
			change = "n/a"
		case *i3.OutputEvent:
			eventType = "output"
			change = ev.(*i3.OutputEvent).Change
		case *i3.WorkspaceEvent:
			eventType = "workspace"
			change = ev.(*i3.WorkspaceEvent).Change
		case *i3.ModeEvent:
			eventType = "mode"
			change = ev.(*i3.ModeEvent).Change
		case *i3.BarconfigUpdateEvent:
			eventType = "barconfig_update"
			change = "n/a"
		case *i3.BindingEvent:
			eventType = "binding"
			change = ev.(*i3.BindingEvent).Change
		case *i3.ShutdownEvent:
			eventType = "shutdown"
			change = ev.(*i3.ShutdownEvent).Change
		}

		fmt.Printf("%s:%s\n", eventType, change)
	}

	recv.Close()
}
