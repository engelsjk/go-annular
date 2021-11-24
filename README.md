# go-annular

```go-annular``` is a Go port of my [node-annulus-bot](https://github.com/engelsjk/node-annulus-bot). This version updates a few things to make the generative art a little more interesting. It can save an .svg file or run as a simple web server.

```bash
go run cmd/image/main.go
// annular.svg saved
```

```bash
go run cmd/web/main.go
// listening at http://localhost:2003/svg and http://localhost:2003/png
```

![](images/1637548980.png)