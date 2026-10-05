# clogger

A small Go HTTP middleware package for structured request logging.

## Installation

```bash
go get github.com/Akene-Uzezi/clogger
```

## Usage

```go
package main

import (
    "log/slog"
    "net/http"

    clogger "github.com/Akene-Uzezi/clogger"
)

func main() {
    logger := clogger.New(slog.Default())

    http.Handle("/", logger.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello"))
    })))

    http.ListenAndServe(":8080", nil)
}
```

## API

### `New(l *slog.Logger) *Logger`

Creates a new `Logger`. If `l` is nil, a default text handler writing to stdout is used.

### `Start(next http.Handler) http.Handler`

Wraps an HTTP handler with logging middleware. Logs method, URL, status code, and duration for each request.

### `Handler(h http.Handler) http.Handler`

Alias for `Start`.

## Log Levels

Log level is derived from the response status:

| Status Range | Level     |
|--------------|-----------|
| 500+         | `Error`   |
| 400+         | `Warn`    |
| 300+         | `Debug`   |
| <300         | `Info`    |

## License

MIT
