package model

type Execution struct {
	ID     string         `json:"id"`
	FlowID string         `json:"flowId"`
	Values map[string]any `json:"values"`
}
