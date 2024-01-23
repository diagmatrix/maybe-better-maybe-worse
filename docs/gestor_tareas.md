# Elección del gestor de tareas

El gestor de tareas es una herramienta que se encarga de automatizar procesos repetitivos como la compilación del código, optimización de los recursos computacionales y ejecución de tests.

## [Make](https://www.gnu.org/software/make/manual/make.html)

Esta herramienta permite, mediante la configuración de un archivo *Makefile*, crear reglas que
automaticen tareas como la compilación y ejecución de tests. Se trata de una herramienta ampliamente
documentada y conocida.

La sintaxis de make no es muy intuitiva y rudimentaria, en la cual el programador necesita escribir 
específicamente todas las órdenes y lógicas de las reglas que describe.

## [Just](https://github.com/casey/just)

Esta herramienta, autodescrita como ''Simplemente un ejecutor de comandos'', es una herramienta de 
dinámica similar a *make*. Permite, como la opción anterior, crear reglas en un archivo *justfile* que 
luego lanzar desde la línea de comandos.

En cuanto a facilidad de uso, no necesita ser lanzado desde el directorio en el que se encuentra el 
archivo de configuración de reglas, detecta errores del *justfile* antes de ejecutar las reglas...

Como desventaja, su sencillez hace que algunas tareas que son automáticas en otros gestores en esta sean manuales, como las *Incremental Builds*.

## [Task](https://taskfile.dev/)

Esta herramienta permite, mediante la creación de un archivo de configuración, crear tareas automáticas.
La configuración se escribe en formato `yaml`, simplificando la comprensión de las reglas. Además, está 
desarrollada en el lenguaje de programación del proyecto.

Esta herramienta permite modularizar las reglas en varios archivos y muchas otras facilidades que encajan
muy bien un proyecto grande. Sin embargo, la creación de reglas es compleja.

## Conclusión

Debido al tamaño del proyecto, la herramienta más adecuada es **just**, puesto que permite trasladar
rápidamente tareas que se realizan desde la línea de comandos a un archivo de configuración, aportando 
las funcionalidades básicas pero necesarias para la automatización de tareas.