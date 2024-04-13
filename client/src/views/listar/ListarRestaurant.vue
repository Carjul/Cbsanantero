<template>
    <div id="app">
      <div class="container">
        <br>
        <div class="card">
          <h5 class="card-header">Panel de Restaurantes</h5>
          <div class="card-body">
            <h5 class="card-title">Creación de Restaurante</h5>
            <p class="card-text">En este espacio creamos nuevos restaurantes para COSTA BRISA</p>
            <button class="btn btn-success mb-4" @click="abrirModalCrear">Crear</button>
          </div>
        </div>
  
        <hr>
        <strong>Todos los Restaurantes</strong>
        <hr>
        <div class="row">
          <div v-for="restaurante in restaurantes" :key="restaurante._id" class="col-md-4 mb-4">
            <div class="card">
              <img :src="restaurante.image" class="card-img-top" alt="Restaurante Image">
              <div class="card-body">
                <h5 class="card-title">{{ restaurante.name }}</h5>
                <p class="card-text">Dirección: {{ restaurante.address }}</p>
                <p class="card-text">Servicios: {{ restaurante.services }}</p>
                <p class="card-text">Precio: {{ restaurante.price }}</p>
  
  
                <button class="btn btn-danger" @click="eliminarRestaurante(restaurante._id)">Eliminar</button>
                <button class="btn btn-primary" @click="abrirModalActualizar(restaurante)">Actualizar</button>
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
              <h5 class="modal-title" id="modalCrearLabel">Crear Restaurante</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalCrear">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="crearRestaurante">
                <div class="form-group">
                  <label for="nombreRestaurante">Nombre:</label>
                  <input type="text" v-model="nuevoRestaurante.name" class="form-control" id="nombreRestaurante" required>
                </div>
                <div class="form-group">
                  <label for="direccionRestaurante">Dirección:</label>
                  <input type="text" v-model="nuevoRestaurante.address" class="form-control" id="direccionRestaurante" required>
                </div>
                <div class="form-group">
                  <label for="imagenRestaurante">Imagen:</label>
                  <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileCrear' id="imagenRestaurante" ref="fileCrear" type="file" style="display:none">
                </div>
                <div class="form-group">
                  <label for="serviciosRestaurante">Servicios:</label>
                  <input type="text" v-model="nuevoRestaurante.services" class="form-control" id="serviciosRestaurante" required>
                </div>
                <div class="form-group">
                  <label for="precioRestaurante">Precio:</label>
                  <input type="text" v-model="nuevoRestaurante.price" class="form-control" id="precioRestaurante" required>
                </div>
  
  
                <button type="submit" class="btn btn-primary">Crear Restaurante</button>
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
              <h5 class="modal-title" id="modalActualizarLabel">Actualizar Restaurante</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalActualizar">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="actualizarRestaurante">
                <div class="form-group">
                  <label for="nombreRestauranteActualizar">Nombre:</label>
                  <input type="text" v-model="restauranteActualizado.name" class="form-control" id="nombreRestauranteActualizar" required>
                </div>
                <div class="form-group">
                  <label for="direccionRestauranteActualizar">Dirección:</label>
                  <input type="text" v-model="restauranteActualizado.address" class="form-control" id="direccionRestauranteActualizar" required>
                </div>
                <div class="form-group">
                  <label for="imagenRestauranteActualizar">Imagen:</label>
                  <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar' id="imagenRestauranteActualizar" ref="fileActualizar" type="file" style="display:none">
                </div>
                <div class="form-group">
                  <label for="serviciosRestauranteActualizar">Servicios:</label>
                  <input type="text" v-model="restauranteActualizado.services" class="form-control" id="serviciosRestauranteActualizar" required>
                </div>
                <div class="form-group">
                  <label for="precioRestauranteActualizar">Precio:</label>
                  <input type="text" v-model="restauranteActualizado.price" class="form-control" id="precioRestauranteActualizar" required>
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
        restaurantes: [],
        restauranteActualizado: {
          _id: null,
          nombre: '',
          services: '',
          image: '',
          phone: '',
          price: '',
          status: '',
          customer_id: '',
         
        },
        file: null,
        fileActualizar: null, 
        nuevoRestaurante: {
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
      this.fetchRestaurantes();
    },
    methods: {
      async fetchRestaurantes() {
        try {
          const response = await axios.get(`${process.env.API}/Restaurante`);
          this.restaurantes = response.data;
        } catch (error) {
          console.error('Error fetching restaurantes:', error);
        }
      },
      async eliminarRestaurante(id) {
        try {
          await axios.delete(`${process.env.API}/Restaurante/${id}`);
          this.fetchRestaurantes();
        } catch (error) {
          console.error('Error deleting restaurante:', error);
        }
      },
      abrirModalActualizar(restaurante) {
        this.restauranteActualizado = { ...restaurante };
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
      async crearRestaurante() {
        let id = localStorage.getItem('customerId');
        try {
          if (id) {
            this.nuevoRestaurante.customer_id = id;
        var formData = new FormData();
        formData.append('image', this.file);
        formData.append('name', this.nuevoRestaurante.name);
        formData.append('address', this.nuevoRestaurante.address);
        formData.append('services', this.nuevoRestaurante.services);
        formData.append('price', this.nuevoRestaurante.price);
        formData.append('customer_id', this.nuevoRestaurante.customer_id);
        const response = await axios.post(`${process.env.API}/Restaurante`, formData);
          console.log(response);
  
          // Reset the form and close the modal
          this.nuevoRestaurante = {
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
          this.fetchRestaurantes();
        }
        } catch (error) {
          console.error('Error creating Restaurante:', error);
        }
      },
      async actualizarRestaurante() {
        if (this.fileActualizar) {
          var formData = new FormData();
          formData.append('imagen', this.fileActualizar);
          formData.append('image', "");
          formData.append('name', this.restauranteActualizado.name);
          formData.append('address', this.restauranteActualizado.address);
          formData.append('services', this.restauranteActualizado.services);
          formData.append('price', this.restauranteActualizado.price);
          formData.append('customer_id', this.restauranteActualizado.customer_id);
          try {
            const response = await axios.put(`${process.env.API}/Restaurante/${this.restauranteActualizado._id}`, formData);
            console.log(response);
            this.fetchRestaurantes();
            this.cerrarModalActualizar();
            this.fileActualizar = null;
          } catch (error) {
            console.error('Error updating Restaurante:', error);
          }
          return;
        }
        try {
          const response = await axios.put(`${process.env.API}/Restaurante/${this.restauranteActualizado._id}`, this.restauranteActualizado);
          console.log(response);
          this.fetchRestaurantes();
          // Close the modal
          this.cerrarModalActualizar();
        } catch (error) {
          console.error('Error updating Restaurante:', error);
        }
      },
      uploadFileCrear(event) {
        this.file = event.target.files[0];
      },
      uploadFileActualizar(event) {
        this.fileActualizar = event.target.files[0];
      },
    },
  };
  </script>
  