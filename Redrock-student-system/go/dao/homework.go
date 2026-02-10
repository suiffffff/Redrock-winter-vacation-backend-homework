package dao

import "system/models"

func AddHomework(homework *models.Homework) error {
	return DB.Create(homework).Error
}
func FindHomework(homework *models.Homework) ([]models.Homework, int64, error) {
	var list []models.Homework
	var total int64

	query := DB.Model(&models.Homework{})

	if homework.Department != "" {
		query = query.Where("department = ?", homework.Department)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Order("id").Find(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
func FindHomeworkByID(homework *models.Homework) error {
	return DB.Preload("Creator").First(&homework, homework.ID).Error
}
