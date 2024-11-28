package routes

type ErrJSON struct {
	Message string `json:"message"`
}

// func CreateRoutes(s *http.ServeMux) {
// 	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		component := templates.Index()
// 		component.Render(context.Background(), w)
// 	})
// }
