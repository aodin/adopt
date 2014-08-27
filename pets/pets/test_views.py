# -*- coding: utf-8 -*
import json

from django.core.urlresolvers import reverse
from django.test import TestCase

from pets.models import Pet


class ViewTests(TestCase):
    """
    Test the context of the views.
    """
    def setUp(self):
        # Add some pets
        Pet.objects.bulk_create([
            Pet(id="1", name="Meow", type="cat"),
            Pet(id="2", name="Bark", type="dog"),
            Pet(id="A3", name="", type="rat"),
        ])

    def test_random(self):
        """
        Test the random pet landing page.
        """
        random_url = reverse('random_pet')
        resp = self.client.get(random_url)
        pet = resp.context['pet']
        
        # The pet should not be None
        self.assertTrue(pet)

        # Get a specific type, it should be case insensitive
        resp = self.client.get(random_url, {'type': "Cat"})
        pet = resp.context['pet']
        self.assertEqual(pet.id, "1")

        # Test JSON output
        resp = self.client.get(random_url, {'type': "Cat", 'format': "JSON"})
        json_resp = json.loads(resp.content)
        self.assertEqual(json_resp['id'], "1")


    def test_all(self):
        """
        Test the list view and its filters.
        """
        all_url = reverse('all_pets')
        resp = self.client.get(all_url)
        pets = resp.context['pets']
        self.assertEqual(len(pets), 3)

        # Get a specific type, it should be case insensitive
        resp = self.client.get(all_url, {'type': "Cat"})
        pets = resp.context['pets']
        self.assertEqual(len(pets), 1)

        # Test JSON output
        resp = self.client.get(all_url, {'format': "JSON"})
        json_resp = json.loads(resp.content)
        self.assertEqual(len(json_resp), 3)
