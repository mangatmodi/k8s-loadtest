k8s-loadtest
========
This project is a distributed load test framework, i.e. load can be generated from multiple generators in parallel. This project is based on  - [locust](https://locust.io/) and [boomer](https://github.com/myzhan/boomer/). The load is generated by slave (boomer), where as master(locust) will coordinate hatch-rate between the slaves. At anytime more slaves could be added and the hatch rate could be changed.

Being distributed in nature, both master and slave could be deployed on Kubernetes. This project enables one to write various slave jobs which can be enabled/disabled from the configuration. An example kubernetes deployment is also given as an example

:warning: _Should be executed only on staging context_

# Usage
GNU Make build system is used to build and deploy the framework. It is split in two apps - (i) master, (ii) slave

`tag=v1 app=master make COMMAND`: This will execute the given command for app master for docker image tagged as `v1`. By default `latest` image is used.

execute `make help` to get list of all commands which are - 

| Command        | Description    
| ------------- |-------------| 
|help              | prints this help.
|build             | build the docker image with latest and the provided tag
|push              | push the docker image with latest tag
|pushTag           | push the docker image with given tag
|composeUp         | compose up the given app
|apply             | apply config for given app for latest or the tagged image      
|delete            | delete config for the given app
|buildApply        | build docker image and deploy the latest version of the image

## How to create new task
- Write your task in slave/task/ in a separate go file
- Enter the task in slave/task/Tasks map 
- Build and push the docker image for slave