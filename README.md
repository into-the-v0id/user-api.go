# User API

Simple JSON API to manage Users

## About

This is a simple test-project to get some experience with go lang. Do not use this project in production!

## Setup

```bash
$ go run main.go
```

## Example

```bash
$ curl 'http://localhost:8080/users/'
[
    {
        "id": "01H31D1G832F40A1J6YCP4J924",
        "name": "max",
        "dateCreated": "2023-06-16T06:07:20.835411646Z",
        "dateUpdated": "2023-06-16T06:07:20.835411646Z"
    }
]
```

## License

Copyright (C) Oliver Amann

This project is licensed under the GNU Affero General Public License Version 3 (AGPL-3.0-only). Please see [LICENSE](./LICENSE) for more information.
