package course_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/JuD4Mo/go_api_web_domain/domain"
	"github.com/JuD4Mo/go_api_web_sdk/course"
	"github.com/ncostamagna/go_http_client/client"
)

var header http.Header
var sdk course.Transport

func TestMain(m *testing.M) {
	header = http.Header{}
	header.Set("Content-Type", "application/json")
	sdk = course.NewHttpClient("base-url", "")
	os.Exit(m.Run())
}

func TestGet_Response404Error(t *testing.T) {
	expectedErr := course.ErrNotFound{Message: "course '1' not found"}
	err := client.AddMockups(&client.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/courses/1",
		RespHTTPCode: 404,
		RespBody: fmt.Sprintf(`{
				"status" : 404,
				"message" : "%s"
		}`, expectedErr.Error()),
	})

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	course, err := sdk.Get("1")
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected nil, got %v", err)
	}

	if course != nil {
		t.Errorf("expected nil, got %v", course)
	}
}

func TestGet_Response500Error(t *testing.T) {
	expectedErr := errors.New("internal server error")

	err := client.AddMockups(&client.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/courses/1",
		RespHTTPCode: 500,
		RespBody: fmt.Sprintf(`{
				"status" : 500,
				"message" : "%s"
		}`, expectedErr.Error()),
	})

	course, err := sdk.Get("1")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("expected nil, got %v", err)
	}

	if course != nil {
		t.Errorf("expected nil, got %v", course)
	}
}

func TestGet_ResponseMarshallError(t *testing.T) {
	expectedErr := errors.New("unexpected end of JSON input")

	err := client.AddMockups(&client.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/courses/1",
		RespHTTPCode: 200,
		RespBody:     `{`,
	})

	course, err := sdk.Get("1")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if course != nil {
		t.Errorf("expected nil, got %v", course)
	}
}

func TestGet_ClientError(t *testing.T) {
	expectedErr := errors.New("client error")
	err := client.AddMockups(&client.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/courses/1",
		RespHTTPCode: 400,
		Err:          expectedErr,
	})

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	course, err := sdk.Get("1")
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if course != nil {
		t.Errorf("expected nil, got %v", course)
	}
}

func TestGet_ResponseSuccess(t *testing.T) {
	expectedCourse := domain.Course{
		ID:   "1",
		Name: "Course 1",
	}
	expectedCourseJson, err := json.Marshal(expectedCourse)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	err = client.AddMockups(&client.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/courses/1",
		RespHTTPCode: 200,
		RespBody: fmt.Sprintf(`{
							"status": 200,
							"message": "success",
							"data": %s
						}`, expectedCourseJson),
	})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	course, err := sdk.Get("1")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if course == nil {
		t.Errorf("expected course, got nil")
	}

	if course.ID != expectedCourse.ID {
		t.Errorf("expected id %v, got %v", expectedCourse.ID, course.ID)
	}

	if course.Name != expectedCourse.Name {
		t.Errorf("expected name %v, got %v", expectedCourse.Name, course.Name)
	}
}
