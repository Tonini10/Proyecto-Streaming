# Proyecto-Streaming ToniFlix 
Nombre: Esteban Tonini
Fecha: 12/13/2024

¨El sistema esta desarrollado en Go y utiliza Gin-Gonic, un framework ligero y rápido para la creación de APIs y aplicaciones web en Go. Ademas el sistema se conecta a una base de datos MySQL para almacenar y gestionar información sobre usuarios, películas y transacciones. Utiliza HTML y CSS para la construcción de vistas, Exposición de servicios mediante APIs RESTful y se implementaron middlewares y configuraciones como el manejo de CORS para seguridad y control de acceso¨.

Funcionalidades:

1. Gestión de Películas:
   
Listar películas disponibles:  Devuelve un listado completo de las películas almacenadas, incluyendo información como título, género y descripciones.
Endpoint: GET /api/peliculas

Consultar detalles de una película: Proporciona detalles específicos de una película, incluyendo enlaces a trailers o contenido relacionado.
Endpoint: GET /api/pelicula/:id

Filtrar películas por usuario: Lista todas las películas asociadas a un usuario registrado, basándose en su correo electrónico.
Endpoint: GET /api/peliculasemail/:email


2. Gestión de Usuarios

Crear usuarios:Permite registrar nuevos usuarios en el sistema, almacenando información básica como nombre y correo electrónico.
Endpoint: POST /crear-usuario

Consultar datos de un usuario por correo electrónico:Busca y devuelve la información de un usuario específico en formato JSON.  
Endpoint: GET /api/usuarioemail/:email

Validaciones de usuarios:
Se implementan controles en el backend para garantizar que los datos sean consistentes y seguros.

3. Sistema de Compras
   
Registrar compras de contenido: Permite que los usuarios adquieran contenido multimedia. Se almacenan los detalles de la transacción como usuario, contenido adquirido y fecha.
Endpoint: POST /api/compra

Ver historial de compras: Lista todas las compras realizadas por un usuario, ofreciendo un historial completo de transacciones.
Endpoint: GET /ver-compras

4. APIs RESTful
   
Endpoints destacados: Exponen datos y funcionalidades del sistema para integraciones con otras aplicaciones o servicios.
GET /api/peliculas: Lista todas las películas.
GET /api/pelicula/:id: Proporciona detalles de una película específica.
GET /api/peliculasemailtotal/:email: Devuelve el número total de películas asociadas a un usuario.
POST /api/compra: Registra una compra.

5. Interfaz de Usuario con Plantillas HTML
Página de inicio: Muestra un catálogo de películas disponible para los usuarios.
Página de detalles de una película: Presenta información detallada de una película seleccionada.
Formulario de registro: Interfaz para que los usuarios creen su cuenta en la plataforma.
Historial de compras: Página que muestra las compras realizadas por un usuario.

6. Sistema de Seguridad

Validaciones de datos:

CORS (Cross-Origin Resource Sharing): Se implementan validaciones en los endpoints para asegurar que las solicitudes sean correctas y seguras.
Configurado para controlar el acceso a las APIs desde diferentes orígenes.
