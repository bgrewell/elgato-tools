package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bgrewell/elgato-tools/internal/types"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	scheme              = "http"
	lightsApiEndpoint   = "elgato/lights"
	infoApiEndpoint     = "elgato/accessory-info"
	settingsApiEndpoint = "elgato/settings"
)

func GetLights(entry *types.ServiceEntry) (lights *types.Lights, err error) {
	url := fmt.Sprintf("%s://%s:%d/%s", scheme, entry.AddrV4, entry.Port, lightsApiEndpoint)
	resp, err := sendGet(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request returned status %d\n", resp.StatusCode)
	}

	// Read the response body
	lights = &types.Lights{}
	err = json.Unmarshal(resp.Body, lights)
	if err != nil {
		return nil, err
	}
	return lights, nil
}

func GetInfo(entry *types.ServiceEntry) (info *types.AccessoryInfo, err error) {
	url := fmt.Sprintf("%s://%s:%d/%s", scheme, entry.AddrV4, entry.Port, infoApiEndpoint)
	resp, err := sendGet(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request returned status %d\n", resp.StatusCode)
	}

	// Read the response body
	info = &types.AccessoryInfo{}
	err = json.Unmarshal(resp.Body, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func PutLights(entry *types.ServiceEntry, lights *types.Lights) (updatedLights *types.Lights, err error) {
	url := fmt.Sprintf("%s://%s:%d/%s", scheme, entry.AddrV4, entry.Port, lightsApiEndpoint)
	if err != nil {
		return nil, err
	}
	resp, err := sendPut(url, lights)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request returned status %d\n", resp.StatusCode)
	}

	// Read the response body
	lights = &types.Lights{}
	err = json.Unmarshal(resp.Body, lights)
	if err != nil {
		return nil, err
	}
	return lights, nil
}

//func (ea *ElgatoApi) GetSettings(entry *types.ServiceEntry) (??, err error) {
//	url := fmt.Sprintf("%s://%s:%d/%s", scheme, entry.AddrV4, entry.Port, settingsApiEndpoint)
//}

// RequestConfig contains the configuration for an HTTP request.
type RequestConfig struct {
	Method      string
	URL         string
	Headers     map[string]string
	QueryParams map[string]string
	Body        interface{}
}

// ResponseData contains the response data from an HTTP request.
type ResponseData struct {
	StatusCode int
	Body       []byte
}

// sendRequest sends an HTTP request with the specified configuration and returns the response data.
func sendRequest(config RequestConfig) (ResponseData, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Populate unused config fields
	config.Headers = make(map[string]string)
	config.QueryParams = make(map[string]string)

	// Set the Accept header to application/json
	config.Headers["Accept"] = "application/json"

	// Create a new request with the specified method, URL, and query parameters
	req, err := http.NewRequest(config.Method, config.URL, nil)
	if err != nil {
		return ResponseData{}, err
	}
	for key, value := range config.QueryParams {
		req.URL.Query().Add(key, value)
	}

	// Set the headers
	for key, value := range config.Headers {
		req.Header.Set(key, value)
	}

	// Set the request body if provided
	if config.Body != nil {
		body, err := json.Marshal(config.Body)
		if err != nil {
			return ResponseData{}, err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		req.ContentLength = int64(len(body))
	}

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		return ResponseData{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := readAll(resp.Body)
	if err != nil {
		return ResponseData{}, err
	}

	// Create the ResponseData struct and return it
	responseData := ResponseData{
		StatusCode: resp.StatusCode,
		Body:       body,
	}
	return responseData, nil
}

// sendGet sends a GET request to the specified URL with the specified query parameters.
func sendGet(url string) (ResponseData, error) {
	config := RequestConfig{
		Method: "GET",
		URL:    url,
	}
	return sendRequest(config)
}

// sendPut sends a PUT request to the specified URL with the specified request body.
func sendPut(url string, body interface{}) (ResponseData, error) {
	config := RequestConfig{
		Method: "PUT",
		URL:    url,
		Body:   body,
	}
	return sendRequest(config)
}

// sendPost sends a POST request to the specified URL with the specified request body.
func sendPost(url string, body interface{}) (ResponseData, error) {
	config := RequestConfig{
		Method: "POST",
		URL:    url,
		Body:   body,
	}
	return sendRequest(config)
}

// sendDelete sends a DELETE request to the specified URL with the specified query parameters.
func sendDelete(url string) (ResponseData, error) {
	config := RequestConfig{
		Method: "DELETE",
		URL:    url,
	}
	return sendRequest(config)
}

// readAll reads all bytes from a reader.
func readAll(reader io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
