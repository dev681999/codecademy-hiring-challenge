package cat

import (
	"catinator-backend/pkg/auth"
	"catinator-backend/pkg/config"
	"catinator-backend/pkg/db/ent"
	"catinator-backend/pkg/db/ent/cat"
	"catinator-backend/pkg/file"
	"catinator-backend/pkg/httpwriter"
	"catinator-backend/pkg/model"
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	pass "github.com/dev681999/go-pass"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service struct {
	pass.Hash
	ent       ent.Client
	tokenAuth *jwtauth.JWTAuth
	cfg       config.Config
	logger    *zap.Logger
	fileSvc   file.Service
}

func NewService(
	ent ent.Client,
	tokenAuth *jwtauth.JWTAuth,
	cfg config.Config,
	fileSvc file.Service,
	logger *zap.Logger,
) *Service {
	return &Service{
		ent:       ent,
		tokenAuth: tokenAuth,
		cfg:       cfg,
		fileSvc:   fileSvc,
		logger:    logger,
	}
}

func (s *Service) MountHandlers(r chi.Router) {
	r.Post("/cats", s.AddCat)
	r.Get("/cats", s.ListCats)
	r.Get("/cat/{catId}", s.GetCat)
	r.Patch("/cat/{catId}", s.UpdateCat)
	r.Delete("/cat/{catId}", s.DeleteCat)
}

func (s *Service) AddCat(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	r.ParseMultipartForm(0)
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		s.logger.Debug("err", zap.Error(err))
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}
	defer file.Close()
	v := &model.AddCat{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
		Tags:        strings.Split(r.FormValue("tags"), ","),
	}

	fileName := uuid.NewString() + filepath.Ext(fileHeader.Filename)
	path := filepath.Join(s.cfg.Server.PublicStorageFolder, fileName)

	err = s.fileSvc.Create(path, file)
	if err != nil {
		s.logger.Debug("err", zap.Error(err))
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	dbCat, err := s.ent.Cat.Create().
		SetName(v.Name).
		SetImageID(fileName).
		SetDescription(v.Description).
		SetTags(v.Tags).
		SetOwnerID(userID).
		Save(r.Context())
	if err != nil {
		s.logger.Debug("err", zap.Error(err))
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	cat := MapDBCatToModel(dbCat)

	httpwriter.Write200JsonResponse(w, cat)
}

func (s *Service) ListCats(w http.ResponseWriter, r *http.Request) {
	userID := auth.GetUserIDFromContext(r.Context())
	order := r.URL.Query().Get("sort")

	dbCats, err := s.ent.Cat.Query().
		Where(
			cat.OwnerIDEQ(userID),
		).
		Order(stringSortToEnt(order, cat.FieldUpdateTime)).
		All(r.Context())
	if err != nil {
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	cats := MapDBCatsToModel(dbCats)

	httpwriter.Write200JsonResponse(w, cats)
}

func (s *Service) GetCat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := auth.GetUserIDFromContext(r.Context())
	catID := chi.URLParam(r, "catId")

	dbCat, err := s.ent.Cat.Get(ctx, catID)
	if err != nil {
		if ent.IsNotFound(err) {
			httpwriter.WriteErrJsonResponse(http.StatusNotFound, w, "cat not found")
			return
		}
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	if userID != dbCat.OwnerID {
		httpwriter.WriteErrJsonResponse(http.StatusForbidden, w, "only owner cat can access a cat")
		return
	}

	cats := MapDBCatToModel(dbCat)

	httpwriter.Write200JsonResponse(w, cats)
}

func (s *Service) UpdateCat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := auth.GetUserIDFromContext(r.Context())
	catID := chi.URLParam(r, "catId")
	catExists, err := s.ent.Cat.Query().
		Where(cat.IDEQ(catID)).
		Exist(ctx)
	if err != nil {
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	if !catExists {
		httpwriter.WriteErrJsonResponse(http.StatusNotFound, w, "cat not found")
		return
	}

	isUserOwner, err := s.ent.Cat.Query().
		Where(
			cat.IDEQ(catID),
			cat.OwnerIDEQ(userID),
		).
		Exist(ctx)
	if err != nil {
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	if !isUserOwner {
		httpwriter.WriteErrJsonResponse(http.StatusForbidden, w, "only owner cat update a cat")
		return
	}

	updater := s.ent.Cat.UpdateOneID(catID)

	r.ParseMultipartForm(0)
	file, fileHeader, err := r.FormFile("image")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}
	if err == nil {
		defer file.Close()
		fileName := uuid.NewString() + filepath.Ext(fileHeader.Filename)
		path := filepath.Join(s.cfg.Server.PublicStorageFolder, fileName)

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
			return
		}

		updater.SetImageID(fileName)
	}

	if len(r.Form["name"]) > 0 {
		updater.SetName(r.Form["name"][0])
	}
	if len(r.Form["description"]) > 0 {
		updater.SetDescription(r.Form["description"][0])
	}
	if len(r.Form["tags"]) > 0 {
		updater.SetTags(r.Form["tags"])
	}

	dbCat, err := updater.Save(ctx)
	if err != nil {
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	cat := MapDBCatToModel(dbCat)

	httpwriter.Write200JsonResponse(w, cat)
}

func (s *Service) DeleteCat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := auth.GetUserIDFromContext(r.Context())
	catID := chi.URLParam(r, "catId")

	dbCat, err := s.ent.Cat.Get(ctx, catID)
	if err != nil {
		if ent.IsNotFound(err) {
			httpwriter.WriteErrJsonResponse(http.StatusNotFound, w, "cat not found")
			return
		}
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	if userID != dbCat.OwnerID {
		httpwriter.WriteErrJsonResponse(http.StatusForbidden, w, "only owner cat delete a cat")
		return
	}

	path := filepath.Join(s.cfg.Server.PublicStorageFolder, dbCat.ImageID)

	err = s.fileSvc.Delete(path)
	if err != nil {
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	err = s.ent.Cat.DeleteOneID(catID).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			httpwriter.WriteErrJsonResponse(http.StatusNotFound, w, "cat not found")
			return
		}
		httpwriter.WriteErrJsonResponse(http.StatusInternalServerError, w, err.Error())
		return
	}

	httpwriter.Write200JsonResponse(w, model.SucessMessage{
		Message: "cat deleted sucessfully",
	})
}
