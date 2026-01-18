# Testing Instructions

## Testing Go Logic
### Endpoints
|Task|Concurrent|Httpie endpoint|
|--|--|--|
|fetch 1 record |single threaded |http get localhost:8080/user |
|fetch 10 records |single threaded|http get localhost:8080/user count==10|
|fetch max 5000 records, process fetched dataset sequentially |single threaded |http get localhost:8080/user count==5000|
|fetch max 5000 records process fetched dataset concurrently |multi threaded |http get localhost:8080/user count==5000 concurrent==true|
|fetch max 5000 records process fetched dataset concurrently, simulate longer processing time for each record processed | multi threaded |http get localhost:8080/user count==5000 concurrent==true sim_long_pro==true|

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
TBA
