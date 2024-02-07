# Pre Prod Fixer

Planificador de planes de rodaje.

## Problema a resolver

Para realizar el rodaje de un producto audiovisual, es necesario realizar un 
**plan de rodaje**, que determina qué planos va a rodar un equipo específico
en una localización específica a unas horas específicas. La creación de este
documento suele plantear múltiples problemas, como el poder cuadrar la
disponibilidad de actores y localizaciones, tener en cuenta posibles
retrasos o cambios en el tiempo atmosférico. Además, muchas veces ese documento
varía a lo largo del rodaje, ya que los retrasos implican reprogramar planos
en otros momentos.

## Solución

A partir de los datos proporcionados por una productora (equipos, actores,
planos, ...) generar diferentes propuestas para planes de rodaje atendiendo
a preferencias específicas y que pueda ser modificado a lo largo del tiempo
conforme va avanzando el rodaje.

## Dependencias

Este proyecto tiene las siguientes dependencias para su creación:

 - [*Go*](https://go.dev/)
 - [*Testify*](https://github.com/stretchr/testify)
 - [*Task*](https://taskfile.dev/)

## Órdenes

Para comprobar el validez sintáctica de los archivos del proyecto, ejecutar:

```Bash
task check
```

Para lanzar los tests del proyecto, ejecutar:

```Bash
task test
```

## Contenedor de pruebas

Existe un contenedor de Docker para realizar pruebas. Para crear el contenedor y
ejecutar las pruebas, ejecutar:

```Bash
docker build -t diagmatrix/maybe-better-maybe-worse . && docker run -u 1001 -t -v `pwd`:/app/test diagmatrix/maybe-better-maybe-worse
```

Para ejecutar únicamente el contenedor (si está construido), ejecutar:

```Bash
docker run -u 1001 -t -v `pwd`:/app/test diagmatrix/maybe-better-maybe-worse
```

## Estado del proyecto

 1. Configuración del repositorio
    - [Configuración de git](docs/git_config.png)
    - [Configuración de clave](docs/ssh_key.png)

 2. Historias de usuario
    - [Personas](docs/personas.md)
    - [Historias de usuario](docs/historias_usuario.md)

 3. [*Milestones*](docs/milestones.md)

 4. Toolchain
    - [Criterios](docs/criterios.md)
    - [Gestor de paquetes](docs/dependencias.md)
    - [Gestor de tareas](docs/gestor_tareas.md)
    - [Tests](docs/tests.md)

 5. [*Imagen de pruebas*](docs/imagen.md)