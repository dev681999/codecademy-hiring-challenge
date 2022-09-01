package auth

import (
	"catinator-backend/pkg/httpwriter"
	"context"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/jwt"
)

type userIDContextKey struct{}

// Authenticator is a default authentication middleware to enforce access from the
// Verifier middleware request context values. The Authenticator sends a 401 Unauthorized
// response for any unverified tokens and passes the good ones through. It's just fine
// until you decide to write something similar and customize your client response.
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token, claims, err := jwtauth.FromContext(ctx)

		if err != nil {
			httpwriter.WriteErrJsonResponse(http.StatusUnauthorized, w, err.Error())
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			httpwriter.WriteErrJsonResponse(http.StatusUnauthorized, w, http.StatusText(http.StatusUnauthorized))
			return
		}

		ctx = context.WithValue(ctx, userIDContextKey{}, claims["id"].(string))

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserIDFromContext(ctx context.Context) string {
	v := ctx.Value(userIDContextKey{})
	return v.(string)
}
