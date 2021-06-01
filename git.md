## Project Description

The project is an web based application that allows user to input text and see the input get hashed with SHA256 in real-time as the user types.

## Languages/Communication Used

NodeJS and Jquery/AJAX were used for the the server/frontend which is used to take user input and display results. Of course HTML and CSS were also used.

Python3 script was used to compile and run a SHA256 application written in Golang.

Golang was used to implement the SHA256 algorithm which was done from scratch.

NodeJS sends the user HTML input to a local Python script which then compiles and calls the SHA256 code in Golang to hash the input and return the result back to the webserver.