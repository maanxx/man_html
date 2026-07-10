# Templ a-h

```bash

- Bước 1: go install github.com/a-h/templ/cmd/templ@latest
- Bước 2: go get github.com/a-h/templ
- Bước 3: mkdir folder -> go mod init -> go get github.com/a-h/templ
- Bước 4: go mod tidy

```

# Go air

## Add những điều cần thiết vào
```bash
include_ext = ["go", "tpl", "tmpl", "html", "templ"]

[build]
cmd = "cmd /c \"templ generate && go build -o ./tmp/main.exe .\""

[build.windows]
cmd = "cmd /c \"templ generate && go build -o ./tmp/main.exe .\""


exclude_regex = ["_test\\.go", ".*_templ\\.go"]


[build]
pre_cmd = ["cmd /c \"taskkill /F /IM main.exe >nul 2>&1 || exit 0\""]

[build.windows]
pre_cmd = ["cmd /c \"taskkill /F /IM main.exe >nul 2>&1 || exit 0\""]


kill_delay = "500ms"
```
## https://github.com/air-verse/air

```bash

- Bước 1: go install github.com/air-verse/air@latest

- Bước 2: go get -tool github.com/air-verse/air@latest

- Check: "go tool air -v"
```

## Cấu trúc a templ file

```bash

package main

    templ hello(name string) {
        <div>Hello, { name }</div>
    }

THEN
templ generate

    package main

    import (
        "context"
        "os"
    )

    func main() {
        component := hello("John")
        component.Render(context.Background(), os.Stdout)
    }

THEN
    go run .
```

# Git ignore

```bash

    *_templ.go
```
