package service

import (
	"system/dao"
	"system/dto"
	"system/models"
)

// 作业附加接口
func FindSubmissionCount(homeworkID uint64) (int64, error) {
	return dao.FindSubmissionCount(homeworkID)
}
func FindMySubmission(homeworkID, userID uint64) (*dto.MySubmissionInfo, error) {
	return dao.FindMySubmission(homeworkID, userID)
}

// 提交接口
func SubmitHomework(req *dto.SubmitHomeworkReq, studentID uint64) (*dto.SubmitHomeworkRes, error) {
	submissionmodel := models.Submission{
		StudentID:  studentID,
		HomeworkID: req.HomeworkID,
		Content:    req.Content,
		FileUrl:    req.FileUrl,
	}
	err := dao.SubmitHomework(&submissionmodel)
	if err != nil {
		return nil, err
	}
	return dao.FindSubmission(&submissionmodel)
}
func FindAllMySubmit(studentID uint64, page, pageSize int) (*dto.FindAllMySubmitRes, error) {
	submissionmodel := models.Submission{
		StudentID: studentID,
	}
	return dao.FindAllMySubmit(&submissionmodel, page, pageSize)
}
func FindAllStudentSubmit(HomeworkID uint64, page, pageSize int) (*dto.FindAllStudentRes, error) {
	submissionmodel := models.Submission{
		HomeworkID: HomeworkID,
	}
	return dao.FindAllStudentSubmit(&submissionmodel, page, pageSize)
}
func FindSubmissionByID(submissionID uint64) (*models.Submission, error) {
	submissionmodel := models.Submission{
		ID: submissionID,
	}
	err := dao.FindSubmissionByID(&submissionmodel)
	if err != nil {
		return nil, err
	}
	return &submissionmodel, nil
}
func CheckHomework(req *dto.CheckHomeworkReq, submissionID uint64) (*dto.CheckHomeworkRes, error) {
	submissionmodel := models.Submission{
		Score:       req.Score,
		Comment:     req.Commit,
		IsExcellent: req.IsExcellent,
		ID:          submissionID,
	}
	return dao.CheckHomework(&submissionmodel)
}
func UpdateExcellent(req *dto.UpdateExcellentReq, submissionId uint64) (*dto.UpdateExcellentRes, error) {
	submissionmodel := models.Submission{
		ID:          submissionId,
		IsExcellent: req.IsExcellent,
	}
	return dao.UpdateExcellent(&submissionmodel)
}
func FindExcellent(req *dto.FindExcellentReq) (*dto.FindExcellentRes, error) {
	return dao.FindExcellent(req.Department, req.Page, req.PageSize)
}
