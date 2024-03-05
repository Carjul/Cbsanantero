<template>
  <div class="container">
    <hr>
    <div class="row">
      
      <div
      
        class="col-md-4"
        v-for="artesania in artesanias"
        :key="artesania.id"
      >
        <div class="card mb-4">
          <router-link  to="/artegalerri">
          <img
          
            :src="artesania.image"
            class="card-img-top"
            alt="Artesanía"
          />
        </router-link>

          <div class="card-body">
            <h5 class="card-title">{{ artesania.name }}</h5>
            <p class="card-text">{{ artesania.description }}</p>
            <p class="card-text"><strong>Dirección:</strong> {{ artesania.address }}</p>
            <p class="card-text"><strong>Teléfono:</strong> {{ artesania.phone }}</p>
           

            
            <!-- Botón "Obtener Servicio" -->
            <button v-if="true" class="btn btn-primary" data-toggle="modal" data-target="#solicitudModal" @click="enviarcid(artesania.customer_id,artesania._id)">Obtener Servicio</button>


          </div>
        </div>
      </div>
   
    </div>
 <!-- Modal -->
 <div class="modal" id="solicitudModal" tabindex="-1" role="dialog" aria-labelledby="solicitudModalLabel" aria-hidden="true">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <!-- Modal Header -->
          <div class="modal-header">
            <h4 class="modal-title" id="solicitudModalLabel">Solicitar Servicio</h4>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="closeModal">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>

          <!-- Modal Body -->
          <div class="modal-body">
            <p>CustomerID Seleccionado: {{ selectedCustomerId }}</p>
            <!-- Formulario con campos solicitados -->
            <form>
              <div class="form-group">
                <label for="nombre">Nombre:</label>
                <input type="text" class="form-control" id="nombre" v-model="formData.nombre">
              </div>

              <!-- ... Otros campos del formulario ... -->

              <!-- Botón de solicitar -->
              <button type="button" class="btn btn-success" @click="submitForm">Solicitar</button>
            </form>
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
      artesanias: [],
      userRole: 'usuario', // Asume que userRole es 'usuario' por defecto. Adáptalo según la información real en tu aplicación.
      formData: {
        _id:null,
        nombre: '',
        apellidos: '',
        descripcion: '',
        celular: '',
        correo: '',
        customerId: null,
      },
    };
  },
  mounted() {
    this.fetchArtesanias();
  },
  methods: {
    async fetchArtesanias() {
      try {
        const response = await axios.get('http://localhost:3000/Artesania');
        this.artesanias = response.data;
        console.log(this.artesanias)
      } catch (error) {
        console.error('Error al obtener las artesanías', error);
      }
    },
    enviarcid(cid,id){
      this.formData.customerId=cid;
      this.formData._id=id;
      console.log(this.formData)
    },
    openModal(customerId) {
      this.selectedCustomerId = customerId;
      this.formData.customerId = customerId;
      this.modalOpen = true;
    },
    closeModal() {
      this.selectedCustomerId = null;
      this.formData = {
        _id:null,
        nombre: '',
        apellidos: '',
        descripcion: '',
        celular: '',
        correo: '',
        customerId: null,
      };
    },
    submitForm() {
      console.log('CustomerID seleccionado:', this.selectedCustomerId);
      console.log('Formulario enviado con datos:', this.formData);
      this.closeModal();
      // Puedes agregar aquí la lógica para enviar el formulario al servidor si es necesario.
    },
  },
};
</script>



<style scoped>
  /* Ajustes generales para las tarjetas */
  .card {
    height: 100%;
  }

  /* Ajustes específicos para las imágenes dentro de las tarjetas */
  .card-img-top {
    height: 200px; /* Ajusta la altura según tus necesidades */
    object-fit: cover; /* Ajusta la forma en que la imagen se ajusta al contenedor */
  }

  /* Ajustes específicos para el cuerpo de la tarjeta */
  .card-body {
    min-height: 150px; /* Ajusta la altura mínima según tus necesidades */
  }

  /* Otras personalizaciones según tus necesidades */
</style>
