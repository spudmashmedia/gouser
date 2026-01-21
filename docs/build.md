# Build and Run instructions

## Building the Go code 

### Using Shell Script
```sh
sh scripts/build-api
```

### Using Make

#### Build
```sh
make
```

```sh
make build
```

#### Run

Debug Mode
```sh
make run ARGS="--env=debug"
```

Dev Mode
```sh
make run ARGS="--env=dev"
```

Test Mode
```sh
make run ARGS="--env=test"
```

Production Mode
```sh
make run ARGS="--env=prod"
```

#### Cleanup
Will clean up the ./bin folder created by make
```sh
make clean
```

## Run

### Development
#### configs
|env|config|
|--|--|
|dev| [/config/gouser_api_config_dev.toml](../config/gouser_api_config_dev.toml)|

#### Command
```sh
go build -o bin/gouserApi cmd/gouserApi/*.go
```
or

```sh
go build -o bin/gouserApi cmd/gouserApi/*.go --env=dev
```

### Test
#### configs
|env|config|
|--|--|
|test| [/config/gouser_api_config_test.toml](../config/gouser_api_config_test.toml)|

#### Command
```sh
go build -o bin/gouserApi cmd/gouserApi/*.go --env=test
```

### Debug
This mode contains the /debug pprof profile that will emit telemetry for troubleshooting

#### configs
|env|config|
|--|--|
|test| [/config/gouser_api_config_debug.toml](../config/gouser_api_config_debug.toml)|

NOTE: Please pay attention to the profiler section of the config

#### Command
```sh
go build -o bin/gouserApi cmd/gouserApi/*.go --env=debug
```

### Production
#### configs
|env|config|
|--|--|
|prod| [/config/gouser_api_config.toml](../config/gouser_api_config.toml)|

#### Command
```sh
go build -o bin/gouserApi cmd/gouserApi/*.go --env=prod
```

