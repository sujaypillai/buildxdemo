FROM golang AS build
COPY hello.go .
RUN go build hello.go

FROM scratch
COPY --from=build /go/hello /hello
CMD ["/hello"]

