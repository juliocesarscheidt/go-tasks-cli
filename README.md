# Go tasks CLI with Cobra

> Install

```bash
curl -L --url https://github.com/juliocesarscheidt/go-tasks-cli/archive/refs/tags/v1.0.0.tar.gz \
  --output go-tasks-cli-v1.0.0.tar.gz
tar -xzvf go-tasks-cli-v1.0.0.tar.gz
mv go-tasks-cli-1.0.0/bin/go-tasks /usr/local/bin/go-tasks && \
  chmod u+x /usr/local/bin/go-tasks
```

> Running locally

```bash
go build -ldflags="-s -w" -o ./bin/go-tasks .

./bin/go-tasks help

./bin/go-tasks --version
# go-tasks version: v1.0.0

./bin/go-tasks tasks
# Available Commands:
# get          get.

./bin/go-tasks tasks --help

./bin/go-tasks tasks create --name TESTE1
./bin/go-tasks tasks create --name TESTE2 --done true

./bin/go-tasks tasks list

./bin/go-tasks tasks get
# Error required flag(s) "name" not set

./bin/go-tasks tasks get --name
# Error flag needs an argument: --name

./bin/go-tasks tasks get --name TESTE
# Task Name :: TESTE
```
