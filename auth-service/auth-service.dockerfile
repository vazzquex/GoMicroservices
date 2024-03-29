#base go image

# FROM golang:1.20-alpine as builder

# RUN mkdir /app

# COPY . /app 

# WORKDIR /app

# RUN CGP_ENABLE=0 go build -o brokerApp ./cmd/api

# RUN chmod +x /app/brokerApp

# build a tiny docker image

FROM alpine:latest

RUN mkdir /app

#COPY --from=builder /app/brokerApp /app
COPY authApp /app
COPY db/users.sql /app/db/users.sql


CMD ["/app/authApp"]