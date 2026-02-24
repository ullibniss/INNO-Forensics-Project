FROM golang:1.22.4-bullseye as builder

RUN mkdir /mac_identifier

WORKDIR /mac_identifier

COPY ./ /mac_identifier

RUN go mod tidy

RUN cd ./cmd && \
    go build -o ../build/mac_identifier

FROM golang:1.22.4-bullseye as runner

RUN mkdir /app

WORKDIR /app

COPY --from=builder /mac_identifier/build/mac_identifier ./

COPY resources/ ./resources/
COPY mac_vendor_map ./mac_vendor_map

ENV VENDOR_MAP_PATH=/app/mac_vendor_map

CMD ["./mac_identifier"]