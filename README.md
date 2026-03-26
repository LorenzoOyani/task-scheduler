# Task Scheduler (Golang)

A domain-driven, greedy-based task scheduling system built in Go.

This project simulates real-world task allocation across workers with constraints such as:

- Worker capacity (daily hours)
- Parallel task limits
- Resource type matching
- Task dependencies
- Priority-based scheduling
- Safe concurrency using goroutines

---

# Features

- Greedy scheduling algorithm
- Domain-driven design (DDD-inspired)
- Worker constraints:
  - Max daily working hours
  - Max parallel tasks
  - Active/inactive status
- Task constraints:
  - Dependencies
  - Duration
  - Resource type
  - Priority
- Concurrency support:
  - Parallel task validation
  - Parallel worker eligibility checks
  - Async schedule persistence
  - Async event publishing
- Extendable architecture (PostgreSQL, Kafka, Redis)

---

# Project Structure
#  Task Scheduler (Golang)

A domain-driven, greedy-based task scheduling system built in Go.

This project simulates real-world task allocation across workers with constraints such as:

- Worker capacity (daily hours)
- Parallel task limits
- Resource type matching
- Task dependencies
- Priority-based scheduling
- Safe concurrency using goroutines

---



#  Project Structure
/domain
├── task.go
├── worker.go
├── schedule.go

/scheduler
├── greedy_scheduler.go
├── pick_worker.go
├── allocate.go

/application
├── build_schedule.go

/infrastructure
├── repository.go
├── event_producer.go
