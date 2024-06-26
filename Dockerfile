FROM golang:1.22.4-bookworm

WORKDIR /app



COPY https-certs ./https-certs

COPY go.* ./

RUN go mod download

COPY *.go ./

RUN go build -o /mzfaqiri

ENV PORT=443

EXPOSE 443

CMD [ "/mzfaqiri" ]
