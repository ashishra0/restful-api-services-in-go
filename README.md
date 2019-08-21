### Building RESTful API and RESTful services in Go

![go server](servergo.png)
---

Part - I <h3>Getting Started</h3>

* Create and start a web server
* Create a default route
* Send message via default route
* Create custom error message

---

Part - II <h3>Making API RESTful</h3>

*Api outline*

* Resource - A list of users
* List item data - name, role, id
* Functionality - CRUD operations
* Endpoints - /user, 
/users/{id}
* Data format - JSON

Create a user
* Method - POST
* Target - collection
* Endpoint - POST/users
* Request content - full item data
* Successful response - 201 created + location

Access users
* Method - GET
* Target - collection
* Endpoint - GET /users
* Request content - none
* Successful response - 200 OK + list of users
 
Access individual user
* Method - GET
* Target - item
* Endpoint - GET /users/{id}
* Request content - none
* Successful response - 200 OK + user data
* Missing resource - 404 Not found

Replace an user
* Method - PUT
* Target - item
* Endpoint - PUT /users/{id}
* Request content - full item data
* Successful response - 200 OK + new item data
* missing resource - 404 not found

Update an user
* Method - PATCH
* Target - item
* Endpoint - PATCH /users/{id}
* Request content - partial item data
* Successful response - 200 OK + new item data
* missing resource - 404 not found

Delete an user
* Method - DELETE
* Target - item
* Endpoint - DELETE /users/{id}
* Request content - none
* Successful response - 200 OK
* missing resource - 404 not found
---