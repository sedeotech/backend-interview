
<p align="center">
  <h1 align="center">
  Sedeo
  </h1>
</p>

<p align="center">
  <b>
  Backend - Take Home Test
  </b><br>
</p>

Welcome and thank you for taking the time to do the take-home test for our backend engineer position.

The assignment consists of two parts :
* Letting you finish a basic backend service. 
* Deploy it on k8s.

We've created for you a simple app that handles carts and items using a GRPC interface. This service is supposed to persist data in a PostgreSQL DB. We provide the GRPC layer, it's up to you to define services and persistence layers.

You are free to improve the app in any way you see fit. **The goal is not to provide you with a list of tasks but to allow you to showcase the skills you think matter.**

Regarding the deployment, you are free to deploy it where you want. (AWS, GCP, locally) with the tools of your choice but we do like IaC (ie : Terraform, Pulumi, ...) and are big fans of Helm to manage changes.

---

There are no time requirements for the assessment but we think you should be able to produce enough material to review in at most 3 hours. Please do not spend more than 5 hours on this task, as this would not be respectful of your time.

If there were features/architecturing you wanted to implement but were not doable in a timely manner we're still curious to know your train of thought and how you would have done it. Simply add them in the readme at the end.

---
**PLEASE FOLLOW THE INSTRUCTIONS BELOW:**

- Clone this Repo, please do not Fork
- Create a private repo and invite [Aurelien](https://github.com/aureliensibiril)
- Make your changes per the requirements of this test and push your changes to your new repo. A summary of your changes is expected in commits.
- Once you are done with the test. Please make sure:
    - All changes have been pushed to your repo.
    - **Please message us with a link to the new Github repository and instructions on how to install / run the application.**

---

## Some guidelines

* You don't have to implement all functions in all services.
* If you think we forgot something, it's not on purpose.
* We try to asses if you work with some clarity in your mind, not if you can produce 2k lines of code in 1 hour. Quality over quantity. 
* Use any tools, event chatGPT if you think it's relevant.
* Our point of view is not absolute, if you think there is a better way, please do (and explain us later on 😁)

## Here are some ideas if you don't know where to start

1) We use DDD as our codebase architecture
2) Some testing ?
3) What about users auth / permissions ?
4) Tracing and instrumentation ?


## Getting Started

### How to start a local stack

__Prerequisite__ : 
* run docker on your computer first

__Then__ : 
In the backend-interview folder, just run :

`make docker.start.local`

It will start all the components needed to run the interview backend :

### Services :

* service GRPC backend : localhost:9111

### Databases :

* Postgres admin : localhost:5050
	* password : admin \
* PostgresDB : localhost:5432
	* user / password : postgres / changeme
	* When creating the new server, Host name / address in Connection tab must be set to "postgres_container"

### To call the GRPC service 

Use [Bloom](https://github.com/bloomrpc/bloomrpc) just load it with the protobuf (so either ```cart.proto``` or ```item.proto```) and set the correct local address and port (protobufs are in ```./protobuf/```)



> **Good luck and if you have any questions, please reach out to us!**
