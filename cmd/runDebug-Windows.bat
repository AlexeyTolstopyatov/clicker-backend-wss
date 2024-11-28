@ECHO OFF
REM         *The server seeks .json or .env file itself!*
REM
REM     No meaning write new scripts for reading configuration file
REM Maybe it's bad idea, but... I suppose, it more comfortable than
REM running server directly from FS (clicking or writing "wsst.EXE")
REM without any mind, that runs without config (uNExPecTEd).

REM Syntax of calling convention:
REM wsst [http://address] [port] [profile]

REM Where you can find working address?
REM ws://address:port/ws
REM Don't worry, this information exists in server's "Start Page"

if exist db.json (
    ECHO env file exists!
    ECHO Command-Line arguments have second priority for building.
    ECHO Aborting...
) else (
    REM Check if startupArgs.json exists
    if exist startupArgs.json (
        ECHO startupArgs.json exists!
        ECHO Command Line arguments have second priority for building.
        ECHO Aborting...
    ) else (
        ECHO Running server...
        REM Your hand-made server's configuration here
        wsst localhost 8080 debug
    )
)

REM The main question is 'priority' of json configuration or directly
REM (Windows) Command-Line arguments...
REM Just because it can be user-friendly...

@PAUSE