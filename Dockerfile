FROM golang:1.21

RUN set -xe && apt update
# Install requirement
RUN apt install -y \
	build-essential \
        psmisc \
        bash \
        bash-completion \
        gcc \
        libc-dev \
        curl \
        openssl \
        git \
        make \
        grep \
        jq \
        autoconf \
        automake \
        libtool \
        gettext \
        g++ \
        texinfo \
    ;

# Fswatch
RUN wget https://github.com/emcrisostomo/fswatch/releases/download/1.14.0/fswatch-1.14.0.tar.gz \
    && tar -xvzf fswatch-1.14.0.tar.gz \
    && cd fswatch-1.14.0/ \
    && ./configure \
    && make \
    && make install;

RUN go install github.com/google/wire/cmd/wire@latest

RUN ldconfig

WORKDIR /var/www/html
