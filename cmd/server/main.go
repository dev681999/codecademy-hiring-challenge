package main

import (
	"catinator-backend/pkg/auth"
	"catinator-backend/pkg/cat"
	"catinator-backend/pkg/config"
	"catinator-backend/pkg/db"
	"catinator-backend/pkg/file"
	"catinator-backend/pkg/fileserve"
	"catinator-backend/pkg/httpwriter"
	"catinator-backend/pkg/log"
	"catinator-backend/pkg/user"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/oklog/run"
	"go.uber.org/zap"
	"moul.io/chizap"
)

var flagConfig = flag.String("config", "./config.yaml", "path to the config file")

func main() {
	flag.Parse()

	var cfg config.Config

	err := config.New(&cfg, config.FromFile(*flagConfig))
	if err != nil {
		log.L.Fatal("", zap.Error(err))
	}

	logger := log.New(cfg.Server.Debug)

	err = os.MkdirAll(cfg.Server.PublicStorageFolder, os.ModePerm)
	if err != nil {
		logger.Debug("", zap.Error(err))
	}

	tokenAuth := jwtauth.New("HS256", []byte(cfg.Server.JWTSecret), nil)

	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	entClient, err := db.OpenEntClient(ctx, cfg.DB)
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	defer entClient.Close()

	fileSvc := file.NewDiskService()

	catSvc := cat.NewService(
		*entClient,
		tokenAuth,
		cfg,
		fileSvc,
		logger.With(zap.String("service", "cat")),
	)
	userSvc := user.NewService(
		*entClient,
		tokenAuth,
		logger.With(zap.String("service", "user")),
	)

	r := chi.NewRouter()

	r.Use(chizap.New(logger, &chizap.Opts{
		// WithReferer:   true,
		// WithUserAgent: true,
	}))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		httpwriter.Write200JsonResponse(w, map[string]string{
			"msg": "ok",
		})
	})

	r.Get("/swagger-ui/*", fileserve.NewHandler("*", cfg.Server.SwaggerUIFolder))

	r.Get("/api.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, cfg.Server.SwaggerAPIFilePath)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(auth.Authenticator)

		catSvc.MountHandlers(r)
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Get("/cat/image/{fileName}", fileserve.NewHandler("fileName", cfg.Server.PublicStorageFolder))
		userSvc.MountHandlers(r)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		httpwriter.WriteErrJsonResponse(http.StatusNotFound, w, "Requested api operation not found")
	})

	var g run.Group
	{
		server := &http.Server{Addr: addr, Handler: r}

		g.Add(func() error {
			logger.Info("server", zap.String("msg", "serving http"), zap.String("addr", addr))
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return err
			}
			return nil
		}, func(error) {
			logger.Info("server", zap.String("msg", "stopping http"))
			/* if err := lis.Close(); err != nil {
				logger.Fatal("", zap.Error(err))
			} */
			ctx, cancel := context.WithTimeout(ctx, time.Second*1)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				logger.Fatal("", zap.Error(err))
			}

			logger.Info("db", zap.String("msg", "stopping db"))
			if err := entClient.Close(); err != nil {
				logger.Fatal("", zap.Error(err))
			}
		})
	}
	{
		// set-up our signal handler
		var (
			cancelInterrupt = make(chan struct{})
			c               = make(chan os.Signal, 2)
		)
		defer close(c)

		g.Add(func() error {
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	logger.Error("exit", zap.Error(g.Run()))
}
