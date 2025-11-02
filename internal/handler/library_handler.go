package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"school-management-backend/internal/model"
	"school-management-backend/internal/service"
)

type LibraryHandler struct {
	service *service.LibraryService
}

func NewLibraryHandler(service *service.LibraryService) *LibraryHandler {
	return &LibraryHandler{service: service}
}

// Book Handlers
func (h *LibraryHandler) CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AddNewBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *LibraryHandler) GetBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := h.service.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// Library Card Handlers
func (h *LibraryHandler) IssueLibraryCard(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	validForYears := 1 // Default validity
	if years, err := strconv.Atoi(c.DefaultQuery("validForYears", "1")); err == nil && years > 0 {
		validForYears = years
	}

	card, err := h.service.IssueLibraryCard(uint(userID), validForYears)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, card)
}

// Book Circulation Handlers
func (h *LibraryHandler) CheckoutBook(c *gin.Context) {
	var request struct {
		BookID uint `json:"bookId" binding:"required"`
		UserID uint `json:"userId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get staff ID from context (set by auth middleware)
	staffID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := h.service.CheckoutBook(request.BookID, request.UserID, staffID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *LibraryHandler) ReturnBook(c *gin.Context) {
	issueID, err := strconv.ParseUint(c.Param("issueId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid issue ID"})
		return
	}

	// Get staff ID from context (set by auth middleware)
	receivedBy, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := h.service.ReturnBook(uint(issueID), receivedBy.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// Fine Handlers
func (h *LibraryHandler) PayFine(c *gin.Context) {
	issueID, err := strconv.ParseUint(c.Param("issueId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid issue ID"})
		return
	}

	var payment model.FinePayment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payment.IssueID = uint(issueID)

	if err := h.service.RecordFinePayment(&payment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
