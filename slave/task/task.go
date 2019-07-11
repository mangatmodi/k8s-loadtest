package task

import (
	"github.com/myzhan/boomer"
)

/**
* Data: used by task to execute
* ctx:  context to be shared across various task functions
* Build: build your dependency here
* Task: Boomer Task. Each task contains a function which is executed once
        per request
*/

type LocustTask struct {
	Task  *boomer.Task
	Data  interface{}
	Build func()
	ctx   map[string]interface{}
}

var TrackerClickTask *LocustTask
var TrackerConvTask *LocustTask
var Tasks map[string]*LocustTask

func init() {
	TrackerClickTask = &LocustTask{
		Task: &boomer.Task{
			Name:   "tracker-click",
			Weight: 1000,
			Fn:     makeClick,
		},
		Data:  []interface{}{},
		Build: buildTrackerClickTask,
		ctx:   map[string]interface{}{},
	}

	TrackerConvTask = &LocustTask{
		Task: &boomer.Task{
			Name:   "tracker-conv",
			Weight: 100,
			Fn:     makeConv,
		},
		Data:  []interface{}{},
		Build: buildTrackerConvTask,
		ctx:   map[string]interface{}{},
	}

	Tasks = map[string]*LocustTask{
		"tracker-click": TrackerClickTask,
		"tracker-conv":  TrackerConvTask,
	}
}
