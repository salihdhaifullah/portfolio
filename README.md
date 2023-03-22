### to build image
```sh
docker build -t portfolio .
```
### to inspect image
```sh
docker run -it --rm --entrypoint /bin/sh portfolio
```
### to run continuer
```sh
docker run -it --rm --name running-portfolio -p 8080:8080 portfolio
```
