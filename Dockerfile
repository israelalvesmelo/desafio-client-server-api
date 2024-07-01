# Usar uma imagem base do Go
FROM golang:1.22-alpine
# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos go.mod e go.sum e baixar as dependências
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copiar o código fonte
COPY server .

# Compilar o aplicativo
RUN go build -o main .

# Executar o aplicativo
CMD ["./main"]