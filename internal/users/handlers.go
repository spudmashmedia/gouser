package users

import (
	"context"
	"github.com/spudmashmedia/gouser/internal/json"
	"log"
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

		log.Printf("User.Handler.UserCtx: Got param count %s", paramCount)

		intCount, err := strconv.Atoi(paramCount)

		if err != nil {
			intCount = 1
			log.Printf("UserCtx: paramCount string to int conversion failed, set to 1")
		}
		q.Count = intCount

		//
		// get concurrent = true/false
		//
		paramConcurrent := r.URL.Query().Get("concurrent")
		log.Printf("User.Handler.UserCtx: Got param concurrent %s", paramConcurrent)
		isConcurrent, err := strconv.ParseBool(paramConcurrent)
		if err != nil {
			isConcurrent = false
			log.Printf("UserCtx: isConcurrent string to bool conversion failed, set to false")
		}
		q.IsConcurrent = isConcurrent

		//
		// get simLongProcess = true/false
		//
		paramSimLongProcess := r.URL.Query().Get("sim_long_proc")
		log.Printf("User.Handler.UserCtx: Got param sim_long_proc %s", paramSimLongProcess)
		isSimLongProcess, err := strconv.ParseBool(paramSimLongProcess)
		if err != nil {
			isSimLongProcess = false
			log.Printf("UserCtx: isSimLongProcess string to bool conversion failed, set to false")
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
	log.Printf("User.Handler.GetUser: Got count %d", count)

	// validation: make sure at least 1 record
	if count == 0 {
		log.Printf("User.Handler.GetUser: Validation - Got 0, set to at least 1 (scenario no count query string param)")
		count = 1
	}

	// validation: make sure max 100
	if count > 5000 {
		log.Printf("User.Handler.GetUser: Got count %d, cap max to 100", count)
		count = 5000
	}

	var response UsersResponse
	var err error

	if q.IsConcurrent {
		log.Printf("User.Handler.GetUser: call GetUserConcurrent")
		response, err = h.service.GetUserConcurrent(ctx, count, q.IsSimLongProcess)
	} else {
		log.Printf("User.Handler.GetUser: call GetUser")
		response, err = h.service.GetUser(count, q.IsSimLongProcess)
	}

	if err != nil {
		log.Println("Something went wrong in handler.GetUser")
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, response)
}
