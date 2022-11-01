package models

import (
	"golang-united-lectures/internal/database"
	"time"

	"github.com/google/uuid"
)

type Lecture struct {
	Id          string `gorm:"primarykey"`
	CourseId    string `gorm:"index"`
	Number      uint32
	Title       string
	Description string
	CreatedAt   time.Time
	CreatedBy   string
	ChangedAt   time.Time
	ChangedBy   string
	DeletedAt   time.Time
	DeletedBy   string
}

func (l *Lecture) Save() error {

	db := database.GetInstance()

	var exists bool

	err := db.Client.Model(l).Select("count(*) > 0").Where("id = ?", l.Id).Find(&exists).Error
	if err != nil {
		return err
	}

	if exists {
		err = db.Client.Updates(l).Error
	} else {
		err = db.Client.Create(l).Error
	}

	if err != nil {
		return err
	}

	return nil

}

func NewLecture() *Lecture {

	lecture := &Lecture{}

	lecture.Id = uuid.New().String()
	lecture.CreatedAt = time.Now()

	return lecture

}

func GetLecture(id string) (*Lecture, error) {

	db := database.GetInstance()

	lecture := Lecture{}

	err := db.Client.First(&lecture, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &lecture, nil

}

func GetLectureList(CourseId string) ([]Lecture, error) {

	db := database.GetInstance()

	lectures := []Lecture{}

	var err error

	if CourseId == "" {
		err = db.Client.Find(&lectures).Error
	} else {
		err = db.Client.Where("course_id = ?", CourseId).Find(&lectures).Error
	}
	if err != nil {
		return nil, err
	}

	return lectures, nil

}
