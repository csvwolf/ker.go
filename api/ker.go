package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/csvwolf/ker.go/constant"
	"github.com/csvwolf/ker.go/model"
	"io"
	"net/http"
	"time"
)

type Ker struct {
	SecretKey string
	Email     string
}

func (ker *Ker) Sign(body []byte) string {
	h := hmac.New(sha256.New, []byte(ker.SecretKey))
	h.Write(body)
	return hex.EncodeToString(h.Sum(nil))
}

func fetch[Input any, Output any](ker *Ker, body *Input) (*Output, error) {
	var (
		err          error
		bodyBytes    []byte
		req          *http.Request
		resp         *http.Response
		responseBody []byte
		retVal       *model.Response[*Output]

		client = http.Client{
			Timeout: 20 * time.Second,
		}
	)
	if bodyBytes, err = json.Marshal(body); err != nil {
		return nil, err
	}

	if req, err = http.NewRequest(http.MethodPost, constant.ApiUrl, bytes.NewReader(bodyBytes)); err != nil {
		return nil, err
	}
	req.Header.Set("Email", ker.Email)
	req.Header.Set("Sign", ker.Sign(bodyBytes))

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

	if !retVal.Success {
		if retVal.Error == "" {
			return nil, errors.New("unknown Error")
		}
		return nil, errors.New(retVal.Error)
	}

	return retVal.Result, nil
}
