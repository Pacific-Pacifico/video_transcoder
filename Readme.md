docker run -d --hostname my-rabbitmq --name my-rabbitmq -p 15672:15672 -p 5672:5672 rabbitmq:3-management  
or  
docker start my-rabbitmq  
export RMQ_ENV="amqp://guest:guest@localhost:5672/"  
