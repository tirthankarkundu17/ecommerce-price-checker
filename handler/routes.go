package handler

import "github.com/tirthankarkundu17/ecommerce-price-checker/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/api/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/api/register", middlewares.SetMiddlewareJSON(s.Register)).Methods("POST")

	// User's Products routes
	s.Router.HandleFunc("/api/products", middlewares.SetMiddlewareAuthentication(s.FetchProductsHandler)).Methods("GET")
	s.Router.HandleFunc("/api/products", middlewares.SetMiddlewareAuthentication(s.CreateUserProductHandler)).Methods("POST")
}
