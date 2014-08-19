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


def clone(url, directory='/srv', user='www-data'):
    """
    Clone the given repo into the /srv directory by default.
    """
    with cd(directory):
        sudo('chown {user}:{user} .'.format(user=user))
        sudo('hg clone {url}'.format(url=url), user=user)


def setup_env(alias, directory='/srv', user='www-data'):
    """
    Create the python virtual environment.
    """
    path = os.path.join(directory, alias, 'env')

    # Create the virtual environment
    sudo('virtualenv {path}'.format(path=path), user=user)

    # TODO dependencies should be in a file
    sudo('{path}/bin/pip install -r {path}/requirements.txt'.format(dir=directory), user=user)


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


def restart_servers():
    """
    Restarts the servers.
    """
    sudo('service hello restart')
    sudo('nginx -t')
    sudo('service nginx restart')


def alter_pg_user(password, user='postgres'):
    """
    Alter the user with the given password.
    """
    sudo("""psql -c "ALTER USER {user} with password '{password}';" """.format(user=user, password=password), user='postgres')


def update():
    """
    """
    pass


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
    clone('https://aodin@github.com/adopt')
    setup_env('adopt')
    server_ln('adopt')
    alter_pg_user('*#(Nrc3wmacwioUHO')

    # Restart
    restart_servers()
