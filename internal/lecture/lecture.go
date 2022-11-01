package lecture

import (
	"context"
	"errors"
	"golang-united-lectures/internal/api"
	"golang-united-lectures/internal/models"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Lecture struct {
	api.UnimplementedLectureServer
}

func (l *Lecture) Create(ctx context.Context, request *api.CreateRequest) (*api.CreateResponse, error) {

	lecture := models.NewLecture()

	lecture.CourseId = request.GetCourseId()
	lecture.Number = request.GetNumber()
	lecture.Title = request.GetTitle()
	lecture.Description = request.GetDescription()
	lecture.CreatedBy = request.GetCreatedBy()

	err := lecture.Save()
	if err != nil {
		return nil, err
	}

	return &api.CreateResponse{Id: lecture.Id}, nil

}

func (l *Lecture) Get(ctx context.Context, request *api.GetRequest) (*api.GetResponse, error) {

	lecture, err := models.GetLecture(request.GetId())
	if err != nil {
		return nil, err
	}

	response := api.GetResponse{}

	response.Id = lecture.Id
	response.CourseId = lecture.CourseId
	response.Number = lecture.Number
	response.Title = lecture.Title
	response.Description = lecture.Description

	if !lecture.CreatedAt.IsZero() {
		response.CreatedAt = timestamppb.New(lecture.CreatedAt)
		response.CreatedBy = lecture.CreatedBy
	}

	if !lecture.ChangedAt.IsZero() {
		response.ChangedAt = timestamppb.New(lecture.ChangedAt)
		response.ChangedBy = lecture.ChangedBy
	}

	if !lecture.DeletedAt.IsZero() {
		response.DeletedAt = timestamppb.New(lecture.DeletedAt)
		response.DeletedBy = lecture.DeletedBy
	}

	return &response, nil

}

func (l *Lecture) Update(ctx context.Context, request *api.UpdateRequest) (*emptypb.Empty, error) {

	lecture, err := models.GetLecture(request.GetId())
	if err != nil {
		return nil, err
	}

	if !lecture.DeletedAt.IsZero() {
		err = errors.New("lecture is deleted")
		return nil, err
	}

	lecture.Number = request.GetNumber()
	lecture.Title = request.GetTitle()
	lecture.Description = request.GetDescription()
	lecture.ChangedAt = time.Now()
	lecture.ChangedBy = request.GetChangedBy()

	err = lecture.Save()
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil

}

func (l *Lecture) Delete(ctx context.Context, request *api.DeleteRequest) (*emptypb.Empty, error) {

	lecture, err := models.GetLecture(request.GetId())
	if err != nil {
		return nil, err
	}

	if !lecture.DeletedAt.IsZero() {
		err = errors.New("lecture is deleted")
		return nil, err
	}

	lecture.DeletedAt = time.Now()
	lecture.DeletedBy = request.GetDeletedBy()

	err = lecture.Save()
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil

}

func (l *Lecture) GetList(ctx context.Context, request *api.GetListRequest) (*api.GetListResponse, error) {

	lectures, err := models.GetLectureList(request.GetCourseId())
	if err != nil {
		return nil, err
	}

	response := &api.GetListResponse{}

	for _, lecture := range lectures {

		if !lecture.DeletedAt.IsZero() {
			continue
		}

		lercureResponse := api.GetResponse{}

		lercureResponse.Id = lecture.Id
		lercureResponse.CourseId = lecture.CourseId
		lercureResponse.Number = lecture.Number
		lercureResponse.Title = lecture.Title
		lercureResponse.Description = lecture.Description

		if !lecture.CreatedAt.IsZero() {
			lercureResponse.CreatedAt = timestamppb.New(lecture.CreatedAt)
			lercureResponse.CreatedBy = lecture.CreatedBy
		}

		if !lecture.ChangedAt.IsZero() {
			lercureResponse.ChangedAt = timestamppb.New(lecture.ChangedAt)
			lercureResponse.ChangedBy = lecture.ChangedBy
		}

		if !lecture.DeletedAt.IsZero() {
			lercureResponse.DeletedAt = timestamppb.New(lecture.DeletedAt)
			lercureResponse.DeletedBy = lecture.DeletedBy
		}

		response.Lectures = append(response.Lectures, &lercureResponse)

	}

	return response, nil

}
