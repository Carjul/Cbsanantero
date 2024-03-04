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
          <img
            :src="artesania.image"
            class="card-img-top"
            alt="Artesanía"
          />
          <div class="card-body">
            <h5 class="card-title">{{ artesania.name }}</h5>
            <p class="card-text">{{ artesania.description }}</p>
            <p class="card-text"><strong>Dirección:</strong> {{ artesania.address }}</p>
            <p class="card-text"><strong>Teléfono:</strong> {{ artesania.phone }}</p>
            
            <!-- Botón "Obtener Servicio" -->
            <button class="btn btn-primary" data-toggle="modal" data-target="#solicitudModal">Obtener Servicio</button>
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
            <!-- Formulario con campos solicitados -->
            <form>
              <div class="form-group">
                <label for="nombre">Nombre:</label>
                <input type="text" class="form-control" id="nombre" v-model="formData.nombre">
              </div>

              <div class="form-group">
                <label for="apellidos">Apellidos:</label>
                <input type="text" class="form-control" id="apellidos" v-model="formData.apellidos">
              </div>

              <div class="form-group">
                <label for="descripcion">Descripción:</label>
                <textarea class="form-control" id="descripcion" v-model="formData.descripcion"></textarea>
              </div>

              <div class="form-group">
                <label for="celular">Celular:</label>
                <input type="tel" class="form-control" id="celular" v-model="formData.celular">
              </div>

              <div class="form-group">
                <label for="correo">Correo Electrónico:</label>
                <input type="email" class="form-control" id="correo" v-model="formData.correo">
              </div>

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
      formData: {
        nombre: '',
        apellidos: '',
        descripcion: '',
        celular: '',
        correo: '',
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
      } catch (error) {
        console.error('Error al obtener las artesanías', error);
      }
    },
    closeModal() {
      this.formData = {
        nombre: '',
        apellidos: '',
        descripcion: '',
        celular: '',
        correo: '',
      };
    },
    submitForm() {
      console.log('Formulario enviado con datos:', this.formData);
      this.closeModal();
      // Puedes agregar aquí la lógica para enviar el formulario al servidor si es necesario.
    },
  },
};
</script>
