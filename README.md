# [sample-htmx-chi][repo]

Sample project using [htmx][htmx], [chi][chi] framework and [liquid][liquid]
[template][go-liquid] engine in [go][go].

Similar experiment to the others seen [here][benchmark].

## Requirements

- go 1.20 or newer
- [htmx 1.9][urn-pkg] or newer (cached in assets folder)

## How to build

```bash
go build
```

## How to run

```bash
go run main.go
```

Or

```bash
./sample-htmx-chi
```

## Noteworthy

- Go has a strong culture of avoid 3rd party libraries imprinted on its
  community. it does have, however, a rich set of libraries to ease the work.
- There is a default template engine already available in the go runtime. It
  does not, however, compare with more complex ones (i.e. no inheritance,
  control flow and so on), so we added one 

[repo]: https://github.com/sombriks/sample-htmx-chi
[htmx]: https://htmx.org
[chi]: https://go-chi.io
[liquid]: https://shopify.github.io/liquid/basics/introduction/
[go-liquid]: https://github.com/osteele/liquid
[go]: https://go.dev/
[urn-pkg]: https://unpkg.com/htmx.org@1.9.11/dist/htmx.min.js
[benchmark]: https://github.com/sombriks/node-vs-kotlin-k6-benchmark
