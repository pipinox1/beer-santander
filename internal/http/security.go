package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type authBody struct {
	UserId     int64
	Role       string
	IsAdmin    bool
	UserStatus string
}

var protectedHeaders = []string{"X-Role", "X-Is-Admin", "X-User-Status"}

func securityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("SKIP_AUTH") != "true" {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				WebResponse(w, 401, "Unauthorized")
				return
			}
			url := fmt.Sprintf("http://auth/token/%s", tokenString)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil || req == nil {
				WebResponse(w, 500, "internal_server_error")
				return
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil || resp.StatusCode != 200 {
				WebResponse(w, 401, "Unauthorized")
				return
			}
			authUser := &authBody{}
			err = json.NewDecoder(resp.Body).Decode(&authUser)
			if err != nil {
				WebResponse(w, 500, "internal_server_error")
			}
			cleanHeader(r)
			addHeader(r, *authUser)
			ctx := context.WithValue(r.Context(), "user_logged", authUser)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			r.Header.Add("X-Is-Admin", "true")
			next.ServeHTTP(w, r)
		}
	})
}

func cleanHeader(request *http.Request) {
	for _, header := range protectedHeaders {
		request.Header.Del(header)
	}
}

func addHeader(request *http.Request, authUser authBody) {
	request.Header.Add("X-User-Status", authUser.UserStatus)
	request.Header.Add("X-Is-Admin", strconv.FormatBool(authUser.IsAdmin))
	request.Header.Add("X-Role", authUser.Role)
}
