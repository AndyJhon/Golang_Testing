#Download the go binary
Copy to /usr/local/go path

#Edit bashrc file
vim .bashrc 

#In case it does not work, configure goroot
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin

Then

source ~/.bashrc

validate go version

go --version


