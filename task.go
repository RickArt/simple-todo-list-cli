package main

const (
	DONE_SYMBOL     = "🟢"
	NOT_DONE_SYMBOL = "🔴"
)

type Task struct {
	Id          int
	Description string
	Done        bool
}

func (t *Task) getDoneSymbol() string {
	if t.Done {
		return DONE_SYMBOL
	}
	return NOT_DONE_SYMBOL
}
