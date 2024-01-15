FROM golang:1.20-alpine

RUN apk update && apk add git

ENV GO111MODULE=on

ENV GOPATH /go
ENV PATH $GOPATH/bin:/user/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/go-fiber-postgres

COPY . .

RUN go mod init github.com/akers1023

# Add missing dependencies
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/go-playground/validator/v10
RUN go get github.com/gofiber/fiber/v2
RUN go get github.com/google/uuid
RUN go get github.com/joho/godotenv
RUN go get golang.org/x/crypto/bcrypt
RUN go get gorm.io/driver/postgres
RUN go get gorm.io/gorm


RUN GOOS=linux go build -o app

ENTRYPOINT [ "./app" ]

EXPOSE 3000