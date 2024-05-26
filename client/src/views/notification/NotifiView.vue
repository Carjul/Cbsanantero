<template>
    <div class="container">
        <br>



  <div v-if="notificaciones.length === 0">
    <p>No tienes notificaciones</p>
  </div>
  <div v-else >
    <div v-for="notificacion in notificaciones" :key="notificacion.id"  @click="cargarNotificacion(notificacion)" data-bs-toggle="modal" data-bs-target="#messageModal" class="shadow-lg p-3 mb-5 bg-body-tertiary rounded message">  
  <div>
    <div class="message-content">
      <i :class="['fas', notificacion.revision ? 'message-icon fa-envelope' : 'fa-envelope-open']"></i>
      <div class="message-content">{{ notificacion.nombre }}</div>
    </div>
  </div>
</div>
</div>


  <!-- Modal -->
  <div class="modal fade" id="messageModal" tabindex="-1" aria-labelledby="messageModalLabel" aria-hidden="true">
    <div class="modal-dialog modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="messageModalLabel">fecha {{ notificacion.hora }}</h5>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close">X</button>
        </div>
        <div class="modal-body">
          <p>{{ notificacion.description }}</p>
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
      notificacion:{
   
  
},
      notificaciones: [],
    };
  },
methods:{
getNotificacion(){
  let idc= localStorage.getItem('customerId');
    axios.get(`${process.env.API}/pedirServicioc/${idc}`)
      .then((response) => {
        this.notificaciones = response.data;
      })
      .catch((error) => {
        console.log(error);
      });
},
cargarNotificacion(obj){
this.notificacion= obj;
if(obj.revision){
axios.put(`${process.env.API}/ServicioStatus/${obj._id}`,{"revision":false})
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
  },
};

</script>

<style>
 .message {
      position: relative;
      display: flex;
      align-items: center;
      padding: 20px;
      margin-bottom: 20px;
      border-radius: 10px;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
      background-color: #f8f9fa; /* bg-body-tertiary */
      transition: background-color 0.3s ease;
    }

    .message:hover {
      background-color: #e2e6ea;
    }

    .message-icon {
      position: relative;
      font-size: 24px;
      margin-right: 15px;
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

    .message-content {
      font-size: 18px;
    }
</style>