package users

import (
	"context"
	"log/slog"
	"runtime"
	"sync"
	"time"

	"github.com/spudmashmedia/gouser/pkg/randomuser"
)

type svc struct {
	randomuserService randomuser.Service
}

type Service interface {
	GetUser(count int, isSimLongProcess bool) (UsersResponse, error)
	GetUserConcurrent(ctx context.Context, count int, isSimLongProcess bool) (UsersResponse, error)
}

func NewService(service randomuser.Service) Service {
	return &svc{
		randomuserService: service,
	}
}

func (s *svc) GetUser(count int, isSimLongProcess bool) (UsersResponse, error) {

	data, err := s.randomuserService.GetUsers(count)

	if err != nil {
		slog.Error(
			"Users.GetUser: Error from external Service",
			"err", err)

		return UsersResponse{}, err
	}

	dto := UsersResponse{
		Results: []User{},
	}

	if data.Results == nil || len(data.Results) == 0 {
		return dto, nil
	}

	slog.Debug(
		"Users.GetUser: Got %d from randomuser",
		"data_Results_length", len(data.Results))

	slog.Debug("Users.GetUser: Processing STARTED...")

	start := time.Now()

	for i := 0; i < len(data.Results); i++ {
		// log.Printf("Task %d: Started", i)
		var usr User
		usr = User{}
		usr.RuToUser(&data.Results[i])
		dto.Results = append(dto.Results, usr)

		if isSimLongProcess {
			time.Sleep(10 * time.Millisecond)
		}
	}

	duration := time.Since(start)

	slog.Debug(
		"Users.GetUser: Processing ENDED.",
		"duration", duration)

	if err != nil {
		slog.Error(
			"Users.GetUser: FromRandomUser failed",
			"error", err)

		return UsersResponse{}, err
	}

	return dto, nil
}

func (s *svc) GetUserConcurrent(ctx context.Context, count int, isSimLongProcess bool) (UsersResponse, error) {

	data, err := s.randomuserService.GetUsers(count)

	if err != nil {
		slog.Error(
			"Users.GetUserConcurrent: Error from external Service",
			"error", err)

		return UsersResponse{}, err
	}

	// response DTO
	dto := UsersResponse{}

	if data.Results == nil || len(data.Results) == 0 {
		return dto, nil
	}

	// Setup Go routines to process randomuser.Results array
	var wg sync.WaitGroup

	slog.Debug("Users.GetUserConcurrent: Processing STARTED...")

	start := time.Now()

	workerPoolSize := runtime.GOMAXPROCS(0)
	jobsBatchSize := 100
	resultsBatchSize := 100
	jobs := make(chan randomuser.User, jobsBatchSize)
	results := make(chan User, resultsBatchSize)

	// create worker pool
	for w := 0; w < workerPoolSize; w++ {
		wg.Add(1)
		go processRuToUserItems(ctx, w, jobs, results, &wg, isSimLongProcess)
	}

	// Feed job pool, nom nom nom...
	go func() {
		defer close(jobs)
		for _, item := range data.Results {
			select {
			// stop feeding job pool, context cancelled
			case <-ctx.Done():
				return

			// Feed job pool, nom nom nom...
			case jobs <- item:
			}
		}
	}()

	// Wait & close off results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Merge results to DTO
	// NOTE: because results is a channel, they will be
	//       streamed sequentially via a range.
	//       this is important to note as performance may
	//       equal the single threaded process.
	//
	//       Test this with isSimLongProcess enabled to
	//       simulate the processRuToUserItems taking longer
	//
	for item := range results {
		dto.Results = append(dto.Results, item)
	}

	duration := time.Since(start)

	slog.Debug(
		"Users.GetUserConcurrent: Processing ENDED.",
		"duration", duration)

	if err != nil {
		slog.Error(
			"Users.ConcurrentUsers: FromRandomUser failed",
			"error", err)

		return UsersResponse{}, err
	}

	// Final Validation, check response size equals original data size
	slog.Debug(
		"Original size (%d), Processed Size (%d)",
		"data_results_length", len(data.Results),
		"dto_results_length", len(dto.Results))

	slog.Debug(
		"Are they the same size? %v",
		"isSameSize", len(data.Results) == len(dto.Results))

	return dto, nil
}

// Worker Go Routine
// Takes a batch of randomuser.User items
// from the jobs channel (batch controlled by parent function)
//
// Converts randomuser.User to users.User
// Sends result back into results channel
func processRuToUserItems(ctx context.Context, index int, jobs <-chan randomuser.User, results chan<- User, wg *sync.WaitGroup, isSimLongProcess bool) {
	defer wg.Done()

	for {
		select {

		// context cancelled, exit worker immediately
		case <-ctx.Done():
			return

		case jobItem, ok := <-jobs:
			if !ok {
				return // jobs channel closed, get out
			}
			usr := User{}
			usr.RuToUser(&jobItem)

			// NOTE: to prevent Log from using mutex between go routines
			//       and slowing down process, comment out debug logging

			// log.Printf(
			// 	"Worker[%], Processing Job[%s:%s]",
			// 	index,
			// 	jobItem.Id.Name,
			// 	jobItem.Id.Value,
			// )

			if isSimLongProcess {
				time.Sleep(10 * time.Millisecond)
			}

			// send processed item back to results + handle context cancellation
			select {
			case results <- usr:
			case <-ctx.Done():
				return
			}
		}
	}
}
