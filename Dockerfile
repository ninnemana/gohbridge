FROM scratch
COPY gohbridge /goh
ENTRYPOINT ["/goh"]
EXPOSE 8080
