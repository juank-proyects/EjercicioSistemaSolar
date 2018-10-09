# EjercicioSistemaSolar
Ejercicio Sistema Solar, dado 3 planetas que se mueven al rededor del sol, calcular el clima, dependiendo este de la posición de los planetas respecto al sol.
# Documentacion
* La Documentacion se encuentra en el siguiente [enlace](https://docs.google.com/document/d/1hUhcSEUILK8SX_d8FT5Otsfzxt5GCJYm9hW94ZpxWuQ/edit#heading=h.kk1966kbedef)
* Las tareas del proyecto se encuentran en el siguiente [enlace](https://github.com/juank-proyects/EjercicioSistemaSolar/projects/1)
# Api Demo en AWS
* El proyecto esta deployado en Amazon, junto con una BD en Postgres en el cual se responden a las preguntas planteadas en el encunciado del problema:
- 1 ¿Cuántos períodos de sequía habrá? <br />
   http://ec2-18-191-229-56.us-east-2.compute.amazonaws.com:8080/periodos_de_sequia
- 2 ¿Cuántos períodos de lluvia habrá y qué día será el pico máximo de lluvia? <br />
   http://ec2-18-191-229-56.us-east-2.compute.amazonaws.com:8080/periodos_de_lluvia
- 3 ¿Cuántos períodos de condiciones óptimas de presión y temperatura habrá? <br />
   http://ec2-18-191-229-56.us-east-2.compute.amazonaws.com:8080/periodos_optimos
- Para obtener el clima de un determinado dia <br />
   http://ec2-18-191-229-56.us-east-2.compute.amazonaws.com:8080/clima?dia=1
- Para ejecutar el cronjob Manualmente 
   http://ec2-18-191-229-56.us-east-2.compute.amazonaws.com:8080/mover_galaxia
# Instalacion
- (Se requiere tener instalado go1.11.1 o superior)
```sh
      go get github.com/go-pg/pg
      go get github.com/gin-gonic/gin
      go get github.com/juank-proyects/EjercicioSistemaSolar
```
- Para levantar la Api del proyecto en localhost:9990 ejecutar: 
```sh
      go run main.go 
```
- Crear una BD en postgres con los siguiente datos: <br />
tambien es posible modificar estas configuraciones en EjercicioSistemaSolar/src/db/db.go

```go
      Addr:     "localhost:5432",
      User:     "postgres",
      Password: "postgres",
      Database: "clima",
```
  
- Finalmente para crear las tablas correspondientes en la BD y ejecutar manualmente el Job correr los siguientes request: <br />
```sh
   localhost:9990/crear_esquema
   localhost:9990/mover_galaxia
```
# Test
- Dentro del folder EjercicioSistemaSolar/src/test/ correr en la consola:
```sh
   go test
```
