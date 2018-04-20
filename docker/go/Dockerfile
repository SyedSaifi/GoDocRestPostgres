FROM varikin/golang-glide-alpine

WORKDIR /home/syed/work/Repository/Go/src/GoDocRestPostgres

RUN ls -la

# copy source code to docker container
COPY / /go/src/GoDocRestPostgres

# Make this as working directory
WORKDIR /go/src/GoDocRestPostgres

RUN ls -la

#install vendor dependencies using glide install
RUN glide install

# build go binaries
RUN go build 

EXPOSE 6000

#CMD ["go", "run", "main.go"]