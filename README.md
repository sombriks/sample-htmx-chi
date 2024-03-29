# [sample-htmx-chi][repo]

Sample project using [htmx][htmx], [chi][chi] framework and [liquid][liquid]
[template][go-liquid] engine in [go][go].

Similar experiment to the others seen [here][benchmark].

## Requirements

- go 1.20 or newer
- [chi][chi] 5.0 or newer
- [(go) liquid][go-liquid] 1.3 or newer
- [gorm][gorm] 1.25 or newer
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

Note that the execution point is important, the server must get the proper path
to the [assets][assets] and [templates][templates] folders.

## Noteworthy

- Go has a strong culture of avoid 3rd party libraries imprinted on its
  community. it does have, however, a rich set of libraries to ease the work.
- There is a default template engine already available in the go runtime. It
  does not, however, compare with more complex ones (i.e. no inheritance,
  control flow and so on), so we added a custom one. 
- To be as close as possible to the other samples, i've added an
  [ORM framework][gorm] so we get some goodies and some overhead.
- Several helpers that comes "for free" in other stacks need to be done
  "by hand" in go. some parameter converters and a few id checks for example. 

[repo]: https://github.com/sombriks/sample-htmx-chi
[htmx]: https://htmx.org
[chi]: https://go-chi.io
[liquid]: https://shopify.github.io/liquid/basics/introduction/
[go-liquid]: https://github.com/osteele/liquid
[gorm]: https://gorm.io/docs/index.html
[go]: https://go.dev/
[urn-pkg]: https://unpkg.com/htmx.org@1.9.11/dist/htmx.min.js
[benchmark]: https://github.com/sombriks/node-vs-kotlin-k6-benchmark
[assets]: ./app/assets
[templates]: ./app/templates
