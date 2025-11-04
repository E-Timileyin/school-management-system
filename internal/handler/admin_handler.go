package handler

import (
	"strconv"

	"github.com/E-Timileyin/school-management-system/internal/models"
	"github.com/E-Timileyin/school-management-system/internal/service"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	userService   *service.UserService
	courseService *service.CourseService
}

func NewAdminHandler(userService *service.UserService, courseService *service.CourseService) *AdminHandler {
	return &AdminHandler{
		userService:   userService,
		courseService: courseService,
	}
}

// User Management
func (h *AdminHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userService.ListUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch users"})
		return
	}
	c.JSON(200, users)
}

func (h *AdminHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid user data"})
		return
	}

	if err := h.userService.CreateUser(&user); err != nil {
		c.JSON(500, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(201, user)
}

func (h *AdminHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, user)
}

func (h *AdminHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}

	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "invalid user data"})
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	// Update user fields
	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName
	user.Email = updateData.Email
	user.Role = updateData.Role

	if err := h.userService.UpdateUser(user); err != nil {
		c.JSON(500, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(200, user)
}

func (h *AdminHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": "failed to delete user"})
		return
	}

	c.Status(204)
}

// Course Management
func (h *AdminHandler) CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(400, gin.H{"error": "invalid course data"})
		return
	}

	if err := h.courseService.CreateCourse(&course); err != nil {
		c.JSON(500, gin.H{"error": "failed to create course"})
		return
	}

	c.JSON(201, course)
}

func (h *AdminHandler) UpdateCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid course ID"})
		return
	}

	var updateData models.Course
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "invalid course data"})
		return
	}

	course, err := h.courseService.GetCourseByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "course not found"})
		return
	}

	// Update course fields
	course.Name = updateData.Name
	course.Description = updateData.Description
	course.TeacherID = updateData.TeacherID

	if err := h.courseService.UpdateCourse(course); err != nil {
		c.JSON(500, gin.H{"error": "failed to update course"})
		return
	}

	c.JSON(200, course)
}

func (h *AdminHandler) DeleteCourse(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid course ID"})
		return
	}

	if err := h.courseService.DeleteCourse(uint(id)); err != nil {
		c.JSON(500, gin.H{"error": "failed to delete course"})
		return
	}

	c.Status(204)
}

func (h *AdminHandler) EnrollStudent(c *gin.Context) {
	courseID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid course ID"})
		return
	}

	var enrollment struct {
		StudentID uint `json:"student_id"`
	}

	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(400, gin.H{"error": "invalid enrollment data"})
		return
	}

	if err := h.courseService.EnrollStudent(uint(courseID), enrollment.StudentID); err != nil {
		c.JSON(500, gin.H{"error": "failed to enroll student"})
		return
	}

	c.Status(201)
}

func (h *AdminHandler) RemoveEnrollment(c *gin.Context) {
	courseID, err := strconv.ParseUint(c.Param("courseId"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid course ID"})
		return
	}

	studentID, err := strconv.ParseUint(c.Param("studentId"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid student ID"})
		return
	}

	if err := h.courseService.RemoveEnrollment(uint(courseID), uint(studentID)); err != nil {
		c.JSON(500, gin.H{"error": "failed to remove enrollment"})
		return
	}

	c.Status(204)
}
