FROM golang:1.16 as builder

#
RUN mkdir -p $GOPATH/src/bitbucket.org/udevs/ur_go_user_service
WORKDIR $GOPATH/src/bitbucket.org/udevs/ur_go_user_service

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    make build && \
    mv ./bin/ur_go_user_service /

FROM alpine
COPY --from=builder ur_go_user_service .
RUN apk add --no-cache tzdata
ENV TZ Asia/Tashkent
ENTRYPOINT ["/ur_go_user_service"]
