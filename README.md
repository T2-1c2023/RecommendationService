# NotificationsService
Walking Skeleton para el microservicio de notificaciones desarrollado con `Golang` y `Gin`.
Incluye `docker`, `docker compose`, tests unitarios con `testify`, linter, pipeline CI-CD con `Github Actions`. 


# Ejecución con Docker

### Instalación de docker
Debe estar instalado docker y docker-compose

### Ejecutar app local en docker
Ejecutar el script (ya está configurado para que compile y ejecute la aplicación - configurado para leer el archivo docker-compose.dev.yml):
```
sh start-docker-dev-app.sh
```

### Comprobar que la app está corriendo con health endpoint:
```
localhost:13002/health
curl localhost:13002/health
```
Obs: Notar que ahora el puerto es el 13002 por el mapeo de puertos que está en el archivo docker-compose.dev.yml


# Comandos auxiliares Docker

### Listar containers que se están ejecutando
```
docker container ps
```

### Ver imágenes de docker creadas
```
docker images
```

### Detener imagen docker
```
docker ps
docker stop <container_id>
```

### Eliminar imagenes de docker
```
docker images
docker rmi -f <image_id>
```

# Ejecución Local

Si se desea ejecutar localmente la aplicación, ya sea para desarrollar o correr los tests


### Instalación
Se debe tener instalado golang (versión 1.20, que es la última versión estable)
```
go mod download
```

### Ejecutar linter:
```
go vet ./...
```

### Formatear código:
```
go fmt ./...
```

### Ejecutar tests unitarios:
```
go test ./__test__
```

### Ejecutar app localmente sin docker:
```
sh start-local-app.sh
```


# Convenciones de codificación

### Idioma
- Comentarios en el código: español
- Comentarios en commit: español
- Nombre de clases, funciones, variables y endpoints: inglés

### Convenciones de nombres
- Archivos: kebab-case
- Structs: PascalCase
- Funciones: camelCase
- Variables: camelCase
- Constantes: MAYÚSCULAS

# Documentación

La documentación se puede ver al ejecutar el microservicio, y llamar al endpoint de `/api-docs/index.html`.
 - De manera local: localhost:3002/api-docs/index.html
 - En docker local: localhost:13002/api-docs/index.html

Para generar nueva documentación, deben agregarse las debidas anotaciones a la función del controller adecuado, y luego correr `sh generate-swagger-files.sh`.

# Convenciones para trabajar/pushear

- Antes de pushear, asegurarse de que pase el linter y las pruebas locales con el siguiente comando
```
sh test-local-app.sh
```

# Modelo de Branching
- Trabajo sobre el repo: Features-Branches
- Es decir, trabjar sobre ramas por features y al terminar el desarrollo del features, mergear contra master. Idealmente, las branches deberían tener corta duración (1 ó 2 días), para mantener la integración continua y que no sea un caos el merge.
