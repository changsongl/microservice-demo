# microservice-dmo
Microservice demos are written by popular Golang microservice framework. 

This is a cli framework to generate the whole project structure. It prepares the 
whole project to run.

### How to use?
````shell script
# load the project files to go file, so we can compile it to the program.
cd -src=your-path/microservice-dmo/tmpl
statik -src=your-path/microservice-dmo/tmpl/file

go build main.go
./kcli init web
````
