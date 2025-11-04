package handler

import (
	"strconv"

	"github.com/E-Timileyin/school-management-system/internal/models"
	"github.com/E-Timileyin/school-management-system/internal/service"
	"github.com/gin-gonic/gin"
)

type CourseHandler struct {
	courseService     *service.CourseService
	enrollmentService *service.EnrollmentService
}

func NewCourseHandler(
	courseService *service.CourseService,
	enrollmentService *service.EnrollmentService,
) *CourseHandler {
	return &CourseHandler{
		courseService:     courseService,
		enrollmentService: enrollmentService,
	}
}

// get all courses
func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	courses, err := h.courseService.ListCourses()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch courses"})
		return
	}
	c.JSON(200, courses)
}

func (h *CourseHandler) GetCourseByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid course ID"})
		return
	}

	course, err := h.courseService.GetCourseByID(uint(id))
	if err != nil {
		c.JSON(404, gin.H{"error": "course not found"})
		return
	}

	c.JSON(200, course)
}

func (h *CourseHandler) GetCourseStudents(c *gin.Context) {
	courseID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid course ID"})
		return
	}

	enrollments, err := h.courseService.GetCourseStudents(uint(courseID))
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch course students"})
		return
	}

	c.JSON(200, enrollments)
}

func (h *CourseHandler) GetMyEnrollments(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	studentID := user.(*models.User).ID
	enrollments, err := h.courseService.GetStudentEnrollments(studentID)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to fetch enrollments"})
		return
	}

	c.JSON(200, enrollments)
}

func (h *CourseHandler) EnrollInCourse(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	courseID, err := strconv.ParseUint(c.Param("courseId"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid course ID"})
		return
	}

	studentID := user.(*models.User).ID
	if err := h.courseService.EnrollStudent(uint(courseID), studentID); err != nil {
		c.JSON(500, gin.H{"error": "failed to enroll in course"})
		return
	}

	c.Status(201)
}

// internal/handler/course_handler.go
func (h *CourseHandler) WithdrawFromCourse(c *gin.Context) {
	enrollmentID, err := strconv.ParseUint(c.Param("enrollmentId"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid enrollment ID"})
		return
	}

	// Use the enrollmentID to find and delete the enrollment
	if err := h.enrollmentService.DeleteEnrollment(uint(enrollmentID)); err != nil {
		c.JSON(500, gin.H{"error": "failed to withdraw from course"})
		return
	}

	c.Status(204)
}
