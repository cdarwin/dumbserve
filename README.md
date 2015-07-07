# dumbserve

`dumbserve` is a really dumb web server written in the Go programming language. It renders static html files from paths passed to it as arguments. You may also specify a port for it to run on. It does not fork or run in the background on it's own. Should you require that `dumbserve` run as a daemon, I suggest using [daemontools](http://cr.yp.to/daemontools.html).

## Requirements

You must have a working [Go build system](http://golang.org/doc/install) available to use `dumbserve`.

## Installation

```shell
go get github.com/cdarwin/dumbserve
go intsall github.com/cdarwin/dumbserve
```

## Usage

```shell
dumbserve -p 9090 -webroot /var/www -altroot /var/alt
```

## Webhook

`dumbserve` makes some assumptions. One of which is that you're serving files which are under version control by [Git](http://git-scm.com/). If `dumbserve` is accessible from the internet, you can use the POST to the `/github` enpoint to update the local copy of your repo. Check out more on [GitHub Post-Receive Hooks](https://help.github.com/articles/post-receive-hooks)
