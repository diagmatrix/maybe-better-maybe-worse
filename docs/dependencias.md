# Elección del gestor de paquetes

Para la gestión de dependencias y paquetes externos, es conveniente utilizar una herramienta que nos ayude a añadir, eliminar y modificarlas de forma automática.

Como opciones para la gestión de dependencias, disponemos de [**Go Modules**](https://go.dev/ref/mod), 
herramienta nativa del lenguaje. Esta es la herramienta que utilizaremos. Al ser nativa, no implica 
ninguna dependencia adicional, además de asegurarnos de que en el futuro próximo no será abandonado su
desarrollo. Existen alternativas, pero están desactualizadas, ya que con la integración de **Go Modules**
al conjunto de herramientas oficiales, se han dejado de desarrollar alternativas.