FROM alpine:3.12 as build

RUN apk --no-cache add go

COPY . /go/src/github.com/safecornerscoffee/employee
WORKDIR /go/src/github.com/safecornerscoffee/employee

RUN go build

FROM alpine:3.12

COPY --from=build /go/src/github.com/safecornerscoffee/employee/employee /usr/bin/employee

CMD [ "/usr/bin/employee" ]
