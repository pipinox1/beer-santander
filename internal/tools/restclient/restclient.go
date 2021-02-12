package restclient

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type RestClient struct {
	HTTPClient *http.Client
}

type Header struct {
	Key   string
	Value string
}

func NewRestClient(time time.Duration) RestClient {
	return RestClient{
		HTTPClient: &http.Client{
			Timeout: time,
		},
	}
}

func (r *RestClient) DoGet(ctx context.Context, url string, response interface{}, additionalHeaders ...Header) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	for _, header := range additionalHeaders {
		req.Header.Add(header.Key, header.Value)
	}

	res, err := r.HTTPClient.Do(req)
	if err != nil {
		return errors.New("asdca")
	}

	//Nos Aseguramos que cierre el body
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		//return errors.NewRestError("error_reading_body", http.StatusInternalServerError)
		return errors.New("asdca")
	}
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		//return errors.NewRestError("rest_client_error", res.StatusCode)
		return errors.New("asdca")
	}
	return nil
}
