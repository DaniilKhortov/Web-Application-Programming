package models

import (
	"database/sql"
	"html/template"
	"log"
)

type StringContent struct {
	Content string
}

type LinkContent struct {
	ID      int
	Link    string
	Content string
}

type QAContent struct {
	ID       int
	Question string
	Answer   string
}
type OptionContent struct {
	ID      int
	Content string
}

type StringContentHTML struct {
	Content template.HTML
}

type Lesson struct {
	LessonID              int
	LessonTitle           string
	LessonAim             string
	LessonIntro           []StringContent
	LessonTheoContent     []StringContent
	LessonLinksContent    []LinkContent
	LessonMaterialContent []LinkContent
	LessonQA              []QAContent
	LessonTask1           string
	LessonTask2           string
	LessonTask3           string
	LessonReq1            string
	LessonReq2            string
	LessonReq3            string
	LessonImage1          string
	LessonImage2          string
	LessonImage3          string
	LessonArch1           string
	LessonProductGithub   string
	LessonProductIdea     string
	LessonProductStruct   string
	LessonProductFeatures string
	LessonProductInteract string
	LessonProductSummary  string
	LessonHomework        []StringContent
	CSRFToken             string
}

// Структура для обліку користувачів
type UserData struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Error    string
	Role     string `json:"role"`
}

// Структура для обліку файлів
type FileData struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Lesson   string `gorm:"not null"`
	FilePath string `gorm:"not null"`
	Date     string
	Mark     string `json:"mark"`
}

// Функція внесення користувача до БД
// Можливо буде усунута...
func AddUser(db *sql.DB, username, email, password string) (int64, error) {
	query := "INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)"
	role := "user"
	res, err := db.Exec(query, username, email, password, role)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// Функція отримання усіх користувачів  БД (Оптимізована)
// Можливо буде усунута...
func GetAllUsers(db *sql.DB) ([]UserData, error) {
	query := "SELECT id, username, email, password FROM users"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Помилка при отриані даних користувачів з БД:", err)
		return nil, err
	}
	defer rows.Close()

	var users []UserData
	for rows.Next() {
		var u UserData
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Password); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
