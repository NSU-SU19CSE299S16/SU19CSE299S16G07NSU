from django.contrib import admin
from django.urls import path, include
from django.conf.urls import url
#from users import views as user_views
from Alumni_Portal import views
from django.contrib.auth import logout
from django.contrib import admin
from django.urls import path

urlpatterns = [
    path('admin/', admin.site.urls),
    path('', include('Alumni_Portal.urls')),
    path('user/', include('user_profile.urls')),
]
