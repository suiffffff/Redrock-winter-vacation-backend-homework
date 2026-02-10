package service

import (
	"system/dao"
	"system/dto"
)

// 作业附加接口
func FindSubmissionCount(homeworkID uint64) (int64, error) {
	return dao.FindSubmissionCount(homeworkID)
}
func FindMySubmission(homeworkID, userID uint64) (*dto.MySubmissionInfo, error) {
	return dao.FindMySubmisssion(homeworkID, userID)
}
