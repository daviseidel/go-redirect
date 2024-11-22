# Imagem base
FROM golang:1.20 as builder

# Configura o diretório de trabalho
WORKDIR /app

# Copia o código para o container
COPY redirect.go .

# Compila o aplicativo
RUN go build -o redirector redirect.go

# Imagem final mínima
FROM alpine:latest

# Copia o binário compilado
COPY --from=builder /app/redirector /redirector

# Porta exposta
EXPOSE 8080

# Comando para iniciar o aplicativo
ENTRYPOINT ["/redirector"]
