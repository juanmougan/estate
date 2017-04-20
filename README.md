# estate
Microservice API to manage real estate. Based on https://thenewstack.io/make-a-restful-json-api-go/  
and also https://www.mongodb.com/blog/post/running-mongodb-queries-concurrently-with-go

Use a valid project structure, check https://golang.org/doc/code.html for details

E. g. `[some folder]/src/github.com/juanmougan/estate` where `estate` is this project's root path

Run the project like this:

    $ go run *.go

Check the API using curl:

    curl 'http://localhost:8080/estates' | jq .

Should get back something like this:

    [
      {
        "name": "Departamento en Palermo",
        "lat": -34.594142,
        "long": -58.422036,
        "created": "0001-01-01T00:00:00Z"
      },
      {
        "name": "Casa en San Isidro",
        "lat": -34.46761,
        "long": -58.510191,
        "created": "0001-01-01T00:00:00Z"
      }
    ]

And if you want to add a new Estate:

    curl -H "Content-Type: application/json" -d '{"name":"Casa de Maria", "Lat": -34.566610, "Long": -58.411191}' http://localhost:8080/estates | jq .

You would get something like this:

    {
      "id": 4,
      "name": "Casa de Maria",
      "lat": -34.56661,
      "long": -58.411191,
      "created": "0001-01-01T00:00:00Z"
    }

Steps missing

1. Add missing endpoint

2. Add MongoDB support

Some data about setting up Mongo: https://docs.mongodb.com/manual/tutorial/install-mongodb-on-os-x/

Start service: `brew services start mongodb`

From now, data from Mykong: https://www.mkyong.com/mongodb/how-to-create-database-or-collection-in-mongodb/
Run `show dbs` from console
Currently, only two databases are available – “admin” and “local“.

Run `use <new-db>` but MangoDB doesn’t create any database yet, until you save something inside.
E.g. `use places`

Define a collection named “users“, and save a dummy document(value) inside.

> db.users.save( {username:"mkyong"} )
> db.users.find()
{ "_id" : ObjectId("4dbac7bfea37068bd0987573"), "username" : "mkyong" }
>
> show dbs
admin   0.03125GB
local   (empty)
mkyongdb        0.03125GB

Mapping chart: https://docs.mongodb.com/manual/reference/sql-comparison/

My example:

> show dbs
admin  0.000GB
local  0.000GB
> use places
switched to db places
> show dbs
admin  0.000GB
local  0.000GB
>
db.estates.save( {id:1, name: "Departamento en Colegiales", lat: -34.575831, long: -58.448499} )
db.estates.save( {id:2, name: "Departamento en Coghlan", lat: -34.562559, long: -58.474026} )

With no args, find() retrieves all documents
db.estates.find()

This retrieves id = 2
db.estates.find( { id: 2 } )

3. Eventually... DB cache?
