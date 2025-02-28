package save

import (
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"urlShortener/internal/config"
	"urlShortener/internal/lib/api/response"
	"urlShortener/internal/lib/logger/sl"
	"urlShortener/internal/lib/random"
	"urlShortener/internal/storage"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	response.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveURL(urlToSave, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.save.New"
		log.With(slog.String("operation", op), middleware.GetReqID(r.Context()))

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to parse request", sl.Err(err))

			render.JSON(w, r, response.Error("failde to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("failed to validate request", sl.Err(err))
			var validationErrors validator.ValidationErrors
			errors.As(err, &validationErrors)

			render.JSON(w, r, response.ValidationError(validationErrors))

			return
		}

		alias := req.Alias
		if alias == "" {
			// todo так можно делать с конфигом?
			alias = random.NewRandomString(config.MustLoad().RandomStringLength)
		}

		id, err := urlSaver.SaveURL(req.URL, alias)
		if err != nil {
			if errors.Is(err, storage.ErrUrlExists) {
				log.Error("url already exist", slog.String("url", req.URL))
				render.JSON(w, r, response.Response{
					Error:  storage.ErrUrlExists.Error(),
					Status: response.StatusError,
				})

				return
			}
			log.Error("failed to save url", sl.Err(err))
			render.JSON(w, r, response.Response{
				Error:  "failed to save url",
				Status: response.StatusError,
			})

			return
		}

		log.Info("url added", slog.Int64("id", id))

		ResponseOK(w, r, alias)
	}
}

func ResponseOK(w http.ResponseWriter, r *http.Request, alias string) {
	render.JSON(w, r, Response{
		Response: response.OK(),
		Alias:    alias,
	})
}
