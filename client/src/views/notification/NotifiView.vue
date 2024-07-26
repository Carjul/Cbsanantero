<template>
    <div class="container">
        <hr>
        <div v-if="notificaciones.length === 0">
            <p>No tienes notificaciones</p>
        </div>
        <div v-else>
            <div v-for="notificacion in notificaciones" :key="notificacion.id" @click="cargarNotificacion(notificacion)" data-bs-toggle="modal" data-bs-target="#messageModal" class="shadow-lg p-3 mb-2 bg-body-tertiary rounded message">  
                <div class="message-content">
                    <i :class="['fas', notificacion.revision ? 'message-icon fa-envelope' : 'fa-envelope-open']"> : {{ notificacion.nombre }} </i>
                    <div class="message-text"> <i>............{{ notificacion.description }}</i></div>
                </div>
            </div>
        </div>

        <!-- Modal -->
        <div class="modal fade" id="messageModal" tabindex="-1" aria-labelledby="messageModalLabel" aria-hidden="true">
            <div class="modal-dialog modal-lg">
                <div class="modal-content">
                    <div class="modal-header">
                        <h5 class="modal-title" id="messageModalLabel">Fecha: {{ notificacion.fecha }}</h5>
                        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close">X</button>
                    </div>
                    <div class="modal-body">
                        <div class="modal-body">
                            <table class="table table-bordered table-striped">
                                <tr>
                                    <td class="title">Solicitante:</td>
                                    <td>{{ notificacion.nombre }}</td>
                                </tr>
                                <tr>
                                    <td class="title">Asunto:</td>
                                    <td>{{ notificacion.description }}</td>
                                </tr>
                                <tr>
                                    <td class="title">Personas:</td>
                                    <td>{{ notificacion.person }}</td>
                                </tr>
                                <tr v-if="notificacion.TipoNegocio !== 'Restaurantes'">
                                    <td class="title">Fecha de entrada:</td>
                                    <td>{{ notificacion.entrada }}</td>
                                </tr>
                                <tr v-if="notificacion.TipoNegocio !== 'Restaurantes'">
                                    <td class="title">Fecha salida:</td>
                                    <td>{{ notificacion.salida }}</td>
                                </tr>
                                <tr>
                                    <td class="title">Correo:</td>
                                    <td>{{ notificacion.correo }}</td>
                                </tr>
                                <tr>
                                    <td class="title">Contacto:</td>
                                    <td>{{ notificacion.celular }}</td>
                                </tr>
                            </table>
                        </div>
                    </div>

                    <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cerrar</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      notificacion: {},
      notificaciones: [],
      tipoNegocioUsuario: '',
      adminId: '669eeb77b1259a7ef20afcab', // ID del admin
    };
  },
  methods: {
    getNotificacion() {
      let idc = localStorage.getItem('customerId');
      let url = `${process.env.API}/pedirServicioc/${idc}`;

      axios.get(url)
        .then((response) => {
          this.notificaciones = response.data || [];
        })
        .catch((error) => {
          console.log(error);
        });
    },
    cargarNotificacion(obj) {
      this.notificacion = obj;
      if (obj.revision) {
        axios.put(`${process.env.API}/ServicioStatus/${obj._id}`, { "revision": false })
          .then((response) => {
            console.log(response.data);
            this.getNotificacion();
          })
          .catch((error) => {
            console.log(error);
          });
      }
    },
  },
  mounted() {
    this.getNotificacion();
    this.tipoNegocioUsuario = localStorage.getItem('tipoNegocioUsuario'); 
  },
};
</script>



<style>
 .message {
    position: relative;
    display: flex;
    align-items: center;
    padding: 20px;
    margin-bottom: 5px; /* Reduced margin to avoid separation */
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    background-color: #f8f9fa; /* bg-body-tertiary */
    transition: background-color 0.3s ease;
}

.message:hover {
    background-color: #e2e6ea;
}

.message-content {
    display: flex;
    align-items: center;
}

.message-icon {
    font-size: 24px;
    margin-right: 40px; /* Increased margin for more separation */
    color: #dc3545; /* bg-danger */
}

.message-icon::after {
    content: '';
    position: absolute;
    top: -5px;
    right: -5px;
    width: 10px;
    height: 10px;
    background-color: #28a745; /* Green indicator */
    border-radius: 50%;
    border: 2px solid white; /* Optional: border to make the dot stand out */
}

.message-text {
    font-size: 18px;
}


/* TEXTO DEL MODAL TIPO FACTURA  */

.table {
    width: 100%;
    max-width: 600px; /* Máximo ancho para evitar que la tabla sea demasiado ancha */
    margin: 0 auto; /* Centrar la tabla horizontalmente */
}

.title {
    font-weight: bold;
    background-color: #f2f2f2; /* Color de fondo para los títulos */
}

.table-striped tbody tr:nth-of-type(odd) {
    background-color: #f9f9f9; /* Color de fondo para filas impares */
}

.table-bordered th,
.table-bordered td {
    border: 1px solid #ddd; /* Borde de la tabla */
    padding: 8px; /* Espaciado interno */
}


</style>