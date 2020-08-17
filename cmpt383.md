## Project Description

The project is an web based application that allows user to input text and see the input get hashed with SHA256 in real-time as the user types.

## Languages/Communication Used

NodeJS and Jquery/AJAX were used for the the server/frontend which is used to take user input and display results. Of course HTML and CSS were also used.

Python3 script was used to compile and run a SHA256 application written in Golang.

Golang was used to implement the SHA256 algorithm which was done from scratch.

As mentioned earlier, NodeJS sends the user HTML input to a local Python script which then compiles and calls the SHA256 code in Golang to hash the input and return the result back to the webserver.

## Starting the project

First,
```sh
vagrant up
vagrant ssh
```
then you want to go to the project directory,
```sh
cd project
```
then start the webserver with:
```sh
node index.js
```
Finally go to http://localhost:8080/

## Troubleshooting

**Port 8080 already in use by host machine**
Please try going into the Vagrantfile and changing the host port number to something else, in the port forwarding line of the file.

**Error with SHA256 hash output**
Make sure the Gopath is set to $HOME/project, the Python script should already do that automatically, but if it doesn't work please try setting the Gopath manually.
