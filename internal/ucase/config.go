package ucase

import (
	"assistant-sf-daemon/internal/service"
	"errors"
	"fmt"
)

var (
	ErrDefineAppPath    = errors.New("error define app path")
	ErrAppPathNotExists = errors.New("app path does not exist")
)

type ConfigUseCase interface {
	GetStatus() error
}

type configUseCase struct{}

func NewConfigUseCase() ConfigUseCase {
	return &configUseCase{}
}

func (uc *configUseCase) GetStatus() error {
	appPath, err := service.GetAppPath()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrDefineAppPath, err.Error())
	}

	pathExists := service.PathExists(appPath)
	if !pathExists {
		return ErrAppPathNotExists
	}
	return nil
}
