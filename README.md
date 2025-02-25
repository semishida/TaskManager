# Task Manager

## Русский

### Описание
Task Manager - приложение для управления задачами с поддержкой многопользовательской авторизации. После авторизации каждый пользователь может создавать, редактировать, удалять и просматривать свои задачи.

### Функционал
- Авторизация и регистрация пользователей.
- Создание, редактирование, удаление задач.
- Каждый пользователь видит только свои задачи.

### Запуск
1. Склонируйте репозиторий:
   ```bash
   git clone https://github.com/<semishida>/task-manager.git
   ```
2. Установите зависимости:
   ```bash
   go mod tidy
   ```
3. Создайте файл `.env` и укажите настройки:
   ```env
   JWT_SECRET=your_jwt_secret
   DB_CONFIG=host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable
   ```
4. Запустите сервер:
   ```bash
   go run main.go
   ```

### Тестовые запросы
- Создание задачи:
  ```bash
  curl -X POST -H "Authorization: Bearer <your_token>" -d '{"group": "Work", "description": "Complete project", "deadline": 20-11-2025, "status": "Pending"}' http://localhost:8080/tasks
  ```
- Просмотр списка задач:
  ```bash
  curl -X GET -H "Authorization: Bearer <your_token>" http://localhost:8080/tasks
  ```

---

## English

### Description
Task Manager is an application for managing tasks with multi-user authentication support. After logging in, each user can create, edit, delete, and view their own tasks.

### Features
- User authentication and registration.
- Creating, editing, deleting tasks.
- Each user can only see their own tasks.

### Running the Application
1. Clone the repository:
   ```bash
   git clone https://github.com/<semishida>/task-manager.git
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Create a `.env` file with the following settings:
   ```env
   JWT_SECRET=your_jwt_secret
   DB_CONFIG=host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable
   ```
4. Start the server:
   ```bash
   go run main.go
   ```

### Example Requests
- Create a task:
  ```bash
  curl -X POST -H "Authorization: Bearer <your_token>" -d '{"group": "Work", "description": "Complete project", "deadline": 11-20-2025, "status": "Pending"}' http://localhost:8080/tasks
  ```
- View tasks list:
  ```bash
  curl -X GET -H "Authorization: Bearer <your_token>" http://localhost:8080/tasks
  ```

