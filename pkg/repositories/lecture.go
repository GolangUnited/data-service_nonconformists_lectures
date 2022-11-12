package repositories

import (
	"errors"
	"golang-united-lectures/pkg/database"
	"log"
	"time"
)

type Lecture struct {
	Id          string `gorm:"type:uuid;primarykey;default:gen_random_uuid()"`
	CourseId    string `gorm:"index"`
	Number      uint64
	Title       string
	Description string
	CreatedAt   time.Time
	CreatedBy   string
	UpdatedAt   time.Time
	UpdatedBy   string
	DeletedAt   time.Time
	DeletedBy   string
}

func CreateLecture(lecture *Lecture) error {

	err := database.DB.Create(lecture).Error
	if err != nil {
		log.Printf("Error on create lecture: %s", err.Error())
		return errors.New("DB error")
	}

	return nil

}

func UpdateLecture(lecture *Lecture) error {

	err := database.DB.Updates(lecture).Error
	if err != nil {
		log.Printf("Error on update lecture: %s", err.Error())
		return errors.New("DB error")
	}

	return nil

}

func DeleteLecture(lecture *Lecture) error {

	err := database.DB.Delete(lecture).Error
	if err != nil {
		log.Printf("Error on delete lecture: %s", err.Error())
		return errors.New("DB error")
	}

	return nil

}

func GetLecture(id string) (*Lecture, error) {

	lecture := Lecture{}

	err := database.DB.First(&lecture, "id = ?", id).Error
	if err != nil {
		log.Printf("Error on select lecture Id = %s: %s", id, err.Error())
		return nil, errors.New("DB error")
	}

	return &lecture, nil

}

func GetLectureList(courseId string, showDeleted bool, limit uint32, offset uint32) (*[]Lecture, error) {

	lectures := []Lecture{}

	query := database.DB.Model(&Lecture{})

	if limit > 0 {
		query.Limit(int(limit))
	}

	if offset > 0 {
		query.Offset(int(offset))
	}

	if courseId != "" {
		query.Where("course_id = ?", courseId)
	}

	if !showDeleted {
		query.Where("deleted_at = '0001-01-01 00:00:00+00'")
	}

	query.Order("created_at asc")

	err := query.Find(&lectures).Error
	if err != nil {
		log.Printf("Error on get list of lectures: %s", err.Error())
		return nil, errors.New("DB error")
	}

	return &lectures, nil

}
