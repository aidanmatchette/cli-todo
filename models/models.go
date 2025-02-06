package models


type Payload struct {
    Command string
    Record int
    Description string
}

type TaskStatus string

const (
    Todo TaskStatus = "todo"
    InProgress TaskStatus = "in-progress"
    Done TaskStatus = "done"
)

type Task struct {
    Id int
    Description string
    Status TaskStatus
}
