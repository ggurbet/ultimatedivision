# BUILDER Image. Used to download all dependenices, etc
FROM golang:1.17.4-alpine3.15 as currencysigner_builder
# Changing root directory
WORKDIR /app
# Copy all files to root directory
COPY . .
# Collect and download dependances
RUN go mod vendor
# Building application
RUN CGO_ENABLED=0 go build -mod vendor -o main ./cmd/currencysigner/main.go

# Result image
FROM alpine:3.15.4

# Volume directorys
ARG APP_DATA_DIR=/data
# Creating volume directory
RUN mkdir -p ${APP_DATA_DIR}
# Creating volumes
VOLUME ["${APP_DATA_DIR}"]

# Copy executable file (builded application) from builder to root directory
COPY --from=currencysigner_builder /app/main .

# Builded application running with config directory as argument
ENTRYPOINT ["/main", "run", "--config=./config"]