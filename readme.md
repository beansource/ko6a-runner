standalone k6 test runner service for ko6a

a go service that uses
- https://github.com/gin-gonic/gin
- https://github.com/hibiken/asynq

prerequisites
- [go](https://go.dev)

development
```bash
git clone https://github.com/beansource/ko6a-runner
```

- install [task](https://taskfile.dev/installation/)
- run server

```bash
task run
```