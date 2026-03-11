package service

import (
	"assistant-sf-daemon/internal/dict"
	"assistant-sf-daemon/internal/dto"
	"assistant-sf-daemon/internal/repository"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Authentication(domain string, login string, password string, configRepo repository.ConfigRepository) (*dto.LoginSuccessResponse, error) {
	url := domain + "/api/auth/login"

	reqBody := dto.LoginRequest{
		Login:    login,
		Password: password,
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	// ошибка по статусу
	if resp.StatusCode != http.StatusOK {
		var er dto.ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&er); err != nil {
			return nil, errors.New("login failed. bad response")
		}
		return nil, errors.New(er.Message)
	}

	// успех
	var sr dto.LoginSuccessResponse
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return nil, err
	}

	//configRepo := repository.GetCreator().Config()
	err = configRepo.Upsert(dict.ConfigKeyToken, sr.Token)
	if err != nil {
		return nil, err
	}
	err = configRepo.Upsert(dict.ConfigKeyRefreshToken, sr.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &sr, nil
}

func RefreshToken(domain string, configRepo repository.ConfigRepository) error {
	url := domain + "/api/auth/refresh-token"

	//configRepo := repository.GetCreator().Config()

	token, err := configRepo.Get(dict.ConfigKeyToken)
	if err != nil {
		return err
	}
	refreshToken, err := configRepo.Get(dict.ConfigKeyRefreshToken)
	if err != nil {

	}

	reqBody := dto.RefreshTokenRequest{
		Token:        token,
		RefreshToken: refreshToken,
	}

	b, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		var er dto.ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&er); err != nil {
			return errors.New("refresh token failed. bad response")
		}
		return errors.New(er.Message)
	}

	var sr dto.LoginSuccessResponse
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return err
	}

	err = configRepo.Upsert(dict.ConfigKeyToken, sr.Token)
	if err != nil {
		return err
	}
	err = configRepo.Upsert(dict.ConfigKeyRefreshToken, sr.RefreshToken)
	if err != nil {
		return err
	}

	return nil
}
