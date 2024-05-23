(Исправил комментарии ко второй дз: добавил скрипты установки istio, настроил, поправил багу в коде сервера (заводилась мапа из стринга в инт при считывании по url-у, ошибка копипасты с кода клиента))

## Desc:

client.go - программа на go, дергающая у service ручку statistics, и записывающая результат в statistics.txt, создавая файл если необходимо

client.yaml - клиентский манифест, говорит откуда брать образ и какой порт вскрывать

dockerfile_client - докерфайл, который запускается из golang:1.22, кладет к себе client.go, и запускает его

dockerfile_server - аналогично

extservice.yaml - заполняем ServiceEntry, как на семинаре

gateway.yaml - заполняем Gateway, как на семинарах

server.go - программа на go, создающая по 5000-му порту ручки /time и /statistics, и запускающая сервер с такими ручками (ТЕПЕРЬ она дергает http://worldtimeapi.org/api/timezone/Europe/Moscow при вызове /time!)

server.yaml - аналогично client.yaml

service.yaml - файл, описывающий сервис в кубернетесе

start.sh - пусковой скрипт

vrservice.yaml - заполняем VirtualService, как на семинарах

## Steps:

0.1. curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.22.0 TARGET_ARCH=x86_64 sh - (качаем istio)

0.2. cd istio-1.22.0 (папка установки)

0.3. export PATH=$PWD/bin:$PATH (поправляем путь, чтобы виден был istio)

0.4. istioctl install --set profile=demo -y --set meshConfig.outboundTrafficPolicy.mode=REGISTRY_ONLY (устанавливаем istio)

0.5. kubectl label namespace default istio-injection=enabled (правим лейбл)

Дальше как в hw_1:

1. minikube start

2. docker login (при запуске заменить пользователя в файлах start.sh, server.yaml, client.yaml)

3. sh start.sh

4. minikube tunnel

5. minikube service service

6. Поспамить пару-тройку раз /time в появившемся окне браузера (или посмотреть url, и вызваться от него явно, если браузер не открылся)

7. kubectl get pods (найти полное название client пода)

8. kubectl exec -it полное_название_client_пода -- cat statistics.txt (убедиться что число вызова ручки /time отображается корректно)
