# Spire

This readme is a WIP

## Developer Setup

### Linux

TODO

### Windows

For Windows development environment, install the following pre-requisites before proceeding with the next steps

#### Windows Pre-Requisites 

* Install [Windows Git](https://git-scm.com/download/win) (For Git Bash)
* Install [Docker](https://docs.docker.com/desktop/windows/install/) you may need to install additional kernel components for WSL2 which will be instructed in the Docker installation

All other necessary software gets automatically installed through the subsequent automated steps

#### Clone

Clone Spire to a directory of your choosing

```
git clone https://github.com/Akkadius/spire.git
```

#### Windows Init

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

#### Development Environment Install

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

At this point the installation should be complete and you should have everything that you need to develop. For good measure and because this is Windows we're talking about, you should probably reboot

#### Running Development Watchers

To run the backend and frontend development servers in Windows; there are simply two top level batch scripts that you can run

* `windows-backend-web.bat` This will run the Golang backend web process on port 3001 (in windows) and will reload when any changes are made to the codebase
* `windows-frontend-web-dev.bat` This will run the NodeJS Webpack watcher which will serve the frontend web development instance and will hot reload any changes made to the frontend codebase on the fly

Both of these scripts are designed to kill an already running instance when it is ran again