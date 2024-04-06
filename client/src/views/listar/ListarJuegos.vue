<template>
    <div id="app">
      <div class="container">
        <br>
        <div class="card">
          <h5 class="card-header">Panel de Recreación</h5>
          <div class="card-body">
            <h5 class="card-title">Creación de Recreación</h5>
            <p class="card-text">En este espacio creamos nuevas recreaciones para COSTA BRISA</p>
            <button class="btn btn-success mb-4" @click="abrirModalCrear">Crear</button>
          </div>
        </div>
  
        <hr>
        <strong>Todas las Recreaciones</strong>
        <hr>
        <div class="row">
          <div v-for="recreacion in recreaciones" :key="recreacion._id" class="col-md-4 mb-4">
            <div class="card">
              <img :src="recreacion.image" class="card-img-top" alt="Recreacion Image">
              <div class="card-body">
                <h5 class="card-title">{{ recreacion.name }}</h5>
                <p class="card-text">Dirección: {{ recreacion.address }}</p>
                <p class="card-text">Servicios: {{ recreacion.services }}</p>
                <p class="card-text">Precio: {{ recreacion.price }}</p>
               
  
                <button class="btn btn-danger" @click="eliminarRecreacion(recreacion._id)">Eliminar</button>
                <button class="btn btn-primary" @click="abrirModalActualizar(recreacion)">Actualizar</button>
              </div>
            </div>
          </div>
        </div>
      </div>
  
      <!-- Modal Crear -->
      <div ref="modalCrear" class="modal fade" id="modalCrear" tabindex="" role="dialog" aria-labelledby="modalCrearLabel" aria-hidden="true">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="modalCrearLabel">Crear Recreación</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalCrear">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="crearRecreacion">
                <div class="form-group">
                  <label for="nombreRecreacion">Nombre:</label>
                  <input type="text" v-model="nuevaRecreacion.name" class="form-control" id="nombreRecreacion" required>
                </div>
                <div class="form-group">
                  <label for="direccionRecreacion">Dirección:</label>
                  <input type="text" v-model="nuevaRecreacion.address" class="form-control" id="direccionRecreacion" required>
                </div>
                <div class="form-group">
                  <label for="imagenRecreacion">Imagen:</label>
                  <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileCrear' id="imagenRecreacion" ref="fileCrear" type="file" style="display:none">
                </div>
                <div class="form-group">
                  <label for="serviciosRecreacion">Servicios:</label>
                  <input type="text" v-model="nuevaRecreacion.services" class="form-control" id="serviciosRecreacion" required>
                </div>
                <div class="form-group">
                  <label for="precioRecreacion">Precio:</label>
                  <input type="text" v-model="nuevaRecreacion.price" class="form-control" id="precioRecreacion" required>
                </div>
               
  
                <button type="submit" class="btn btn-primary">Crear Recreación</button>
              </form>
            </div>
          </div>
        </div>
      </div>
  
      <!-- Modal Actualizar -->
      <div ref="modalActualizar" class="modal fade" id="modalActualizar" tabindex="-1" role="dialog" aria-labelledby="modalActualizarLabel" aria-hidden="true">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="modalActualizarLabel">Actualizar Recreación</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalActualizar">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="actualizarRecreacion">
                <div class="form-group">
                  <label for="nombreRecreacionActualizar">Nombre:</label>
                  <input type="text" v-model="recreacionActualizada.name" class="form-control" id="nombreRecreacionActualizar" required>
                </div>
                <div class="form-group">
                  <label for="direccionRecreacionActualizar">Dirección:</label>
                  <input type="text" v-model="recreacionActualizada.address" class="form-control" id="direccionRecreacionActualizar" required>
                </div>
                <div class="form-group">
                  <label for="imagenRecreacionActualizar">Imagen:</label>
                  <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar()' id="imagenRecreacionActualizar" ref="fileActualizar" type="file" style="display:none">
                </div>
                <div class="form-group">
                  <label for="serviciosRecreacionActualizar">Servicios:</label>
                  <input type="text" v-model="recreacionActualizada.services" class="form-control" id="serviciosRecreacionActualizar" required>
                </div>
                <div class="form-group">
                  <label for="precioRecreacionActualizar">Precio:</label>
                  <input type="text" v-model="recreacionActualizada.price" class="form-control" id="precioRecreacionActualizar" required>
                </div>
               
                <button type="submit" class="btn btn-primary">Guardar Cambios</button>
              </form>
            </div>
            <div class="modal-footer">
              <p v-if="successMessage" class="text-success">{{ successMessage }}</p>
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
      recreaciones: [],
      recreacionActualizada: {
        _id: null,
        nombre: '',
        services: '',
        image: '',
        phone: '',
        price: '',
        status: '',
       
      },
      file: null,
      nuevaRecreacion: {
        nombre: '',
        services: '',
        image: '',
        phone: '',
        price: '',
        status: '',
        customer_id: '',
      },
    };
  },
  mounted() {
    this.fetchRecreaciones();
  },
  methods: {
    async fetchRecreaciones() {
      try {
        const response = await axios.get(`${process.env.API}/Recreacion`);
        this.recreaciones = response.data;
      } catch (error) {
        console.error('Error fetching recreaciones:', error);
      }
    },
    async eliminarRecreacion(id) {
      try {
        await axios.delete(`${process.env.API}/Recreacion/${id}`);
        this.fetchRecreaciones();
      } catch (error) {
        console.error('Error deleting recreacion:', error);
      }
    },
    abrirModalActualizar(recreacion) {
      this.recreacionActualizada = { ...recreacion };
      const modalElement = this.$refs.modalActualizar;
      if (modalElement) {
        modalElement.classList.add('show');
        modalElement.style.display = 'block';
      }
    },
    cerrarModalActualizar() {
      const modalElement = this.$refs.modalActualizar;
      if (modalElement) {
        modalElement.classList.remove('show');
        modalElement.style.display = 'none';
      }
    },
    abrirModalCrear() {
      const modalElement = this.$refs.modalCrear;
      if (modalElement) {
        modalElement.classList.add('show');
        modalElement.style.display = 'block';
      }
    },
    cerrarModalCrear() {
      const modalElement = this.$refs.modalCrear;
      if (modalElement) {
        modalElement.classList.remove('show');
        modalElement.style.display = 'none';
      }
    },
    async crearRecreacion() {
      let id = localStorage.getItem('customerId');
      try {
        if (id) {
          this.nuevaRecreacion.customer_id = id;
      var formData = new FormData();
      formData.append('image', this.file);
      formData.append('name', this.nuevaRecreacion.name);
      formData.append('address', this.nuevaRecreacion.address);
      formData.append('services', this.nuevaRecreacion.services);
      formData.append('price', this.nuevaRecreacion.price);
      formData.append('customer_id', this.nuevaRecreacion.customer_id);
      const response = await axios.post(`${process.env.API}/Recreacion`, formData);
        console.log(response);

        // Reset the form and close the modal
        this.nuevaRecreacion = {
          nombre: '',
          services: '',
          imagen: '',
          phone: '',
          price: '',
          status: '',
          customer_id: '',
        };

        // Close the modal
        this.cerrarModalCrear();
        this.fetchRecreaciones();
      }
      } catch (error) {
        console.error('Error creating Recreacion:', error);
      }
    },
    async actualizarRecreacion() {
      try {
        const response = await axios.put(`${process.env.API}/Recreacion/${this.recreacionActualizada._id}`, this.recreacionActualizada);
        console.log(response);
        this.fetchRecreaciones();
        // Close the modal
        this.cerrarModalActualizar();
      } catch (error) {
        console.error('Error updating Recreacion:', error);
      }
    },
    uploadFileCrear(event) {
      this.file = event.target.files[0];
    },
    uploadFileActualizar(event) {
      this.file = event.target.files[0];
    },
  },
};
</script>
