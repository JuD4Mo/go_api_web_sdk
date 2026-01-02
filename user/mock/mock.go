package mock

import (
	"errors"

	"github.com/JuD4Mo/go_api_web_domain/domain"
)

type UserSdkMock struct {
	GetMock func(id string) (*domain.User, error)
}

func (m *UserSdkMock) Get(id string) (*domain.User, error) {
	if m.GetMock == nil {
		return nil, errors.New("Get mock is not set")
	}

	return m.GetMock(id)
}
