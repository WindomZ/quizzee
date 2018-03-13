FROM golang:1.9-alpine
MAINTAINER WindomZ "git.windomz@gmail.com"

# Copy the directory into the container.
RUN mkdir -p /go/src/github.com/WindomZ/quizzee/
ADD . /go/src/github.com/WindomZ/quizzee/
WORKDIR /go/src/github.com/WindomZ/quizzee/server/

# Download and install the required third party dependencies.
#RUN apk add --no-cache git \
#&& go-wrapper download \
#&& go-wrapper install \
#&& apk del git

RUN go-wrapper install

# Expose port
EXPOSE 8080

# Run server
CMD ["go-wrapper", "run"]
