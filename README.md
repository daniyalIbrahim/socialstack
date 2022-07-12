<h1 align="center">🎉 Social Stack🎉</h1>
This project is a SaaS running in a serverless envoirment in google cloud platform. It is a authentification service written in go lang and built with docker

<img src="https://img.shields.io/github/license/daniyalibrahim/socialstack" />
<a href="https://github.com/daniyalibrahim/socialstack/issues">
<img src="https://img.shields.io/github/issues/daniyalibrahim/socialstack" />
</a>
<img src="https://img.shields.io/github/languages/count/daniyalibrahim/socialstack?style=flat-square"/>


## What's Inside

- Pure Golang, rest-api only using standard go packages
- Rest Api with [net/http](https://golang.org/pkg/net/http/)
- CRUD using Mongo Golang Driver and Mongo Atlas Distributed Database
- [crypto/bcrypt](https://golang.org/x/crypto/bcrypt) for password hashing
- Dockerfile
- CI/CD to Google Cloud Run

## How Robust is This Service
- Using a NO/SQL DB storing json data
- Not using a JWT,but session management system
  - at each login the user receives a newly generated session id
  - at each logout the user deletes the session id from the database
  - at each request the service checks if the session id is valid and present in the DB
  - if not, the service returns a 401 error
- Can not register with the same email twice
- All fields are validated using Regex expressions
- All HTTP request errors are handled using a custom error handler
- All HTTP request are logged using a custom logger
- All User Endpoints are inaccessible if the user logs out

## How to Use This Service
- Register with your email and password
- Login with your email and password
- Get a new session id
- Use the session id to access all the endpoints
- Logout to delete the session id

## Live demo

[API URL](https://techonomy-labs-o2k3wv2fsq-uc.a.run.app/api/v1/)

[Documentation](https://documenter.getpostman.com/view/21725756/UzJHRdXy)

[Join Us](https://app.getpostman.com/join-team?invite_code=40a4a16810b9f88648390722e98b8e79)

[Our Website](https.//daniyalibrahim.com)

