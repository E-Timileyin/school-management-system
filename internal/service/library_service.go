package service

import (
	"errors"
	"time"

	"github.com/E-Timileyin/school-management-system/internal/model"
	"github.com/E-Timileyin/school-management-system/internal/repository"
)

type LibraryService struct {
	repo *repository.LibraryRepository
}

func NewLibraryService(repo *repository.LibraryRepository) *LibraryService {
	return &LibraryService{repo: repo}
}

// Book Management
func (s *LibraryService) GetBookByID(id uint) (*model.Book, error) {
	return s.repo.GetBookByID(id)
}

func (s *LibraryService) AddNewBook(book *model.Book) error {
	// Set default values
	if book.TotalCopies == 0 {
		book.TotalCopies = 1
	}
	book.AvailableCopies = book.TotalCopies
	book.IsActive = true

	return s.repo.CreateBook(book)
}

func (s *LibraryService) UpdateBookDetails(book *model.Book) error {
	// Prevent updating available copies directly
	existingBook, err := s.repo.GetBookByID(book.ID)
	if err != nil {
		return err
	}

	// Calculate new available copies based on total copies change
	if book.TotalCopies != existingBook.TotalCopies {
		diff := book.TotalCopies - existingBook.TotalCopies
		book.AvailableCopies = existingBook.AvailableCopies + diff
		if book.AvailableCopies < 0 {
			book.AvailableCopies = 0
		}
	}

	return s.repo.UpdateBook(book)
}

// Library Card Management
func (s *LibraryService) IssueLibraryCard(userID uint, validForYears int) (*model.LibraryCard, error) {
	// Check if user already has an active card
	existingCard, err := s.repo.GetLibraryCardByUserID(userID)
	if err == nil && existingCard.Status == "active" {
		return nil, errors.New("user already has an active library card")
	}

	card := &model.LibraryCard{
		UserID:     userID,
		CardNumber: generateLibraryCardNumber(), // Implement this function
		IssueDate:  time.Now(),
		ExpiryDate: time.Now().AddDate(validForYears, 0, 0),
		Status:     "active",
		MaxBooks:   5, // Default value
	}

	if err := s.repo.CreateLibraryCard(card); err != nil {
		return nil, err
	}

	return card, nil
}

// Book Circulation
func (s *LibraryService) CheckoutBook(bookID, userID, staffID uint) error {
	// Get user's library card
	card, err := s.repo.GetLibraryCardByUserID(userID)
	if err != nil {
		return errors.New("no active library card found")
	}

	// Check if user has reached max books limit
	// (Implementation depends on your repository methods)

	// Create book issue record
	issue := &model.BookIssue{
		BookID:    bookID,
		CardID:    card.ID,
		UserID:    userID,
		IssuedBy:  staffID,
		Status:    "issued",
		IssueDate: time.Now(),
		DueDate:   time.Now().AddDate(0, 0, 14), // 14 days from now
	}

	return s.repo.CheckoutBook(issue)
}

func (s *LibraryService) ReturnBook(issueID, receivedBy uint) error {
	return s.repo.ReturnBook(issueID, receivedBy)
}

// Fine Management
func (s *LibraryService) CalculateFine(issueID uint) (float64, error) {
	// Get the issue record
	// Calculate fine based on due date and current date
	// Return the fine amount
	return 0, nil // Implement this
}

func (s *LibraryService) RecordFinePayment(payment *model.FinePayment) error {
	// Validate payment amount
	// Record the payment
	return s.repo.RecordFinePayment(payment)
}

// Helper function to generate library card number
func generateLibraryCardNumber() string {
	// Implement your card number generation logic
	// Example: LIB-{timestamp}-{random}
	return "LIB-12345678"
}
