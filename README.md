Pics
---

Use this project to share a slideshow of pictures. Pictures can also be uploaded
by visiting /add. docker-compose will mount static/pics as the directory of the
pictures for the slideshow.

It uses docker to compile the binaries and the main Dockerfile adds the linux
binary to the busybox image to create an extremely small final image

Building
---

```bash
go build -o pics_linux-amd64 -tags netgo *.go
# Or
./script/build
```
> `-tags netgo` will help you achieve static binaries :)

Running
---

```bash
./pics_linux-amd64
docker run --rm -ti pdxjohnny/pics
```

Changing The Name
---

```bash
./script/change-name $GITHUB_USERNAME $PROJECT_NAME
```


- John Andersen
