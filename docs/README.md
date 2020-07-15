# API

## Package

### getAllPackages

> GET All packages information

`GET /api/v2/packages`

```json
[
  {
    "name":"frp",
    "subname":["frpc","frps"],
    "version":"0.33.0-1",
    "users":["a-wing"],
    "log":{
      "1586079115":{"duration":106,"version":"0.32.1-2","status":"successful"},
      "1587980666":{"duration":104,"version":"0.33.0-1","status":"successful"}
    }
  },
  {
    "name":"wps-office-cn",
    "version":"11.1.0.9604-1",
    "users":["Universebenzene","MarvelousBlack"],
    "log":{
      "1589106405":{"duration":4,"version":"None-None","status":"failed"},
      "1591733260":{"duration":415,"version":"11.1.0.9522-1","status":"successful"},
      "1594329341":{"duration":611,"version":"11.1.0.9604-1","status":"successful"}
    }
  }
]
```

### getPackage

`GET /api/v2/packages/{name}`

Field | Type | Description
----- | ---- | -----------
name      | string   | pkgbase
subname   | []string | Opt: pkgname
version   | string   | version
users     | []string | Maintainer
{timestamp} | string | Opt
{timestamp}.duration  | uint | unit: `s`
{timestamp}.version | string | version
{timestamp}.status  | string | `successful`, `failed`, `skiped`

### getBuildLog

`GET /api/v2/packages/{name}/logs/{timestamp}`

## User

### getAllUsers

`GET /api/v2/users`

```json
[
  {
    "name":"Dr-Incognito",
    "packages":["v2ray-desktop"]
  },
  {
    "name":"cuihaoleo",
    "packages":["android-studio","intel-opencl-runtime","qtwebkit","intel-opencl-sdk","fcitx-sogoupinyin","tinc-pre"]
  }
]
```

### getUser

`GET /api/v2/users/{name}`

## Webhooks

### sync

`POST /api/v2/webhooks/sync`

