(Исправил комментарии по первому дз - развел дз по разным веткам, скрыл порты клиента, использую golang:1.22-alpine)

## Desc:

client.go - программа на go, дергающая у service ручку statistics, и записывающая результат в statistics.txt, создавая файл если необходимо

client.yaml - клиентский манифест, говорит откуда брать образ и какой порт вскрывать

dockerfile_client - докерфайл, который запускается из golang:1.22, кладет к себе client.go, и запускает его

dockerfile_server - аналогично

server.go - программа на go, создающая по 5000-му порту ручки /time и /statistics, и запускающая сервер с такими ручками

server.yaml - аналогично client.yaml

service.yaml - файл, описывающий сервис в кубернетесе

start.sh - пусковой скрипт

## Steps:

1. minikube start

2. docker login (при запуске заменить пользователя в файлах start.sh, server.yaml, client.yaml)

3. sh start.sh

4. minikube tunnel

5. minikube service service

6. Поспамить пару-тройку раз /time в появившемся окне браузера (или посмотреть url, и вызваться от него явно, если браузер не открылся)

7. kubectl get pods (найти полное название client пода)

8. kubectl exec -it полное_название_client_пода -- cat statistics.txt (убедиться что число вызова ручки /time отображается корректно)
