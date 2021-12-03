# crowdfunding-go

## Description
This project is created during the [Kemendikbud's Kampus Merdeka Program](https://kampusmerdeka.kemdikbud.go.id/) that the owner do on since August 2021. [Alterra Academy](https://academy.alterra.id/) is where the owner do the Independent Study (SI). Other than following a 3 hours daily live session, 3 hours daily self-learning and a daily tasks, participants are instructed to develop a RESTful API implementing all of the material that have been studied.

In short, the RESTful API is about a crowdfunding app where users can have help each other via campaign. The project is written in Go Programming Language and implements the Clean Architecture Code Structure.

Some of the technologies and libraries implemented on the projects are as follows:
* [Echo Labstack](echo.labstack.com): A Golang Web Framework
* [GORM](https://gorm.io/): A library for implementing ORM
* [tkanos/gonfig](https://github.com/tkanos/gonfig): Used for environment configurations
* [vektra/mockery](https://github.com/vektra/mockery): Used for creating mocks for unit testing
* [stretchr/testify](https://github.com/stretchr/testify): Used for unit testing
* [sonarqube](https://www.sonarqube.org): Used for static analysis
* [jaeger](http://jaegertracing.io): Used for tracing
* [prometheus](https://prometheus.io): Used for application metrics
* [grafana](https://grafana.com): Used for visualize data from jaegertracing & prometheus
* [traefik](https://traefik.io): Used for edge router & service proxy
* [gomail](https://github.com/go-gomail/gomail): Used for send mail  
* [midtrans](https://midtrans.com): Used for payment gateway
* [swarmpit](https://swarmpit.io): Used for infrastructure monitoring


The database management system used are both SQL and NoSQL:
* MySQL: The main database used are relational database, this is where the API stores the entity


The deployment of the project uses:
* Docker Swarm; and
* Digital Ocean

Continuous Integration and Continuous Deployment (CI/CD) using GitHub Actions is implemented to automate the deployment process.

## Clean Architecture
As previously mentioned, the project implements Clean Architecture. The four layers on the project are:
* Domain Layer
* Repository Layer
* Usecase Layer
* Delivery Layer

