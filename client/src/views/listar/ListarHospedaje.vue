<template>
    <div id="app">
      <div class="container">
        <br>
        <div class="card">
          <h5 class="card-header">Panel de Hospedaje</h5>
          <div class="card-body">
            <h5 class="card-title">Creacion de Hospedaje</h5>
            <p class="card-text">En este espacio creamos nuevos hospedajes para COSTA BRISA</p>
            <button class="btn btn-success mb-4" @click="abrirModalCrear">Crear</button>
          </div>
        </div>
  
        <hr>
        <strong>Todos los Hospedajes</strong>
        <hr>
        <div class="row">
          <div v-for="hospedaje in hospedajes" :key="hospedaje._id" class="col-md-4 mb-4">
            <div class="card">
              <img :src="hospedaje.image" class="card-img-top" alt="Hospedaje Image">
              <div class="card-body">
                <h5 class="card-title">{{ hospedaje.name }}</h5>
                <p class="card-text">Dirección: {{ hospedaje.address }}</p>
                <p class="card-text">Descripción: {{ hospedaje.description }}</p>
                <p class="card-text">Teléfono: {{ hospedaje.phone }}</p>
  
                <button class="btn btn-danger" @click="eliminarHospedaje(hospedaje._id)">Eliminar</button>
                <button class="btn btn-primary" @click="abrirModalActualizar(hospedaje)">Actualizar</button>
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
        <h5 class="modal-title" id="modalCrearLabel">Crear Hospedaje</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalCrear">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">
        <form @submit.prevent="crearHospedaje">
          <div class="form-group">
            <label for="nombreHospedaje">Nombre:</label>
            <input type="text" v-model="nuevoHospedaje.name" class="form-control" id="nombreHospedaje" required>
          </div>
          <div class="form-group">
            <label for="direccionHospedaje">Dirección:</label>
            <input type="text" v-model="nuevoHospedaje.address" class="form-control" id="direccionHospedaje" required>
          </div>
          <div class="form-group">
            <label for="imagenHospedaje">Imagen:</label>
            <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileCrear' id="imagenHospedaje" ref="fileCrear" type="file" style="display:none">
          </div>
          <div class="form-group">
            <label for="telefonoHospedaje">Teléfono:</label>
            <input type="text" v-model="nuevoHospedaje.phone" class="form-control" id="telefonoHospedaje" required>
          </div>
          
          <button type="submit" class="btn btn-primary">Crear Hospedaje</button>
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
        <h5 class="modal-title" id="modalActualizarLabel">Actualizar Hospedaje</h5>
        <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalActualizar">
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
      <div class="modal-body">
        <form @submit.prevent="actualizarHospedaje">
          <div class="form-group">
            <label for="nombreHospedajeActualizar">Nombre:</label>
            <input type="text" v-model="hospedajeActualizado.name" class="form-control" id="nombreHospedajeActualizar" required>
          </div>
          <div class="form-group">
            <label for="direccionHospedajeActualizar">Dirección:</label>
            <input type="text" v-model="hospedajeActualizado.address" class="form-control" id="direccionHospedajeActualizar" required>
          </div>
          <div class="form-group">
            <label for="imagenHospedajeActualizar">Imagen:</label>
            <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar' id="imagenHospedajeActualizar" ref="fileActualizar" type="file" style="display:none">
          </div>
          <div class="form-group">
            <label for="telefonoHospedajeActualizar">Teléfono:</label>
            <input type="text" v-model="hospedajeActualizado.phone" class="form-control" id="telefonoHospedajeActualizar" required>
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
      hospedajes: [],
      hospedajeActualizado: {
        _id: null,
        name: '',
        address: '',
        image: '',
        phone: '',
    
      },
      file: null,
      fileActualizar: null,
      nuevoHospedaje: {
        name: '',
        address: '',
        image: '',
        phone: '',
        customer_id: '',
      },
    };
  },
  mounted() {
    this.fetchHospedajes();
  },
  methods: {
    async fetchHospedajes() {
      try {
        const response = await axios.get(`${process.env.API}/Hospedaje`);
        this.hospedajes = response.data;
      } catch (error) {
        console.error('Error fetching hospedajes:', error);
      }
    },
    async eliminarHospedaje(id) {
      try {
        await axios.delete(`${process.env.API}/Hospedaje/${id}`);
        this.fetchHospedajes();
      } catch (error) {
        console.error('Error deleting hospedaje:', error);
      }
    },
    abrirModalActualizar(hospedaje) {
      this.hospedajeActualizado = { ...hospedaje };
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
    async crearHospedaje() {
      let id = localStorage.getItem('customerId');
      try {
        if (id) {
          this.nuevoHospedaje.customer_id = id;
          const formData = new FormData();
          formData.append('image', this.file);
          formData.append('name', this.nuevoHospedaje.name);
          formData.append('address', this.nuevoHospedaje.address);
          formData.append('phone', this.nuevoHospedaje.phone);
          formData.append('customer_id', this.nuevoHospedaje.customer_id);
          const response = await axios.post(`${process.env.API}/Hospedaje`, formData);
        
        console.log(response);

        // Reset the form and close the modal
        this.nuevoHospedaje = {
          name: '',
          address: '',
          image: '',
          phone: '',
          customer_id: '',
        };

        // Close the modal
        this.cerrarModalCrear();
        this.fetchHospedajes();
      }
      } catch (error) {
        console.error('Error creating Hospedaje:', error);
      }
    },
    async actualizarHospedaje() {
      if (this.fileActualizar) {
        const formData = new FormData();
        formData.append('name', this.hospedajeActualizado.name);
        formData.append('address', this.hospedajeActualizado.address);
        formData.append('phone', this.hospedajeActualizado.phone);
        formData.append('customer_id', this.hospedajeActualizado.customer_id);
        formData.append('imagen', this.fileActualizar);
        formData.append('image', "");
        try {
          const response = await axios.put(`${process.env.API}/Hospedaje/${this.hospedajeActualizado._id}`, formData);
          console.log(response);
          this.fetchHospedajes();
          this.cerrarModalActualizar();
          this.fileActualizar = null;
          
        } catch (error) {
          console.error('Error updating Hospedaje:', error);
        }
      }else{
      try {
        const response = await axios.put(`${process.env.API}/Hospedaje/${this.hospedajeActualizado._id}`, this.hospedajeActualizado);
        console.log(response);
        this.fetchHospedajes();
        // Close the modal
        this.cerrarModalActualizar();
      } catch (error) {
        console.error('Error updating Hospedaje:', error);
      }
    }
    },
    uploadFileCrear(event){
      this.file = event.target.files[0];
    },
    uploadFileActualizar(event) {
      this.fileActualizar = event.target.files[0];
    },
  
  },
};
</script>
