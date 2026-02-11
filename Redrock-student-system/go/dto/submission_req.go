package dto

// 作业接口
type FindSubmissionReq struct {
	HomeworkID uint64 `json:"homework_id"`
}

// 提交接口
type SubmitHomeworkReq struct {
	HomeworkID uint64 `json:"homework_id" binding:"required"`
	Content    string `json:"content" binding:"required"`
	FileUrl    string `json:"file_url"`
}
type CheckHomeworkReq struct {
	Score       *int   `json:"score"`
	Commit      string `json:"commit" binding:"required"`
	IsExcellent bool   `json:"is_excellent"`
}
type UpdateExcellentReq struct {
	IsExcellent bool `json:"is_excellent" binding:"required"`
}
type FindExcellentReq struct {
	Department string `json:"department"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
}
