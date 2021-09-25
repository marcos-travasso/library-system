
# library-system 📚
This library system was developed using Go and SQLite, its purpose is to serve a future application through a web service API. It was used the [Gin web framework](https://github.com/gin-gonic/gin) to build the route requests.

## Contents
- [library-system](#library-system-)
	- [Contents](#contents)
	- [API endpoints](#api-endpoints)
		- [/users](#users)
		- [/books](#books)
		- [/lendings](#lendings)
	- [Struct's format](#structs-format)
		- [User](#user)
		- [Book](#book)
		- [Lending](#lending)

## API endpoints
Once the api.go code is build, the following routes are opened:
>See below the JSON format used in the request's body.

###  `/users`
 - `GET` - Get a list of all created users, returned as JSON.
 - `POST` - Add a new user from request data sent as JSON.
 - `PATCH` - Update a specific created user information with the new JSON data.
### `/users/:id`
- `GET`- Get a user by its ID, returned as JSON.
- `DELETE`- Delete a user by its ID.

###  `/books`
 - `GET` - Get a list of all books, returned as JSON.
 - `POST` - Add a new book from request data sent as JSON.
 - `PATCH` - Update a specific created book information with the new JSON data.
### `/books/:id`
- `GET`- Get a book by its ID, returned as JSON.
- `DELETE`- Delete a book by its ID.

###  `/lendings`
 - `GET` - Get a list of all lendings, returned as JSON.
 - `POST` - Add a new lending from request data sent as JSON.
### `/lendings/:id`
- `GET`- Get a lending by its ID, returned as JSON.
- `PATCH` - Return the sent ID's lending.


## Struct's format
Some requests must include in its body the JSON with data using this formats:
>Users and books `PATCH` requests must include the ID in the JSON.
>
### User
    {
	    "id":0,
		"person":{
			"id":0,
			"name":"",
			"gender":"",
			"birthday":""
		},
		"cellNumber":"",
		"phoneNumber":"",
		"cpf":"",
		"email":"",
		"address":{
			"id":0,
			"number":0,
			"cep":"",
			"city":"",
			"neighborhood":"",
			"street":"",
			"complement":""
		}
	}

### Book

    {
    	"year":0,
    	"pages":0,
    	"title":"",
    	"author":{
    		"person":{
    			"name":"",
    			"gender":"",
    			"birthday":""
    		}
    	},
    	"genre":{
    		"name":""
    	}
    }
### Lending

    {
    	"user":{
    		"id":0
    	},
    	"book":{
    		"id":0
    	},
    	"lendDay":"",
    	"returned":0,
    	"devolution":{
    		"id":0,
    		"date":""
    	}
    }
