# Imagen base del contenedor

Utilizaremos una herramienta de compilación y administración de imágenes de contenedores para crear una imagen que facilite la ejecución de pruebas unitarias en la aplicación en desarrollo. En este proceso, es crucial seleccionar la imagen base como punto de partida.

## Criterios de Selección

Al optar por la imagen base para el proyecto, se evaluarán las opciones según los siguientes criterios:

 - **Tamaño**: Buscaremos minimizar el tamaño de la imagen para optimizar el espacio.
 - **Popularidad**: Buscaremos usar las imágenes más populares posibles, basándonos en los pulls de Docker Hub.
 - **Mantenimiento**: Se descartarán las imágenes obsoletas, eligiendo aquellas que reciban actualizaciones frecuentes.

Se descartarán las imágenes que no cuenten con una instalación de Go por defecto, ya que implican un trabajo extra que no considero necesario teniendo en cuenta la gran
cantidad de opciones.

## [Bitnami](https://hub.docker.com/r/bitnami/golang)

Esta imagen, de la organización *Bitnami*, promete ser segura y pequeña, ya que se basa en una distribución llamada Minideb. Veamos como se comporta en relación a nuestros criterios:

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Tamaño* | La última versión de la imagen pesa 210.28 MB comprimida. |
| *Popularidad* | La imagen tiene más de medio millón de pulls en Docker Hub. |
| *Mantenimiento* | La última imagen es de hace 15h (visitado el 7 de febrero de 2024). |

## [Golang Oficial](https://hub.docker.com/_/golang)

El lenguaje de programación tiene varias imágenes oficiales, de las cuales discutiremos la basada en Debian-bullseye y Alpine.

En general, podemos decir que estas opciones son muy populares, con más de mil millos de pulls en Docker Hub.

#### bullseye

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Tamaño* | La última versión de la imagen pesa 270.01 MB comprimida. |
| *Mantenimiento* | La última imagen es de hace 5h (visitado el 7 de febrero de 2024). |

Es importante notar que esta imagen tiene 87 vulnerabilidades leves y 2 sin especificar conocidas.

#### alpine3.18

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Tamaño* | La última versión de la imagen pesa 67.59 MB comprimida. |
| *Mantenimiento* | La última imagen es de hace 5h (visitado el 7 de febrero de 2024). |

## Elección

Por los criterios descritos, vamos a utilizar la imagen oficial de alpine, puesto que no tiene vulnerabilidades conocidas, es la más pequeña y es muy popular.