Manual Updates
==============

For bootstrapping and manual updates, download the `HTML` of the following three links:

1. http://www.petharbor.com/results.asp?searchtype=ADOPT&friends=1&samaritans=1&nosuccess=0&rows=1000&imght=120&imgres=thumb&view=sysadm.v_animal_short&fontface=arial&fontsize=10&zip=80209&miles=10&shelterlist=%27ARAP%27,%27AURO%27,%27DNVR%27,%27DDFL%27,%2783615%27,%2780454%27,%2779367%27,%2782294%27,%2777298%27,%2784657%27,%2769972%27,%2784715%27,%2779780%27,%2777803%27,%2776338%27,%2785330%27,%2776065%27,%2778397%27,%2786214%27,%2785252%27,%2774805%27,%2773867%27,%2782242%27,%2781793%27,%2772856%27,%2773086%27,%2782431%27,%2786406%27,%2774867%27,%2783241%27,%2772907%27,%2774328%27,%2786813%27,%2771436%27,%2782755%27,%2782206%27,%2776134%27&atype=&PAGE=1&where=type_DOG
2. http://www.petharbor.com/results.asp?searchtype=ADOPT&friends=1&samaritans=1&nosuccess=0&rows=1000&imght=120&imgres=thumb&view=sysadm.v_animal_short&fontface=arial&fontsize=10&zip=80209&miles=10&shelterlist=%27ARAP%27,%27AURO%27,%27DNVR%27,%27DDFL%27,%2783615%27,%2780454%27,%2779367%27,%2782294%27,%2777298%27,%2784657%27,%2769972%27,%2784715%27,%2779780%27,%2777803%27,%2776338%27,%2785330%27,%2776065%27,%2778397%27,%2786214%27,%2785252%27,%2774805%27,%2773867%27,%2782242%27,%2781793%27,%2772856%27,%2773086%27,%2782431%27,%2786406%27,%2774867%27,%2783241%27,%2772907%27,%2774328%27,%2786813%27,%2771436%27,%2782755%27,%2782206%27,%2776134%27&atype=&PAGE=1&where=type_CAT
3. http://www.petharbor.com/results.asp?searchtype=ADOPT&friends=1&samaritans=1&nosuccess=0&rows=1000&imght=120&imgres=thumb&view=sysadm.v_animal_short&fontface=arial&fontsize=10&zip=80209&miles=10&shelterlist=%27ARAP%27,%27AURO%27,%27DNVR%27,%27DDFL%27,%2783615%27,%2780454%27,%2779367%27,%2782294%27,%2777298%27,%2784657%27,%2769972%27,%2784715%27,%2779780%27,%2777803%27,%2776338%27,%2785330%27,%2776065%27,%2778397%27,%2786214%27,%2785252%27,%2774805%27,%2773867%27,%2782242%27,%2781793%27,%2772856%27,%2773086%27,%2782431%27,%2786406%27,%2774867%27,%2783241%27,%2772907%27,%2774328%27,%2786813%27,%2771436%27,%2782755%27,%2782206%27,%2776134%27&atype=&PAGE=1&where=type_OO

Use the go script in `cmd` to load all three files at the same time. It is important that thye all be loaded at once!

    go run load_file.go ../_data/others.html ../_data/cats.html ../_data/dogs.html
