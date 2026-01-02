package mock

import (
	"errors"

	"github.com/JuD4Mo/go_api_web_domain/domain"
)

type CourseSdkMock struct {
	GetMock func(id string) (*domain.Course, error)
}

func (m *CourseSdkMock) Get(id string) (*domain.Course, error) {
	if m.GetMock == nil {
		return nil, errors.New("Get mock is not set")
	}

	return m.GetMock(id)
}
