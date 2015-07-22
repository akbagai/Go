# Go Server

Serve a GeoJSON file as a REST Endpoint. Allows cross domain requests.

You must create an Environment Variable called JSONServerFile with the full filename and path (ie.. C:/myfiles/JSON/somedata.geojson). See [Here](http://www.computerhope.com/issues/ch000549.htm) for instructions on how to do that. For the /sql route, you need a DB and a connection.

You can modify the souce code and build your own as well.
To serve more files, create another route and attach another handler.

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
Added a subset route that returns the specified number of records
```
http://localhost:8080/10
```
