# Rover control api project

### Background 
Rover will read a file and move follow instruction in file.
and prepare API for get position of rover on every step.

#### 3 API end point request example: 
```
    | GET      "/status"                    | get staus of server is ready to used other end-point.
    | GET      "/map-info"                  | get map info size and last setup of rover move.
    | GET      "/rover-status?step=1"       | get rover position from request `step`.
```
#### Return data example in JSON format:
```
    /status         | {"server_status":"ok"}
    /map-info       | {"map_size":24,"rover_step":8}
    /rover-status   | {"status":{"direction":"E","position_x":1,"position_y":0},"current_step":2,"status_text_format":"E:1,0"}
```

### How to set up project local 
in this project is have make file for convenient to run command that useful
of this project. and also used docker for make sure all machine that run this project get the same result. 
#### Start new project
```
   1| git clone and cd into project folder
   2| run command 'make newfile' for make new instructions.txt file for rover to move follow this file.
   3| run command 'make docker' for make new or rebuild docker images and start. 
```
#### Need new route 
```
   1| edit file instructions_file/instructions.txt and save make sure command is long only 1 
   2| run command 'make docker' for reload project.
   3| test with /rover-status api and see position of rever is change after change route and reload project.
```