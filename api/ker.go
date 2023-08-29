package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/csvwolf/ker.go/constant"
	"io"
	"net/http"
	"time"
)

type Ker struct {
	SecretKey string
	Email     string
}

func (ker *Ker) sign(body []byte) string {
	h := hmac.New(sha256.New, []byte(ker.SecretKey))
	h.Write(body)
	return hex.EncodeToString(h.Sum(nil))
}

func (ker *Ker) fetch(body map[string]interface{}) (map[string]interface{}, error) {
	var (
		err          error
		bodyBytes    []byte
		req          *http.Request
		resp         *http.Response
		responseBody []byte
		retVal       map[string]interface{}

		client = http.Client{
			Timeout: 5 * time.Second,
		}
	)
	if bodyBytes, err = json.Marshal(body); err != nil {
		return nil, err
	}

	if req, err = http.NewRequest(http.MethodPost, constant.ApiUrl, bytes.NewReader(bodyBytes)); err != nil {
		return nil, err
	}
	req.Header.Set("Email", ker.Email)
	req.Header.Set("Sign", ker.sign(bodyBytes))

	if resp, err = client.Do(req); err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if responseBody, err = io.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	if err = json.Unmarshal(responseBody, &retVal); err != nil {
		return nil, err
	}
	return retVal, nil
}
