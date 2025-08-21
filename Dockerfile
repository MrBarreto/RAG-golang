# --- Estágio 1: Build da Aplicação ---
# Usa uma imagem Go para compilar o código
FROM golang:1.25 AS builder

# Define o diretório de trabalho dentro do container
WORKDIR /app

# Copia o go.mod e go.sum para gerenciar as dependências
COPY go.mod go.sum ./

# Baixa as dependências do projeto
RUN go mod download

# Copia todos os arquivos da sua aplicação Go
COPY . .

# Compila a aplicação. O CGO_ENABLED=0 cria um executável estaticamente linkado.
# Isso garante que o executável não precise de bibliotecas C na imagem final.
# O -o myapp especifica o nome do executável
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o myapp .

# --- Estágio 2: Imagem Final ---
# Usa uma imagem minimalista ou "scratch" para a imagem de produção
FROM alpine:latest

# Define o diretório de trabalho na imagem final
WORKDIR /app

# Copia o executável compilado do estágio "builder" para a imagem final
COPY --from=builder /app/myapp .

# Expõe a porta que sua aplicação usará (ajuste se necessário)
EXPOSE 8080

# Define o comando para executar a aplicação
CMD ["./myapp"]