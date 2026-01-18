package users

import (
	"context"
	"github.com/spudmashmedia/gouser/internal/json"
	"log/slog"
	"net/http"
	"strconv"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// create QueryData object for context
		q := QueryData{}

		// count := chi.URLParam(r, "count") // use chi for path variables i.e. /user/{count}

		// use r http.request to get query strings
		// i.e. /user?count={int}
		paramCount := r.URL.Query().Get("count")

		slog.Debug(
			"User.Handler.UserCtx: Got param count ",
			"paramCount", paramCount)

		intCount, err := strconv.Atoi(paramCount)

		if err != nil {
			intCount = 1
			slog.Debug("UserCtx: paramCount string to int conversion failed, set to 1")
		}
		q.Count = intCount

		//
		// get concurrent = true/false
		//
		paramConcurrent := r.URL.Query().Get("concurrent")

		slog.Debug(
			"User.Handler.UserCtx: Got param concurrent ",
			"paramConcurrent", paramConcurrent)

		isConcurrent, err := strconv.ParseBool(paramConcurrent)
		if err != nil {
			isConcurrent = false

			slog.Debug("UserCtx: isConcurrent string to bool conversion failed, set to false")
		}

		q.IsConcurrent = isConcurrent

		//
		// get simLongProcess = true/false
		//
		paramSimLongProcess := r.URL.Query().Get("sim_long_proc")

		slog.Debug("User.Handler.UserCtx: Got param sim_long_proc",
			"paramSimLongProcess", paramSimLongProcess)

		isSimLongProcess, err := strconv.ParseBool(paramSimLongProcess)
		if err != nil {
			isSimLongProcess = false
			slog.Debug("UserCtx: isSimLongProcess string to bool conversion failed, set to false")
		}
		q.IsSimLongProcess = isSimLongProcess

		ctx := context.WithValue(r.Context(), "queryData", q)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	q, ok := ctx.Value("queryData").(QueryData)

	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	count := q.Count
	slog.Debug(
		"User.Handler.GetUser: Got count",
		"count", count)

	// validation: make sure at least 1 record
	if count == 0 {
		slog.Debug("User.Handler.GetUser: Validation - Got 0, set to at least 1 (scenario no count query string param)")
		count = 1
	}

	// validation: make sure max 100
	if count > 5000 {
		slog.Debug(
			"User.Handler.GetUser: Got count %d, cap max to 100",
			"count", count)

		count = 5000
	}

	var response UsersResponse
	var err error

	if q.IsConcurrent {
		slog.Debug("User.Handler.GetUser: call GetUserConcurrent")
		response, err = h.service.GetUserConcurrent(ctx, count, q.IsSimLongProcess)
	} else {
		slog.Debug("User.Handler.GetUser: call GetUser")
		response, err = h.service.GetUser(count, q.IsSimLongProcess)
	}

	if err != nil {
		slog.Error(
			"Something went wrong in handler.GetUser",
			"error", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, response)
}
