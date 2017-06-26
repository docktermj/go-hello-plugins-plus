# go-hello-plugins-plus

Build `go-hello-plugins-plus-M.m.P-I.x86_64.rpm`
and   `go-hello-plugins-plus_M.m.P-I_amd64.deb`
where "M.m.P-I" is Major.minor.Patch-Iteration.

## Usage

A simple "hello world" style program.
The purpose of the repository it to show how to:

1. Build go executable locally
1. Build go executable via Docker
1. Build RPM / DEB installation via Docker.
1. Show how to use HashiCorp's go-plugin

### Invocation

```console
go-hello-plugins-plus
```

## Development

### Dependencies

#### Set environment variables

```console
export GOPATH="${HOME}/go"
export PATH="${PATH}:${GOPATH}/bin:/usr/local/go/bin"
export PROJECT_DIR=${GOPATH}/src/github.com/docktermj
```

#### Download project

```console
mkdir -p ${PROJECT_DIR}
cd ${PROJECT_DIR}
git clone git@github.com:docktermj/go-hello-plugins-plus.git
```

#### Download dependencies

```console
go get github.com/docopt/docopt-go
go get github.com/hashicorp/go-plugin
```

### Build

#### Local build

Build plugins

```console
go install github.com/docktermj/go-hello-plugins-plus/plugin/greeter/greeter-english
go install github.com/docktermj/go-hello-plugins-plus/plugin/greeter/greeter-german
go install github.com/docktermj/go-hello-plugins-plus/plugin/greeter/greeter-italian

go install github.com/docktermj/go-hello-plugins-plus/plugin/hello/hello-english
go install github.com/docktermj/go-hello-plugins-plus/plugin/hello/hello-german
go install github.com/docktermj/go-hello-plugins-plus/plugin/hello/hello-italian
```

Build program

```console
go install github.com/docktermj/go-hello-plugins-plus
```

The results will be in the `${GOPATH}/bin` directory.

#### Docker build

```console
cd ${PROJECT_DIR}/go-hello-plugins-plus
make build
```

The results will be in the `.../target` directory.

### Install

#### RPM-based

Example distributions: openSUSE, Fedora, CentOS, Mandrake

##### RPM Install

Example:

```console
sudo rpm -ivh go-hello-plugins-plus-M.m.P-I.x86_64.rpm
```

##### RPM Update

Example: 

```console
sudo rpm -Uvh go-hello-plugins-plus-M.m.P-I.x86_64.rpm
```

#### Debian

Example distributions: Ubuntu

##### Debian Install / Update

Example:

```console
sudo dpkg -i go-hello-plugins-plus_M.m.P-I_amd64.deb
```

### Cleanup

```console
make clean
```
