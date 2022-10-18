FROM golang:1.18
WORKDIR /app
COPY . .
COPY ./go.* .
RUN go mod download
COPY . .

EXPOSE 8000

ENTRYPOINT [ "make", "run" ]
