#!/bin/bash
dockerize -wait tcp://db:3306 -timeout 40s
php artisan key:generate
php artisan migrate
php-fpm