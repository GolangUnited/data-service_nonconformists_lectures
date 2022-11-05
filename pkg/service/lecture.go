package service

import (
	"context"
	"golang-united-lectures/pkg/api"
	"golang-united-lectures/pkg/repositories"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Lecture struct {
	api.UnimplementedLectureServer
}

func (l *Lecture) Create(ctx context.Context, request *api.CreateRequest) (*api.CreateResponse, error) {

	lecture := &repositories.Lecture{
		Id:          uuid.New().String(),
		CourseId:    request.GetCourseId(),
		Number:      request.GetNumber(),
		Title:       request.GetTitle(),
		Description: request.GetDescription(),
		CreatedBy:   request.GetCreatedBy(),
		UpdatedBy:   request.GetCreatedBy(),
	}

	err := repositories.CreateLecture(lecture)
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	return &api.CreateResponse{Id: lecture.Id}, nil

}

func (l *Lecture) Get(ctx context.Context, request *api.GetRequest) (*api.GetResponse, error) {

	lecture, err := repositories.GetLecture(request.GetId())
	if err != nil {
		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	response := &api.GetResponse{
		Id:          lecture.Id,
		CourseId:    lecture.CourseId,
		Number:      lecture.Number,
		Title:       lecture.Title,
		Description: lecture.Description,
		CreatedBy:   lecture.CreatedBy,
		UpdatedBy:   lecture.UpdatedBy,
		DeletedBy:   lecture.DeletedBy,
		CreatedAt:   timestamppb.New(lecture.CreatedAt),
		UpdatedAt:   timestamppb.New(lecture.UpdatedAt),
	}

	if !lecture.DeletedAt.IsZero() {
		response.DeletedAt = timestamppb.New(lecture.DeletedAt)
	}

	return response, nil

}

func (l *Lecture) Update(ctx context.Context, request *api.UpdateRequest) (*emptypb.Empty, error) {

	lecture, err := repositories.GetLecture(request.GetId())
	if err != nil {
		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	if !lecture.DeletedAt.IsZero() {
		return nil, status.New(codes.Aborted, "lecture is deleted").Err()
	}

	lecture.Number = request.GetNumber()
	lecture.Title = request.GetTitle()
	lecture.Description = request.GetDescription()
	lecture.UpdatedBy = request.GetUpdatedBy()

	err = repositories.UpdateLecture(lecture)
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	return &emptypb.Empty{}, nil

}

func (l *Lecture) Delete(ctx context.Context, request *api.DeleteRequest) (*emptypb.Empty, error) {

	lecture, err := repositories.GetLecture(request.GetId())
	if err != nil {
		return nil, status.New(codes.NotFound, err.Error()).Err()
	}

	if !lecture.DeletedAt.IsZero() {
		return nil, status.New(codes.Aborted, "lecture is deleted").Err()
	}

	lecture.DeletedAt = time.Now()
	lecture.DeletedBy = request.GetDeletedBy()

	err = repositories.UpdateLecture(lecture)
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	return &emptypb.Empty{}, nil

}

func (l *Lecture) GetList(ctx context.Context, request *api.ListRequest) (*api.ListResponse, error) {

	lectures, err := repositories.GetLectureList(request.GetCourseId(), request.GetShowDeleted(), request.GetLimit(), request.GetOffset())
	if err != nil {
		return nil, status.New(codes.Internal, err.Error()).Err()
	}

	response := &api.ListResponse{}

	response.Lectures = make([]*api.GetResponse, 0, len(*lectures))

	for _, lecture := range *lectures {

		lercureResponse := &api.GetResponse{
			Id:          lecture.Id,
			CourseId:    lecture.CourseId,
			Number:      lecture.Number,
			Title:       lecture.Title,
			Description: lecture.Description,
			CreatedAt:   timestamppb.New(lecture.CreatedAt),
			CreatedBy:   lecture.CreatedBy,
			UpdatedAt:   timestamppb.New(lecture.UpdatedAt),
			UpdatedBy:   lecture.UpdatedBy,
			DeletedBy:   lecture.DeletedBy,
		}

		if !lecture.DeletedAt.IsZero() {
			lercureResponse.DeletedAt = timestamppb.New(lecture.DeletedAt)
		}

		response.Lectures = append(response.Lectures, lercureResponse)

	}

	return response, nil

}
