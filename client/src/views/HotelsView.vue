<template>
  <div class="container">
    <hr>
    <div class="row">
      <div
        class="col-md-4"
        v-for="hotel in hoteles"
        :key="hotel.id"
      >
        <div class="card mb-4">
          <img
            @click="enviarIdGaleria(hotel._id)"
            :src="hotel.image"
            class="card-img-top"
            alt="Hotel"
          />
          <div class="card-body">
            <h5 class="card-title">{{ hotel.name }}</h5>
            <p class="card-text">{{ hotel.address }}</p>
            <p class="card-text"><strong>Email:</strong> {{ hotel.email }}</p>
            <p class="card-text"><strong>Teléfono:</strong> {{ hotel.phone }}</p>
            <p class="card-text"><strong>Precio:</strong> {{ hotel.price }}</p>

            <!-- Add the button for requesting service -->
            <button class="btn btn-primary" data-toggle="modal" data-target="#solicitudModal" @click="enviarcid(hotel.customer_id, hotel._id)">Obtener Servicio</button>
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
            <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="closeModal"></button>
          </div>

          <!-- Modal Body -->
          <div class="modal-body">
            <b>
              <p>Creando servicio:</p>
            </b>
          
            <!-- Formulario con campos solicitados -->
            <form>
              <div class="form-group">
                <label for="nombre">Nombres:</label>
                <input type="text" class="form-control" id="nombre" placeholder="Describa nombre y apellido" v-model="formData.nombre">
              </div>

              <div class="form-group">
                <label for="correo">Correo:</label>
                <input type="email" class="form-control" placeholder="Escriba su correo" id="apellidos" v-model="formData.correo">
              </div>

              <div class="form-group">
                <label for="celular">Numero de celular:</label>
                <input type="text" class="form-control" id="celular" placeholder="eje +57 304 000 4445" v-model="formData.celular">
              </div>

              <div class="form-group">
                <label for="description">Descripcion:</label>
                <textarea id="description" placeholder="Describa su solicitud" rows="3" v-model="formData.descripcion"></textarea>
              </div>

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
      hoteles: [],
      formData: {
        _id: null,
        nombre: '',
        celular: '',
        correo: '',
        descripcion: '',
        customerId: null,
      },
    };
  },
  mounted() {
    this.fetchHoteles();
  },
  methods: {
    async fetchHoteles() {
      try {
        const response = await axios.get('http://localhost:3000/Hotel');
        this.hoteles = response.data;
      } catch (error) {
        console.error('Error al obtener los hoteles', error);
      }
    },
    enviarIdGaleria(id) {
      localStorage.setItem('negocioId', id);
      this.$router.push({ path: '/artegaleria' });
    },
    enviarcid(cid, id) {
      this.formData.customerId = cid;
      this.formData._id = id;
      console.log(this.formData);
    },
    closeModal() {
      this.formData = {
        _id: null,
        nombre: '',
        celular: '',
        correo: '',
        descripcion: '',
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
.card {
  transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
}

.card:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}
</style>
