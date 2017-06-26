FROM centos:7.1.1503

ENV REFRESHED_AT 2017-06-03

ARG PROGRAM_NAME="unknown"
ARG BUILD_VERSION=0.0.0
ARG BUILD_ITERATION=0

# --- YUM installs ------------------------------------------------------------

# Avoid "Error: libselinux conflicts with fakesystemd-1-17.el7.centos.noarch"
RUN yum -y swap fakesystemd systemd && \
    yum -y install systemd-devel

RUN yum -y update

# --- Install Go --------------------------------------------------------------

ENV GO_VERSION=1.8.3

# Install dependencies.
RUN yum -y install \
    git \
    tar \
    wget

# Install "go".
RUN wget https://storage.googleapis.com/golang/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local/ -xzf go${GO_VERSION}.linux-amd64.tar.gz

# --- Install Ruby 2.4.0 ------------------------------------------------------

# Install dependencies.
RUN yum -y install \
    curl \
    gcc \
    make \
    rpm-build \
    ruby-devel \
    which

# Install Ruby Version Manager (RVM)
RUN gpg --keyserver hkp://keys.gnupg.net --recv-keys 409B6B1796C275462A1703113804BB82D39DC0E3
RUN curl -L get.rvm.io | bash -s stable

# Install Ruby 2.4.0
ENV PATH /usr/local/rvm/gems/ruby-2.4.0/bin:/usr/local/rvm/gems/ruby-2.4.0@global/bin:/usr/local/rvm/rubies/ruby-2.4.0/bin:/usr/local/rvm/bin:$PATH
RUN rvm install 2.4.0

# --- Install Effing Package Manager (FPM) ------------------------------------

RUN gem install fpm --version 1.8.1

# --- Compile go program ------------------------------------------------------

ENV HOME="/root"
ENV GOPATH="${HOME}/gocode"
ENV PATH="${PATH}:/usr/local/go/bin:${GOPATH}/bin"
ENV GO_PACKAGE="github.com/docktermj/${PROGRAM_NAME}"

# Install dependencies.
RUN go get github.com/docopt/docopt-go
RUN go get github.com/hashicorp/go-plugin

# Copy local files from the Git repository.
COPY . ${GOPATH}/src/${GO_PACKAGE}

# Build plugins
RUN go install ${GO_PACKAGE}/plugin/greeter/greeter-english
RUN go install ${GO_PACKAGE}/plugin/greeter/greeter-german
RUN go install ${GO_PACKAGE}/plugin/greeter/greeter-italian
RUN go install ${GO_PACKAGE}/plugin/hello/hello-english
RUN go install ${GO_PACKAGE}/plugin/hello/hello-german
RUN go install ${GO_PACKAGE}/plugin/hello/hello-italian

# Build go program.
RUN go install \
    -ldflags "-X main.programName=${PROGRAM_NAME} -X main.buildVersion=${BUILD_VERSION} -X main.buildIteration=${BUILD_ITERATION}" \
    ${GO_PACKAGE}

# --- Package as RPM and DEB --------------------------------------------------

WORKDIR /output

# RPM package.
RUN fpm \
  --input-type dir \
  --output-type rpm \
  --name ${PROGRAM_NAME} \
  --version ${BUILD_VERSION} \
  --iteration ${BUILD_ITERATION} \
  /root/gocode/bin/=/usr/bin

# DEB package.
RUN fpm \
  --input-type dir \
  --output-type deb \
  --name ${PROGRAM_NAME} \
  --version ${BUILD_VERSION} \
  --iteration ${BUILD_ITERATION} \
  /root/gocode/bin/=/usr/bin

# --- Epilog ------------------------------------------------------------------

RUN yum clean all

CMD ["/bin/bash"]
