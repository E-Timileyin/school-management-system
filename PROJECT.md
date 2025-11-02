
---

# 🧩 PROJECT: **Student Management System (Backend Core)**

### 🎯 Goal:

A backend-style Go project (CLI or REST-ready) that teaches:

* Data structures (structs, maps, slices)
* Error handling
* Functions, methods, interfaces
* File I/O and JSON (for backup)
* Database connection (Postgres/MySQL/SQLite)
* CRUD operations
* Concurrency basics (auto backup or notifications)

---

## ⚙️ **PHASE 1 — CORE STRUCTURE**

### **🗂 Folder Layout (clean & scalable)**

```
student-mgmt/
│
├── main.go                # Entry point
├── go.mod
│
├── config/                # DB connection + app configs
│   └── config.go
│
├── models/                # Structs & DB models
│   └── student.go
│
├── repository/            # Database operations
│   └── student_repo.go
│
├── services/              # Business logic
│   └── student_service.go
│
├── handlers/              # CLI or HTTP handlers (optional)
│   └── student_handler.go
│
├── utils/                 # Helper functions (logging, errors, etc.)
│   └── logger.go
│
└── data/                  # Backup files (JSON)
    └── students_backup.json
```

---

## 🧱 **CORE FEATURES TO IMPLEMENT**

### **1. Student CRUD**

* ✅ Create new student
* ✅ View all students
* ✅ Update student info (age, department, etc.)
* ✅ Delete student

**Database fields:**

| Field      | Type      | Description                |
| ---------- | --------- | -------------------------- |
| id         | int       | Auto increment primary key |
| name       | string    | Student’s full name        |
| email      | string    | Unique email               |
| age        | int       | Student’s age              |
| department | string    | Department name            |
| created_at | timestamp | Creation time              |
| updated_at | timestamp | Update time                |

---

### **2. Course Management**

* ✅ Create course
* ✅ Assign students to courses (many-to-many relationship)
* ✅ View students per course
* ✅ Remove student from course

**Tables:**

* `courses`: id, name, unit
* `student_courses`: student_id, course_id (junction table)

---

### **3. Grade Tracking**

* ✅ Add grades for a student in a course
* ✅ Calculate GPA
* ✅ View student transcript

**Table:**

* `grades`: id, student_id, course_id, score, grade_letter

---

### **4. Search & Filtering**

* ✅ Search student by name, email, or department
* ✅ Filter students by department or grade level

**Concepts used:**

* Query params or filter functions
* SQL `LIKE` or `WHERE` conditions
* Struct filtering if in-memory

---

### **5. Backup & Restore**

* ✅ Export all students to JSON file (`data/students_backup.json`)
* ✅ Import from backup JSON
* ✅ Auto-backup feature using GoRoutines

**Concepts used:**

* `encoding/json`
* File I/O (`os.WriteFile`, `os.ReadFile`)
* Goroutines (`go func(){}`)

---

### **6. Logging & Error Handling**

* ✅ Log every create/update/delete action
* ✅ Centralized error handling in `utils/logger.go`
* ✅ Custom error struct for business logic

**Concepts used:**

* Go’s `log` package or custom logger
* Struct-based error return (like `type AppError struct`)

---

### **7. DB Connection Layer (config/)**

* ✅ Load DB config from `.env`
* ✅ Use `database/sql` or `gorm.io/gorm` for ORM
* ✅ Central connection management

**Example:**

```go
package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	DB = db
}
```

---

### **8. Business Logic (services/)**

* Validate requests (no duplicate emails)
* Calculate GPA logic
* Handle student-course relationships

---

### **9. Concurrency (optional but powerful)**

* ✅ Auto backup data to file every 30s using Goroutine
* ✅ Use a channel to gracefully stop background jobs on shutdown

---

### **10. CLI or API Interface**

You can pick one of two interfaces:

#### **Option A — CLI (for Fundamentals)**

* Use `fmt.Scanln` to interact
* Run options: add student, view student, backup, exit

#### **Option B — API (when you start Gin)**

* Expose routes: `/students`, `/courses`, `/grades`
* Migrate handlers → Gin later with minimal refactor

---

## 🧠 **WHAT YOU’LL MASTER FROM THIS SINGLE PROJECT**

| Concept              | How You’ll Use It                         |
| -------------------- | ----------------------------------------- |
| Variables, Types     | Input handling and struct fields          |
| Functions            | CRUD logic, utilities                     |
| Structs              | Student, Course, Grade models             |
| Maps/Slices          | In-memory caches                          |
| Methods              | Attach logic to structs                   |
| Interfaces           | Abstract repository/service               |
| Pointers             | Pass and update structs                   |
| Error Handling       | Return `error` gracefully                 |
| JSON                 | Backup/export system                      |
| File I/O             | Read/write backups                        |
| Concurrency          | Background tasks (auto-save)              |
| Database             | GORM or SQL-level CRUD                    |
| Project Architecture | Config → Model → Repo → Service → Handler |

---

## 🧭 LEARNING STRATEGY (6-Week Build Plan)

| Week | Focus                                | Deliverable                           |
| ---- | ------------------------------------ | ------------------------------------- |
| 1    | Basic Go syntax, structs, loops, I/O | CLI scaffold (menu + input)           |
| 2    | Functions, slices, maps              | CRUD in-memory                        |
| 3    | GORM + DB models                     | Persist data to Postgres              |
| 4    | Service + Repo layer                 | Clean separation of logic             |
| 5    | File handling + backup               | Auto-save feature                     |
| 6    | Goroutines + polish                  | Background jobs, error handling, docs |

---