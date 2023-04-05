# DEVOPS demo project «Build, Ship, Run» 

Software development lifecycle

## 1 - Planning

Choose technologies with which we will work. The goal is dynamic display of content.

![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white) ![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white) ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

1. SVG animation (library [Vivus](https://maxwellito.github.io/vivus/) );
2. Web server [Golang](https://go.dev/)
3.  [Docker](https://docs.docker.com/) - distribution of the product in the form of a container image for further deployment on [Kubernetes](https://kubernetes.io/) clusters.

## 2 - Coding 

Added [svg image](./html/frame1.svg),[web server](./src/main.go) implementation.

## 3 - Build

```bash
go build -o bin/app src/main.go
``` 

"go build" - compiles the code, the result of which will be placed according to the instructions of the "-o" option - the "bin/app " directory, which does not yet exist. The last argument "src/main.go" - we specify where exactly our code is located. &#10145; As a result, we get a binary code file.

## 4 - Manual QA
run code

```bash
bin/app
``` 
open web page in browser  http://127.0.0.1:8082/


## 5 - Delivery ( container image )
Added [Dockerfile](./Dockerfile) (read comments in file  &#128196;).

Move in 'src' directoty and generate a dependency file for our project
```bash
go mode init devops-types
``` 
move in root directory devops-types and start the process of building the container image with the command:

```bash
docker build .
``` 
docker returns a file archive that includes the server binary, content, and metadata.

 Image example: sha256:0645c9bceac82854187010edca7d2ce598a8e7d0a6591c41803a5fda665e3d70 


&#10067;&#10067; How to run a process in an isolated environment?&#10067;&#10067;

```bash
docker run -p 8082:8082 sha256:0645c9bceac82854187010edca7d2ce598a8e7d0a6591c41803a5fda665e3d70 
``` 
Open in browser http://127.0.0.1:8082/ and check how is it working.
At first glance, nothing has changed, but the process was completed in an isolated environment.


### Distribution of software.
Let's make the product publicly available. &#128483;

```bash
docker tag 7440997f4474 nalkav/app:v1.0.0
docker push natalkav/app:v1.0.0
``` 
check here https://hub.docker.com/


Create a Kubernetes cluster for the application.
```bash
k3d cluster create devops-types
``` 
Сreate a Kubernetes deployment with a reference to the project in Docker Hub. Kubernetes will extract the image and deploy the container.

```bash
kubectl create deploy devops-types --image natalkav/app:v1.0.0
``` 
Let's do local port forwarding to the container and check how it works.

```bash
kubectl get po -w
kubectl port-forward deploy/devops-types 8082
``` 

Changes or features:
1. Something has been changed.
2. Let's start the build process and upgrade the application version.
```bash
docker build . -t natalkav/app:v1.0.1
``` 
3. Push new version on Docker hub.
```bash
docker push docker.io/natalkav/app:v1.0.1
``` 
4. Zero downtime deployment
```bash
kubectl get deploy devops-types -o wide
kubectl get po -w
kubectl set image deploy devops-types app=natalkav/app:v1.0.1
``` 
5. Check new version in browser:
```bash
kubectl port-forward deploy/devops-types 8082
``` 
http://127.0.0.1:8082/ 