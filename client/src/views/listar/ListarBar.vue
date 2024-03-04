<template>
  <div id="app">
    <div class="container">
      <br>
      <div class="card">
        <h5 class="card-header">Panel de Bar</h5>
        <div class="card-body">
          <h5 class="card-title">Creación de Bar</h5>
          <p class="card-text">En este espacio creamos nuevos Bares para COSTA BRISA</p>
          <button class="btn btn-success mb-4" @click="abrirModalCrear">Crear</button>
        </div>
      </div>
      
      <hr>
      <strong>Todos los Bares</strong>
      <hr>
      <div class="row">
        <div v-for="bar in bares" :key="bar._id" class="col-md-4 mb-4">
          <div class="card">
            <img :src="bar.image" class="card-img-top" alt="Bar Image">
            <div class="card-body">
              <h5 class="card-title">{{ bar.name }}</h5>
              <p class="card-text">Dirección: {{ bar.address }}</p>
              <p class="card-text">Descripción: {{ bar.description }}</p>
              <p class="card-text">Teléfono: {{ bar.phone }}</p>
              <p class="card-text">Estado: {{ bar.status }}</p>

              <button class="btn btn-danger" @click="eliminarBar(bar._id)">Eliminar</button>
              <button class="btn btn-primary" @click="abrirModalActualizar(bar)">Actualizar</button>
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
            <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFile()' id="imagenBar" ref="file" type="file" style="display:none">
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
            <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar()' id="imagenBarActualizar" ref="fileActualizar" type="file" style="display:none">
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
        status: '',
        customer_id: '',
      },
      file: null,
      nuevoBar: {
        name: '',
        address: '',
        image: '',
        phone: '',
        description: '',
        status: '',
        customer_id: '',
      },
    };
  },
  mounted() {
    this.fetchBares();
  },
  methods: {
    async fetchBares() {
      try {
        const response = await axios.get('http://localhost:3000/Bar');
        this.bares = response.data;
      } catch (error) {
        console.error('Error fetching bares:', error);
      }
    },
    async eliminarBar(id) {
      try {
        await axios.delete(`http://localhost:3000/Bar/${id}`);
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
      let id = localStorage.getItem('customerId');
      try {
        if (id) {
          this.nuevoBar.customer_id = id;
        }
        const response = await axios.post('http://localhost:3000/Bar', this.nuevoBar);
        console.log(response);

        // Reset the form and close the modal
        this.nuevoBar = {
          name: '',
          address: '',
          image: '',
          phone: '',
          description: '',
          status: '',
          customer_id: '',
        };

        // Close the modal
        this.cerrarModalCrear();
        this.fetchBares();
      } catch (error) {
        console.error('Error creating Bar:', error);
      }
    },
    async actualizarBar() {
      try {
        const response = await axios.put(`http://localhost:3000/Bar/${this.barActualizado._id}`, this.barActualizado);
        console.log(response);
        this.fetchBares();
        // Close the modal
        this.cerrarModalActualizar();
      } catch (error) {
        console.error('Error updating Bar:', error);
      }
    },
    uploadFileCrear() {
      this.uploadFile('fileCrear', 'nuevoBar');
    },
    uploadFileActualizar() {
      this.uploadFile('fileActualizar', 'barActualizado');
    },
     uploadFile() {
      this.file = this.$refs.file.files[0];
      var sizeByte = this.$refs.file.files[0].size;
      var extencion = this.$refs.file.files[0].type;
      /*
      console.log(this.$refs.file.files) */
      if (extencion !== 'image/jpeg' && extencion !== 'image/jpg' && extencion !== 'image/png') {
        console.log('La imagen no es valida');
        alert('La imagen no es valida');
      } 
      var siezekiloByte = parseInt(sizeByte / 1024);
      if (siezekiloByte > 2048) {
        console.log('La imagen es muy grande');
        alert('La imagen es muy grande: ' + sizeByte + ' bytes');
      } else {
        this.createImage(this.file);
      }
    },
    createImage(file) {
      var reader = new FileReader();
      reader.onload = (e) => {
        this.nuevoBar.image = e.target.result;
      };
      reader.readAsDataURL(file);
    },
  },
};
</script>