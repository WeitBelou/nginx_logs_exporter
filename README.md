## nginx_logs_exporter

DISCLAIMER: This repo is in PoC state, please, do not use it.

Parses nginx access logs in json format and exports http metrics in prometheus format.

### Usage

#### Build
```bash
go build nginx_logs_exporter.go
```

#### Run
```bash
./nginx_logs_exporter -i ./examples/nginx.access.json -p 8989
```

Now metrics server is running and recording log events from `examples/nginx.access.json` (TODO anonymize real logs)


You can observe collected metrics at `127.0.0.1:8989/metrics`.

### Nginx logs config to cover all labels
```nginx
log_format json_combined escape=json
    '{'
      '"http_host":"$http_host",'
      '"uri":"$uri",'
      '"status":$status,'
      '"request_time":$request_time,'
      '"request_method":"$request_method",'
    '}'
```