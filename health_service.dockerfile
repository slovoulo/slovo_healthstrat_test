#build a tiny docker image
FROM golang:1.19-alpine as builder

RUN mkdir /app

#copy fromthe first broker image to this one

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o healthBrokerService ./cmd

RUN chmod +x /app/healthBrokerService


#build a tiny docker image
FROM alpine:latest

RUN mkdir /app

#copy fromthe first broker image to this one
COPY --from=builder app/healthBrokerService /app

# Expose the port on which the application listens
EXPOSE 8081

#Build the first docker image, create a  much smaller docker image then copy the executable from first to second smaller image
CMD [ "/app/healthBrokerService" ]