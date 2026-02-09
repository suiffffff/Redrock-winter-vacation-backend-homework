package service

import (
	"system/dao"
	"system/dto"
	"system/models"
)

func AddHomework(req dto.AddHomeworkReq, userID uint64) (*models.Homework, error) {
	homeworkmodel := models.Homework{
		Title:       req.Title,
		Description: req.Description,
		Department:  req.Department,
		Deadline:    req.Deadline,
		AllowLate:   req.AllowLate,
		CreatorID:   userID,
	}
	err := dao.AddHomework(&homeworkmodel)
	if err != nil {
		return nil, err
	}
	return &homeworkmodel, nil
}
func FindHomework(req dto.FindHomeworkReq) ([]models.Homework, int64, error) {
	homeworkmodel := models.Homework{
		Department: req.Department,
	}
	list, total, err := dao.FindHomework(&homeworkmodel)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
