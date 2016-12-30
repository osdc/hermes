# hermes
Social Network for JIIT

## Frontend

### Setup Dependencies
Run either `yarn install` or simply `yarn` to setup dependencies. Visit yarn's official [installation guide](https://yarnpkg.com/en/docs/install) for installating yarn.

### Run Locally
The project already contains a fully-configured `webpack-dev-server` to run the frontend locally. Simply start it by firing up `yarn start` and visit `localhost:3000` on your browser.

### Testing (Yet to be implemented)
`yarn test`

### Linting using ESLint
`yarn eslint`

## Backend

###Install Go
Download the archive from https://golang.org/dl/ and extract it into /usr/local, creating a Go tree in /usr/local/go. For example:
`tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz`
Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:
`export PATH=$PATH:/usr/local/go/bin`
Create a directory to contain your workspace, $HOME/work for example, and set the GOPATH environment variable to point to that location.
`$ export GOPATH=$HOME/work`
For more Information visit the official installation guide https://golang.org/doc/install

###Install glide
The easiest way to install the latest release on Mac or Linux is with the following script:
curl https://glide.sh/get | sh
On Ubuntu Precise(12.04), Trusty (14.04), Wily (15.10) or Xenial (16.04)
`sudo add-apt-repository ppa:masterminds/glide && sudo apt-get update`
`sudo apt-get install glide`
On Mac OS X you can also install the latest release via Homebrew:
`$ brew install glide`

###Install Packages and Dependencies
`$ glide install`

###Install and configure PostgreSQL
`$ sudo apt-get update`
`$ sudo apt-get install postgresql postgresql-contrib`
Log in using the postgres account:
'$ sudo -i -u postgres
Create Superuser:
`postgres@server:~$ createuser --interactive`
Output:
`Enter name of role to add: <yoursuperuser>
Shall the new role be a superuser? (y/n) y`
Now go to db_setup.go and replace <yoursuperuser> and <yoursudopassword> with respective values.
Create a local database for Hermes:
'postgres@server:~$ createdb hermes'
For verification you can type psql hermes to check if the database is created.

###Setup Database
`$ go build db_setup.go`
`$ $GOPATH/<pathtohermes>/db_setup`

###Running Server
`$ go build server.go`
`$ $GOPATH/<pathtohermes>/server`
visit `localhost:1323` on your browser.

###Format using gofmt
`$ gofmt -w <yourfilename.go>`
