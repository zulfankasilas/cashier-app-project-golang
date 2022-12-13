package api

import (
	"a21hc3NpZ25tZW50/model"
	"context"
	"encoding/json"
	"net/http"
)

func (api *API) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_token")

		if err != nil {
			response, _ := json.Marshal(model.ErrorResponse{Error: "http: named cookie not present"})
			w.WriteHeader(401)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(response))
			return
		}

		sessionToken := c.Value

		sessionFound, err := api.sessionsRepo.CheckExpireToken(sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "username", sessionFound.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {

			errJson, _ := json.Marshal(model.ErrorResponse{Error: "Method is not allowed!"})

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(405)
			w.Write(errJson)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (api *API) Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {

			errJson, _ := json.Marshal(model.ErrorResponse{Error: "Method is not allowed!"})

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(405)
			w.Write(errJson)
			return
		}
		next.ServeHTTP(w, r)
	})
}
