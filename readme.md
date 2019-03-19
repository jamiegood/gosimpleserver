go docker setup:
https://medium.com/@rrgarciach/bootstrapping-a-go-application-with-docker-47f1d9071a2a

Go simple server:
https://itnext.io/building-a-web-http-server-with-go-6554029b4079


go run src/main.go

docker build . -t my-golang--app-image2
docker run -p 3000:3000 -it --rm --name my-golang--app-run2 my-golang--app-image2

https://ropenscilabs.github.io/r-docker-tutorial/04-Dockerhub.html