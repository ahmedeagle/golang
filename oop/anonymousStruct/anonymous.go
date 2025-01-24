package main

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Question model
type Question struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	UID          string         `gorm:"type:varchar(100);" json:"uid" form:"uid"`
	SurveyID     uint           `gorm:"not null" json:"-"`
	QuestionText string         `gorm:"type:text;not null" json:"questionText"`
	QuestionType string         `gorm:"size:50" json:"questionType"`
	Description  string         `gorm:"type:text;" json:"description"`
	Active       int            `gorm:"type:tinyint(1);default:0;index:questions_idx_active" json:"-"`
	Required     int            `gorm:"type:tinyint(1);default:0;" json:"required"`
	CreatedBy    int            `json:"-"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedBy    int            `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}

// ResponseDetail model
type ResponseDetail struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	UID              string         `gorm:"type:varchar(100);" json:"uid" form:"uid"`
	ResponseID       uint           `gorm:"not null" json:"responseId"`
	QuestionID       uint           `gorm:"not null" json:"questionId"`
	Answer           string         `gorm:"type:text" json:"answer"`
	CreatedAt        time.Time      `json:"createdAt"`
	DeletedAt        gorm.DeletedAt `json:"-"`
	PlainText        string         `gorm:"type:text" json:"plainText"`
	ActionCategoryID uint           `gorm:"foreignKey:QuestionActionCategoryID" json:"actionCategoryId"`
}

func main() {
	// Sample data for questions
	questions := []Question{
		{
			ID:           1,
			UID:          "Q001",
			SurveyID:     101,
			QuestionText: "What is your favorite color?",
			QuestionType: "Multiple Choice",
			Description:  "Select one of the colors listed.",
			Required:     1,
			CreatedBy:    1,
			CreatedAt:    time.Now(),
		},
		{
			ID:           2,
			UID:          "Q002",
			SurveyID:     101,
			QuestionText: "How often do you exercise?",
			QuestionType: "Text",
			Description:  "Describe your exercise routine.",
			Required:     0,
			CreatedBy:    1,
			CreatedAt:    time.Now(),
		},
	}

	// Sample data for response details
	responseDetails := []ResponseDetail{
		{
			ID:               1,
			UID:              "R001",
			ResponseID:       201,
			QuestionID:       1,
			Answer:           "Blue",
			CreatedAt:        time.Now(),
			PlainText:        "Blue",
			ActionCategoryID: 101,
		},
		{
			ID:               2,
			UID:              "R002",
			ResponseID:       201,
			QuestionID:       2,
			Answer:           "I exercise 3 times a week.",
			CreatedAt:        time.Now(),
			PlainText:        "I exercise 3 times a week.",
			ActionCategoryID: 102,
		},
	}

	// Anonymous struct for response
	response := struct {
		Questions []Question       `json:"questions"`
		Answers   []ResponseDetail `json:"answers"`
	}{
		Questions: questions,
		Answers:   responseDetails,
	}

	// Print the response
	fmt.Printf("%+v\n", response)
}
