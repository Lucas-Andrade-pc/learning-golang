FROM golang:1.23.3 AS build

ENV CGO_ENABLED=1

WORKDIR /project/

COPY go.* /project/

RUN go mod download

COPY . /project/

RUN GOOS=linux go build -o /app -a -ldflags '-linkmode external -extldflags "-static"' .

FROM scratch

WORKDIR /api

COPY api.db /api/

COPY --from=build /app /api/

ENTRYPOINT [ "/api/app" ]

