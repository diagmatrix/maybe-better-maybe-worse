# Imagen base del contenedor

Utilizaremos una herramienta de compilación y administración de imágenes de contenedores para crear una imagen que facilite la ejecución de pruebas unitarias en la aplicación en desarrollo. En este proceso, es crucial seleccionar la imagen base como punto de partida.

## Criterios de Selección

Al optar por la imagen base para el proyecto, se evaluarán las opciones según los siguientes criterios:

 - **Tamaño**: Buscaremos minimizar el tamaño de la imagen para optimizar el espacio.
 - **Popularidad**: Buscaremos usar las imágenes más populares posibles, basándonos en los pulls de Docker Hub.
 - **Mantenimiento**: Se descartarán las imágenes obsoletas, eligiendo aquellas que reciban actualizaciones frecuentes.
 - **Seguridad**: Se procurará elegir una imagen con la menor cantidad de vulnerabilidades posibles.

Se descartarán las imágenes que no cuenten con una instalación de Go por defecto, ya que implican un trabajo extra que no considero necesario teniendo en cuenta la gran
cantidad de opciones.

## [Bitnami](https://hub.docker.com/r/bitnami/golang)

Esta imagen, de la organización *Bitnami*, promete ser segura y pequeña, ya que se basa en una distribución llamada Minideb. Veamos como se comporta en relación a nuestros criterios:

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Popularidad* | La imagen tiene más de medio millón de pulls en Docker Hub. |
| *Mantenimiento* | La última imagen es de hace 15h (visitado el 7 de febrero de 2024). |
| *Seguridad* | En DockerHub no aparecen publicadas las vulnerabilidades de la imagen, pero en el [README](https://github.com/bitnami/containers?tab=readme-ov-file#vulnerability-scan-in-bitnami-container-images) del repositorio de sus contenedores describen que utilizan dos herramientas para escanear sus imágenes antes de publicarlas. |

## [Golang Oficial](https://hub.docker.com/_/golang)

El lenguaje de programación tiene varias imágenes oficiales, de las cuales discutiremos la basada en Debian-bullseye y Alpine.

En general, podemos decir que estas opciones son muy populares, con más de mil millos de pulls en Docker Hub.

#### bullseye

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Mantenimiento* | La última imagen es de hace 5h (visitado el 7 de febrero de 2024). |
| *Seguridad* | La última versión de esta imagen contiene más de 80 vulnerabilidades, siendo 2 de ellas severas (visitado el 11 de febrero de 2024). |

#### alpine

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Mantenimiento* | La última imagen es de hace 5h (visitado el 7 de febrero de 2024). |
| *Seguridad* | La última versión de esta imagen no tiene vulnerabilidades conocidas en DockerHub. |

## [Jumpserver](https://hub.docker.com/r/jumpserver/golang)

Esta imagen, del grupo *Jumpserver*, es una imagen mantenida por la comunidad, construida sobre Debian.
Está patrocinada por Docker a través de su programa de patrocinio a proyectos de Software Libre.

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Popularidad* | La imagen tiene alrededor de 1600 pulls en Docker Hub. |
| *Mantenimiento* | La última imagen es de hace 3 meses (visitado el 11 de febrero de 2024). |
| *Seguridad* | No hay ninguna información acerca de la seguridad de esta imagen. |

## Elección

Por los criterios descritos, vamos a utilizar la imagen oficial de alpine, puesto que no tiene vulnerabilidades conocidas, es la más pequeña y es muy popular.