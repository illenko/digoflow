package core

import "github.com/google/uuid"

type Execution struct {
	ID     string         `json:"id"`
	FlowID string         `json:"flowId"`
	Values map[string]any `json:"values"`
	Output map[string]any `json:"output"`
}

func NewExecution(flowId string) Execution {
	return Execution{
		ID:     uuid.New().String(),
		FlowID: flowId,
		Values: map[string]any{},
		Output: map[string]any{},
	}
}
