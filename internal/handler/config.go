package handler

import (
	"assistant-sf-daemon/internal/dto"
	"assistant-sf-daemon/internal/locale"
	"assistant-sf-daemon/internal/ucase"
	"encoding/json"
	"net/http"
)

type ConfigHandler struct {
	useCase ucase.ConfigUseCase
}

func NewConfigHandler(uCase ucase.ConfigUseCase) *ConfigHandler {
	return &ConfigHandler{
		useCase: uCase,
	}
}

func (h *ConfigHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	langRequest := locale.GetLangFromContext(r.Context())

	err := h.useCase.GetStatus()
	if err != nil {
		SendErrorResponse(w, buildErrorMessage(langRequest, err), http.StatusUnprocessableEntity, 0)
		return
	}
	SendResponse(w, http.StatusOK, nil)
}

func (h *ConfigHandler) Init(w http.ResponseWriter, r *http.Request) {
	langRequest := locale.GetLangFromContext(r.Context())
	var requestDTO dto.Config

	err := json.NewDecoder(r.Body).Decode(&requestDTO)
	if err != nil {
		SendErrorResponse(w, locale.T(langRequest, "error_reading_request_body"), http.StatusBadRequest, 0)
		return
	}

	err = h.useCase.Init(requestDTO)
	if err != nil {
		SendErrorResponse(w, buildErrorMessage(langRequest, err), http.StatusUnprocessableEntity, 0)
		return
	}
	SendResponse(w, http.StatusOK, nil)
}
