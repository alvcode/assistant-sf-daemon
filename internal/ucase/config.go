package ucase

import (
	"assistant-sf-daemon/internal/dict"
	"assistant-sf-daemon/internal/dto"
	"assistant-sf-daemon/internal/repository"
	"assistant-sf-daemon/internal/service"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var (
	ErrDefineAppPath      = errors.New("error define app path")
	ErrAppPathNotExists   = errors.New("app path does not exist")
	ErrDBFileNotExists    = errors.New("db file not exist")
	ErrCreateAppDirectory = errors.New("error create app directory")
	ErrAuth               = errors.New("error auth")
)

type ConfigUseCase interface {
	GetStatus() error
	Init(in dto.Config) error
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

	dbExists := service.FileExists(filepath.Join(appPath, dict.DBName))
	if !dbExists {
		return ErrDBFileNotExists
	}

	return nil
}

func (uc *configUseCase) Init(in dto.Config) error {
	appPath, err := service.GetAppPath()
	if err != nil {
		return fmt.Errorf("%w: %s", ErrDefineAppPath, err.Error())
	}

	pathExists := service.PathExists(appPath)
	if !pathExists {
		if err := os.MkdirAll(filepath.Dir(appPath), 0o755); err != nil {
			return fmt.Errorf("error create directory: %v", err)
		}
	}

	configRepo := repository.GetCreator().Config()
	err = configRepo.CreateTableIfNotExists()
	if err != nil {
		return err
	}

	err = configRepo.Upsert(dict.ConfigKeyDomain, in.DriveDomain)
	if err != nil {
		return err
	}
	err = configRepo.Upsert(dict.ConfigKeyFolderPath, in.FolderPath)
	if err != nil {
		return err
	}

	_, err = service.Authentication(in.DriveDomain, in.AssistantLogin, in.AssistantPassword, configRepo)
	if err != nil {
		return ErrAuth
	}
	return nil
}
