package user_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/JuD4Mo/go_api_web_domain/domain"
	"github.com/JuD4Mo/go_api_web_sdk/user"
	"github.com/ncostamagna/go_http_client/client"
)

var header http.Header
var sdk user.Transport

func TestMain(m *testing.M) {
	header = http.Header{}
	header.Set("Content-Type", "application/json")
	sdk = user.NewHttpClient("base-url", "")
	os.Exit(m.Run())
}

func TestGet_Response404Error(t *testing.T) {
	expectedErr := user.ErrNotFound{Message: "user '1' not found"}

	err := client.AddMockups(&client.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/users/1",
		RespHTTPCode: 404,
		RespBody: fmt.Sprintf(`{
							"status": 404,
							"message": "%s"
						}`, expectedErr.Error()),
	})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	user, err := sdk.Get("1")
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if user != nil {
		t.Errorf("expected nil, got %v", user)
	}
}

func TestGet_Response500Error(t *testing.T) {
	expectedErr := errors.New("internal server error")

	err := client.AddMockups(&client.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/users/1",
		RespHTTPCode: 500,
		RespBody: fmt.Sprintf(`{
							"status": 500,
							"message": "%s"
						}`, expectedErr.Error()),
	})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	user, err := sdk.Get("1")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if user != nil {
		t.Errorf("expected nil, got %v", user)
	}
}

func TestGet_ResponseMarshalError(t *testing.T) {
	expectedErr := errors.New("unexpected end of JSON input")

	err := client.AddMockups(&client.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/users/1",
		RespHTTPCode: 200,
		RespBody:     `{`,
	})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	user, err := sdk.Get("1")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if user != nil {
		t.Errorf("expected nil, got %v", user)
	}
}

func TestGet_ClientError(t *testing.T) {
	expectedErr := errors.New("client error")

	err := client.AddMockups(&client.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/users/1",
		RespHTTPCode: 400,
		Err:          expectedErr,
	})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	user, err := sdk.Get("1")
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if user != nil {
		t.Errorf("expected nil, got %v", user)
	}
}

func TestGet_ResponseSuccess(t *testing.T) {
	expectedUser := &domain.User{
		ID:        "1",
		FirstName: "Nahuel",
		LastName:  "Costamagna",
		Email:     "nahuel@nahuel.com",
	}
	expectedUserJson, err := json.Marshal(expectedUser)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	err = client.AddMockups(&client.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/users/1",
		RespHTTPCode: 200,
		RespBody: fmt.Sprintf(`{
							"status": 200,
							"message": "success",
							"data": %s
						}`, expectedUserJson),
	})
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	user, err := sdk.Get("1")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if user == nil {
		t.Errorf("expected user, got nil")
	}

	if user.ID != expectedUser.ID {
		t.Errorf("expected id %v, got %v", expectedUser.ID, user.ID)
	}

	if user.FirstName != expectedUser.FirstName {
		t.Errorf("expected first name %v, got %v", expectedUser.FirstName, user.FirstName)
	}

	if user.LastName != expectedUser.LastName {
		t.Errorf("expected last name %v, got %v", expectedUser.LastName, user.LastName)
	}

	if user.Email != expectedUser.Email {
		t.Errorf("expected email %v, got %v", expectedUser.Email, user.Email)
	}
}
