package handlers

import (
	"net/http"

	"github.com/abdullahnettoor/tastybites/internal/usecases"
	"github.com/abdullahnettoor/tastybites/internal/utils"
)

type menuHandler struct {
	MenuUsecase usecases.MenuIUsecase
}

func NewMenuHandler(menuUsecase usecases.MenuIUsecase) *menuHandler {
	return &menuHandler{
		MenuUsecase: menuUsecase,
	}
}

func (h *menuHandler) GetAllMenuItems(w http.ResponseWriter, r *http.Request) {

	allMenuItems, err := h.MenuUsecase.GetAllMenuItems(r.Context())
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, allMenuItems)
}

func (h *menuHandler) GetAllTables(w http.ResponseWriter, r *http.Request) {
	allTables, err := h.MenuUsecase.GetAvailableTables(r.Context())
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, allTables)
}
