![image](https://user-images.githubusercontent.com/3319450/192067289-4cf2fe7f-25ab-47be-ae36-d7be8398ddfa.png)  

<p align="center">A rich, portable server editing and development toolkit for EverQuest Emulator servers</p>

<hr>

<p align="center">
	<a href="https://go.dev/" target="_blank" rel="noopener noreferrer">
	    <img src="https://user-images.githubusercontent.com/3319450/192068038-60fb9928-5bc1-47e8-b348-b0b9ef9ced38.png" width="300"/>
	</a>

  <a href="https://vuejs.org" target="_blank" rel="noopener noreferrer">
	  <img width="200" src="https://vuejs.org/images/logo.png" alt="Vue logo">
  </a>
</p>

<p align="center">Powered by Vue and Golang</p>

<hr>

![image](https://user-images.githubusercontent.com/3319450/192069391-16aef5e5-5675-40f7-9545-ca5083d0d20e.png)

![image](https://user-images.githubusercontent.com/3319450/192069476-7f76ebe2-b331-4f72-96fd-59224759ccd8.png)

![image](https://user-images.githubusercontent.com/3319450/192069569-9acc9563-cd49-4972-8318-673d10fec3ee.png)

![image](https://user-images.githubusercontent.com/3319450/192069627-f99cf2dd-85c2-4563-b19a-e082de26efda.png)

![image](https://user-images.githubusercontent.com/3319450/192069348-59bd8e7f-35c1-44dc-81ee-b09644e3a910.png)

<hr>

- [Why Spire?](#why-spire)
- [Using Spire - Locally](#using-spire---locally)
- [Using Spire - Hosted](#using-spire---hosted)
- [Using Spire - Locally, but Remote](#using-spire---locally-but-remote)
- [Feature Requests](#feature-requests)
- [Reporting Bugs](#reporting-bugs)
- [Developer Setup](#developer-setup)
- [Linux  Development Setup](#linux--development-setup)
  - [Linux - Clone](#linux---clone)
  - [Linux - Install](#linux---install)
  - [Linux - Install - What happens](#linux---install---what-happens)
  - [Linux - Running Development Watchers](#linux---running-development-watchers)
- [Windows Development Setup](#windows-development-setup)
  - [Windows - Pre-Requisites](#windows---pre-requisites)
  - [Windows - Docker](#windows---docker)
  - [Windows - Clone](#windows---clone)
  - [Windows - Init](#windows---init)
  - [Windows - Install Environment](#windows---install-environment)
  - [Windows - Running Development Watchers](#windows---running-development-watchers)

## Why Spire?

The motive for Spire is simple, to empower creativity in the super fans of EverQuest creating content on emulated servers.

Rich, deep, tooling that leaves no stone un-turned for quality and intuitiveness.

Built for the long haul with code generation to make keeping things up to date far easier.

## Using Spire - Locally
  
Download the [latest release](https://github.com/Akkadius/spire/releases). for your operating system.

Place the executable in your EQEmulator Server directory and simply run it. 

That's it. No dependencies, no installations, no extra steps. 

Spire on your development server instantly.

![image](https://user-images.githubusercontent.com/3319450/192069875-ba916482-d28f-4b56-8819-7ce971781e87.png)

![image](https://user-images.githubusercontent.com/3319450/192070079-48a7ed8c-fd7e-4ae5-aa03-db9e5af8d677.png)  

## Using Spire - Hosted

If you don't want to install Spire, but you just want to have it connected to your database over the internet similar to how EOC worked (If you're familiar with it) - you can do so by using the hosted Spire at http://spire.akkadius.com/ and navigate to the login page @ http://spire.akkadius.com/login

![image](https://user-images.githubusercontent.com/3319450/192075339-69893324-0c0c-45a6-827e-4deaa0a338bb.png)

Once you've signed in with your Github account, you can manage your server connections. Navigate to **Create New** to create a new connection. Once you have connections created it is easy to switch between your connections and they are displayed at the bottom left.

![image](https://user-images.githubusercontent.com/3319450/192075441-6e2e6caf-671b-462c-9d44-9ad13ee726fd.png)


## Using Spire - Locally, but Remote

If you want to run Spire without an EQEmu server installation, place it in an empty folder alongside a `eqemu_config.json` with database information to connect to your MySQL server and it will connect remotely to the server specified in the configuration.
  
**Spire is currently only to be used as a local development tool, it is not safe to host publicly without protected access until authentication and roles are implemented. (On the roadmap)**  
  
## Feature Requests

Interested in a feature in Spire? Please file an [issue](https://github.com/Akkadius/spire/issues) in the issue tracker with the prefix `[Feature Request]`
 
## Reporting Bugs

Found a bug? Please file an [issue](https://github.com/Akkadius/spire/issues) in the issue tracker with the prefix `[Bug]`

## Contributing

Want to help contribute to Spire? Anyone can submit [pull requests](https://github.com/Akkadius/spire/pulls) however learning the skills required to work in this project might require some extra help and learning resources.

### Contributing - Learning Resources

| Area | Topic | Resource |
|--|--|--|
| Frontend | Learning Vue | [Learn Vue 2 - Step by Step](https://laracasts.com/series/learn-vue-2-step-by-step) |
| Backend | Learning Go | [Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU) |
| Backend | Go CLI Library used in Spire| [Cobra Github](https://github.com/spf13/cobra) |
| Backend | Web Framework Used in Spire | [Echo Web Labstack Docs](https://echo.labstack.com/) 
| Backend | Database Interaction (ORM) | [Gorm - The fantastic ORM library for Golang](https://gorm.io/) |
| Backend | Google Wire - Dependency Injection | [Google Wire Github](https://github.com/google/wire) - [Tutorial](https://github.com/google/wire/blob/main/_tutorial/README.md)
| Backend | Go Dot Env (.env file loader) | [joho/godotenv](https://github.com/joho/godotenv) |
| Backend | Echo Swagger - Used to serve swagger docs | [Echo Swagger Github](https://github.com/swaggo/echo-swagger) |
| Backend | Swaggo Generator - Used to code generate swagger documentation | [Swaggo](https://github.com/swaggo/swag)




  

## Developer Setup  
  
These are instructions for those who are looking to develop on Spire. If you are just trying to use the tool then see the **using Spire** sections.
  
## Linux  Development Setup

![enter image description here](https://user-images.githubusercontent.com/3319450/192069126-b1daf88a-b728-4e9f-90eb-6715dd49e924.png)

These instructions assume you have **git**, **node,** **docker** already installed. All of the dependencies are taken care of within the docker environment.

* Git - Install with your package manager
* [Node](https://nodejs.org/en/download/package-manager/)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)

### Linux - Clone

First clone Spire, copy the base `.env.dev` file to the `.env` used by Spire in local development and run `make install` in one line below.

```
git clone https://github.com/Akkadius/spire.git
```  

###  Linux - Install

```
cd spire && cp .env.dev .env && make install && make install-frontend
```

### Linux - Install - What happens

* MariaDB container gets initialized with credentials held in `.env` (make mysql-init)
* A ProjectEQ database gets seeded into the database container from http://db.projecteq.net/api/v1/dump/latest (make seed-peq-database) to a database called `peq`
* Spire tables get installed to a separate `spire` database (make seed-spire-tables)
* Installs static assets (icons, images, preview images) (make install-assets) from https://github.com/Akkadius/eq-asset-preview
* Installs `frontend/.env.example` which is required in development to properly route API requests to the development backend (VUE_APP_BACKEND_BASE_URL=http://localhost:3000) (make install-frontend)

### Linux - Running Development Watchers

Conveniently, there are two make commands that run the development watchers for both the back-end and the front-end.

#### Linux - Backend Web API Development Watcher

Re-compiles and runs the backend webserver when changes are detected in `*.go` files.

```
make watch-be
```

![image](https://user-images.githubusercontent.com/3319450/192075174-f08ee34a-3ebb-48a6-8c5a-c0b35ec101d3.png)

#### Linux - Frontend Vue Development Watcher

Re-compiles and runs the backend webserver when changes are detected in the `./frontend` folder (Ignores `node_modules` and `public`)

```
make watch-fe
```

![image](https://user-images.githubusercontent.com/3319450/192075149-5a7e0c42-384b-4d7f-a60c-a54696cf6c6f.png)


#### Linux - Browse

At this point, you should be able to browse to http://localhost:8080 to view the development server.

![image](https://user-images.githubusercontent.com/3319450/192075588-2c7f4383-deab-4881-ab2f-a216dbccb7f7.png)

## Windows Development Setup
 
![image](https://user-images.githubusercontent.com/3319450/192071001-ca2c314a-6436-4f37-a02b-a55595906b12.png)

For Windows development environment, install the following pre-requisites before proceeding with the next steps  
  
### Windows - Pre-Requisites   
  
* Install [Windows Git](https://git-scm.com/download/win) (For Git Bash)  
* Install [Docker](https://docs.docker.com/desktop/windows/install/) you may need to install additional kernel components for WSL2 which will be instructed in the Docker installation  
  
All other necessary software gets automatically installed through the subsequent automated steps  
 
### Windows - Docker   

Windows for Docker can be terrible on performance, for this reason a lot of the development environment operations are done at the host level to keep things simple, fast. 

You **can** use the workspace bash container but do know it comes with a performance hit. The containerized MariaDB database is also a convenience as well. 

You **can** point the `.env` to a separate eqemu server installation if you want to avoid Docker altogether and run everything on the host.  

### Windows - Clone  
  
Clone Spire to a directory of your choosing  
  
```  
git clone https://github.com/Akkadius/spire.git  
```  
  
### Windows - Init  
  
Once you have your pre-requisites installed you will need to run `windows-init.bat` on the top level folder as **administrator**  
  
This init step will perform the following **automatically**  
  
* Install [Choco](https://chocolatey.org/) a package manager for Windows  
* Install [Golang](https://golang.org/)  
* Install [NodeJS](https://nodejs.org/en/) LTS version  
* Install [Make](https://www.gnu.org/software/make/) for make commands  
* Copies `.bashrc` `.bash_profile` files to ~/ home directory  
* Copies `.wslconfig` to ~/ home directory  
* Initializes the Frontend `./frontend/.env.example.windows` as `frontend/.env` (Tells the development webserver where to route API calls for development)  
* Initializes the Backend `.env.dev` to `.env`  
* Launches a Git Bash (MinGW) shell when done for the following steps  
  
### Windows - Install Environment
  
In a MinGW shell, which you should have after Windows Init batch file is done running - you run the following command

```  
make install  
```  
  
If you don't have a MinGW window already open from the previous step; either click the `windows-bash.bat` alias or launch a "Git Bash" instance yourself through the Windows Start Menu and cd to the Spire folder  
  
Make install will do the following things automatically  
  
* Build the `workspace` docker image; the workspace contains Go and many other utilities installed and fully working out of the box. For windows users we will try to run as many things on the host as much as possible to avoid performance or compatibility issues. This can be bashed into using `windows-workspace-bash.bat`
* Build the `mysql` docker image which will contain a basic `mariadb` instance with a relatively tuned database configuration  
* Initializes a local MariaDB instance that you can access from localhost port `33066` (Note the extra 6 so we don't conflict with a local install)  
* Creates local databases `peq` and `spire`  
* Seeds the latest ProjectEQ database for development purposes to the local `peq` database  
* Seeds the local Spire database tables to the `spire` database  
  
At this point the installation should be complete and you should have everything that you need to develop. For good measure and because this is Windows, you should probably reboot  
  
### Windows - Running Development Watchers  
  
To run the backend and frontend development servers in Windows; there are simply two top level batch scripts that you can run  
  
* `windows-backend-web.bat` This will run the Golang backend web process on port 3001 (in windows) and will reload when any changes are made to the codebase  
* `windows-frontend-web-dev.bat` This will run the NodeJS Webpack watcher which will serve the frontend web development instance and will hot reload any changes made to the frontend codebase on the fly  
  
Both of these scripts are designed to kill an already running instance when it is ran again
