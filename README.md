# filter_main
https://medium.com/the-godev-corner/how-to-create-a-go-private-module-with-docker-b705e4d195c4
Now, let’s add the go module to our simple Go application via our project’s root directory:
pushd ~/src/reports/volume/go/create-go-module
go mod init github.com/brentgroves/filter_main
👉 The preceding will generate a go.mod file that will track our dependencies.

Next, within the root directory, create a main.go file and paste the content below into it:

https://go.dev/doc/modules/managing-dependencies
pushd ~/src/reports/volume/go/create-go-module/filter_main
go get .
ls ~/go/pkg/mod/github.com/brentgroves
this is where the module dependancy is.
go mod tidy // this removes unneeded dependancies
go run .
current directory is contained in a module that is not one of the workspace modules listed in go.work. You can add the module to the workspace using:
        go work use .