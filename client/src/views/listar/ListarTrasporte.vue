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
                  <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFile' id="imagenTrasporte" ref="fileCrear" type="file" style="display:none">
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
                  <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar' id="imagenTrasporteActualizar" ref="fileActualizar" type="file" style="display:none">
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
          customer_id: '',
        },
        file: null,
        fileActualizar: null,
        nuevoTransporte: {
          tipo: '',
          image: '',
          placa: '',
          conductor: '',
          celular: '',
          customer_id: '',
        },
        // Agregar estado para controlar la visibilidad de los modales
        mostrarModalActualizar: false,
        mostrarModalCrear: false,
      };
    },
    mounted() {
      this.fetchTransportes();
    },
    methods: {
      async fetchTransportes() {
      let id = localStorage.getItem('customerId');

        try {
          const response = await axios.get(`${process.env.API}/Trasporte`);
          let Filtro = response.data.filter((customer) => customer._id === id);
          this.transportes = Filtro || [];
        } catch (error) {
          console.error('Error fetching transportes:', error);
          // Mostrar mensaje de error al usuario
        }
      },
      async eliminarTransporte(id) {
        try {
          await axios.delete(`${process.env.API}/Trasporte/${id}`);
          this.fetchTransportes();
        } catch (error) {
          console.error('Error deleting transporte:', error);
          // Mostrar mensaje de error al usuario
        }
      },
      abrirModalActualizar(transporte) {
        this.transporteActualizado = { ...transporte };
        this.mostrarModalActualizar = true;
      },
      cerrarModalActualizar() {
        this.mostrarModalActualizar = false;
      },
      abrirModalCrear() {
        this.mostrarModalCrear = true;
      },
      cerrarModalCrear() {
        this.mostrarModalCrear = false;
      },
      async crearTransporte() {
        let id = localStorage.getItem('customerId');
        try {
          if (id) {
            this.nuevoTransporte.customer_id = id;
            var formData = new FormData();
            formData.append('image', this.file);
            formData.append('tipo', this.nuevoTransporte.tipo);
            formData.append('placa', this.nuevoTransporte.placa);
            formData.append('conductor', this.nuevoTransporte.conductor);
            formData.append('celular', this.nuevoTransporte.celular);
            formData.append('customer_id', this.nuevoTransporte.customer_id);
            const response = await axios.post(`${process.env.API}/Trasporte`, formData);

            console.log(response);

            // Resetear el formulario y cerrar el modal
            this.resetearFormulario('nuevoTransporte');
            this.cerrarModalCrear();
            this.fetchTransportes();
          }
        } catch (error) {
          console.error('Error creating Trasporte:', error);
          // Mostrar mensaje de error al usuario
        }
      },
      async actualizarTransporte() {
        if (this.fileActualizar) {
          var formData = new FormData();
          formData.append('tipo', this.transporteActualizado.tipo);
          formData.append('placa', this.transporteActualizado.placa);
          formData.append('conductor', this.transporteActualizado.conductor);
          formData.append('celular', this.transporteActualizado.celular);
          formData.append('customer_id', this.transporteActualizado.customer_id);
          formData.append('imagen', this.fileActualizar);
          formData.append('image', "");
          try {
            const response = await axios.put(`${process.env.API}/Trasporte/${this.transporteActualizado._id}`, formData);
            console.log(response);
            this.fetchTransportes();
            this.cerrarModalActualizar();
            this.fileActualizar = null;
          } catch (error) {
            console.error('Error updating Trasporte:', error);
            // Mostrar mensaje de error al usuario
          }
          return;
        }
        try {
          const response = await axios.put(`${process.env.API}/Trasporte/${this.transporteActualizado._id}`, this.transporteActualizado);
          console.log(response);
          this.fetchTransportes();
          // Cerrar el modal
          this.cerrarModalActualizar();
        } catch (error) {
          console.error('Error updating Trasporte:', error);
          // Mostrar mensaje de error al usuario
        }
      },
    
    uploadFile(event) {
      this.file = event.target.files[0];
    },
    uploadFileActualizar(event) {
      this.fileActualizar = event.target.files[0];
    },
      resetearFormulario(formulario) {
        // Resetear los valores del formulario
        this[formulario] = {
          tipo: '',
          image: '',
          placa: '',
          conductor: '',
          celular: '',
          customer_id: '',
        };
      }
    },
  };
</script>
