Adopt
=====

Find adoptable pets in Denver!

Visit a deployed version of the site at [pets.codefordenver.org](http://pets.codefordenver.org/)!

This site creates an alternative API for pets in [shelters near Denver](http://www.petharbor.com/pick_shelter.asp?searchtype=ADOPT&friends=1&samaritans=1&nosuccess=0&rows=100&imght=120&imgres=thumb&view=sysadm.v_animal_short&fontface=arial&fontsize=10&zip=80209&miles=10). It is written in Python on the [Django](https://www.djangoproject.com/) web framework. It is backed by a custom robot written in [Go](http://golang.org/).


### Website

The website is deployed and updated through [Fabric](http://www.fabfile.org/) commands. Deploying to a Ubuntu 14.04 server can be done with:

    fab -H user@server deploy

This will install the necessary software on the server, including [PostgreSQL](http://www.postgresql.org/), [nginx](http://nginx.org/), and [uWSGI](http://projects.unbit.it/uwsgi/).

Any updates can also be performed by the `fabfile`:

    fab -H user@server update


### Robot

The database is populated by an automated Go process. To run the robot, including any tests or commands, first create a `local_settings.json` file in the `robot` sub-directory with valid database credentials. An example:

```json
{
    "database": {
        "driver": "postgres",
        "host": "localhost",
        "port": 5432,
        "name": "adopt",
        "user": "user",
        "password": "pass" 
    }
}
```

To update the animals in the database, simply run the `get_pets.go` script in the `cmd` directory:

    go run go run get_pets.go

A database can also be bootstrapped with `html` files downloaded from the source website. These are loaded with the `load_file.go` script in the `cmd` directory. For example:

    go run load_file.go others.html cats.html dogs.html

It is important to run all file in one command. Subsequent operations will be considered separate batches: any animals not in a batch will be marked as removed by the database.

-aodin, 2014
