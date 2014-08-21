Adopt
=====

Visit the site at [pets.codefordenver.org](http://pets.codefordenver.org/)!

API for adoptable pets in Denver.

[View the shelters near Denver](http://www.petharbor.com/pick_shelter.asp?searchtype=ADOPT&friends=1&samaritans=1&nosuccess=0&rows=100&imght=120&imgres=thumb&view=sysadm.v_animal_short&fontface=arial&fontsize=10&zip=80209&miles=10)!

For testing the `robot` package, create a `local_settings.json` file in the `robot` sub-directory with valid database credentials. An example:

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

The `pets` table should also be created using the sql in the `docs` sub-directory.
