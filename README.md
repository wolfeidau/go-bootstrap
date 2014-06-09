# go-bootstrap

Small bootstrap project for golang projects.

You will need to update the files in with the owner and project
name, and register with [bintray](https://bintray.com/) to upload
binaries.

# get started

Modify the following files to update the name of the service, this will be automated in the future.

* `.gitignore`
* `Makefile`
* `command/version.go`
* `command/agent/command.go`

Code changes are located as follows, found using ack:

```
$ ack appname
command/agent/command.go
42:	return "Runs a appname agent"
47:Usage: appname agent [options]
49:  Starts the appname agent and runs until an interrupt is received.

command/version.go
23:	fmt.Fprintf(&versionString, "appname v%s", c.Version)
37:	return "Prints the appname version"

main.go
35:		HelpFunc: cli.BasicHelpFunc("appname"),
```
