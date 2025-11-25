# Dockerfile
FROM golang:1.24-alpine

# Directorio de trabajo dentro del contenedor
WORKDIR /app

# Instalar dependencias necesarias para PostgreSQL y compilación
RUN apk add --no-cache git gcc musl-dev

# Copiamos go.mod y go.sum primero para aprovechar caching
COPY go.mod go.sum ./
RUN go mod download

# Copiamos el resto del proyecto
COPY . .

# Compilamos la app
RUN go build -o kanban-backend ./cmd

# Puerto que exponemos
EXPOSE 8080

# Comando para correr la app
CMD ["./kanban-backend"]