# -*- coding: utf-8 -*
from django.shortcuts import render

from pets.models import Pet    


def random_pet(request):
    """
    Pull a random pet from the database. The pet's removed field must be NULL.
    """
    try:
        pet = Pet.objects.filter(removed__isnull=True).order_by('?')[0]
    except IndexError:
        pet = None

    attrs = {
        'pet': pet,
    }

    return render(request, 'random.html', attrs)
