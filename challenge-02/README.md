# Challenge - 02

Files provided:
```
├── challenge-02
│   ├── README.md
│   └── golang
│       ├── file.txt
│       └── server.go
```

## Ask
- Please find `golang` folder which has simple web server implementation in go. Server reads file.txt and displays the contents as home page.

- server can be built with CLI command. This will provide executable file `server`
```
#cd golang 
#go build server.go
```
- server can be started by CLI command
```
#./server 
```
- you will also find two files, which should instead be read based on the Deployment environment 
```
file-dev.txt --> file read for DEV environment
file-prood.txt --> file read for PROD environment
```
- your task is to work on Deployment files for this server. Start with containerizing the app and then use `docker-compose` to deploy.

- make sure at deployment time you can pass Environment as parameter to select value DEV/PROD. Based on Environment, server should display contents of either file-dev.txt or file-prod.txt

- Hints:
  1) you can't change server.go but can manipulate file*.* files in your deployment.
  2) make sure your process includes golang build and run steps accordingly.
  
- Submit/commit your code as a new file/s to your personal repo folder named challenge-02

- Please provide complete written instructions on how to use your solution files. Like how to run, build, deploy.

## Solution 
???



