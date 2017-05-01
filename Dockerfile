FROM golang:1.8
Add gohbridge /app
ENTRYPOINT ["/goh"]
EXPOSE 8080
