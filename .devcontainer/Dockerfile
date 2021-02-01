FROM golang

ARG USER=bfreire
ARG UID=1000
ARG GID=1000

RUN groupadd -g $GID $USER \
    && useradd -u $UID -g $GID $USER
