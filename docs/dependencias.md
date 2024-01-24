# Elección del gestor de paquetes

Para la gestión de dependencias y paquetes externos, es conveniente utilizar una herramienta que nos ayude a añadir, eliminar y modificarlas de forma automática.

Como opciones para la gestión de dependencias, disponemos de [**Go Modules**](https://go.dev/ref/mod), 
herramienta nativa del lenguaje. Esta es la herramienta que utilizaremos. Al ser nativa, no implica 
ninguna dependencia adicional, además de asegurarnos de que en el futuro próximo no será abandonado su
desarrollo.

## Alternativas

Existen alternativas, pero están desactualizadas, ya que con la integración de **Go Modules** al conjunto de herramientas oficiales, se han dejado de desarrollar alternativas. Como posibles candidatas, de ser actualizadas, podemos encontrar:

 - [**gopm**](https://github.com/gpmgo/gopm): Herramienta sencilla implementada en Go. Al estar implementada en el propio lenguaje de programación del proyecto, puede instalarse de forma sencilla sin necesidad de otros paquetes. Los comandos son muy similares a los de **Go Modules**.

 - [**glide**](https://github.com/Masterminds/glide): Esta herramienta también está implementada en Go. Además, tiene un uso muy similar al gestor de paquetes de Node.js, *npm*, lo cual hace que resulte muy sencilla de aprender.

 - [**gpm**](https://github.com/pote/gpm): Esta herramienta es un script de bash, por lo que no es 
 difícil de instalar y no contiene dependencias extra para la gran mayoría de sistemas Linux. Contiene 
 cuatro comandos, lo cual la hace muy sencilla de utilizar.

Todas estas alternativas a **Go Modules** llevan al menos 5 años sin desarrollarse, debido a que la 
mayoría de usuarios han optado por utilizar la herramienta oficial.