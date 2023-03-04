
```bash

go build -ldflags="-s -w" -o ./go-tasks .

./go-tasks help

./go-tasks --version
# go-tasks version: v1.0.0

./go-tasks tasks
# Available Commands:
# get          get.

./go-tasks tasks --help


./go-tasks tasks create --name TESTE1
./go-tasks tasks create --name TESTE2 --done true


./go-tasks tasks list


./go-tasks tasks get
# Error required flag(s) "name" not set

./go-tasks tasks get --name
# Error flag needs an argument: --name

./go-tasks tasks get --name TESTE
# Task Name :: TESTE

```
