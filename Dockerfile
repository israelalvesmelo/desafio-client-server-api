# Usar uma imagem base do Go
FROM golang:1.22-alpine
# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos go.mod e go.sum e baixar as dependências
COPY go.mod .
COPY go.sum .

# Copiar o código fonte
COPY server .

RUN go mod tidy


# Compilar o aplicativo
RUN go build -o main .

# Executar o aplicativo
CMD ["./main"]