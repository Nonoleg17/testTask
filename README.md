# Go Test Issue

Задача:Необходимо разработать сервис и спроектировать DB (PostgreSQL).

### Инструкция для запуска
Для развертывания сервиса необходимо прописать в терминале корневой папке команду: 
```console
sudo docker-compose up
```
Если необходимо развернуть только базу данных, то необходимо прописать в терминале корневой папке команду:
```console
sudo docker-compose  -f docker-compose.envonly.yml up
```
### Структура проекта
UML Диаграмма БД представлена на рисунке:


На рисунке представлена связи между сущностями. Между users и friendships связь many to many, также между orders и product связь many to many
![](pictures/Screenshot from 2023-01-29 12-39-04.png)

