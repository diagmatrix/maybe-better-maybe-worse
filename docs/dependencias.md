# Elección del gestor de paquetes

Para la gestión de dependencias y paquetes externos, es conveniente utilizar una herramienta que nos ayude a añadir, eliminar y modificarlas de forma automática.

Como opciones para la gestión de dependencias, disponemos de algunas herramientas, donde la más destaca es [**Go Modules**](https://go.dev/ref/mod).

## Go Modules

**Go Modules** es la herramienta oficial del lenguaje para la gestión de paquetes. Esta herramienta viene ya instalada con el lenguaje, por lo que implica el uso de una dependencia menos. Todos los paquetes y librerías del lenguaje pueden ser gestionados mediante esta herramienta, que es capaz también de gestionar conflictos entre las mismas. Otra ventaja es la facilidad de aprendizaje, puesto que al formar parte del conjunto oficial de herramientas del lenguaje, el propio tour del lenguaje, así como los tutoriales oficiales del mismo, contienen información de uso de esta herramienta.

## Alternativas

Existen alternativas, pero están desactualizadas, ya que con la integración de **Go Modules** al conjunto de herramientas oficiales, se han dejado de desarrollar alternativas. Como posibles candidatas, de ser actualizadas, podemos encontrar:

 - [**gopm**](https://github.com/gpmgo/gopm): Herramienta sencilla implementada en Go. Al estar implementada en el propio lenguaje de programación del proyecto, puede instalarse de forma sencilla sin necesidad de otros paquetes. Los comandos son muy similares a los de **Go Modules**.
 - [**glide**](https://github.com/Masterminds/glide): Esta herramienta también está implementada en Go. Además, tiene un uso muy similar al gestor de paquetes de Node.js, *npm*, lo cual hace que resulte muy sencilla de aprender.
 - [**gpm**](https://github.com/pote/gpm): Esta herramienta es un script de bash, por lo que no es difícil de instalar. Contiene cuatro comandos, lo cual la hace muy sencilla de utilizar. 

## Conclusión

Todas las alternativas a **Go Modules** llevan al menos 5 años sin desarrollarse, debido a que la mayoría de usuarios han optado por utilizar la herramienta oficial. Aunque todas ellas fuesen mejores que **Go Modules**, no tiene sentido utilizar una herramienta desfasada.