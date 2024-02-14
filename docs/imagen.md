# Imagen base del contenedor

Utilizaremos una herramienta de compilación y administración de imágenes de contenedores para crear una imagen que facilite la ejecución de pruebas unitarias en la aplicación en desarrollo. En este proceso, es crucial seleccionar la imagen base como punto de partida.

Tenemos en cuenta que el lenguaje utiliza la librería de *enlazado* de C que tenga el SO, por lo que
no nos fijaremos en la librería de C de las imágenes base.

## Criterios de Selección

Al optar por la imagen base para el proyecto, se evaluarán las opciones según los siguientes criterios:

 - **Popularidad**: Buscaremos usar las imágenes más populares posibles, basándonos en los pulls de Docker Hub.
 - **Mantenimiento**: Se descartarán las imágenes obsoletas, eligiendo aquellas que reciban actualizaciones frecuentes.
 - **Seguridad**: Se procurará elegir una imagen con la menor cantidad de vulnerabilidades posibles.
 - **Tamaño**: Buscaremos minimizar el tamaño de la imagen para optimizar el espacio. 
 - **Paquetes innecesarios**: Se intentarán elegir distribuciones que no incluyan demasiados paquetes ajenos a nuestras necesidades.

El tamaño vamos a discutirlo en su propia sección, puesto que tiene más sentido hacer una comparativa total.

## [Bitnami](https://hub.docker.com/r/bitnami/golang)

