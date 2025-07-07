package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/abdullahnettoor/tastybites/internal/api/dto"
	"github.com/abdullahnettoor/tastybites/internal/models"
	"github.com/abdullahnettoor/tastybites/internal/usecases"
	"github.com/abdullahnettoor/tastybites/internal/utils"
)

type orderHandler struct {
	OrderUsecase usecases.OrderIUsecase
	UserUsecase  usecases.UserIUsecase
}

func NewOrderHandler(orderUsecase usecases.OrderIUsecase, userUsecase usecases.UserIUsecase) *orderHandler {
	return &orderHandler{
		OrderUsecase: orderUsecase,
		UserUsecase:  userUsecase,
	}
}

func (h *orderHandler) AdminGetAllOrders(w http.ResponseWriter, r *http.Request) {

	allOrders, err := h.OrderUsecase.GetAllOrders(r.Context())
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, allOrders)
}

func (h *orderHandler) GetOrderByTableId(w http.ResponseWriter, r *http.Request) {

	tableIdStr := r.URL.Query().Get("tableId")
	if tableIdStr == "" {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Table ID is required")
		return
	}

	tableId, err := strconv.Atoi(tableIdStr)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid table ID format")
		return
	}

	allOrders, err := h.UserUsecase.GetOrderByTableId(r.Context(), tableId)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, allOrders)
}

func (h *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var orderReq dto.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	// Get user ID from authenticated context
	userID, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	oItems := make([]models.OrderItem, len(orderReq.Items))
	for i, item := range orderReq.Items {
		oItems[i] = models.OrderItem{
			MenuItemID: item.ItemID,
			Quantity:   item.Quantity,
			Price:      item.Price,
		}
	}

	order := models.Order{
		UserID:  userID,
		ItemsID: orderReq.ItemsID,
		Items:   oItems,
		TableID: orderReq.TableID,
		Status:  models.OrderStatusPending,
	}

	order.CalculateTotalPrice()

	orderId, err := h.OrderUsecase.CreateOrder(r.Context(), order)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteSuccessResponse(w, http.StatusCreated, "Order created successfully", orderId)
}

func (h *orderHandler) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.GetUserIDFromContext(r)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	allOrders, err := h.OrderUsecase.GetOrdersByUser(r.Context(), userId)
	if err != nil && errors.Is(err, models.ErrIsEmpty) {
		utils.WriteErrorResponse(w, http.StatusNotFound, "No orders found")
	} else if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, allOrders)
}
