FROM php:5.6-apache
MAINTAINER Jetsada Machom jim@imjim.im

RUN apt-get update && \
    apt-get install -y curl && \
    docker-php-ext-install mysqli && \
    curl -L https://www.adminer.org/latest-mysql-en.php > index.php && \
    apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

