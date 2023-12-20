# ToDoApp

## Описание
Сервис **ToDoApp** предназначен для работы с задачами (list) и подзадачами к ним (task).

Для запуска в терминале выполнить следующую команду:  
`docker-compose up --build todo-app`

При первом запуске необходимо сделать миграцию базы данных с помощью команды:  
`migrate -path ./schema -database 'postgres://postgres:password@localhost:5436/postgres?sslmode=disable' up`

## Функционльность проекта
### Функциональность эндпонта /auth

**POST /auth/sign-up** - регистрация пользователя и генерация токена доступа для него

**POST auth/sign-in** - авторизация пользователя с использованием токена доступа 

### Функциональность эндпонта /api/lists

**GET /api/lists/** - получение всех задач пользователя

**GET /api/lists/:id** - получение задачи по его id

**POST /api/lists/** - создание новой задачи

**PUT /api/lists/:id** - обновление задачи по его id

**DELETE /api/lists/:id** - удаление задачи по его id

**GET /api/lists/:id/tasks/** - получение всех подзадач задачи по id задачи

**POST /api/lists/:id/tasks/** -  добавление подзадачи в задачу по id задачи


### Функциональность эндпонта /api/lists/tasks

**GET /api/lists/tasks/:id** - получение подзадачи задачи по id подзадачи

**PUT /api/lists/tasks/:id** - обновление подзадачи задачи по id подзадачи

**DELETE /api/lists/tasks/:id** - удаление подзадачи задачи по id подзадачи


## Схема базы данных

![todoapp_DB_scheme](https://github.com/ciiska5/todo-app/todo_app_diagram.png)

Она состоит из 5 таблиц:
| Название таблицы | Описание |
| --- | --- |
| **users** | таблица с пользователями |
| **users_lists** | таблица с пользователями и задачами (lists) |
| **lists** | таблица с задачами (lists) |
| **tasks** | | таблица с подзадачами (tasks) |
| **lists_tasks** | таблица с задачами (lists) и подзадачами (tasks) |
