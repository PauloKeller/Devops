FROM postgres

# specify real user with -U to prevent errors in log

HEALTHCHECK --interval=5s --timeout=3s \
CMD pg_isready -U postgres || exit 1