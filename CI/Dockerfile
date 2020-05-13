FROM golang:alpine3.11 as build

WORKDIR /go/src/soma
#Copiando o fonte, compilado com a opção CGO_ENABLED=0 (CGO_ENABLED=0  go build -o ./build/main main.go)
COPY ./src/soma/ /go/src/soma

RUN go install .
#Executa o executável
ENTRYPOINT [ "./soma" ]

