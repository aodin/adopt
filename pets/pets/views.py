# -*- coding: utf-8 -*
import json

from django.core.serializers.json import DjangoJSONEncoder
from django.db.models import Q
from django.forms.models import model_to_dict
from django.http import HttpResponse
from django.shortcuts import render, get_object_or_404

from pets.models import Pet    


def json_response(q):
    """
    Return a Django queryset as a JSON encoded response.
    """
    return HttpResponse(
        json.dumps(q, cls=DjangoJSONEncoder),
        content_type="application/json",
    )


def random_pet(request):
    """
    Pull a random pet from the database. The pet's removed field must be NULL.
    Allow filtering through GET parameters.
    """
    q = Pet.objects.filter(removed__isnull=True)
    types = request.GET.get('type')
    if types:
        where = None
        for typ in types.split(','):
            if where:
                where |= Q(type__icontains=typ)
            else:
                where = Q(type__icontains=typ)
        q = q.filter(where)

    try:
        pet = q.order_by('?')[0]
    except IndexError:
        pet = None

    format = request.GET.get('format', "")
    if format.lower() == "json":
        return json_response(model_to_dict(pet))
    
    attrs = {
        'pet': pet,
    }
    return render(request, 'random.html', attrs)


def all_pets(request):
    """
    List all pets currently available. Allow filtering through GET parameters.
    """
    pets = Pet.objects.filter(removed__isnull=True)
    types = request.GET.get('type')
    if types:
        where = None
        for typ in types.split(','):
            if where:
                where |= Q(type__icontains=typ)
            else:
                where = Q(type__icontains=typ)
        pets = pets.filter(where)

    format = request.GET.get('format', "")
    if format.lower() == "json":
        return json_response([model_to_dict(p) for p in pets])
    
    attrs = {
        'pets': pets,
    }
    return render(request, 'list.html', attrs)


def single_pet(request, pet_id):
    """
    Display a single pet by id. Allow removed pets to be displayed.
    """
    pet = get_object_or_404(Pet, id=pet_id)

    attrs = {
        'pet': pet,
    }

    return render(request, 'random.html', attrs)