My opinionated view on how a generic Golang service project(along with its code skeleton) should look like. 


# Directory structure explanation

```bash
.
├── Dockerfile
├── Makefile
├── README.md
├── bin                 # All build artifacts go here.
├── deploy              # Various deployment scripts(if any) go here.
│   └── README.md
├── docker-compose.yml
├── main.go             # The main entry point of the app. Keep this file as the only .go file at this level.
├── pkg                 # All go app-related packages should go here
│   ├── README.md
│   ├── cli             # A package to incapsulate all CLI commands and the entrypoint for command invocation.
│   │   ├── cli.go
│   │   └── create.go
│   └── srv             # A package to incapsulate all the code about server side.
│       └── srv.go
├── scripts             # All build/admin related scripts for _this_ project go here.
    └── build.sh
```

# Not covered yet
 - [ ] Dependency management(e.g. glide)
 - [ ] Coverage tool(go-coverage)
 - [ ] Codegen examples
