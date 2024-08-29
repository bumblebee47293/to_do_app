## Go Todo App Backend

This project implements the backend for a simple todo application written in Go. It allows users to create, view, mark as complete, and delete tasks.

### Features

* Create new tasks
* View all existing tasks
* Mark tasks as complete
* Delete tasks

### Dependencies

* This project requires Go installed on your system. Download and install from the official website: [https://go.dev/doc/install](https://go.dev/doc/install)

### Running the application

1.  **Clone the repository:**

```bash
git clone https://github.com/bumblebee47293/to_do_app.git
```

2.  **Navigate to the project directory:**

```bash
cd to_do_app
```

3.  **Install dependencies:**

```bash
go mod download
```

4.  **Run the application:**

```bash
go run main.go
```

By default, the application will run on port 8080. You can access the API endpoints at `http://localhost:8080/tasks` (depending on the configuration in your code).

**Note:** This is a basic example and might require further configuration depending on your specific implementation.

### API Endpoints

* **GET /tasks:** Retrieves a list of all tasks.
* **POST /tasks:** Creates a new task. The request body should be in JSON format containing the task description.
* **PUT /tasks/:id:** Marks a task as complete. The `:id` represents the unique identifier of the task.
* **DELETE /tasks/:id:** Deletes a task. The `:id` represents the unique identifier of the task.

**Note:** These are common endpoints for a basic todo application. Your specific implementation may have different endpoints or require additional information in the request body.

### Development

Feel free to modify the code to add additional functionalities like:

* User authentication
* Task prioritization
* Due dates for tasks
* Persistence using a database

### Contributing

We welcome contributions to this project! Please create a pull request on the Github repository with your changes.

### License

This project is licensed under the MIT License. See the LICENSE file for details.
