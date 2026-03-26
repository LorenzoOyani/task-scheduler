package engine

import (
	"sort"
	"task-scheduler/cmd/internal/domain"
	"time"
)


func SortReadyTask(tasks []domain.Task, now time.Time){

	sort.Slice(tasks, func(i, j int) bool { ///a callback func that sort the tasks
		si := scoreTask(tasks[i], now)
		sj := scoreTask(tasks[j], now)

		if si == sj {
			return tasks[i].CreatedAt.Before(tasks[j].CreatedAt)	
			}

			return si >sj
	})
}

func scoreTask(task domain.Task, now time.Time) int {
	score := 0

	switch task.Priority {
	case domain.Critical:
		score += 1000
	case domain.High:
		score += 700
	case domain.Medium:
		score += 400
	case domain.Low:
		score += 100
	}

	score += task.GravityScore * 10

	expectedProfit := int(task.ExpectedProfit() / 100)
	score += expectedProfit

	if task.DueDate != nil {
		
		hoursUntilDue := int(task.DueDate.Sub(now).Hours())
		switch {
		case hoursUntilDue <= 24:
			score += 500
		case hoursUntilDue <= 72:
			score += 250
		case hoursUntilDue <= 168:
			score += 100
		}
	}

	return score
}