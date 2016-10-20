FROM php:5.6-apache
MAINTAINER Jetsada Machom jim@imjim.im

RUN apt-get update && \
    apt-get install -y curl \
    postgresql-server-dev-all && \
    docker-php-ext-install mysqli \
    pgsql && \
    curl -L https://www.adminer.org/latest-en.php > index.php && \
    apt-get remove -y postgresql-server-dev-all && \
    apt-get autoclean && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

