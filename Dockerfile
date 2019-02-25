
FROM golang:1.11.1-alpine3.8 AS build-env
ADD . /src
RUN cd /src && CGO_ENABLED=0 go build -o app

FROM golang:1.11.1-alpine3.8
WORKDIR /app
COPY --from=build-env /src/app /app/
COPY --from=build-env /src/templates /app/templates
COPY --from=build-env /src/assets /app/assets
ENTRYPOINT ./app
