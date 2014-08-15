# -*- coding: utf-8 -*
from django.conf.urls import patterns, include, url
from django.contrib import admin
from rest_framework import viewsets, routers

from pets.api import PetViewSet


admin.autodiscover()

router = routers.DefaultRouter()
router.register(r'pets', PetViewSet)

urlpatterns = patterns('',
    url(r'^', include(router.urls)),
    url(r'^admin/', include(admin.site.urls)),
)
