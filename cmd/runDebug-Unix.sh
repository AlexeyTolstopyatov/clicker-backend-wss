#!/usr/bin/env bash

#        *The server seeks .json or .env file itself!*
#    No meaning write new scripts for reading configuration file
# Maybe it's bad idea, but... I suppose, it more comfortable than
# running server directly from FS (clicking or writing "wsst.EXE")
# without any mind, that runs without config (uNExPecTEd).
# Syntax of calling convention:
# wsst [http://address] [port] [profile]
# Where you can find working address?
# ws://address:port/ws
#
# Don't worry, this information exists in server's "Start Page"

if [ -f db.json ]; then
    echo "env file exists!"
    echo "Command-Line arguments have second priority for building."
    echo "Aborting..."
else
    # Check if server.json exists
    if [ -f server.json ]; then
        echo "startupArgs.json exists!"
        echo "Command Line arguments have second priority for building."
        echo "Aborting..."
    else
        echo "Running server..."
        # your hand-made server's configuration here
        wsst localhost 8080 debug
    fi
fi