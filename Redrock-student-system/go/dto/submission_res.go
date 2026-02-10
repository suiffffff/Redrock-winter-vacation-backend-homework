package dto

// 作业额外接口
type MySubmissionInfo struct {
	ID          uint64 `json:"id"`
	Score       *int   `json:"score"`
	IsExcellent bool   `json:"is_excellent"`
}
