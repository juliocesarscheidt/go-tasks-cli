# Go tasks CLI with Cobra

> Install

```bash
curl -L --url https://github.com/juliocesarscheidt/go-tasks-cli/archive/refs/tags/v1.0.0.tar.gz \
  --output go-tasks-cli-v1.0.0.tar.gz
tar -xzvf go-tasks-cli-v1.0.0.tar.gz
mv go-tasks-cli-1.0.0/bin/go-tasks /usr/local/bin/go-tasks && \
  chmod u+x /usr/local/bin/go-tasks
```

> Running

```bash
go-tasks help

go-tasks --version
# go-tasks version: v1.0.0

go-tasks tasks --help

go-tasks tasks create --name TEST1
go-tasks tasks create --name TEST2 --done true

go-tasks tasks list
+-------+-------+-------------------------------------+-------------------------------------+------------+
| NAME  | DONE  | CREATED AT                          | UPDATED AT                          | DELETED AT |
+-------+-------+-------------------------------------+-------------------------------------+------------+
| TEST2 | true  | 2023-03-04T14:09:19.81755573-03:00  | 2023-03-04T14:09:19.81755573-03:00  |            |
| TEST1 | false | 2023-03-04T14:09:17.671884679-03:00 | 2023-03-04T14:09:17.671884679-03:00 |            |
+-------+-------+-------------------------------------+-------------------------------------+------------+

go-tasks tasks get --name TEST1
+-------+-------+-------------------------------------+-------------------------------------+------------+
| NAME  | DONE  | CREATED AT                          | UPDATED AT                          | DELETED AT |
+-------+-------+-------------------------------------+-------------------------------------+------------+
| TEST1 | false | 2023-03-04T14:09:17.671884679-03:00 | 2023-03-04T14:09:17.671884679-03:00 |            |
+-------+-------+-------------------------------------+-------------------------------------+------------+
```

> Build locally

```bash
go build -ldflags="-s -w" -o go-tasks .
```

> Uninstall

```
rm -f /usr/local/bin/go-tasks
rm -f /tmp/go_tasks.db
```

