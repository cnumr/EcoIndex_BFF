FROM golang:1.24 AS build

WORKDIR /app
ADD ./ ./
RUN go mod download
RUN go build -o /ecoindex-bff


FROM gcr.io/distroless/base-debian12

WORKDIR /
COPY --from=build /ecoindex-bff /ecoindex-bff
EXPOSE 3000
USER nonroot:nonroot
ENTRYPOINT ["/ecoindex-bff"]
