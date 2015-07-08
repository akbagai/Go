# Go Server

Serve a GeoJSON file as a REST Endpoint. Allows cross domain requests. Change path to the GeoJSON file to use your own. Create an executable using Build and Install.

To add a new file, create another handler route.

Code taken from [Making a RESTful JSON API in Go](http:////thenewstack.io/make-a-restful-json-api-go/) and modified to read json file and allow cross domain requests.

Server runs on
```
http://localhost:8080
```
Added a Query route that currently only works on strings. Takes format of /Key/Value
```
http://localhost:8080/ART_CODE/101
http://localhost:8080/YEAR/1985
```
