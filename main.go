// Copyright (C) Oliver Amann
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License version 3 as
// published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/into-the-v0id/user-api.go/route"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	stdlog "log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func registerRoutes(router *mux.Router) {
	router.NotFoundHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, "404 Not Found", http.StatusNotFound)
	})

	route.RegisterUserRoutes(router.PathPrefix("/users").Subrouter())
}

// See https://stackoverflow.com/a/28746725
func RecoverHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			logger := zerolog.Ctx(request.Context())

			defer func() {
				rawErr := recover()
				if rawErr != nil {
					var err error
					switch castError := rawErr.(type) {
					case string:
						err = errors.New(castError)
					case error:
						err = castError
					default:
						err = errors.New(fmt.Sprintf("unknown error of type %T", rawErr))
					}

					logger.Error().Err(err).Msg("Panic")

					http.Error(writer, "500 Internal Server Error", http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(writer, request)
		})
	}
}

func NewServer() *http.Server {
	router := mux.NewRouter()

	router.Use(hlog.NewHandler(log.Logger))
	router.Use(hlog.ProtoHandler("protocol"))
	router.Use(hlog.URLHandler("url"))
	router.Use(hlog.MethodHandler("method"))
	router.Use(hlog.RefererHandler("referer"))
	router.Use(hlog.RequestIDHandler("requestId", "Request-Id"))
	router.Use(hlog.CustomHeaderHandler("correlationId", "X-Correlation-ID"))
	router.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().Int("status", status).Int("size", size).Dur("duration", duration).Msg("Request")
	}))

	router.Use(RecoverHandler())

	router.Use(handlers.HTTPMethodOverrideHandler)

	registerRoutes(router)

	return &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}
}

func init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Stack().Logger()
	stdlog.SetFlags(0)
	stdlog.SetOutput(log.Logger)
}

func main() {
	exitCode := 0

	server := NewServer()

	go func() {
		fmt.Printf("Listening on http://%s ...\n", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Server failed")
			if exitCode == 0 {
				exitCode = 1
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Server shutdown failed")
		if exitCode == 0 {
			exitCode = 1
		}
	}

	os.Exit(exitCode)
}
