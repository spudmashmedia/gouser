FROM golang:1.26.2-trixie AS build

# 1 - Update base image
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    zip \
    unzip \
    tar \
    ca-certificates \
    cmake \
    && rm -rf /var/lib/apt/lists/*

# 2 - Setup Vcpkg
WORKDIR /usr

# 3 - Prep build directory and get vcpkg packages
WORKDIR /usr/gouser
COPY . .

# 4 - Build Project
# RUN chmod +x ./scripts/build-api.sh && ./scripts/build-api.sh
RUN make build

# 5 - Run Unit Test
RUN make test

# Use this for small final size
FROM gcr.io/distroless/cc-debian13 AS app

# Use this for debugging
# FROM debian:trixie-slim AS app

WORKDIR /app
COPY --from=build /usr/gouser/bin/gouser .
COPY --from=build /usr/gouser/config/gouser_api_config.toml .
EXPOSE 8080
ENTRYPOINT ["./gouser", "--env=prod"]
