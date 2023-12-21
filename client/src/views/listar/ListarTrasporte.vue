<template>
    <div id="app">
      <div class="container">
        <br>
        <div class="card">
          <h5 class="card-header">Panel de Trasporte</h5>
          <div class="card-body">
            <h5 class="card-title">Creación de Trasporte</h5>
            <p class="card-text">En este espacio creamos nuevos transportes para COSTA BRISA</p>
            <button class="btn btn-success mb-4" @click="abrirModalCrear">Crear</button>
          </div>
        </div>
  
        <hr>
        <strong>Todos los Trasportes</strong>
        <hr>
        <div class="row">
          <div v-for="transporte in transportes" :key="transporte._id" class="col-md-4 mb-4">
            <div class="card">
              <img :src="transporte.image" class="card-img-top" alt="Trasporte Image">
              <div class="card-body">
                <h5 class="card-title">{{ transporte.tipo }}</h5>
                <p class="card-text">Placa: {{ transporte.placa }}</p>
                <p class="card-text">Conductor: {{ transporte.conductor }}</p>
                <p class="card-text">Celular: {{ transporte.celular }}</p>
  
                <button class="btn btn-danger" @click="eliminarTransporte(transporte._id)">Eliminar</button>
                <button class="btn btn-primary" @click="abrirModalActualizar(transporte)">Actualizar</button>
              </div>
            </div>
          </div>
        </div>
      </div>
  
      <!-- Modal Crear -->
      <div ref="modalCrear" class="modal fade" id="modalCrear" tabindex="-1" role="dialog" aria-labelledby="modalCrearLabel" aria-hidden="true">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="modalCrearLabel">Crear Trasporte</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalCrear">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="crearTransporte">
                <div class="form-group">
                  <label for="tipoTrasporte">Tipo:</label>
                  <input type="text" v-model="nuevoTransporte.tipo" class="form-control" id="tipoTrasporte" required>
                </div>
                <div class="form-group">
                  <label for="imagenTrasporte">Imagen:</label>
                  <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileCrear()' id="imagenTrasporte" ref="fileCrear" type="file" style="display:none">
                </div>
                <div class="form-group">
                  <label for="placaTrasporte">Placa:</label>
                  <input type="text" v-model="nuevoTransporte.placa" class="form-control" id="placaTrasporte" required>
                </div>
                <div class="form-group">
                  <label for="conductorTrasporte">Conductor:</label>
                  <input type="text" v-model="nuevoTransporte.conductor" class="form-control" id="conductorTrasporte" required>
                </div>
                <div class="form-group">
                  <label for="celularTrasporte">Celular:</label>
                  <input type="text" v-model="nuevoTransporte.celular" class="form-control" id="celularTrasporte" required>
                </div>
  
                <button type="submit" class="btn btn-primary">Crear Trasporte</button>
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
              <h5 class="modal-title" id="modalActualizarLabel">Actualizar Trasporte</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalActualizar">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="actualizarTransporte">
                <div class="form-group">
                  <label for="tipoTrasporteActualizar">Tipo:</label>
                  <input type="text" v-model="transporteActualizado.tipo" class="form-control" id="tipoTrasporteActualizar" required>
                </div>
                <div class="form-group">
                  <label for="imagenTrasporteActualizar">Imagen:</label>
                  <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar()' id="imagenTrasporteActualizar" ref="fileActualizar" type="file" style="display:none">
                </div>
                <div class="form-group">
                  <label for="placaTrasporteActualizar">Placa:</label>
                  <input type="text" v-model="transporteActualizado.placa" class="form-control" id="placaTrasporteActualizar" required>
                </div>
                <div class="form-group">
                  <label for="conductorTrasporteActualizar">Conductor:</label>
                  <input type="text" v-model="transporteActualizado.conductor" class="form-control" id="conductorTrasporteActualizar" required>
                </div>
                <div class="form-group">
                  <label for="celularTrasporteActualizar">Celular:</label>
                  <input type="text" v-model="transporteActualizado.celular" class="form-control" id="celularTrasporteActualizar" required>
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
        transportes: [],
        transporteActualizado: {
          _id: null,
          tipo: '',
          image: '',
          placa: '',
          conductor: '',
          celular: '',
        },
        nuevoTransporte: {
          tipo: '',
          image: '',
          placa: '',
          conductor: '',
          celular: '',
          customer_id: '',
        },
      };
    },
    mounted() {
      this.fetchTransportes();
    },
    methods: {
      async fetchTransportes() {
        try {
          const response = await axios.get('http://localhost:3000/Trasporte');
          this.transportes = response.data;
        } catch (error) {
          console.error('Error fetching transportes:', error);
        }
      },
      async eliminarTransporte(id) {
        try {
          await axios.delete(`http://localhost:3000/Trasporte/${id}`);
          this.fetchTransportes();
        } catch (error) {
          console.error('Error deleting transporte:', error);
        }
      },
      abrirModalActualizar(transporte) {
        this.transporteActualizado = { ...transporte };
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
      async crearTransporte() {
        let id = localStorage.getItem('customerId');
        try {
          if (id) {
            this.nuevoTransporte.customer_id = id;
          }
          const response = await axios.post('http://localhost:3000/Trasporte', this.nuevoTransporte);
          console.log(response);
  
          // Reset the form and close the modal
          this.nuevoTransporte = {
            tipo: '',
            image: '',
            placa: '',
            conductor: '',
            celular: '',
            customer_id: '',
          };
  
          // Close the modal
          this.cerrarModalCrear();
          this.fetchTransportes();
        } catch (error) {
          console.error('Error creating Trasporte:', error);
        }
      },
      async actualizarTransporte() {
        try {
          const response = await axios.put(`http://localhost:3000/Trasporte/${this.transporteActualizado._id}`, this.transporteActualizado);
          console.log(response);
          this.fetchTransportes();
          // Close the modal
          this.cerrarModalActualizar();
        } catch (error) {
          console.error('Error updating Trasporte:', error);
        }
      },
      uploadFileCrear() {
        this.uploadFile('fileCrear', 'nuevoTransporte');
      },
      uploadFileActualizar() {
        this.uploadFile('fileActualizar', 'transporteActualizado');
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
  