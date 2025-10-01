# auth

## Build

> [!NOTE]  
> SERVICE_PORT environment variable not supported now (for REST).  

Example `.env` file:  

```env
ALPINE_IMAGE_VERSION=3.22
SERVICE_PORT=<your_service_port>
GRPC_PORT=<your_grpc_service_port>
JWT_SECRET=<your_secret_key_for_jwt>
CONFIG_PATH=/etc/auth_service/conf.yaml

BACKEND_USER=backend
BACKEND_USER_PASSWORD=<your_password_for_backend_user>

REDIS_VERSION=8.2.1
REDIS_PORT=6379
REDIS_PASSWORD=<your_redis_password>
REDIS_USER=$BACKEND_USER
REDIS_USER_PASSWORD=$BACKEND_USER_PASSWORD

POSTGRES_VERSION=18.0
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=<postgres_password>
POSTGRES_DB=users
```
