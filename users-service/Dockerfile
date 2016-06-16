# Use Node v4 as the base image.
FROM golang:1.6-alpine

# Add everything in the current directory to our image, in the 'app' folder.
ADD . /go/app

# Install dependencies
RUN go install /go/app

# Expose our server port.
EXPOSE 3000

# Run our app.
ENTRYPOINT ["/go/bin/app"]
