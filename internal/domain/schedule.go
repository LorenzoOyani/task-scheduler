package domain

import (
	"time"
)

type TaskSegment struct {
	TASKID    string
	WORKERID  string
	StartTime time.Time
	EndTime   time.Time
}

type WorkSchedule struct {
	WorkerId    string
	AvailableAt time.Time
}

type ScheduleResult struct {
	Segments         []TaskSegment
	ScheduleTasks    []Task
	UnscheduledTasks []Task
}