Esta imagen, de la organización *Bitnami*, promete ser segura y pequeña, ya que se basa en una distribución llamada Minideb. Veamos como se comporta en relación a nuestros criterios:

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Popularidad* | La imagen tiene más de medio millón de pulls en Docker Hub. |
| *Mantenimiento* | La última imagen es de hace 15h (visitado el 7 de febrero de 2024). |
| *Seguridad* | En DockerHub no aparecen publicadas las vulnerabilidades de la imagen, pero en el [README](https://github.com/bitnami/containers?tab=readme-ov-file#vulnerability-scan-in-bitnami-container-images) del repositorio de sus contenedores describen que utilizan dos herramientas para escanear sus imágenes antes de publicarlas. |
| *Paquetes innecesarios* | Si vemos la construcción de sus [imágenes](https://github.com/bitnami/containers/blob/main/bitnami/golang/1.22/debian-11/Dockerfile), podemos ver que añaden paquetes que no vamos a usar, como `git` o `unzip`, además de crear una estructura de ficheros innecesaria en nuestro caso. |

## [Golang Oficial](https://hub.docker.com/_/golang)

El lenguaje de programación tiene varias imágenes oficiales, de las cuales discutiremos la basada en Debian-bullseye y Alpine.

En general, podemos decir que estas opciones son muy populares, con más de mil millones de pulls en Docker Hub. Ambas imágenes instalan por defecto los paquetes necesarios para utilizar [cgo](https://pkg.go.dev/cmd/cgo), que no se utiliza en el proyecto, entre otros.

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
| *Paquetes innecesarios* | Por lo que podemos ver [aquí](https://github.com/wojiushixiaobai/docker-library-loong64/blob/master/golang/1.21/buster/Dockerfile), esta imagen instala paquetes innecesarios para nuestro proyecto (los mismos que las oficiales). |

## [Archlinux](https://hub.docker.com/_/archlinux)

Archlinux es una distribución de linux famosa por ser *ligthweight* y tener una gran comunidad.

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Popularidad* | La imagen tiene alrededor de 10 millones de pulls en Docker Hub. |
| *Mantenimiento* | La última imagen es de hace un mes (visitado el 12 de febrero de 2024). |
| *Seguridad* | Esta imagen contiene dos vulnerabilidades, siendo una considerada severa (visitado el 12 de febrero de 2024). |
| *Paquetes innecesarios* | En el [README](https://gitlab.archlinux.org/archlinux/archlinux-docker/-/blob/master/README.md?ref_type=heads#dependencies) del repositorio en el cual se crean las imágenes nos explica qué paquetes se descargan, y de ellos sólamente `devtools` es necesario. | 

En cuanto a la seguridad, la imagen recomienda actualizar los paquetes para evitar el mayor número de vulnerabilidades, utilizando el hecho de que Archlinux es una *rolling release*, es decir, va actualizando sus paquetes de manera regular, sin esperar a versiones.

## [Alpine](https://hub.docker.com/_/archlinux)

Alpine es una distribución de Linux pequeña, ideal para generar entornos seguros, ya que la seguridad es
uno de sus pilares base.

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Popularidad* | La imagen tiene alrededor de 1000 millones de pulls en Docker Hub. |
| *Mantenimiento* | La última imagen es de hace 18 días (visitado el 13 de febrero de 2024). |
| *Seguridad* | Esta imagen no contiene vulnerabilidades. (visitado el 13 de febrero de 2024). De hecho, parece que la imagen se actualiza cuando se descubren y arreglan vulnerabilidades (Como [aquí](https://github.com/docker-library/official-images/pull/16144)). |
| *Paquetes innecesarios* | Es complicado saber qué paquetes instala la imagen, puesto que el Dockerfile contiene 3 líneas (añade un `.tar`). Si inspeccionamos ese archivo comprimido, parece instalar únicamente paquetes necesarios para el SO (`apk`, `sysctl`, ...) | 

## Tamaño

Para comprobar el tamaño de las imágenes, se ha ejecutado `docker pull <nombre_imagen>` sobre cada
una de las discutidas, y se ha ejecutado el comando `docker images` para visualizarlas, obteniendo como resultado:

```Bash
diagmatrix@gibbon ~ $ docker images
REPOSITORY                            TAG        IMAGE ID       CREATED        SIZE
bitnami/golang                        latest     a53c682817d0   26 hours ago   611MB
golang                                alpine     a2742f74d90f   5 days ago     230MB
golang                                bullseye   3af37afcedb7   5 days ago     755MB
archlinux                             latest     69f38d8f6347   5 weeks ago    436MB
jumpserver/golang                     latest     d4c090e16c3f   3 months ago   688MB
alpine                                latest     05455a08881e   2 weeks ago    7.37MB
```

Podemos comprobar que la más pequeña es la de **alpine**, seguida por la oficial de **go/alpine**.

## Elección

Por los criterios descritos, parece que tanto **alpine** como **archlinux** y **go/alpine** parecen ser las más correctas. Para ello, probamos las imágenes. El Dockerfile de la imagen de archlinux se encuentra [aquí](docs/archlinux_dockerfile.md), y la de go/alpine [aquí](docs/go_alpine_dockerfile.md).
Estas imágenes no están ''pulidas'' como la final del [Dockerfile](Dockerfile), pero nos sirven para probar.

Durante la creación de la imagen de **archlinux**, aparecen algunos errores en la actualización de los paquetes, posiblemente relacionados con el hecho de que en la imagen se eliminan por seguridad la clave
`lsign` de `pacman` (el gestor de paquetes de la distribución).

Obtenemos como tamaño final de las imágenes (habiendo utilizado previamente `docker tag` para diferenciarlas):

```Bash
diagmatrix@gibbon ~ $ docker images
REPOSITORY                            TAG         IMAGE ID       CREATED              SIZE
diagmatrix/maybe-better-maybe-worse   archlinux   64b91dbeea80   About a minute ago   1.14GB
diagmatrix/maybe-better-maybe-worse   go-alp      1e45f71c51c0   18 hours ago         236MB
diagmatrix/maybe-better-maybe-worse   alpine      326111d58e24   2 minutes ago        337MB
```

Podemos ver que, al tener que instalar el lenguaje de Go con todas sus dependencias, el tamaño de la imagen de **archlinux** se duplica, mientras que el de **go/alpine** solo aumenta un poco debido a que solo ha sido necesario instalar *just*. En cuanto a la de **alpine**, es 1MB más grande que la oficial, igual producto de errores de redondeo. Sin embargo, son bastante similares en tamaño. En cuanto a velocidad de ejecución, las tres imágenes tardan lo mismo, por lo que no nos sirve para compararlas.

Por los criterios descritos, y teniendo en cuenta que ambas candidatas ejecutan los tests a la misma velocidad, vamos a utilizar la imagen de **alpine**, puesto que no tiene vulnerabilidades conocidas, es igual de pequeña que la oficial y es muy popular. Además, nos hemos asegurado de no tener paquetes.

# Justificación de las órdenes en el Dockerfile

Para la creación del contenedor de pruebas del proyecto, se han tomado las siguientes decisiones en el
Dockerfile:

1. *Creación de un usuario para test.* Mediante la sentencia `adduser -D ppf-tests`, creamos un nuevo usuario para las pruebas que no sea el usuario `root` del contenedor. 

2. *Instalación de dependencias.* Instalamos las dependencias del proyecto con el usuario creado, pero 
las dependencias ajenas (*just* en nuestro caso) las instalamos desde con privilegios de *root*, puesto
que no es posible de otra manera.

3. *Modificación de las variables de entorno.* Puesto que `go test` necesita escribir en archivos (tiene que "compilar" temporalmente los archivos), y `app/test` está montada sin esos permisos, es necesario cambiar las variables de entorno de Go a directorios a los que tenga acceso el usuario creado, además de dar permisos de escritura sobre el directorio de caché.
