# Imagen base del contenedor

Utilizaremos una herramienta de compilación y administración de imágenes de contenedores para crear una imagen que facilite la ejecución de pruebas unitarias en la aplicación en desarrollo. En este proceso, es crucial seleccionar la imagen base como punto de partida.

## Criterios de Selección

Al optar por la imagen base para el proyecto, se evaluarán las opciones según los siguientes criterios:

 - **Tamaño**: Buscaremos minimizar el tamaño de la imagen para optimizar el espacio.
 - **Popularidad**: Buscaremos usar las imágenes más populares posibles, basándonos en los pulls de Docker Hub.
 - **Mantenimiento**: Se descartarán las imágenes obsoletas, eligiendo aquellas que reciban actualizaciones frecuentes.
 - **Seguridad**: Se procurará elegir una imagen con la menor cantidad de vulnerabilidades posibles.

El tamaño vamos a discutirlo en su propia sección, puesto que tiene más sentido hacer una comparativa total.

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

## [Archlinux](https://hub.docker.com/_/archlinux)

Archlinux es una distribución de linux famosa por ser *ligthweight* y tener una gran comunidad.

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Popularidad* | La imagen tiene alrededor de 10 millones de pulls en Docker Hub. |
| *Mantenimiento* | La última imagen es de hace un mes (visitado el 12 de febrero de 2024). |
| *Seguridad* | Esta imagen contiene dos vulnerabilidades, siendo una considerada severa (visitado el 12 de febrero de 2024). |

En cuanto a la seguridad, la imagen recomienda actualizar los paquetes para evitar el mayor número de vulnerabilidades, utilizando el hecho de que Archlinux es una *rolling release*, es decir, va actualizando sus paquetes de manera regular, sin esperar a versiones.

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
```

Podemos comprobar que la más pequeña es la de alpine, seguida por la de archlinux.

## Elección

Por los criterios descritos, parece que tanto **alpine** como **archlinux** parecen ser correctas. Para ello, vamos a comprobar qué imagen ejecuta más rápido los tests, y comprobaremos el tamaño de la imagen final también. El Dockerfile de la imagen de archlinux se encuentra [aquí](docs/archlinux_dockerfile.md)

Durante la creación de la imagen de **archlinux**, aparecen algunos errores en la actualización de los paquetes, posiblemente relacionados con el hecho de que en la imagen se eliminan por seguridad la clave
`lsign` de `pacman` (el gestor de paquetes de la distribución).

Obtenemos como tamaño final de las imágenes (habiendo utilizado previamente `docker tag` para diferenciarlas):

```Bash
diagmatrix@gibbon ~ $ docker images
REPOSITORY                            TAG         IMAGE ID       CREATED              SIZE
diagmatrix/maybe-better-maybe-worse   archlinux   64b91dbeea80   About a minute ago   1.14GB
diagmatrix/maybe-better-maybe-worse   alpine      1e45f71c51c0   18 hours ago         236MB
```

Podemos ver que, al tener que instalar el lenguaje de Go con todas sus dependencias, el tamaño de la imagen de **archlinux** se duplica, mientras que el de **alpine** solo aumenta un poco debido a que solo
ha sido necesario instalar *just*.

Veamos ahora la velocidad a la que se ejecutan los tests. Para ello, utilizamos el comando `time`:

Para **archlinux**:

```Bash
diagmatrix@gibbon ~/git/maybe-better-maybe-worse $ time docker run -u 1001 -t -v `pwd`:/app/test diagmatrix/maybe-better-maybe-worse:archlinux
Ejecutando los tests...
go test -v ./tests/
=== RUN   TestMemberTypeFromString
--- PASS: TestMemberTypeFromString (0.00s)
=== RUN   TestMemberTypeFromStringInvalid
--- PASS: TestMemberTypeFromStringInvalid (0.00s)
=== RUN   TestParseTimeTable
--- PASS: TestParseTimeTable (0.00s)
=== RUN   TestParseTimeTableInvalidFile
--- PASS: TestParseTimeTableInvalidFile (0.00s)
=== RUN   TestParseTimeTableInvalidHeader
--- PASS: TestParseTimeTableInvalidHeader (0.00s)
=== RUN   TestParseTimeTableInvalidAvailability
--- PASS: TestParseTimeTableInvalidAvailability (0.00s)
=== RUN   TestParseTechnicalGuide
--- PASS: TestParseTechnicalGuide (0.00s)
=== RUN   TestParseTechnicalGuideInvalidFile
--- PASS: TestParseTechnicalGuideInvalidFile (0.00s)
=== RUN   TestParseTechnicalGuideInvalidHeader
--- PASS: TestParseTechnicalGuideInvalidHeader (0.00s)
=== RUN   TestParseTechnicalGuideInvalidCastCrew
--- PASS: TestParseTechnicalGuideInvalidCastCrew (0.00s)
PASS
ok  	github.com/diagmatrix/maybe-better-maybe-worse/tests	0.003s

real	0m26,353s
user	0m0,011s
sys	0m0,011s
```

Para **alpine**:

```Bash
diagmatrix@gibbon ~/git/maybe-better-maybe-worse $ time docker run -u 1001 -t -v `pwd`:/app/test diagmatrix/maybe-better-maybe-worse:alpine
Ejecutando los tests...
go test -v ./tests/
=== RUN   TestMemberTypeFromString
--- PASS: TestMemberTypeFromString (0.00s)
=== RUN   TestMemberTypeFromStringInvalid
--- PASS: TestMemberTypeFromStringInvalid (0.00s)
=== RUN   TestParseTimeTable
--- PASS: TestParseTimeTable (0.00s)
=== RUN   TestParseTimeTableInvalidFile
--- PASS: TestParseTimeTableInvalidFile (0.00s)
=== RUN   TestParseTimeTableInvalidHeader
--- PASS: TestParseTimeTableInvalidHeader (0.00s)
=== RUN   TestParseTimeTableInvalidAvailability
--- PASS: TestParseTimeTableInvalidAvailability (0.00s)
=== RUN   TestParseTechnicalGuide
--- PASS: TestParseTechnicalGuide (0.00s)
=== RUN   TestParseTechnicalGuideInvalidFile
--- PASS: TestParseTechnicalGuideInvalidFile (0.00s)
=== RUN   TestParseTechnicalGuideInvalidHeader
--- PASS: TestParseTechnicalGuideInvalidHeader (0.00s)
=== RUN   TestParseTechnicalGuideInvalidCastCrew
--- PASS: TestParseTechnicalGuideInvalidCastCrew (0.00s)
PASS
ok  	github.com/diagmatrix/maybe-better-maybe-worse/tests	0.003s

real	0m24,478s
user	0m0,009s
sys	0m0,010s
```

Podemos comprobar que en tiempo de ejecución casi idénticos.

Por los criterios descritos, y teniendo en cuenta que ambas candidatas ejecutan los tests a la misma velocidad, vamos a utilizar la imagen oficial de alpine, puesto que no tiene vulnerabilidades conocidas, es más pequeña y es muy popular.