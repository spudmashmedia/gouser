# Testing Instructions

## Testing Go Logic
### Endpoints
|Task|Concurrent|Httpie endpoint|
|--|--|--|
| health check| single threaded | htttp get localhost:8080/health|
|fetch 1 record |single threaded |http get localhost:8080/user |
|fetch 10 records |single threaded|http get localhost:8080/user count==10|
|fetch max 5000 records, process fetched dataset sequentially |single threaded |http get localhost:8080/user count==5000|
|fetch max 5000 records process fetched dataset concurrently |multi threaded |http get localhost:8080/user count==5000 concurrent==true|
|fetch max 5000 records process fetched dataset concurrently, simulate longer processing time for each record processed | multi threaded |http get localhost:8080/user count==5000 concurrent==true sim_long_pro==true|

## ALL TEST
### Run (-force run incase cached)
```sh
> go test -v ./... -count=1
```

## Integration Tests

Integration tests are located in /tests folder.

### Run

```sh
> go test ./tests/...
```

With Verbose
```sh
> go test -v ./tests/...
```

## Unit Tests

Unit tests are coupled to the file under test and located in the same folder.

### Run
```sh
> go test -v ./cmd/... ./internal/... ./pkg/...
```

## Profiling

Chi routers have middleware that will output a pprof profile for performance analysis. the **/debug** endpoint will be available for environments (--env):
- debug
- dev
- test

### How to setup

#### Terminal 1 - Start API

- env==debug
```sh
  go run cmd/gouserApi/*.go --env=debug
```

- env==dev
```sh
  go run cmd/gouserApi/*.go --env=dev
```

- env==test
```sh
  go run cmd/gouserApi/*.go --env=test
```

NOTE: Once started, a DEBUG message will appear in the logs:
"***WARNING:***: PPROF /debug endpoint is enabled"

#### Terminal 2 - Launch pprof profile generator

```sh
  go tool pprof http://localhost:8080/debug/pprof/profile
```

You can also set into background mode or create a daemon.
Once run, any HTTP traffic to the API triggered will trigger the middleware to collect api and system telemetry (cpu, memory, calls etc...)

NOTE: Wait for this terminal to return "Entering interactive mode" before proceeding with Terminal 3.

#### Terminal 3 - Launch pprof profile aggregator + Web UI
Running the follow command will aggregate the data and open a Web UI on localhost:9090 automatically

```sh
  go tool pprof -http:localhost:9090 http://localhost:8080/debug/pprof/profile
```

NOTE: if you get the error "Could not enable CPU profiling: cpu profiling already in use", just wait for Terminal 2 to enter Interactive mode.

