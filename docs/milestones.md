# Milestones (hitos)

En este documento aparecen y se irán añadiendo las milestones del proyecto.

#### [M0] Estructura de datos para representar los diferentes elementos de un plan de rodaje

Esta primera milestone es el mínimo necesario para poder modelar un plan de 
rodaje. Será la milestone interna puesto que no genera ningún producto mínimo
y es necesaria para las siguientes. Los elementos típicos suelen ser:

 - **Actores**: Qué actores participan del rodaje, que rol o roles tienen en la
   película, qué disponibilidad tienen...
 - **Localizaciones**: Si son de interior o exterior, qué disponibilidad tienen
   (por ejemplo si es alquilada), ...
 - **Equipo**: Qué tipo de equipo es (cámara, foco, ...), qué disponibilidad
   tienen, ...
 - **Equipo técnico**: Qúe rol tiene, como director, cámara, técnico de sonido,
   ...

#### [M1] Estructura de datos para representar un plan de rodaje

Se trata de crear la estructura necesaria para modelar un plan de rodaje, que
no es más que una lista ordenada de planos. Un plano se suele componer de:

 - **Identificador del plano**: Por ejemplo, plano 2 de la escena 3.
 - **Localización del plano**: Lugar en el que se debe grabar el plano.
 - **Actores que participan**: Actores que participan en el plano.
 - **Equipo técnico**: Miembros del rodaje necesarios para la grabación.
 - **Equipo**: Equipo necesario (cámaras, focos, ...) para la grabación.
 - **Restricciones horarias**: Puede ser que un plano deba ser grabado
   en un cierto tramo horario, como por la noche.

#### [M2] Generar un plan de rodaje a partir de los planos

Esta milestone incorpora el funcionamiento más básico de la lógica de negocio,
que trata de generar un plan de rodaje a partir de los planos del producto
audiovisual con el que se desea construir.