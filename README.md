# todo-list

## __Desktop API__

Приложение для составления задач (todo-list). 
Стек: 
- Go
- JS
- Wails (framework)

## __Для запуска приложения__ 

`make run` 

> [!WARNING]
> В случае, если команда не сработала, надо будет по очередно выполнить следующие комманды:
> `docker-compose up -d`
> `go run ./cmd`
> `cd myproject`
> `wails dev`
 
## __Архитектура приложения__

![image](https://github.com/user-attachments/assets/3e170281-02bc-4223-a029-0f3fb9d83cc4)

Приложение реализована с помощью гексагональной архитектуры, которую часто применяют в микросервисных приложениях (см. gRPC microservices). 
В качестве базы данных использовался postgres, запускаемый через docker-compose, и миграции реализованы с помощью библиотеки goose.


## __Use cases__

Внешний вид API

![image](https://github.com/user-attachments/assets/7183954a-bdf8-4d14-9124-5eb61ddf10b6)


