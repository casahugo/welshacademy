FROM golang:1.21

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
