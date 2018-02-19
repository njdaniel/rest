REST GO EXAMPLE

Docker Setup:

$ env GOOS=linux GOARCH=amd64 go build --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'

$ sudo docker build -t <tag/imagename:version> .

$ sudo docker run -d <tag/imagename:version>

$ sudo docker inspect <tag/imagename:version>

$ curl <container IP>
