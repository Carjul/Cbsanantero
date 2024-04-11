<template>
  <div id="app">
    <div class="container">
      <br>
      <div class="card">
        <h5 class="card-header">Panel de Bares</h5>
        <div class="card-body">
          <h5 class="card-title">Creación de Bares</h5>
          <p class="card-text">En este espacio creamos nuevos bares para COSTA BRISA</p>
          <button class="btn btn-success mb-4" @click="abrirModalCrear">Crear</button>
          <router-link to="/bargaleria">Galería</router-link>
        </div>
      </div>
      
      <hr>
      <strong>Todos los bares</strong>
      <hr>
      <div class="row">
        <div v-for="bar in bares" :key="bar._id" class="col-md-4 mb-4">
          <div class="card">
            <img :src="bar.image" class="card-img-top" alt="Imagen del Bar">
            <div class="card-body">
              <h5 class="card-title">{{ bar.name }}</h5>
              <p class="card-text">Dirección: {{ bar.address }}</p>
              <p class="card-text">Descripción: {{ bar.description }}</p>
              <p class="card-text">Teléfono: {{ bar.phone }}</p>

              <button class="btn btn-danger" @click="eliminarBar(bar._id)">Eliminar</button>
              <button class="btn btn-primary" @click="abrirModalActualizar(bar)">Actualizar</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal Crear -->
    <div ref="modalCrear" class="modal fade" id="modalCrear" tabindex="" role="dialog" aria-labelledby="modalCrearLabel" aria-hidden="true">
      <!-- Código del modal para crear un Bar -->
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="modalCrearLabel">Crear Bar</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalCrear">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="crearBar">
              <div class="form-group">
                <label for="nombreBar">Nombre:</label>
                <input type="text" v-model="nuevoBar.name" class="form-control" id="nombreBar" required>
              </div>
              <div class="form-group">
                <label for="direccionBar">Dirección:</label>
                <input type="text" v-model="nuevoBar.address" class="form-control" id="direccionBar" required>
              </div>
              <div class="form-group">
                <label for="imagenBar">Imagen:</label>
                <input name="image" class="form-control" accept="image/jpeg, image/jpg, image/png" @change="uploadFile" id="imagenBar" ref="file" type="file" multiple style="display:none">
              </div>
              <div class="form-group">
                <label for="telefonoBar">Teléfono:</label>
                <input type="text" v-model="nuevoBar.phone" class="form-control" id="telefonoBar" required>
              </div>
              <div class="form-group">
                <label for="descripcionBar">Descripción:</label>
                <textarea v-model="nuevoBar.description" class="form-control" id="descripcionBar" required></textarea>
              </div>
              <button type="submit" class="btn btn-primary">Crear Bar</button>
            </form>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal Actualizar -->
    <div ref="modalActualizar" class="modal fade" id="modalActualizar" tabindex="-1" role="dialog" aria-labelledby="modalActualizarLabel" aria-hidden="true">
      <!-- Código del modal para actualizar un Bar -->
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="modalActualizarLabel">Actualizar Bar</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalActualizar">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="actualizarBar">
              <div class="form-group">
                <label for="nombreBarActualizar">Nombre:</label>
                <input type="text" v-model="barActualizado.name" class="form-control" id="nombreBarActualizar" required>
              </div>
              <div class="form-group">
                <label for="direccionBarActualizar">Dirección:</label>
                <input type="text" v-model="barActualizado.address" class="form-control" id="direccionBarActualizar" required>
              </div>
              <div class="form-group">
                <label for="imagenBarActualizar">Imagen:</label>
                <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change="uploadFileActualizar" id="imagenBarActualizar" ref="fileActualizar" type="file" style="display:none">
              </div>
              <div class="form-group">
                <label for="telefonoBarActualizar">Teléfono:</label>
                <input type="text" v-model="barActualizado.phone" class="form-control" id="telefonoBarActualizar" required>
              </div>
              <div class="form-group">
                <label for="descripcionBarActualizar">Descripción:</label>
                <textarea v-model="barActualizado.description" class="form-control" id="descripcionBarActualizar" required></textarea>
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
      bares: [],
      barActualizado: {
        _id: null,
        name: '',
        address: '',
        image: '',
        phone: '',
        description: '',
        customer_id: '',
      },
      file: null,
      fileActualizar: null,
      nuevoBar: {
        name: '',
        address: '',
        image: '',
        phone: '',
        description: '',
        customer_id: '',
      },
    };
  },
  mounted() {
    this.fetchBares();
    this.nuevoBar.customer_id = localStorage.getItem('customerId');

  },
  methods: {
    async fetchBares() {
      try {
        const response = await axios.get(`${process.env.API}/Bar`);
        this.bares = response.data;
      } catch (error) {
        console.error('Error fetching bares:', error);
      }
    },
    async eliminarBar(id) {
      try {
        await axios.delete(`${process.env.API}/Bar/${id}`);
        this.fetchBares();
      } catch (error) {
        console.error('Error deleting bar:', error);
      }
    },
    abrirModalActualizar(bar) {
      this.barActualizado = { ...bar };
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
    async crearBar() {
      try {
        const formData = new FormData();

        formData.append('name', this.nuevoBar.name);
        formData.append('address', this.nuevoBar.address);
        formData.append('phone', this.nuevoBar.phone);
        formData.append('description', this.nuevoBar.description);
        formData.append('customer_id', this.nuevoBar.customer_id);
        formData.append('image', this.file);
    
        // Agregar datos del nuevo Bar al FormData
        const response = await axios.post(`${process.env.API}/Bar`, formData);
        console.log(response);

        // Limpiar formulario y cerrar modal
        this.nuevoBar = {
          name: '',
          address: '',
          image: '',
          phone: '',
          description: '',
          customer_id: '',
        };
        this.file= null;
        this.cerrarModalCrear();

        // Actualizar la lista de bares
        this.fetchBares();
      } catch (error) {
        console.error('Error creating Bar:', error);
      }
    },
    async actualizarBar() {
      if (this.fileActualizar) {
        const formData = new FormData();
        formData.append('name', this.barActualizado.name);
        formData.append('address', this.barActualizado.address);
        formData.append('phone', this.barActualizado.phone);
        formData.append('description', this.barActualizado.description);
        formData.append('customer_id', this.barActualizado.customer_id);
        formData.append('imagen', this.fileActualizar);
        formData.append('image', "");


        try {
          const response = await axios.put(`${process.env.API}/Bar/${this.barActualizado._id}`, formData);
          console.log(response);

          // Actualizar la lista de bares y cerrar modal
          this.fetchBares();
          this.cerrarModalActualizar();
          this.fileActualizar = null;
        } catch (error) {
          console.error('Error updating Bar:', error);
        }
        return;
      }
      try {
        const response = await axios.put(`${process.env.API}/Bar/${this.barActualizado._id}`, this.barActualizado);
        console.log(response);

        // Actualizar la lista de bares y cerrar modal
        this.fetchBares();
        this.cerrarModalActualizar();
      } catch (error) {
        console.error('Error updating Bar:', error);
      }
    },
    uploadFile(event) {
      this.file = event.target.files[0];
    },
    uploadFileActualizar(event) {
      this.fileActualizar = event.target.files[0];
    },
  },
};
</script>
