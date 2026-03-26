package engine

import (
	"fmt"
	"task-scheduler/cmd/internal/domain"
)

func ValidateDependencies(tasks []domain.Task) error {
	taskMap := make(map[string]domain.Task, len(tasks))
	for _, task := range tasks {
		taskMap[task.ID] = task
	}

	for _, task := range tasks {
		for _, dep := range task.Dependencies {
			if _, exists := taskMap[dep]; !exists {
				return fmt.Errorf("task %s depends on unknown task %s", task.ID, dep)
			}
		}
	}

	visited := make(map[string]int) // 0 = unvisited, 1 = visiting, 2 = visited

	var dfs func(taskID string) error
	dfs = func(taskID string) error {
		if visited[taskID] == 1 {
			return fmt.Errorf("cyclic dependency detected at task %s", taskID)
		}
		if visited[taskID] == 2 {
			return nil
		}

		visited[taskID] = 1
		for _, dep := range taskMap[taskID].Dependencies {
			if err := dfs(dep); err != nil {
				return err
			}
		}
		visited[taskID] = 2
		return nil
	}

	for _, task := range tasks {
		if err := dfs(task.ID); err != nil { 
			return err
		}
	}

	return nil
}
