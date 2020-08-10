# kiss2ugo


kiss2u V2

## Dev run

```sh
go run -tags=dev ./main.go
```

## Build

* golang >= 1.12.x

```sh
go generate

go build
```

## Run

```sh
./kiss2u
```

## Environment Variables

Variable Name | Description                                      | Default Value
------------- | ------------------------------------------------ | ------------------------
`LISTEN_ADDR` | Address to listen                                | `127.0.0.1:22333`
`PORT`        | Override `LISTEN_ADDR` to `0.0.0.0:$PORT` (PaaS) | None
`LILAC_LOG`   | Lilac log path                                   | `~/.lilac`
`LILAC_REPO`  | Lilac repo path                                  | `~/Code/archlinuxcn/repo`
`REPO_NAME`   | Repo name                                        | `archlinuxcn`

## API

[Document](./api/api.go)

For archlinux
- `name` is `pkgbase`
- `subname` is `pkgname`

