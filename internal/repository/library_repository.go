package repository

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"school-management-backend/internal/model"
)

type LibraryRepository struct {
	db *gorm.DB
}

func NewLibraryRepository(db *gorm.DB) *LibraryRepository {
	return &LibraryRepository{db: db}
}

// Book Methods
func (r *LibraryRepository) CreateBook(book *model.Book) error {
	return r.db.Create(book).Error
}

func (r *LibraryRepository) GetBookByID(id uint) (*model.Book, error) {
	var book model.Book
	err := r.db.Preload("Category").First(&book, id).Error
	return &book, err
}

func (r *LibraryRepository) UpdateBook(book *model.Book) error {
	return r.db.Save(book).Error
}

func (r *LibraryRepository) DeleteBook(id uint) error {
	return r.db.Delete(&model.Book{}, id).Error
}

// Library Card Methods
func (r *LibraryRepository) CreateLibraryCard(card *model.LibraryCard) error {
	return r.db.Create(card).Error
}

func (r *LibraryRepository) GetLibraryCardByUserID(userID uint) (*model.LibraryCard, error) {
	var card model.LibraryCard
	err := r.db.Where("user_id = ?", userID).First(&card).Error
	return &card, err
}

// Book Issue Methods
func (r *LibraryRepository) CheckoutBook(issue *model.BookIssue) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check book availability
	var book model.Book
	if err := tx.First(&book, issue.BookID).Error; err != nil {
		tx.Rollback()
		return err
	}

	if book.AvailableCopies <= 0 {
		tx.Rollback()
		return errors.New("no copies available")
	}

	// Update book available copies
	book.AvailableCopies--
	if err := tx.Save(&book).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Create issue record
	issue.IssueDate = time.Now()
	issue.DueDate = time.Now().AddDate(0, 0, 14) // 14 days from now
	issue.Status = "issued"

	if err := tx.Create(issue).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *LibraryRepository) ReturnBook(issueID, receivedBy uint) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get the issue record
	var issue model.BookIssue
	if err := tx.Preload("Book").First(&issue, issueID).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update book available copies
	var book model.Book
	if err := tx.First(&book, issue.BookID).Error; err != nil {
		tx.Rollback()
		return err
	}

	book.AvailableCopies++
	if err := tx.Save(&book).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update issue record
	now := time.Now()
	issue.ReturnDate = &now
	issue.Status = "returned"
	receivedByUint := receivedBy
	receivedByUint = receivedBy
	issue.ReceivedBy = &receivedByUint

	// Calculate fine if any
	if now.After(issue.DueDate) {
		daysOverdue := int(now.Sub(issue.DueDate).Hours() / 24)
		issue.FineAmount = float64(daysOverdue) * 5.0 // $5 per day fine
	}

	if err := tx.Save(&issue).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Fine Payment Methods
func (r *LibraryRepository) RecordFinePayment(payment *model.FinePayment) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Update the issue record
	if err := tx.Model(&model.BookIssue{}).
		Where("id = ?", payment.IssueID).
		UpdateColumns(map[string]interface{}{
			"fine_paid":   true,
			"fine_amount": 0, // Reset fine amount after payment
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Record the payment
	payment.PaymentDate = time.Now()
	if err := tx.Create(payment).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Query Methods
func (r *LibraryRepository) GetBooksByCategory(categoryID uint) ([]model.Book, error) {
	var books []model.Book
	err := r.db.Where("category_id = ?", categoryID).Find(&books).Error
	return books, err
}

func (r *LibraryRepository) GetOverdueBooks() ([]model.BookIssue, error) {
	var issues []model.BookIssue
	err := r.db.Where("status = 'issued' AND due_date < ?", time.Now()).
		Preload("Book").
		Preload("User").
		Find(&issues).Error
	return issues, err
}

func (r *LibraryRepository) GetBorrowingHistory(userID uint) ([]model.BookIssue, error) {
	var issues []model.BookIssue
	err := r.db.Where("user_id = ?", userID).
		Preload("Book").
		Order("issue_date DESC").
		Find(&issues).Error
	return issues, err
}
