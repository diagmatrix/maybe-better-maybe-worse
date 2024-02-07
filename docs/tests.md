# Elección de herramientas de test

Para elegir las herramientas de test, usaremos los siguientes criterios, en este orden:

 1. Se descartarán automáticamente herramientas que no estén actualmente en desarrollo.
 2. Se preferirán las herramientas que necesiten la instalación del menor número de bibliotecas externas posibles.
 3. Se considerará favorablemente la popularidad de la herramienta, basándonos en el nº de estrellas del repositorio en GitHub.
 4. Tendremos en cuenta la puntuación dada por la herramienta Synk a la herramienta.

## Biblioteca de aserciones

El lenguaje Go no contiene biblioteca de aserciones propia en el lenguaje, por lo que es necesario utilizar una biblioteca externa para ello.
Posibles bilbiotecas de aserciones a tener en cuenta y que están en desarrollo actualmente pueden ser las siguientes:

#### [Testify](https://github.com/stretchr/testify)

Testify es una biblioteca de aserciones que proporciona a su vez la posibilidad de *mockear* objetos. 

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Popularidad en GitHub* | Testify tiene más de 21.000 estrellas en GitHub. |
| *Puntuación en Synk* | Testify tiene una puntuación de 86/100 y no tiene vulnerabilidades conocidas. |

#### [Go Convey](https://github.com/smartystreets/goconvey)

 Se trata de otra biblioteca que ofrece una sintaxis de aserción más expresiva y legible. Permite escribir pruebas de manera más descriptiva y visual, lo que puede facilitar la comprensión de las pruebas. Los tests pueden ser lanzados a través de un navegador web, funcionando así como test runner.

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Popularidad en GitHub* | Go Convey tiene más de 8.000 estrellas en GitHub. |
| *Puntuación en Synk* | [El paquete en sí de aserciones de Go Convey](https://github.com/smarty/assertions) tiene una puntuación de 71/100 y no tiene vulnerabilidades conocidas. |

#### [Go Omega](https://github.com/onsi/gomega)

Gomega es conocida por proporcionar una sintaxis fluida y expresiva que facilita la escritura de pruebas legibles y comprensibles en Go. Esta biblioteca se utiliza comúnmente con el test runner *Ginkgo*, pero también puede ser utilizada de forma independiente.

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Popularidad en GitHub* | Go Omega tiene 2.000 estrellas en GitHub, aunque esto puede ser consecuencia de formar parte del ecosistema *Ginko*, que sí es más popular. |
| *Puntuación en Synk* | Go Omega tiene una puntuación de 88/100 y no tiene vulnerabilidades conocidas. |

### Decisión

Teniendo en cuenta los criterios descritos, vamos a utilizar la herramienta **Testify**, puesto que es la más popular, está en desarrollo y tiene una puntuación de Synk suficientemente alta.

## Test runner / Testing framework

Go tiene su propio test runner que viene por defecto con la instalación del lenguaje: `go test`. Este paquete es el que usaremos como test runner, ya que no implica la instalación
de una biblioteca externa, es el runner más popular y, al formar parte del paquete estandar, posee mucho soporte y actualizaciones.


Como alternativas a este paquete, tenemos **Go Convey**, a su vez biblioteca de aserciones. Cabe destacar también el siguiente test runner, ya que hemos mencionado **Go Omega**.

#### [Ginko](https://github.com/onsi/ginkgo)

Ginkgo es un testing framework para Go que se utiliza comúnmente para realizar pruebas BDD (*Behavior-Driven Development*). Permite, de forma sencilla, organizar los tests de forma
que se ejecuten bloques de código antes y después de cada prueba.

| Criterio              | Ajuste en herramienta | 
|------------------------|-----------------------|
| *Popularidad en GitHub* | Ginko tiene casi 8.000 estrellas en GitHub. |
| *Puntuación en Synk* | Ginko tiene una puntuación de 91/100 y no tiene vulnerabilidades conocidas. |
