package api

import "github.com/go-chi/chi/v5"

func (api *Api) BindRoutes() {
	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/user", func(r chi.Router) {
				r.Post("/signup", api.HandleSignup)
				r.Post("/login", api.HandleLogin)
				r.Post("/logout", api.HandleLogout)
			})
		})
	})
}
