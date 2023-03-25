# Домашнее задание

## Настройки
Для работы сервиса необходимо добавить файл checkout/config.yml. Пример файла расположен в той же директории. На месте поля токен должен быть уникальный токен для доступа к сервису продуктов

## Общие команды
Общие команды запускаются из корня.
```bash
    make prepare-proto # Запускает процесс генерации прото-файлов
    make build-all # Билдит все сервисы в монорепе
    make run-all # Запускает все сервисы с помощью docker-compose
```
# TODO

* create ci file
* take env variables from arguments and .env
* run database on remote server