# go-bootstrap

Small bootstrap project for golang projects.

You will need to update the files in with the owner and project
name, and register with [bintray](https://bintray.com/) to upload
binaries.

# get started

* Modify the name of the service in the following places, found using ack:

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

* Add code to the agent, this is where the long running service logic lives.

* Add any status, action commands you need.

Note, this will be automated in the future.

# Plan

* Add an example REST endpoint using [negroni](https://github.com/codegangsta/negroni).
* Add an RPC service for commands to query status / metrics on the local system.
* Add some example metrics for the aforementioned features.
* Authentication of RPC requests.
* Build an example using REST and static JS site to show a working single page app.
