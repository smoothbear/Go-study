# 1. Simple restful api by using gorilla/mux

### Gorilla/mux?

Gorilla/mux 패키지는 요청 라우터(Request Router)와 디스패쳐(dispatcher)가 구현되어 있는 패키지이다.

* [Gorilla/mux github page](https://github.com/gorilla/mux)

-----

### mux.NewRouter()?

mux 패키지의 가장 핵심이 되는 Router를 만든다.

------------

### 기본적인 Router 사용법

- Router 선언

```go
r := mux.NewRouter()
r.Host("www.example.com")
r.Host("{subdomain:[a-z]+}.example.com")
```

* URI prefix 정의

```go
r.PathPrefix("/products/")
```

* Http Method 정의

```go
r.Methods("GET", "POST")
```

* URL scheme 정의

```go
r.Schemes("https")
```

* Header Value 정의

```go
r.Headers("X-Requested-With", "XMLHttpRequest")
```

* Query Value 정의

```go
r.Queries("key", "value")
```

이것들을 기본적으로 사용한다면, 하나의 route를 만들 수 있다.

```go
r.HandleFunc("/products", ProductsHandler).
  Host("www.example.com").
  Methods("GET").
  Schemes("http")
```

