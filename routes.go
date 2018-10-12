package main

func (s *Server) routes() {
	api := s.router.PathPrefix("/api/").Subrouter()

	// Servers entity
	api.HandleFunc("/servers", ListOfServersHandler).Methods("GET")
}
