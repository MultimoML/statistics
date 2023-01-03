FROM golang:alpine AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init

RUN CGO_ENABLED=0 go build -o /stats main.go

FROM gcr.io/distroless/static-debian11:latest

COPY --from=build /stats /stats

ENV PORT=6004
EXPOSE $PORT

ENTRYPOINT ["/stats"]
