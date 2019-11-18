gotask
-------

**gotask** is a easy to use command line interface tool for adding, removing and editing daily tasks in a database stored in your local filesystem.

Installation
------------

**gotask** is completely written in go, so all you have to do is install git, go and execute below command.

    $ go get github.com/thinkofher/gotask

Make sure you have `$GOPATH/bin` added to your `$PATH` variable.

Usage
----

    NAME:
       gotask - Add, remove and edit tasks in your local database.

    USAGE:
       gotask [global options] command [command options] [arguments...]

    VERSION:
       0.0.1

    AUTHOR:
       Beniamin Dudek <beniamin.dudek@yahoo.com>

    COMMANDS:
         add, a   Add task to your tasks list
         show, s  Show tasks in your tasks list
         done, d  Complete task with given id.
         help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --help, -h     show help
       --version, -v  print the version

Used technologies
-----------------

- [mitchellh/go-homedir](https://github.com/mitchellh/go-homedir)
- [boltdb](https://github.com/boltdb/bolt)
- [urfave/cli](https://github.com/urfave/cli)
