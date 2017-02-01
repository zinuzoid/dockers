FROM debian:jessie
MAINTAINER Jetsada Machom <zinuzoid@gmail.com>

RUN apt-get update
RUN apt-get install -y vim subversion build-essential libncurses5-dev zlib1g-dev gawk flex quilt git-core unzip libssl-dev wget python file

WORKDIR /
RUN git clone https://github.com/WereCatf/source.git

COPY .config /source/

ENV FORCE_UNSAFE_CONFIGURE 1

WORKDIR /source
RUN make help
RUN make tools/install
RUN make toolchain/install

CMD tail -f /dev/null

