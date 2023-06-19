# Conversor de Temperatura

Este es un microservicio simple desarrollado en Go que permite convertir temperaturas entre diferentes unidades. Puedes enviar una solicitud POST con la temperatura y la unidad de origen, y obtendrás la temperatura convertida en Fahrenheit, Celsius y Kelvin.

## Instalación y Ejecución

1. Asegúrate de tener Go instalado en tu sistema.
2. Clona este repositorio en tu máquina local.
3. Navega hasta el directorio del proyecto.
4. Ejecuta el comando `go run main.go` para iniciar el microservicio.
5. El microservicio estará disponible en `http://localhost:8080/convert`.

## Uso

Puedes enviar una solicitud POST a `http://localhost:8080/convert` para obtener la conversión de temperatura. Los parámetros deben incluirse en el cuerpo de la solicitud utilizando el formato `x-www-form-urlencoded`. Los parámetros necesarios son:

- `temperature`: La temperatura a convertir.
- `unit`: La unidad de origen (por ejemplo, "celsius", "fahrenheit", "kelvin").

Aquí tienes un ejemplo de cómo hacer una solicitud POST utilizando cURL:

```shell
curl -X POST -d "temperature=25&unit=celsius" http://localhost:8080/convert