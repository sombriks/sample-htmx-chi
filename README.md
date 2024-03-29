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


## Benchmark results

The following results for the same test applied
[against the other versions][benchmark].

```bash
          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

     execution: local
        script: benchmark-koa.js
        output: -

     scenarios: (100.00%) 1 scenario, 10 max VUs, 1m0s max duration (incl. graceful stop):
              * default: 10 looping VUs for 30s (gracefulStop: 30s)


     ✓ 200 ok

     checks.........................: 100.00% ✓ 93442       ✗ 0    
     data_received..................: 377 MB  13 MB/s
     data_sent......................: 7.5 MB  249 kB/s
     http_req_blocked...............: avg=2.52µs  min=1.03µs   med=2.29µs  max=904.75µs p(90)=3.05µs  p(95)=3.45µs 
     http_req_connecting............: avg=26ns    min=0s       med=0s      max=310.49µs p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=3.11ms  min=757.8µs  med=2.99ms  max=10.38ms  p(90)=4.31ms  p(95)=4.79ms 
       { expected_response:true }...: avg=3.11ms  min=757.8µs  med=2.99ms  max=10.38ms  p(90)=4.31ms  p(95)=4.79ms 
     http_req_failed................: 0.00%   ✓ 0           ✗ 93442
     http_req_receiving.............: avg=41.33µs min=14.34µs  med=39.33µs max=3.16ms   p(90)=53.41µs p(95)=59.64µs
     http_req_sending...............: avg=10.99µs min=5.14µs   med=10.25µs max=782.44µs p(90)=14.9µs  p(95)=17.2µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=3.06ms  min=724.42µs med=2.93ms  max=10.33ms  p(90)=4.26ms  p(95)=4.73ms 
     http_reqs......................: 93442   3114.471413/s
     iteration_duration.............: avg=3.19ms  min=823.97µs med=3.07ms  max=10.48ms  p(90)=4.4ms   p(95)=4.87ms 
     iterations.....................: 93442   3114.471413/s
     vus............................: 10      min=10        max=10 
     vus_max........................: 10      min=10        max=10 


running (0m30.0s), 00/10 VUs, 93442 complete and 0 interrupted iterations
default ✓ [======================================] 10 VUs  30s
```

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
