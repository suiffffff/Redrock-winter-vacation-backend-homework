package dao

import (
	"system/dto"
	"system/models"
)

// 作业附加接口
func FindSubmissionCount(homeworkID uint64) (int64, error) {
	var count int64
	err := DB.Model(&models.Submission{}).
		Where("homework_id = ?", homeworkID).
		Count(&count).Error
	return count, err
}
func FindMySubmisssion(homeworkID, userID uint64) (*dto.MySubmissionInfo, error) {
	var result dto.MySubmissionInfo
	err := DB.Model(&models.Submission{}).
		Where("homework_id = ? AND student_id = ?", homeworkID, userID).
		Select("id, score, is_excellent").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}
	if result.ID == 0 {
		return nil, nil
	}

	return &result, nil
}
