package handler

import (
	"strconv"
	"system/dto"
	"system/pkg"
	"system/service"
	"time"

	"github.com/gin-gonic/gin"
)

func SubmitHomework(c *gin.Context) {
	var req dto.SubmitHomeworkReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.Error(c, 400, "参数错误")
		return
	}
	userID, err := pkg.GetUserID(c)
	if err != nil {
		pkg.ErrorWithStatus(c, 401, pkg.CodeAuthError, err.Error())
		return
	}
	user, err := service.GetProfile(userID)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "查询身份失败")
		return
	}
	if user.Role != "student" {
		pkg.Error(c, pkg.CodeNoPermission, "老登就不要交作业了，亲")
		return
	}
	homework, err := service.FindHomeworkByID(req.HomeworkID)
	if err != nil {
		pkg.Error(c, pkg.CodeNotFound, "作业不存在")
		return
	}
	if time.Now().After(homework.Deadline) && !homework.AllowLate {
		pkg.Error(c, pkg.CodeNoPermission, "已截止，禁止提交")
		return
	}
	resp, err := service.SubmitHomework(&req, userID)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "提交失败")
		return
	}
	pkg.Success(c, "提交成功", resp)
}
func FindAllMySubmit(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}
	userID, err := pkg.GetUserID(c)
	if err != nil {
		pkg.ErrorWithStatus(c, 401, pkg.CodeAuthError, err.Error())
		return
	}
	user, err := service.GetProfile(userID)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "查询身份失败")
		return
	}
	if user.Role != "student" {
		pkg.Error(c, pkg.CodeNoPermission, "可以看看去年的作业哦，亲")
		return
	}
	resp, err := service.FindAllMySubmit(userID, page, pageSize)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "查询错误")
		return
	}
	pkg.Success(c, "success", resp)
}
func FindAllStudentSubmit(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}
	idStr := c.Param("homework_id")
	homeworkID, _ := strconv.ParseUint(idStr, 10, 64)
	Homework, err := service.FindHomeworkByID(homeworkID)
	if err != nil {
		pkg.Error(c, pkg.CodeNotFound, "作业不存在")
		return
	}
	userID, err := pkg.GetUserID(c)
	if err != nil {
		pkg.ErrorWithStatus(c, 401, pkg.CodeAuthError, err.Error())
		return
	}
	user, err := service.GetProfile(userID)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "查询身份失败")
		return
	}
	if user.Role != "admin" && user.Department != Homework.Department {
		pkg.Error(c, pkg.CodeNoPermission, "你无权限修改哦，亲")
		return
	}
	resp, err := service.FindAllStudentSubmit(homeworkID, page, pageSize)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "查询错误")
		return
	}
	pkg.Success(c, "success", resp)
}
func CheckHomework(c *gin.Context) {
	var req dto.CheckHomeworkReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.Error(c, pkg.CodeParamError, "参数错误")
		return
	}
	idStr := c.Param("id")
	submissionID, _ := strconv.ParseUint(idStr, 10, 64)
	submission, err := service.FindSubmissionByID(submissionID)
	if err != nil {
		pkg.Error(c, pkg.CodeNotFound, "未查询到作业")
		return
	}
	userID, err := pkg.GetUserID(c)
	if err != nil {
		pkg.ErrorWithStatus(c, 401, pkg.CodeAuthError, err.Error())
		return
	}
	user, err := service.GetProfile(userID)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "查询身份失败")
		return
	}
	if user.Role != "admin" && user.Department != submission.Homework.Department {
		pkg.Error(c, pkg.CodeNoPermission, "你无权限修改哦，亲")
		return
	}
	resp, err := service.CheckHomework(&req, submissionID)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "批改失败")
	}
	pkg.Success(c, "批改成功", resp)
}

// 感觉有点多余
func UpdateExcellent(c *gin.Context) {
	var req dto.UpdateExcellentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.Error(c, pkg.CodeParamError, "参数错误")
		return
	}
	idStr := c.Param("id")
	submissionID, _ := strconv.ParseUint(idStr, 10, 64)
	submission, err := service.FindSubmissionByID(submissionID)
	if err != nil {
		pkg.Error(c, pkg.CodeNotFound, "未查询到作业")
		return
	}
	userID, err := pkg.GetUserID(c)
	if err != nil {
		pkg.ErrorWithStatus(c, 401, pkg.CodeAuthError, err.Error())
		return
	}
	user, err := service.GetProfile(userID)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "查询身份失败")
		return
	}
	if user.Role != "admin" && user.Department != submission.Homework.Department {
		pkg.Error(c, pkg.CodeNoPermission, "你无权限修改哦，亲")
		return
	}
	resp, err := service.UpdateExcellent(&req, submissionID)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "修改错误")
		return
	}
	pkg.Success(c, "标记成功", resp)
}
func FindExcellent(c *gin.Context) {
	var req dto.FindExcellentReq
	if err := c.ShouldBindQuery(&req); err != nil {
		pkg.Error(c, pkg.CodeParamError, "参数错误")
		return
	}
	resp, err := service.FindExcellent(&req)
	if err != nil {
		pkg.Error(c, pkg.CodeSystemError, "查询失败")
		return
	}
	pkg.Success(c, "success", resp)
}
