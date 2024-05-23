docker build -t sexualwhitechocolate/server -f dockerfile_server .
docker build -t sexualwhitechocolate/client -f dockerfile_client .

docker push sexualwhitechocolate/server
docker push sexualwhitechocolate/client

kubectl apply -f server.yaml
kubectl apply -f service.yaml
kubectl apply -f client.yaml

kubectl apply -f gateway.yaml
kubectl apply -f vrservice.yaml
kubectl apply -f extservice.yaml
