package course

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/JuD4Mo/go_api_web_domain/domain"
	"github.com/ncostamagna/go_http_client/client"
)

type (
	DataResponse struct {
		Message string      `json:"message"`
		Code    int         `json:"code"`
		Data    interface{} `json:"data"`
		Meta    interface{} `json:"meta"`
	}

	Transport interface {
		Get(id string) (*domain.Course, error)
	}

	clientHttp struct {
		client client.Transport
	}
)

func NewHttpClient(baseURL, token string) Transport {
	header := http.Header{}

	if token != "" {
		header.Set("Authorization", token)
	}

	return &clientHttp{
		client: client.New(header, baseURL, 5000*time.Millisecond, true),
	}
}

func (c clientHttp) Get(id string) (*domain.Course, error) {
	dataResponse := DataResponse{Data: &domain.Course{}}

	u := url.URL{}

	u.Path += fmt.Sprintf("/courses/%s", id)

	reps := c.client.Get(u.String())

	if reps.Err != nil {
		return nil, reps.Err
	}

	if reps.StatusCode > 404 {
		return nil, ErrNotFound{fmt.Sprintf("%s", reps)}
	}

	if reps.StatusCode > 299 {
		return nil, fmt.Errorf("%s", reps)
	}

	err := reps.FillUp(&dataResponse)
	if err != nil {
		return nil, err
	}

	return dataResponse.Data.(*domain.Course), nil
}
