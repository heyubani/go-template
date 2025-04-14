package base

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	appErrors "github.com/heyubani/go-template/core/app-errors"
	"github.com/heyubani/go-template/interfaces"
)

const (
	// DefaultTimeout is the default timeout for http requests
	DefaultTimeout = 40 * time.Second
)

func InArray[T comparable](arr []T, pin T) bool {

	for i := 0; i < len(arr); i++ {

		if arr[i] == pin {
			return true
		}
	}

	return false
}

func IndexOf[T comparable](collection []T, el T) int {
	for i, x := range collection {
		if x == el {
			return i
		}
	}
	return -1
}

func ForEach[T comparable](arr []T, funcValue func(val T, i int)) {
	for i := 0; i < len(arr); i++ {
		funcValue(arr[i], i)
	}
}

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HttpGetRequest(url string, headers map[string]string, logger interfaces.ILogger, query url.Values) ([]byte, error) {
	return HttpRequest(url, http.MethodGet, nil, headers, logger, query)
}

func HttpPostRequest(url string, body []byte, headers map[string]string, logger interfaces.ILogger) ([]byte, error) {
	return HttpRequest(url, http.MethodPost, body, headers, logger, nil)
}

func HttpPatchRequest(url string, body []byte, headers map[string]string, logger interfaces.ILogger) ([]byte, error) {
	return HttpRequest(url, http.MethodPatch, body, headers, logger, nil)
}

func HttpPostRequestWithStatusCode(url string, body []byte, headers map[string]string, logger interfaces.ILogger) (int, []byte, error) {
	return HttpRequestWithStatusCode(url, http.MethodPost, body, headers, DefaultTimeout, nil, logger)
}

func HttpPutRequestWithStatusCode(url string, body []byte, headers map[string]string, logger interfaces.ILogger) (int, []byte, error) {
	return HttpRequestWithStatusCode(url, http.MethodPut, body, headers, DefaultTimeout, nil, logger)
}
func HttpPatchRequestWithStatusCode(url string, body []byte, headers map[string]string, logger interfaces.ILogger) (int, []byte, error) {
	return HttpRequestWithStatusCode(url, http.MethodPatch, body, headers, DefaultTimeout, nil, logger)
}

func HttpDeleteRequestWithStatusCode(url string, body []byte, headers map[string]string, logger interfaces.ILogger) (int, []byte, error) {
	return HttpRequestWithStatusCode(url, http.MethodDelete, body, headers, DefaultTimeout, nil, logger)
}

func HttpGetRequestWithStatusCode(url string, headers map[string]string, query url.Values, logger interfaces.ILogger) (int, []byte, error) {
	return HttpRequestWithStatusCode(url, http.MethodGet, nil, headers, DefaultTimeout, query, logger)
}

// HttpRequest will run a http request against a url returning the response bytes and a custom defined
// error that defines what happened.
// When the external service returns an error, then that error should be the response, and error string should be returned too
func HttpRequest(url, method string, body []byte, headers map[string]string, logger interfaces.ILogger, query url.Values) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))

	if err != nil {
		logger.Error("cannot initialize http request. Error: " + err.Error())
		return nil, errors.New("cannot initialize http request")
	}

	// by default add "Content-Type", "application/json"
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	for key, element := range headers {
		req.Header.Set(key, element)
	}

	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	res, err := (&http.Client{Timeout: DefaultTimeout}).Do(req)

	if err != nil {
		logger.Error("oops, error encountered while making an api call. Error: " + err.Error() + ", url is: " + url)
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		logger.Error("oops, error reading api response into buffer. Error: " + err.Error())
		return nil, err
	}

	// it's an error if error code >= 400
	if res.StatusCode >= http.StatusBadRequest {
		logger.Warnf("Error during API request. Url: %s,  Error: %s", url, string(resBody))
		return resBody, errors.New(fmt.Sprintf("API error with status code %d", res.StatusCode))
	}

	return resBody, nil
}

func BuildQueryParams(queryParams map[string]string) url.Values {
	if len(queryParams) == 0 {
		return nil
	}
	query := url.Values{}
	for key, value := range queryParams {
		query.Add(key, value)
	}
	return query
}

func HttpRequestWithStatusCode(url, method string, body []byte, headers map[string]string, timeout time.Duration, query url.Values, logger interfaces.ILogger) (int, []byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))

	if err != nil {
		logger.Errorf("cannot initialize http request. Error: " + err.Error())

		return http.StatusInternalServerError, nil, errors.New("cannot intialize http request")
	}

	// by default add "Content-Type", "application/json"
	req.Header.Set("Content-Type", "application/json")
	for key, element := range headers {
		req.Header.Set(key, element)
	}

	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	// LogRequestDetails(req, "outbound", nil)

	res, err := (&http.Client{Timeout: timeout}).Do(req)

	if err != nil {
		logger.Errorf("oops, error encountered while making an api call. Error: " + err.Error() + ", url is: " + url)

		// If the error is a timeout error, set the error message accordingly.
		if os.IsTimeout(err) {
			return http.StatusGatewayTimeout, nil, errors.New("slow network connection, please try again in a few minutes")
		}

		var statusCode int = http.StatusInternalServerError
		if res != nil {
			statusCode = res.StatusCode
		}

		return statusCode, nil, errors.New("error from external API call. Server might be unaiavailable")
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	if res.StatusCode == http.StatusUnauthorized {
		return http.StatusUnauthorized, nil, errors.New("invalid authentication")
	}

	if res.StatusCode >= http.StatusBadRequest && res.StatusCode < http.StatusInternalServerError {
		logger.Warnf("Bad Request during API request. Url: %s,  Error: %s", url, string(resBody))
		return res.StatusCode, resBody, fmt.Errorf("API error with status code %d", res.StatusCode)
	}

	if res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusOK && res.StatusCode != http.StatusAccepted {
		return res.StatusCode, resBody, errors.New("api error")
	}

	if err != nil {
		log.Println("oops, error reading api response into buffer. Error: " + err.Error())

		return http.StatusInternalServerError, nil, errors.New("internal server error. Cannot read response")
	}

	return res.StatusCode, resBody, nil
}

func DebunkError(resBody []byte, statusCode int, method string) interfaces.IAppError {
	var failedResponse HttpError

	err := json.Unmarshal(resBody, &failedResponse)
	if err != nil {
		log.Printf(fmt.Sprintf("%v::Unable to unmarshal response:", method), err.Error())
		return appErrors.InternalErrorMsg("Unable to unmarshal failed response")
	}
	// log the details here
	log.Println(fmt.Sprintf("%v::failed request while in %v method:", method, method), failedResponse)
	switch statusCode {
	case http.StatusBadRequest:
		return appErrors.BadRequestError(failedResponse.Message)
	case http.StatusUnauthorized:
		return appErrors.UnauthorizedError(failedResponse.Message)
	default:
		return appErrors.InternalErrorMsg(failedResponse.Message)
	}
}
