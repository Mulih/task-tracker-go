# Task Tracker CLI Application

The **Task Tracker CLI Application** is a command-line tool for managing tasks in a to-do list. It allows users to add, list, update, delete, and retrieve tasks. Tasks are stored in a JSON file for persistence.

## Features

- **Add Tasks**: Create new tasks with a description and status.
- **List Tasks**: View all tasks with their details.
- **Update Tasks**: Modify the description or status of an existing task.
- **Delete Tasks**: Remove tasks by their ID.
- **Get Task by ID**: Retrieve a specific task by its unique ID.

## File Structure

```go
task-tracker-go/
├── main.go               # Entry point of the application
├── tasks/
│   ├── add.go            # Logic for adding tasks
│   ├── delete.go         # Logic for deleting tasks
│   ├── list.go           # Logic for listing and retrieving tasks
│   ├── update.go         # Logic for updating tasks
│   ├── models.go         # Task and TaskList definitions
├── tasks.json            # JSON file for storing tasks
├── go.mod                # Go module file
├── README.md             # Documentation
```

## Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd task-tracker-go
   ```

2. Install Go (if not already installed). Refer to the [official Go installation guide](https://golang.org/doc/install).

3. Run the application:

   ```bash
   go run main.go <command>
   ```

## Usage

### Commands

1. **Add a Task**

   ```bash
   go run main.go add <description> <status>
   ```

   - `<description>`: A brief description of the task.
   - `<status>`: The status of the task (`todo`, `in progress`, or `done`).

   Example:

   ```bash
   go run main.go add "Buy groceries" "todo"
   ```

2. **List All Tasks**

   ```bash
   go run main.go list
   ```

3. **Update a Task**

   ```bash
   go run main.go update <id> <description> <status>
   ```

   - `<id>`: The ID of the task to update.
   - `<description>`: The new description of the task.
   - `<status>`: The new status of the task.

   Example:

   ```bash
   go run main.go update 1 "Buy groceries and snacks" "in progress"
   ```

4. **Delete a Task**

   ```bash
   go run main.go delete <id>
   ```

   - `<id>`: The ID of the task to delete.

   Example:

   ```bash
   go run main.go delete 1
   ```

5. **Get a Task by ID**

   ```bash
   go run main.go get <id>
   ```

   - `<id>`: The ID of the task to retrieve.

   Example:

   ```bash
   go run main.go get 1
   ```

## Task Data Format

Tasks are stored in a JSON file (`tasks.json`) in the following format:

```json
{
  "tasks": [
    {
      "id": 1,
      "description": "Buy groceries",
      "status": "todo",
      "created_at": "2025-04-02T08:00:06.180792033+03:00",
      "updated_at": "2025-04-02T08:00:06.180792346+03:00"
    }
  ],
  "next_id": 2
}
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
