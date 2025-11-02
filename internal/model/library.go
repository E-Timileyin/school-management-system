package model

import (
	"time"
)

// BookCategory represents a category for books
type BookCategory struct {
    Base
    Name        string `gorm:"size:100;not null;uniqueIndex" json:"name"`
    Description string `gorm:"type:text" json:"description,omitempty"`
    IsActive    bool   `gorm:"default:true" json:"is_active"`
}

// Book represents a book in the library
type Book struct {
    Base
    ISBN            string       `gorm:"size:20;uniqueIndex;not null" json:"isbn"`
    Title           string       `gorm:"size:255;not null" json:"title"`
    Author          string       `gorm:"size:255;not null" json:"author"`
    Publisher       string       `gorm:"size:255" json:"publisher,omitempty"`
    PublicationYear int          `gorm:"type:smallint" json:"publication_year"`
    Edition         string       `gorm:"size:50" json:"edition,omitempty"`
    CategoryID      uint         `gorm:"not null" json:"category_id"`
    Price           float64      `gorm:"type:decimal(10,2)" json:"price"`
    Pages           int          `gorm:"default:0" json:"pages"`
    Description     string       `gorm:"type:text" json:"description,omitempty"`
    CoverImage      string       `gorm:"size:255" json:"cover_image,omitempty"`
    TotalCopies     int          `gorm:"default:1" json:"total_copies"`
    AvailableCopies int          `gorm:"default:1" json:"available_copies"`
    RackNumber      string       `gorm:"size:20" json:"rack_number,omitempty"`
    IsActive        bool         `gorm:"default:true" json:"is_active"`
    
    // Relationships
    Category *BookCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

// LibraryCard represents a library membership
type LibraryCard struct {
    Base
    UserID      uint       `gorm:"not null;uniqueIndex" json:"user_id"`
    CardNumber  string     `gorm:"size:50;unique;not null" json:"card_number"`
    IssueDate   time.Time  `gorm:"not null" json:"issue_date"`
    ExpiryDate  time.Time  `gorm:"not null" json:"expiry_date"`
    Status      string     `gorm:"type:varchar(20);default:'active'" json:"status"` // active, expired, blocked
    MaxBooks    int        `gorm:"default:3" json:"max_books"`
    FineAmount  float64    `gorm:"type:decimal(10,2);default:0" json:"fine_amount"`
    
    // Relationships
    User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// BookIssue represents a book checkout
type BookIssue struct {
    Base
    BookID        uint       `gorm:"not null" json:"book_id"`
    CardID        uint       `gorm:"not null" json:"card_id"`
    UserID        uint       `gorm:"not null" json:"user_id"` // For quick access
    IssueDate     time.Time  `gorm:"not null" json:"issue_date"`
    DueDate       time.Time  `gorm:"not null" json:"due_date"`
    ReturnDate    *time.Time `gorm:"index" json:"return_date,omitempty"`
    Status        string     `gorm:"type:varchar(20);default:'issued'" json:"status"` // issued, returned, overdue, lost
    FineAmount    float64    `gorm:"type:decimal(10,2);default:0" json:"fine_amount"`
    FinePaid      bool       `gorm:"default:false" json:"fine_paid"`
    IssuedBy      uint       `gorm:"not null" json:"issued_by"` // Staff ID who issued the book
    ReceivedBy    *uint      `gorm:"index" json:"received_by,omitempty"` // Staff ID who received the book
    
    // Relationships
    Book        *Book        `gorm:"foreignKey:BookID" json:"book,omitempty"`
    LibraryCard *LibraryCard `gorm:"foreignKey:CardID" json:"library_card,omitempty"`
    User        *User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
    Issuer      *User        `gorm:"foreignKey:IssuedBy" json:"issuer,omitempty"`
    Receiver    *User        `gorm:"foreignKey:ReceivedBy" json:"receiver,omitempty"`
}

// FinePayment represents fine payments
type FinePayment struct {
    Base
    IssueID     uint      `gorm:"not null" json:"issue_id"`
    Amount      float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
    PaymentDate time.Time `gorm:"not null" json:"payment_date"`
    ReceivedBy  uint      `gorm:"not null" json:"received_by"` // Staff ID who received the payment
    PaymentMode string    `gorm:"type:varchar(20);not null" json:"payment_mode"` // cash, card, online
    ReferenceNo string    `gorm:"size:100" json:"reference_no,omitempty"`
    
    // Relationships
    BookIssue *BookIssue `gorm:"foreignKey:IssueID" json:"book_issue,omitempty"`
    Receiver  *User      `gorm:"foreignKey:ReceivedBy" json:"receiver,omitempty"`
}
