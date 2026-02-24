package dao

import (
	"errors"
	"system/dto"
	"system/models"
	"system/pkg"
	"time"

	"gorm.io/gorm"
)

// 作业附加接口
func FindSubmissionCount(homeworkID uint64) (int64, error) {
	var count int64
	err := DB.Model(&models.Submission{}).
		Where("homework_id = ?", homeworkID).
		Count(&count).Error
	return count, err
}
func FindMySubmission(homeworkID, userID uint64) (*dto.MySubmissionInfo, error) {
	var result dto.MySubmissionInfo
	err := DB.Model(&models.Submission{}).
		Where("homework_id = ? AND student_id = ?", homeworkID, userID).
		Select("id, score, is_excellent").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}
	if result.ID == 0 {
		return nil, errors.New("submission not found")
	}

	return &result, nil
}

// 提交接口
func SubmitHomework(submission *models.Submission) error {
	var existing models.Submission
	err := DB.Where("homework_id = ? AND student_id = ?", submission.HomeworkID, submission.StudentID).
		First(&existing).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if submission.SubmittedAt.IsZero() {
			submission.SubmittedAt = time.Now()
		}
		return DB.Create(submission).Error
	}

	if err != nil {
		return err
	}

	updates := map[string]interface{}{
		"content":      submission.Content,
		"file_url":     submission.FileUrl,
		"submitted_at": time.Now(),
		"is_late":      submission.IsLate,
		"updated_at":   time.Now(),
	}
	return DB.Model(&existing).Updates(updates).Error
}
func FindSubmission(submission *models.Submission) (*dto.SubmitHomeworkRes, error) {
	var result dto.SubmitHomeworkRes
	err := DB.Model(&models.Submission{}).Where("homework_id = ? AND student_id = ?", submission.HomeworkID, submission.StudentID).
		Select("id, homework_id, submitted_at, is_late").
		Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func FindAllMySubmit(submission *models.Submission, page, pageSize int) (*dto.FindAllMySubmitRes, error) {
	var submissions []models.Submission
	var total int64
	query := DB.Model(&models.Submission{}).Where("student_id = ?", submission.StudentID)

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	err := query.Preload("Homework").
		Order("submitted_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&submissions).Error
	if err != nil {
		return nil, err
	}

	list := make([]dto.SubmissionItem, 0)
	for _, s := range submissions {
		list = append(list, dto.SubmissionItem{
			ID:          s.ID,
			Score:       s.Score,
			Comment:     s.Comment,
			IsExcellent: s.IsExcellent,
			SubmittedAt: s.SubmittedAt,
			Homework: dto.HomeworkMsg{
				ID:              s.Homework.ID,
				Title:           s.Homework.Title,
				Department:      s.Homework.Department,
				DepartmentLabel: pkg.GetDepartmentLabel(s.Homework.Department),
			},
		})
	}

	return &dto.FindAllMySubmitRes{
		List:     list,
		Total:    uint64(total),
		Page:     uint64(page),
		PageSize: uint64(pageSize),
	}, nil
}
func FindAllStudentSubmit(submission *models.Submission, page, pageSize int) (*dto.FindAllStudentRes, error) {
	var submissions []models.Submission
	var total int64
	query := DB.Model(&models.Submission{}).Where("homework_id = ?", submission.HomeworkID)
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}
	err := query.Preload("Student").
		Order("submitted_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&submissions).Error
	if err != nil {
		return nil, err
	}
	list := make([]dto.CommitItem, 0)
	for _, s := range submissions {
		list = append(list, dto.CommitItem{
			ID:          s.ID,
			Content:     s.Content,
			IsLate:      s.IsLate,
			Score:       s.Score,
			Comment:     s.Comment,
			SubmittedAt: s.SubmittedAt,
			Student: dto.StudentItem{
				ID:              s.Student.ID,
				NickName:        s.Student.Nickname,
				Department:      s.Student.Department,
				DepartmentLabel: pkg.GetDepartmentLabel(s.Student.Department),
			},
		})
	}
	return &dto.FindAllStudentRes{
		List:     list,
		Total:    uint64(total),
		Page:     uint64(page),
		PageSize: uint64(pageSize),
	}, nil
}
func FindSubmissionByID(submission *models.Submission) error {
	return DB.Preload("Homework").First(&submission, submission.ID).Error
}
func CheckHomework(submission *models.Submission) (*dto.CheckHomeworkRes, error) {
	err := DB.Model(&models.Submission{}).
		Where("id=?", submission.ID).
		Updates(map[string]interface{}{
			"score":        submission.Score,
			"comment":      submission.Comment,
			"is_excellent": submission.IsExcellent,
			"reviewer_id":  submission.ReviewerID,
			"reviewed_at":  time.Now(),
		}).Error
	if err != nil {
		return nil, err
	}
	result := dto.CheckHomeworkRes{
		ID:          submission.ID,
		Score:       submission.Score,
		Comment:     submission.Comment,
		IsExcellent: submission.IsExcellent,
		ReviewedAt:  submission.ReviewedAt,
	}
	return &result, err
}
func UpdateExcellent(submission *models.Submission) (*dto.UpdateExcellentRes, error) {
	err := DB.Model(&models.Submission{}).
		Where("id=?", submission.ID).
		Updates(map[string]interface{}{
			"is_excellent": submission.IsExcellent,
		}).Error
	if err != nil {
		return nil, err
	}
	result := dto.UpdateExcellentRes{
		ID:          submission.ID,
		IsExcellent: submission.IsExcellent,
	}
	return &result, err
}
func FindExcellent(department string, page, pageSize int) (*dto.FindExcellentRes, error) {
	var submissions []models.Submission
	var total int64

	query := DB.Model(&models.Submission{}).Where("is_excellent = ?", true)

	if department != "" {
		query = query.Joins("Homework").
			Where("homeworks.department = ?", department)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	err := query.Preload("Homework").Preload("Student").
		Order("submitted_at DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&submissions).Error

	if err != nil {
		return nil, err
	}

	list := make([]dto.ExcellentList, 0)
	for _, s := range submissions {
		list = append(list, dto.ExcellentList{
			ID: s.ID,
			Homework: dto.ExcellentHomeworkItem{
				ID:              s.Homework.ID,
				Title:           s.Homework.Title,
				Department:      s.Homework.Department,
				DepartmentLabel: pkg.GetDepartmentLabel(s.Homework.Department),
			},
			Student: dto.ExcellentStudentItem{
				ID:       s.Student.ID,
				NickName: s.Student.Nickname,
			},
			Score:   s.Score,
			Comment: s.Comment,
		})
	}

	return &dto.FindExcellentRes{
		List:     list,
		Total:    uint64(total),
		Page:     uint64(page),
		PageSize: uint64(pageSize),
	}, nil
}
