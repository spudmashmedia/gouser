# Build and Run instructions

## Building the Go code 

```
> sh scripts/build-api
```

## Run

### Development
#### configs
|env|config|
|--|--|
|dev| [/config/gouser_api_config_dev.toml](../config/gouser_api_config_dev.toml)|

#### Command
```sh
> go build -o bin/gouserApi cmd/gouserApi/*.go
```
or

```sh
> go build -o bin/gouserApi cmd/gouserApi/*.go --env=dev
```

### Test
#### configs
|env|config|
|--|--|
|test| [/config/gouser_api_config_test.toml](../config/gouser_api_config_test.toml)|

#### Command
```sh
> go build -o bin/gouserApi cmd/gouserApi/*.go --env=test
```

### Production
#### configs
|env|config|
|--|--|
|prod| [/config/gouser_api_config.toml](../config/gouser_api_config.toml)|

#### Command
```sh
> go build -o bin/gouserApi cmd/gouserApi/*.go --env=prod
```

