FROM node:16-alpine as react_builder
WORKDIR /app
COPY client/package.json ./
RUN npm install
COPY client/ ./
RUN npm run build

FROM golang:1.17-alpine as go_builder
WORKDIR /server
COPY server/go.mod ./
# COPY server/go.mod server/go.sum ./
# RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server server/main.go

FROM alpine:latest
COPY --from=react_builder /app/dist ./client/dist
COPY --from=go_builder /server/server/main ./main
EXPOSE 8080
CMD ["./main"]
