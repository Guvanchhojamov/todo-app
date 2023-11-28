FROM go:1.21
LABEL authors="guvanch"

WORKDIR /usr/src/app

COPY . .

RUN go mod tidy
