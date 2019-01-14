FROM scratch

ENV SERVICE_PORT 8000

EXPOSE $SERVICE_PORT

COPY k8s-introduction /

CMD ["/k8s-introduction"]

# build app for container
# k8s-introduction$ GOOS=linux CGO_ENABLED=0 go build

# build container
# k8s-introduction$ docker build -t k8s-introduction -f ./Dockerfile .

# run container with port forwarding
# k8s-introduction$ docker run -p 8000:8000 k8s-introduction  
