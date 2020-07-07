# kiss2ugo


kiss2u V2

## Build

* golang >= 1.11.x

```sh
go build
```

## Run

```sh
./kiss2u -migrate

./kiss2u
```

## Environment Variables

Variable Name        | Description                                      | Default Value
-------------------- | ------------------------------------------------ | ---------------------------------------
`DATABASE_URL`       | LevelDB path                                     | `leveldb`
`LISTEN_ADDR`        | Address to listen                                | `127.0.0.1:22333`
`PORT`               | Override `LISTEN_ADDR` to `0.0.0.0:$PORT` (PaaS) | None
`LILAC_LOG`          | lilac log path                                   | `/home/lilydjwg/.lilac`
`LILAC_REPO`         | lilac repo path                                  | `/data/archgitrepo-webhook/archlinuxcn`

## API

[Document](./api/api.go)

For archlinux
- `name` is `pkgbase`
- `subname` is `pkgname`

