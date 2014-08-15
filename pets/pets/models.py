# -*- coding: utf-8 -*
from django.db import models


class Pet(models.Model):
    """
    A single pet.
    """
    id = models.CharField(max_length=64, primary_key=True)
    name = models.CharField(max_length=64)
    type = models.CharField(max_length=64)
    gender = models.CharField(max_length=64)
    color = models.CharField(max_length=64)
    breed = models.CharField(max_length=64)
    age = models.CharField(max_length=64)
    location = models.CharField(max_length=128)
    detail_url = models.CharField(max_length=256)
    image_url = models.CharField(max_length=256)
    added = models.DateTimeField(auto_now_add=True)
    removed = models.DateTimeField(null=True)

    def __unicode__(self):
        return "{name}: {gender} {color} {breed}".format(
            name = self.name,
            gender = self.gender,
            color = self.color,
            breed = self.breed,
        )

    class Meta:
        db_table = "pets"
