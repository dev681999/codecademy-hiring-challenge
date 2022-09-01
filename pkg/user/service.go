package user

import (
	"catinator-backend/pkg/db/ent"
	"catinator-backend/pkg/db/ent/user"
	"catinator-backend/pkg/httpwriter"
	"catinator-backend/pkg/model"
	"net/http"
	"time"

	pass "github.com/dev681999/go-pass"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type Service struct {
	pass.Hash
	ent       ent.Client
	tokenAuth *jwtauth.JWTAuth
	logger    *zap.Logger
}

func NewService(
	ent ent.Client,
	tokenAuth *jwtauth.JWTAuth,
	logger *zap.Logger,
) *Service {
	return &Service{
		ent:       ent,
		tokenAuth: tokenAuth,
		logger:    logger,
		Hash:      pass.Hash{},
	}
}

func (s *Service) MountHandlers(r chi.Router) {
	r.Post("/auth/register", s.Register)
	r.Post("/auth/login", s.Login)
}

func (s *Service) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	v := &model.Registration{}
	err := render.Decode(r, v)
	if err != nil {
		httpwriter.WriteErrJsonResponse(http.StatusBadRequest, w, err.Error())
		return
	}
	v.Password, err = s.Generate(v.Password)
	if err != nil {
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}
	err = s.ent.User.Create().
		SetEmail(v.Email).
		SetName(v.Name).
		SetPassword(v.Password).Exec(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			httpwriter.WriteErrJsonResponse(http.StatusBadRequest, w, "user already exists")
			return
		}
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	httpwriter.Write200JsonResponse(w, model.SucessMessage{
		Message: "user created sucessfully",
	})
}

func (s *Service) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	v := &model.Login{}
	err := render.Decode(r, v)
	if err != nil {
		httpwriter.WriteErrJsonResponse(http.StatusBadRequest, w, err.Error())
		return
	}
	dbUser, err := s.ent.User.Query().
		Where(
			user.Email(v.Email),
		).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			httpwriter.WriteErrJsonResponse(http.StatusNotFound, w, "user not found")
			return
		}
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	err = s.Compare(dbUser.Password, v.Password)
	if err != nil {
		httpwriter.WriteErrJsonResponse(http.StatusBadRequest, w, "bad user credentials")
		return
	}

	userClaims := map[string]interface{}{
		"id": dbUser.ID,
	}

	jwtauth.SetExpiryIn(userClaims, time.Hour*24*365)

	_, token, err := s.tokenAuth.Encode(userClaims)
	if err != nil {
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	httpwriter.Write200JsonResponse(w, model.LoginDetails{
		Token: token,
	})
}
