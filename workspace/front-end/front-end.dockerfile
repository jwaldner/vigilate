FROM alpine:latest

RUN mkdir /app

ADD ipe /app/ipe
ADD migrations /app/migrations
ADD views /app/views
ADD static /app/static
COPY frontApp /app

CMD [ "/app/frontApp"]







