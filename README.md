# go-api
A RESTful todo list API written in Go on top of the [Revel web framework](https://revel.github.io/).

# running
Start the app with `revel run`. Unless changed in [app.conf](go-api/conf/app.conf), the app will start on port 9000 by default.

The following endpoints are available. Use a tool like [Postman](https://www.getpostman.com/) to exercise them.

* `GET /up` returns a static HTML page to show that the app is up and running.
* `GET /todos` gets all of the todo list items.
* `GET /todos/<uuid>` gets a specific todo item by its ID.
* `POST /todos` creates a todo item. At minimum, the `Title` must be supplied in the body of the request.
* `PATCH /todos/<uuid>` edits the specified todo item.
* `DEL /todos/<uuid>` removes a todo list item by ID.

# testing
Run `revel test github.com/natejenson/go-api` to run the unit tests.
