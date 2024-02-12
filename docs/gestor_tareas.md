# Elección del gestor de tareas

El gestor de tareas es una herramienta que se encarga de automatizar procesos repetitivos como la compilación del código, optimización de los recursos computacionales y ejecución de tests.

## [Make](https://www.gnu.org/software/make/)

Esta herramienta permite, mediante la configuración de un archivo *Makefile*, crear reglas que
automaticen tareas como la compilación y ejecución de tests. Se trata de una herramienta ampliamente
documentada y conocida.

 | Criterio              | Ajuste en herramienta |     |
|------------------------|-----------------------| --- |
| ***Soporte en Go***    | Make permite crear reglas para cualquier lenguaje de programación. | ✓ |
| ***Dependencias***    | Make es una herramienta que únicamente tiene como dependencias librerías de C necesarias para un núcleo Linux. Sin embargo, es un programa externo, añadiéndose como dependencia. | - |
| ***Popularidad***      | Make es una herramienta muy conocida y utilizada. | ✓ |
| ***Desarrollo***       | Make está en desarrollo y tiene soporte. La última versión fue lanzada en febrero de 2023 | ✓ |

Al no ser una herramienta nativa de Go, no he encontrado ninguna puntuación de la misma en Synk.

## [Just](https://github.com/casey/just)

Esta herramienta, autodescrita como ''Simplemente un ejecutor de comandos'', es una herramienta de 
dinámica similar a *make*. Permite, como la opción anterior, crear reglas en un archivo *justfile* que 
luego lanzar desde la línea de comandos.

 | Criterio              | Ajuste en herramienta |     |
|------------------------|-----------------------| --- |
| ***Soporte en Go***    | Just permite crear reglas para cualquier lenguaje de programación. | ✓ |
| ***Dependencias***    | Just está escrito en Rust, un lenguaje de programación que está en desarrollo y con una popularidad creciente pero que no suele estar instalado por defecto. Sin embargo, tiene binarios para la gran mayoría de distribuciones utilizadas en las imágenes de los contenedores. | ✓ |
| ***Popularidad***      | Just es conocida y bastante utilizada, sobretodo en proyectos de tamaño pequeño o medio, como es el caso. El repositorio tiene más de 15000 estrellas. | ✓ |
| ***Desarrollo***       | Just está en desarrollo y tiene soporte. La última versión fue lanzada en enero de 2024 | ✓ |

Al no ser una herramienta nativa de Go, no he encontrado ninguna puntuación de la misma en Synk.

## [Task](https://taskfile.dev/)

Esta herramienta permite, mediante la creación de un archivo de configuración, crear tareas automáticas.
La configuración se escribe en formato `yaml`, simplificando la comprensión de las reglas. Se trata de una
visión 'moderna' de *make*.

 | Criterio              | Ajuste en herramienta |     |
|------------------------|-----------------------| --- |
| ***Soporte en Go***    | Task permite crear reglas para cualquier lenguaje de programación. | ✓ |
| ***Dependencias***    | Task está escrito en Go y puede ser instalado desde el propio gestor de paquetes del lenguaje. Por tanto, no supone ninguna dependencia adicional. | ✓ |
| ***Popularidad***      | Task es conocida, aunque no tanto como las demás. Su repositorio tiene más de 9000 estrellas. | ✓ |
| ***Desarrollo***       | Task está en desarrollo y tiene soporte. La última versión fue lanzada en diciembre de 2023 | ✓ |

La herramienta Synk la puntúa con 63/100 puntos, alegando que se encuentra inactiva. Sin embargo, podemos
comprobar en el repositorio de la misma que no es así.

## Conclusión

Teniendo en cuenta los criterios, la herramienta más adecuada es **just**, puesto que no añade dependencias externas, es suficientemente popular y está en desarrollo acutalmente. Además, tiene la ventaja de ser sencilla.
