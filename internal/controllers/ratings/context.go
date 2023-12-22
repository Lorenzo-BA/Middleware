package ratings

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/rating/internal/models"
	"net/http"
)

func CtxSongId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		songId, err := uuid.FromString(chi.URLParam(r, "song_id"))
		if err != nil {
			logrus.Errorf("error : %d - Parsing error : %s", http.StatusUnprocessableEntity, err.Error())
			customError := &models.CustomError{
				Message: fmt.Sprintf("cannot parse id (%s) as UUID", chi.URLParam(r, "id")),
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
			return
		}

		ctx := context.WithValue(r.Context(), "songID", songId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func CtxRatingId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ratingId, err := uuid.FromString(chi.URLParam(r, "rating_id"))
		if err != nil {
			logrus.Errorf("error : %d - Parsing error : %s", http.StatusUnprocessableEntity, err.Error())
			customError := &models.CustomError{
				Message: fmt.Sprintf("cannot parse id (%s) as UUID", chi.URLParam(r, "id")),
				Code:    http.StatusUnprocessableEntity,
			}
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
			return
		}

		ctx := context.WithValue(r.Context(), "ratingId", ratingId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
