# todo-list

## __Desktop API__

Приложение для составления задач (todo-list). 
Стек: 
- Go
- JS
- Wails (framework)

## __Для запуска приложения__ 

Введите: `make run` 

> [!WARNING]
> В случае, если команда не сработала, надо будет по очередно выполнить следующие комманды:
> `docker-compose up -d`
> `go run ./cmd/`
> `cd myproject`
> `wails dev`
 
## __Архитектура приложения__

![image](https://github.com/user-attachments/assets/3e170281-02bc-4223-a029-0f3fb9d83cc4)

Приложение реализована с помощью гексагональной архитектуры, которую часто применяют в микросервисных приложениях (см. gRPC microservices). 
В качестве базы данных использовался postgres, запускаемый через docker-compose, и миграции реализованы с помощью библиотеки goose.


## __Use cases__

Внешний вид API

![image](https://github.com/user-attachments/assets/7183954a-bdf8-4d14-9124-5eb61ddf10b6)

__После нажатия add task появляется надпись: "Задача успешно добавлена"__

![image](https://github.com/user-attachments/assets/0bc86669-61af-4997-8f6f-71c1425d9123)

__Логи на сервере:__

![image](https://github.com/user-attachments/assets/b3bbf9c4-3e1a-4dc0-a2fa-d111c86ce959)

__После нажатия на show tasks, можно будет посмотреть на добавленные задачи:__

![image](https://github.com/user-attachments/assets/bec33b08-a910-4311-8c5a-164a44f3d76f)

__Задачу можно изменить на выполненное, если нажать на кнопку "Mark as done" или удалить задачу, нажав "Delete"__
