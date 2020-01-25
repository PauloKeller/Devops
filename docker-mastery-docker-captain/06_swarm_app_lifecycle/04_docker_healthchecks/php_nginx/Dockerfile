FROM your-nginx-php-fpm-combo-image

HEALTHCHECK --interval=5s --timeout=3s \
CMD curl -f http://localhost/ping || exit 1