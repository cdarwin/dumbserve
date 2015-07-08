# dumbserve

`dumbserve` is a really dumb web server written in the Go programming language. It renders static html files from paths passed to it as arguments. You may also specify a port for it to run on. It does not fork or run in the background on it's own. Should you require that `dumbserve` run as a daemon, I suggest using [daemontools](http://cr.yp.to/daemontools.html).

## Installation

### Native

If you plam to use this tool natively, you must have a working 
[Go build system](http://golang.org/doc/install) available.

```shell
luser@lolcathost:~$ go get github.com/cdarwin/dumbserve
luser@lolcathost:~$ go intsall github.com/cdarwin/dumbserve
```

### Docker

If you plan to use this tool in a Docker environment, you just need to pull
the latest image from the hub.

```shell
luser@lolcathost:~$ docker pull cdarwin/dumbserve
```

## Usage

### Native

To run the service natively, you may set variables through command line
arguments or environment variables.

```shell
luser@lolcathost:~$ # as env vars
luser@lolcathost:~$ export PORT=9090
luser@lolcathost:~$ export WEBROOT=/var/www
luser@lolcathost:~$ export ALTROOT=/var/alt
luser@lolcathost:~$ dumbserve
luser@lolcathost:~$ # as cmd line args
luser@lolcathost:~$ dumbserve -p 9090 -webroot /var/www -altroot /var/alt
```

### Docker

Running the service as a Docker container is only slightly differnt.

```shell
luser@lolcathost:~$ # as env vars
luser@lolcathost:~$ docker run -d -e WEBROOT=/var/www -v /home/luser/www:/var/www -p 8080:8080 cdarwin/dumbserve --
luser@lolcathost:~$ # as cmd line args
luser@lolcathost:~$ docker run -d -v /home/luser/www:/var/www -p 8080:8080 cdarwin/dumbserve -webroot /var/www
```

## Webhook

`dumbserve` makes some assumptions. One of which is that you're serving files which are under version control by [Git](http://git-scm.com/). If `dumbserve` is accessible from the internet, you can use the POST to the `/github` enpoint to update the local copy of your repo. Check out more on [GitHub Post-Receive Hooks](https://help.github.com/articles/post-receive-hooks)
