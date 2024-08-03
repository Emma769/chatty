package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/emma769/chatty/internal/model"
)

type UserClient struct {
	url string
}

func (cl *UserClient) CreateUser(in model.UserIn) (*model.User, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/api/users", cl.url)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("non-2** response: %v", resp)
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	user := &model.User{}

	if err := json.Unmarshal(payload, user); err != nil {
		return nil, err
	}

	return user, nil
}
