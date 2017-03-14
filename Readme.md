A stupidly simple example of an HTTP service in Go.

## Build & run

```bash
docker build -t golang-numberwang .
docker run -it --rm -p 8080:8080 --name golang-numberwang-1 golang-numberwang
```