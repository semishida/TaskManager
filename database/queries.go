package database

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task/models"
)

func GetHandler(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser := user.(models.User)

	var tasks []models.Task
	if err := Db.Where("user_id = ?", currentUser.ID).Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.Report{
			Status:  "Error",
			Message: "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func PostHandler(c *gin.Context) {
	var task models.Task

	user, _ := c.Get("user")
	currentUser := user.(models.User)

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, models.Report{
			Status:  "Error",
			Message: "Could not create task",
		})
		return
	}

	task.UserID = uint(currentUser.ID)

	if err := Db.Create(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, models.Report{
			Status:  "Error",
			Message: "Could not create task in database",
		})
		return
	}
	c.JSON(http.StatusOK, models.Report{
		Status:  "Success",
		Message: "Task created successfully",
	})
}

func DeleteHandler(c *gin.Context) {
	user, _ := c.Get("user")
	currentUser := user.(models.User)

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Report{
			Status:  "Error",
			Message: "Could not parse id",
		})
		return
	}

	var task models.Task
	if err := Db.Where("id = ? AND user_id = ?", id, currentUser.ID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, models.Report{
			Status:  "Error",
			Message: "Task not found",
		})
		return
	}

	Db.Delete(&task)
	c.JSON(http.StatusOK, models.Report{
		Status:  "Success",
		Message: "Task deleted successfully",
	})
}

func PatchHandler(c *gin.Context) {
	var UpdatedTask models.Task
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Report{
			Status:  "Error",
			Message: "Could not parse id",
		})
		return
	}

	if err := c.ShouldBindJSON(&UpdatedTask); err != nil {
		c.JSON(http.StatusBadRequest, models.Report{
			Status:  "Error",
			Message: "Could not update task",
		})
		return
	}

	result := Db.Model(&models.Task{}).Where("id = ?", id).Updates(UpdatedTask)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Report{
			Status:  "Error",
			Message: "Could not update task in database",
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.Report{
			Status:  "Error",
			Message: "Task not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.Report{
		Status:  "Success",
		Message: "Task updated successfully",
	})
}
