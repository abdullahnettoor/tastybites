package api

import (
	"net/http"

	"github.com/abdullahnettoor/tastybites/internal/api/handlers"
	"github.com/abdullahnettoor/tastybites/internal/api/middlewares"
	"github.com/abdullahnettoor/tastybites/internal/usecases"
)

type RouteGroup struct {
	mux         *http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

func (app *application) NewRouteGroup(middleware ...func(http.Handler) http.Handler) *RouteGroup {
	return &RouteGroup{
		mux:         app.Mux,
		middlewares: middleware,
	}
}

func (rg *RouteGroup) Handle(pattern string, handler http.HandlerFunc) {
	wrappedHandler := middlewares.MiddlewareChain(
		http.HandlerFunc(handler),
		rg.middlewares...,
	)
	rg.mux.Handle(pattern, wrappedHandler)
}

func (rg *RouteGroup) HandleFunc(pattern string, handler http.HandlerFunc) {
	rg.Handle(pattern, handler)
}

// Alternative approach using route groups
func (app *application) InitializeRoutes(
	userUsecase usecases.UserIUsecase,
	orderUsecase usecases.OrderIUsecase,
	menuUsecase usecases.MenuIUsecase,
	tableUsecase usecases.TableIUsecase,
) {

	publicGroup := app.NewRouteGroup()
	adminGroup := app.NewRouteGroup(middlewares.Authenticate, middlewares.AuthorizeAdmin) // Admin protected routes
	userGroup := app.NewRouteGroup(middlewares.Authenticate)                              // User protected routes

	// Handlers
	userHandler := handlers.NewUserHandler(userUsecase, tableUsecase)
	menuHandler := handlers.NewMenuHandler(menuUsecase)
	orderHandler := handlers.NewOrderHandler(orderUsecase, userUsecase)

	// Public routes
	publicGroup.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})
	publicGroup.HandleFunc("POST /login", userHandler.UserLogin)
	publicGroup.HandleFunc("POST /register", userHandler.UserRegister)
	publicGroup.HandleFunc("GET /menu", menuHandler.GetAllMenuItems)
	publicGroup.HandleFunc("GET /tables", userHandler.GetAvailableTables)

	// Authenticated user routes
	userGroup.HandleFunc("GET /orders", orderHandler.GetUserOrders)
	userGroup.HandleFunc("POST /orders", orderHandler.CreateOrder)

	// Admin routes
	adminGroup.HandleFunc("GET /admin/orders", orderHandler.AdminGetAllOrders)
	adminGroup.HandleFunc("GET /admin/tables/", orderHandler.GetOrderByTableId)

}
