# -*- coding: utf-8 -*
from rest_framework import viewsets

from pets.models import Pet


class PetViewSet(viewsets.ModelViewSet):
    model = Pet
