# http2log

[![http2log](https://img.shields.io/badge/GitHub-http2log-blueviolet?style=flat-square&logo=github)](https://github.com/fernvenue/http2log)
[![http2log](https://img.shields.io/badge/GitLab-http2log-orange?style=flat-square&logo=gitlab)](https://gitlab.com/fernvenue/http2log)

A simple tool to write log file over HTTP protocol.

## Features

- Write file over HTTP protocol;
- HTTP header `X-API-KEY` authentication;

## Flags

- `-u`: URL path of log interface, will use `/http2log` by default;
- `-k`: HTTP header 'X-API-KEY' content, will read the environment variable if not set;
- `-f`: Path to log file, will use `/var/log/http2log.log` by default;
- `-a`: Bind address, will use `127.0.0.1` by default;
- `-p`: Bind port, will use TCP port `8283` by default;
