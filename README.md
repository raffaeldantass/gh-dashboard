# Github Repositories Dashboard

## Overview

This project is a dashboard for viewing repositories from a Github user's personal account and organizations. It provides a paginated list of repositories, with options to view more repositories per page and navigate through pages.

## Requirements

- Docker Desktop
- Node.js (v18+)
- Go (v1.22+)

## Project Structure

```
gh-dashboard/
├── backend/
│   ├── config/
│   │   ├── config.go
│   ├── handlers/
│   │   ├── auth.go
│   │   ├── repositories.go
│   ├── models/
│   │   ├── repository.go
│   ├── services/
│   │   ├── github.go
│   ├── Dockerfile
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
├── frontend/
│   ├── app/
│   │   ├── components/
│   │   │   ├── GithubAuth
│   │   │   ├── Loading
│   │   │   ├── Pagination
│   │   │   ├── ResultList
│   │   │   ├── RepositoryCard
│   │   ├── lib/
│   │   │   ├── hooks/
│   │   │   │   ├── useGithubAuth
│   │   │   │   ├── useGithubRepositories
│   │   │   │   ├── useAuth
│   │   │   ├── utils/
│   │   │   │   ├── index.ts
│   │   │   ├── repositories/
│   │   │   │   ├── index.ts
│   │   │   ├── styles/
│   │   │   │   ├── globals.css
│   │   │   ├── types/
│   │   │   │   ├── github.ts
│   │   │   ├── page.tsx
│   │   │   ├── layout.tsx
│   ├── package.json
│   ├── Dockerfile
├── docker-compose.yml
├── .env.sample.development
├── .env.sample.production
├── .gitignore
```

### Backend

Backend is using Air to run the application to make it easier to develop. It's a tool that allows you to run the application with hot reloading.

#### Config

- Environment variables are stored in the .env.development and .env.production files and loaded through docker compose.

#### Handlers

- Handlers are the functions that handle the requests from the frontend to the backend. They are stored in the handlers folder.
Such as login, repositories, etc.

#### Models

- Models are the structs that represent the data of the application, specifically the data from the Github API.

#### Services

- Services are the functions that handle the comunication with the Github APP API.

___

### Frontend

#### App/components

- Components to handle pagination, loading, authentication, etc.

#### App/lib

- The lib folder contains utility functions and hooks to handle the application logic and comunication with the backend.

#### App/page.tsx

- The page.tsx file is the main component of the application.

#### App/repositories

-  The repositories folder contains the components to handle the repositories list and all of the details of the repository card.

#### App/types

- The types folder contains the types of the whole application.


## Installation

Clone the repository using Github CLI:

```bash
gh repo clone raffaeldantass/gh-dashboard 
```
**For the sake of simplicity, the .pem file is include in this repo. Also the env vars are stored in .env.sample files.**

In a real scenario, they should be stored in a secrets manager like AWS Secrets Manager or Azure Key Vault.

Go to the project directory and create a *.env.development* and also *.env.production* file and copy the contents of *.env.sample.development* and *.env.sample.production* to it.

Go to the frontend folder and run the following command to install the dependencies:

```
$ ~/gh-dashboard/frontend npm install
```

Go to the backend folder and run the following command to update and install the dependencies **This is not required, it's optional and the goal is to make it easier to update the packages if necessary**:

```
$ ~/gh-dashboard/backend go mod tidy
```

Open Docker Desktop and run:

For Local Development (In your gh-dashboard root folder):

```
docker compose --env-file .env.development up frontend-dev backend-dev
```

For Production (In your gh-dashboard root folder):

```
docker compose --env-file .env.production up frontend-dev backend-dev
```

## Usage

### Open the application in your browser

```
http://localhost:3000
```
It may take sometime to load the page in the first time. 
It will be opened the Frontend App. 

The backend will be running on http://localhost:8080.

### Page Routes

- **/:**
It's the main page. It's a login page. There's a button to login with Github.

- **/repositories:**
It's the repositories page. It's a paginated list of repositories. It's a list of cards with the repository name, description, last update and owner.
If the user is not logged in, a button to login with Github will be shown.

### API Routes

- **/api/login:**
It's the login route. It'll redirect the user to Github's OAuth page.

- **/api/repositories:**
It's the repositories route. It's used to get the repositories of the user. It'll return a paginated list of repositories if the user is logged in.

#### Running the backend tests

To run the backend tests, you can use the following command:

```
$ ~/gh-dashboard/backend go test .
```

If you want the verbose output, you can use the following command:

```
$ ~/gh-dashboard/backend go test -v .
```

To run the tests for a specific file, you can use the following command:

```
$ ~/gh-dashboard/backend go test -v ./<folder>/<file_test>.go
```

It was added a unit test for the github service to test the integration with the Github API. 

## Philosophy

The application is a monorepo with a frontend and a backend and a docker compose file to run them. The reason for this is to have a single repository to manage the whole application and to have a single place to make changes to the whole application, making it easier to deploy and maintain.

Also I decided to use a Next.js as framework for the frontend because it's a popular framework with good documentation and a large community. And it's a framework that I'm familiar with. Easier to make it work.
Also it's easy to use as client side application, which is the case of this application. 

Backend is using gin framework because it's almost zero configuration and it's very easy to use. Speeds up the development process.
For the Github APP Api, I'm using the [go-github](https://github.com/google/go-github) library with the [oauth2](https://github.com/golang-jwt/jwt/v4) library to handle the OAuth2 flow. It speeds up integration time.


The application is containerized with Docker to make it easier to deploy and run locally. It's using a docker-compose file to handle all of the configs to centralize the environment variables and tooling.
