package mock

import (
	"testing"

	"github.com/JuD4Mo/go_api_web_sdk/course"
)

func TestMock_Course(t *testing.T) {
	t.Run("test course mock", func(t *testing.T) {
		var _ course.Transport = (*CourseSdkMock)(nil)
	})
}
