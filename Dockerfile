FROM busybox
WORKDIR /pics
ADD ./static /pics/static
ADD ./pics_linux-amd64 /pics/run
CMD ["/pics/run"]
