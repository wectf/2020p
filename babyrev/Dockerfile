FROM php:7-apache

RUN docker-php-source extract
RUN mkdir /usr/src/php/ext/babyrev
COPY babyrev.c /usr/src/php/ext/babyrev/babyrev.c
COPY config.m4 /usr/src/php/ext/babyrev/config.m4
COPY php_babyrev.h /usr/src/php/ext/babyrev/php_babyrev.h

RUN docker-php-ext-install babyrev && docker-php-source delete
COPY index.php /var/www/html/
RUN chmod 755 /var/www/html/index.php
CMD ["apache2-foreground"]
