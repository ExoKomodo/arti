Onto -build= you got to put "go build -mod=mod -gcflags=\"all=-N-l\""

Onto -command= you put "dlv --listen=:40000 --headless=true --api-version=2 --accept-multiclient exec $COMMAND --continue -- $COMMAND_ARGS".

Were $COMMAND would be your /main and $COMMAND_ARGS some parameters if you take them.

Finally, I also use the following CompileDaemon options: -command-stop=true and -graceful-kill=true because I want the subprocess of my application that is spun with dlv to be killed, and internally, CompileDaemon uses os.process.Kill() which doesn't kill the entire process tree, just the dlv one.
