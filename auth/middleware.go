package auth

import (
	"context"
	"fmt"
	"net/http"

	//"strconv"

	"github.com/El-Hendawy/gograph/internal/users"
	"github.com/El-Hendawy/gograph/jwt"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			// Allow unauthenticated users in
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token
			tokenStr := header
			//fmt.Println(tokenStr)

			email, err := jwt.ParseToken(tokenStr)
			//fmt.Println(email)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// create user and check if user exists in db
			user := users.User{Email: email}

			id, err := usersRepo.GetUserIdByUsername(email)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			user.ID = id
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)
			fmt.Println(user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}

var usersRepo users.UsersRepository = users.NewValidate()
