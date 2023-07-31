# GoMicroservices 💻

¡Hola! Bienvenido a este proyecto en el que he estado trabajando. Es el resultado de mi aprendizaje sobre microservicios en Go. En este proyecto encontrarás una serie de microservicios interconectados, todos escritos en Go, utilizando algunas bibliotecas como Gin y otras herramientas. 

## 🚀 Microservicios desarrollados

- **broker-service**: Punto de entrada único opcional para conectarse a todos los servicios desde un solo lugar (acepta JSON, envía JSON, hace llamadas a través de gRPC y empuja a RabbitMQ).
- **authentication-service**: Autentica a los usuarios contra una base de datos de Postgres (acepta JSON).
- **logger-service**: Registra eventos importantes en una base de datos MongoDB (acepta RPC, gRPC y JSON).
- **queue-listener-service**: Consume mensajes de AMQP (RabbitMQ) e inicia acciones basadas en la carga útil (envía a través de RPC).
- **mail-service**: Envía correos electrónicos (acepta JSON).

Todos los servicios (excepto el broker) registran sus URLs de acceso con etcd y renuevan sus arrendamientos automáticamente. Esto nos permite implementar un sistema de descubrimiento de servicios sencillo, donde todas las URLs de los servicios son accesibles con "mapas de servicios" en el tipo Config utilizado para compartir la configuración de la aplicación en el servicio de intermediario.

## Herramientas Utilizadas 🛠️

Algunas de las herramientas y tecnologías utilizadas en este proyecto incluyen:

- **Go:** Lenguaje de programación principal para todo el proyecto.
- **PostgreSQL:** Sistema de gestión de bases de datos para almacenar y gestionar los datos.
- **Docker:** Utilizado para contenerizar las bases de datos y otros servicios que se necesiten.
- **RabbitMQ:** Mensajería de cola utilizada para manejar la comunicación entre los servicios.
- **Gin:** Marco de trabajo web HTTP de alto rendimiento utilizado en varias partes del proyecto.
- **gRPC-go:** Marco de trabajo utilizado para implementar la comunicación entre los servicios.
- **Mailhog:** utilizado como un servidor de correo falso para trabajar con el servicio de correo.
- **etcd:** Sistema de almacenamiento de clave-valor utilizado para el descubrimiento y configuración de servicios.
- **MongoDB:** Base de datos NoSQL utilizada para almacenar los registros de los eventos importantes.

## 📦 Ejecución del proyecto

```shell
# Clona el repositorio
git clone https://github.com/vazzquex/GoMicroservices.git

# Ve al directorio del proyecto
cd GoMicroservices
```

Desde el nivel raíz del proyecto, ejecuta este comando (esto asume que tienes [GNU make](https://www.gnu.org/software/make/) y una versión reciente de [Docker](https://www.docker.com/products/docker-desktop), [Docker-Compose](https://docs.docker.com/compose/)  instalada en tu sistema):

```shell
make up_build
```
Si el código no ha cambiado, las ejecuciones posteriores pueden ser simplemente make up.

Luego inicia el front end:
```shell
make start
```
Puedes acceder al front end con tu navegador web en http://localhost:8080.

Para detener todo:

```shell
make stop
make down
```

Mientras trabajas en el código, puedes reconstruir solo el servicio en el que estás trabajando ejecutando
```shell
make auth
```

Donde auth es uno de los servicios:
- **auth**
- **broker**
- **logger**
- **listener**
- **mail**

Cualquier cosa puedes entrar al archivo makefile donde podés ver todos los comandos

## Contribuciones 🤝
Las contribuciones a este proyecto son muy bienvenidas. Si tienes ideas para mejorar el código, agregar nuevas funcionalidades, o simplemente quieres practicar tus habilidades con Go, no dudes en abrir un 'issue' o un 'pull request'.
