package hooks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"restaurant-flow-auth/pkg/util"
)

const apiURL = "http://localhost:3333"

func PreRegister(userId, email, username, name, password string) error {
	token, err := util.GenerateJWT("uwEats")

	if err != nil {
		return err
	}

	postBody, err := json.Marshal(map[string]string{
		"userId":   userId,
		"email":    email,
		"username": username,
		"name":     name,
		"password": password,
	})

	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", apiURL, "internal/register"), bytes.NewBuffer(postBody))

	if err != nil {
		return err
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	request.Header.Add("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}
	if !(response.StatusCode >= 200 && response.StatusCode < 300) {
		return errors.New("failed to contact API")
	}

	return nil
}
