FROM golang:1.20

#RUN apt-get update -q \
#    && apt-get install --no-install-recommends \
#        git \
#    && apt-get clean

#RUN groupadd -g 1000 appuser \
#    && useradd -u 1000 -g appuser appuser

#USER appuser

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY . .

RUN go mod download

#RUN go build -o /gortener

EXPOSE 3000

#CMD ["/gortener"]
#CMD ["bash"]