<h1 align="center"><strong>GoStrap</strong></h1>

<div align="center">
<a href="https://pkg.go.dev/github.com/piigyy/gostrap"><img src="https://pkg.go.dev/badge/github.com/piigyy/gostrap.svg" alt="Go Reference"></a>
<a href="https://goreportcard.com/report/github.com/piigyy/gostrap"><img src="https://goreportcard.com/badge/github.com/piigyy/gostrap" alt="Go Report Card"></a>
</div>

**GoStrap** probably will going to be the easiest way for you to create new Go project using your existing project template.

This CLI actually only do:
- Clone you Go Project templates to new directory.
- Re init the git.
- Replace the old go module name with the new one.

YES! You can use any project template you wanted!

```txt
please tell me if you have easier way than this CLI. I need it too!
```

## Install
If you using Go version 1.17 or newer, you can use this command
```bash
go install github.com/piigyy/gostrap@latest
```
If you using elder version of Go (<1.17), well tbh I don't know if this CLI support it. But a try wouldn't hurt, right?
```bash
go install github.com/piigyy/gostrap@latest
```

## How To Use
After installing GoStrap, simple run this command
```bash
gostrap new <project-name> -m <you/golang/module/name> -t <your-golang-project-template> -p <your-go-module-placeholder>

# Example
gostrap new go-authorization -m github.com/piigyy/authorization -t https://github.com/golang-standards/project-layout -p github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME
```

## How It Work?
GoStrap work basically copy your Go project template into new directory, remove the `.git` directory, re init new git configuration, and replace the existing go module name with your new go module name.

## Configuration File
GoStrap configuration file use `.yaml` file. The default `.gostrap.yaml` location is in you home directory (`$HOME/.gostrap.yaml`).

Configuration file will look like this:
```yaml
template:
gomoduleplaceholder:
```

You can set this file by manually editing `.gostrap.yaml` file or using **GoStrap** command: 
```bash
gostrap set [key] [value]

## example
gostrap set template https://github.com/golang-standards/project-layout
gostrap set github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME github.com/piigyy/gostrap
```

Once you setup **GoStrap** configuration file you don't need to pass template and placeholder flag to `new` command

```bash
gostrap new microservice -m github.com/piigyy/microservice
```

The `new` command will automatically set template and placeholder from you `.gostrap.yaml` file.
