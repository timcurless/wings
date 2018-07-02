FROM scratch
EXPOSE 8080
ENTRYPOINT ["/wings"]
COPY ./bin/ /