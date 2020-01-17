FROM postgres:11

WORKDIR /var/run

ADD build/crunchy-proxy ./
ADD tests/pgbench/* ./

CMD ["./run-all-tests"]
