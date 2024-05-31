package model

type Execution struct {
	ID          string         `json:"id"`
	FlowID      string         `json:"flowId"`
	CurrentTask int            `json:"currentTask"`
	Values      map[string]any `json:"values"`
}
