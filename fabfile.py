# -*- coding: utf-8 -*
import os

from fabric.api import local, settings, abort, run, cd, env
from fabric.contrib.console import confirm
from fabric.operations import sudo

# Use the local SSH config file
env.use_ssh_config = True


def restart():
    """
    Restart the server
    """
    sudo('shutdown -r now')


def update_ubuntu():
    """
    Perform updates for ubuntu.
    """
    sudo('apt-get -q -y update')
    sudo('apt-get -q -y dist-upgrade')
    sudo('apt-get -q -y autoremove')


def setup_ubuntu():
    """
    Install the required Ubuntu dependencies from the package repository.
    These packages are for Ubuntu 14.04
    """
    sudo('apt-get -q -y install make gcc python-dev git mercurial nginx python-pip python-setuptools python-virtualenv')
    sudo('apt-get -q -y install postgresql-9.3 postgresql-server-dev-9.3 python-psycopg2')


def upgrade_pip():
    """
    Upgrade PIP by using PIP!
    """
    sudo('pip install -q -U pip')
    sudo('pip install -q -U virtualenv')


def clone(url, directory='/srv'):
    """
    Clone the given repo into the /srv directory by default.
    """
    with cd(directory):
        sudo('git clone {url}'.format(url=url))


def update_git(alias, app, remote='origin', reset=False, syncdb=False):
    """
    Update the git repository.
    """
    # TODO os.path.join these various file paths
    src = '/srv/{alias}/'.format(alias=alias)
    python = '/srv/{alias}/env/bin/python'.format(alias=alias)
    manage = '/srv/{alias}/{app}/manage.py'.format(alias=alias, app=app)

    with cd(src):
        if reset:
            sudo('git reset --hard HEAD')
        sudo('git pull {remote} master'.format(remote=remote))

    if syncdb:
        sudo('{python} {manage} syncdb --verbosity=0 --noinput'.format(python=python, manage=manage))

    sudo('{python} {manage} collectstatic --verbosity=0 --noinput'.format(python=python, manage=manage))


def setup_env(alias, directory='/srv'):
    """
    Create the python virtual environment.
    """
    requirements = os.path.join(directory, alias, 'requirements.txt')
    path = os.path.join(directory, alias, 'env')

    # Create the virtual environment
    sudo('virtualenv {path}'.format(path=path))

    # And install the requirements
    sudo('{path}/bin/pip install -r {requirements}'.format(path=path, requirements=requirements))


def server_ln(alias):
    """
    Create the symbolic links for the server configurations.

    To the new init job without restarting, call:
    initctl reload-configuration

    View all services:
    service --status-all
    """
    server_src = '/srv/{alias}/conf/{alias}.conf'.format(alias=alias)
    server_dest = '/etc/init/{alias}.conf'.format(alias=alias)

    nginx_src = '/srv/{alias}/conf/{alias}.nginx'.format(alias=alias)
    nginx_dest = '/etc/nginx/sites-enabled/{alias}'.format(alias=alias)

    # Remove any existing links at the destinations
    with settings(warn_only=True):
        sudo('rm {dest}'.format(dest=server_dest))
        sudo('rm {dest}'.format(dest=nginx_dest))

        # Remove the default site from nginx
        sudo('rm /etc/nginx/sites-enabled/default')

    # Create the symbolic links
    sudo('ln -s {src} {dest}'.format(src=server_src, dest=server_dest))
    sudo('ln -s {src} {dest}'.format(src=nginx_src, dest=nginx_dest))
    sudo('initctl reload-configuration')


def restart_servers(alias):
    """
    Restarts the application and static servers.
    """
    sudo('service {alias} restart'.format(alias=alias))
    sudo('nginx -t')
    sudo('service nginx restart')


def create_pg_db(db):
    """
    Create the given database.
    """
    # TODO test to see if database already exists
    sudo('createdb {db}'.format(db=db), user='postgres')


def alter_pg_user(password, user='postgres'):
    """
    Alter the user with the given password.
    """
    sudo("""psql -c "ALTER USER {user} with password '{password}';" """.format(user=user, password=password), user='postgres')


def setup_django(alias, app):
    """
    Run django's syncdb and collectstatic in its virtualenv
    """
    python = '/srv/{alias}/env/bin/python'.format(alias=alias)
    manage = '/srv/{alias}/{app}/manage.py'.format(alias=alias, app=app)

    # Do not accept input during syncdb, we will create a superuser ourselves
    sudo('{python} {manage} syncdb --verbosity=0 --noinput'.format(python=python, manage=manage))
    sudo('{python} {manage} collectstatic --verbosity=0 --noinput'.format(python=python, manage=manage))


def update(reset=False, syncdb=False):
    """
    Update the pets application.
    """
    update_git('adopt', 'pets', reset=reset, syncdb=syncdb)
    restart_servers('adopt')


def deploy(upgrade=False):
    """
    The master deploy script.
    Examples:
    fab -H ubuntu@54.244.224.30 deploy
    fab -H root@104.131.132.143 deploy
    """
    # Ubuntu setup
    if upgrade:
        update_ubuntu()
        setup_ubuntu()
        upgrade_pip()

    # Fun starts here
    clone('https://github.com/aodin/adopt.git')
    setup_env('adopt')
    server_ln('adopt')
    create_pg_db('adopt')
    setup_django('adopt', 'pets')

    # Restart
    restart_servers('adopt')
