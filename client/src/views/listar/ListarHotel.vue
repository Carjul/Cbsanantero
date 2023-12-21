<template>
    <div id="app">
      <div class="container">
        <br>
        <div class="card">
          <h5 class="card-header">Panel de Hotel</h5>
          <div class="card-body">
            <h5 class="card-title">Creación de Hotel</h5>
            <p class="card-text">En este espacio creamos nuevos hoteles para COSTA BRISA</p>
            <button class="btn btn-success mb-4" @click="abrirModalCrear">Crear</button>
          </div>
        </div>
  
        <hr>
        <strong>Todos los Hoteles</strong>
        <hr>
        <div class="row">
          <div v-for="hotel in hoteles" :key="hotel._id" class="col-md-4 mb-4">
            <div class="card">
              <img :src="hotel.image" class="card-img-top" alt="Hotel Image">
              <div class="card-body">
                <h5 class="card-title">{{ hotel.name }}</h5>
                <p class="card-text">Dirección: {{ hotel.address }}</p>
                <p class="card-text">Descripción: {{ hotel.description }}</p>
                <p class="card-text">Teléfono: {{ hotel.phone }}</p>
  
                <button class="btn btn-danger" @click="eliminarHotel(hotel._id)">Eliminar</button>
                <button class="btn btn-primary" @click="abrirModalActualizar(hotel)">Actualizar</button>
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
              <h5 class="modal-title" id="modalCrearLabel">Crear Hotel</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalCrear">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="crearHotel">
                <div class="form-group">
                  <label for="nombreHotel">Nombre:</label>
                  <input type="text" v-model="nuevoHotel.name" class="form-control" id="nombreHotel" required>
                </div>
                <div class="form-group">
                  <label for="direccionHotel">Dirección:</label>
                  <input type="text" v-model="nuevoHotel.address" class="form-control" id="direccionHotel" required>
                </div>
                <div class="form-group">
                  <label for="imagenHotel">Imagen:</label>
                  <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileCrear()' id="imagenHotel" ref="fileCrear" type="file" style="display:none">
                </div>
                <div class="form-group">
                  <label for="telefonoHotel">Teléfono:</label>
                  <input type="text" v-model="nuevoHotel.phone" class="form-control" id="telefonoHotel" required>
                </div>
  
                <button type="submit" class="btn btn-primary">Crear Hotel</button>
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
              <h5 class="modal-title" id="modalActualizarLabel">Actualizar Hotel</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalActualizar">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="actualizarHotel">
                <div class="form-group">
                  <label for="nombreHotelActualizar">Nombre:</label>
                  <input type="text" v-model="hotelActualizado.name" class="form-control" id="nombreHotelActualizar" required>
                </div>
                <div class="form-group">
                  <label for="direccionHotelActualizar">Dirección:</label>
                  <input type="text" v-model="hotelActualizado.address" class="form-control" id="direccionHotelActualizar" required>
                </div>
                <div class="form-group">
                  <label for="imagenHotelActualizar">Imagen:</label>
                  <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar()' id="imagenHotelActualizar" ref="fileActualizar" type="file" style="display:none">
                </div>
                <div class="form-group">
                  <label for="telefonoHotelActualizar">Teléfono:</label>
                  <input type="text" v-model="hotelActualizado.phone" class="form-control" id="telefonoHotelActualizar" required>
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
        hoteles: [],
        hotelActualizado: {
          _id: null,
          name: '',
          address: '',
          image: '',
          phone: '',
        },
        nuevoHotel: {
          name: '',
          address: '',
          image: '',
          phone: '',
          customer_id: '',
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
          console.error('Error fetching hoteles:', error);
        }
      },
      async eliminarHotel(id) {
        try {
          await axios.delete(`http://localhost:3000/Hotel/${id}`);
          this.fetchHoteles();
        } catch (error) {
          console.error('Error deleting hotel:', error);
        }
      },
      abrirModalActualizar(hotel) {
        this.hotelActualizado = { ...hotel };
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
      async crearHotel() {
        let id = localStorage.getItem('customerId');
        try {
          if (id) {
            this.nuevoHotel.customer_id = id;
          }
          const response = await axios.post('http://localhost:3000/Hotel', this.nuevoHotel);
          console.log(response);
  
          // Reset the form and close the modal
          this.nuevoHotel = {
            name: '',
            address: '',
            image: '',
            phone: '',
            customer_id: '',
          };
  
          // Close the modal
          this.cerrarModalCrear();
          this.fetchHoteles();
        } catch (error) {
          console.error('Error creating Hotel:', error);
        }
      },
      async actualizarHotel() {
        try {
          const response = await axios.put(`http://localhost:3000/Hotel/${this.hotelActualizado._id}`, this.hotelActualizado);
          console.log(response);
          this.fetchHoteles();
          // Close the modal
          this.cerrarModalActualizar();
        } catch (error) {
          console.error('Error updating Hotel:', error);
        }
      },
      uploadFileCrear() {
        this.uploadFile('fileCrear', 'nuevoHotel');
      },
      uploadFileActualizar() {
        this.uploadFile('fileActualizar', 'hotelActualizado');
      },
      uploadFile(refName, dataProperty) {
        const fileInput = this.$refs[refName];
        if (fileInput && fileInput.files.length > 0) {
          const file = fileInput.files[0];
          var sizeByte = file.size;
          var sizeKiloByte = parseInt(sizeByte / 1024);
  
          if (sizeKiloByte > 2048) {
            console.log('La imagen es muy grande');
          } else {
            this.createImage(file, dataProperty);
          }
        }
      },
      createImage(file, dataProperty) {
        var reader = new FileReader();
        reader.onload = (e) => {
          this[dataProperty].image = e.target.result;
        };
        reader.readAsDataURL(file);
      },
    },
  };
  </script>
  