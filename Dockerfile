FROM golang:1.19
ENV GOPROXY="https://goproxy.cn,direct"

WORKDIR /root/toolkit
COPY . .
RUN go build -o app

EXPOSE 4396

CMD ["./app"]