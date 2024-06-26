FROM golang:1.22.4-bookworm

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY *.go ./

RUN go build -o /mzfaqiri

ENV PORT=80

EXPOSE 80

CMD [ "/mzfaqiri" ]
