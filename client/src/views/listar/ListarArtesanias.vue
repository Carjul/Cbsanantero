<template>
    <div id="app">
      <div class="container">
        <br>
        <div class="card">
          <h5 class="card-header">Panel de Artesania</h5>
          <div class="card-body">
            <h5 class="card-title">Creacion de Artesania</h5>
            <p class="card-text">En este espacio creamos nuevas artesanías para COSTA BRISA</p>
            <button class="btn btn-success mb-4" @click="abrirModalCrear">Crear</button>
            <router-link to="/artegaleria">Galeria</router-link>
          </div>
        </div>
        
        <hr>
        <strong>Todas las artesanías</strong>
        <hr>
        <div class="row">
          <div v-for="artesania in artesanias" :key="artesania._id" class="col-md-4 mb-4">
            <div class="card">
              <img :src="artesania.image" class="card-img-top" alt="Artesania Image">
              <div class="card-body">
                <h5 class="card-title">{{ artesania.name }}</h5>
                <p class="card-text">Dirección: {{ artesania.address }}</p>
                <p class="card-text">Descripción: {{ artesania.description }}</p>
                <p class="card-text">Teléfono: {{ artesania.phone }}</p>
  
                <button class="btn btn-danger" @click="eliminarArtesania(artesania._id)">Eliminar</button>
                <button class="btn btn-primary" @click="abrirModalActualizar(artesania)">Actualizar</button>
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
              <h5 class="modal-title" id="modalCrearLabel">Crear Artesania</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalCrear">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="crearArtesania">
                <div class="form-group">
                  <label for="nombreArtesania">Nombre:</label>
                  <input type="text" v-model="nuevaArtesania.name" class="form-control" id="nombreArtesania" required>
                </div>
                <div class="form-group">
                  <label for="direccionArtesania">Dirección:</label>
                  <input type="text" v-model="nuevaArtesania.address" class="form-control" id="direccionArtesania" required>
                </div>
                <div class="form-group">
                  <label for="imagenArtesania">Imagen:</label>
                  <input name="image" class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFile' id="imagenArtesania" ref="file" type="file" multiple style="display:none">
                </div>
                <div class="form-group">
                  <label for="telefonoArtesania">Teléfono:</label>
                  <input type="text" v-model="nuevaArtesania.phone" class="form-control" id="telefonoArtesania" required>
                </div>
                <div class="form-group">
                  <label for="descripcionArtesania">Descripción:</label>
                  <textarea v-model="nuevaArtesania.description" class="form-control" id="descripcionArtesania" required></textarea>
                </div>
                <button type="submit" class="btn btn-primary">Crear Artesania</button>
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
              <h5 class="modal-title" id="modalActualizarLabel">Actualizar Artesania</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalActualizar">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="actualizarArtesania">
                <div class="form-group">
                  <label for="nombreArtesaniaActualizar">Nombre:</label>
                  <input type="text" v-model="artesaniaActualizada.name" class="form-control" id="nombreArtesaniaActualizar" required>
                </div>
                <div class="form-group">
                  <label for="direccionArtesaniaActualizar">Dirección:</label>
                  <input type="text" v-model="artesaniaActualizada.address" class="form-control" id="direccionArtesaniaActualizar" required>
                </div>
                <div class="form-group">
                  <label for="imagenArtesaniaActualizar">Imagen:</label>
                  <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar()' id="imagenArtesaniaActualizar" ref="fileActualizar" type="file" style="display:none">
                </div>
                <div class="form-group">
                  <label for="telefonoArtesaniaActualizar">Teléfono:</label>
                  <input type="text" v-model="artesaniaActualizada.phone" class="form-control" id="telefonoArtesaniaActualizar" required>
                </div>
                <div class="form-group">
                  <label for="descripcionArtesaniaActualizar">Descripción:</label>
                  <textarea v-model="artesaniaActualizada.description" class="form-control" id="descripcionArtesaniaActualizar" required></textarea>
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
       artesanias: [],
       artesaniaActualizada: {
         _id: null,
         name: '',
         address: '',
         image: '',
         phone: '',
         description: '',
       },
       file: null,
       nuevaArtesania: {
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
     this.fetchArtesanias();
   },
   methods: {
     async fetchArtesanias() {
       try {
         const response = await axios.get(`${process.env.API}/Artesania`);
         this.artesanias = response.data;
       } catch (error) {
         console.error('Error fetching artesanias:', error);
       }
     },
     async eliminarArtesania(id) {
       try {
         await axios.delete(`${process.env.API}/Artesania/${id}`);
         this.fetchArtesanias();
       } catch (error) {
         console.error('Error deleting artesania:', error);
       }
     },
     abrirModalActualizar(artesania) {
       this.artesaniaActualizada = { ...artesania };
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
     async crearArtesania() {
      try {
     var formData = new FormData(); 
      this.nuevaArtesania.customer_id = localStorage.getItem('customerId');
      
       
     if(this.nuevaArtesania.customer_id != '') {
           formData.append('image', this.file)
           formData.append('name', this.nuevaArtesania.name)
           formData.append('address', this.nuevaArtesania.address)
           formData.append('phone', this.nuevaArtesania.phone)
            formData.append('description', this.nuevaArtesania.description)
           formData.append('customer_id', this.nuevaArtesania.customer_id)
          
         
  
         
        const response = await axios.post(`${process.env.API}/Artesania`, formData);
         console.log(response);
 
          
         this.nuevaArtesania = {
           name: '',
           address: '',
           image: '',
           phone: '',
           description: '',
           customer_id: '',
         }; 
        
       
         this.cerrarModalCrear();
         this.fetchArtesanias(); 
        } else {
          alert('No tienes permisos para crear una artesanía')
        }
       } catch (error) {
         console.error('Error creating Artesania:', error);
       } 
     },
     async actualizarArtesania() {
       try {
         const response = await axios.put(`${process.env.API}/Artesania/${this.artesaniaActualizada._id}`, this.artesaniaActualizada);
         console.log(response);
         this.fetchArtesanias();
         // Close the modal
         this.cerrarModalActualizar();
       } catch (error) {
         console.error('Error updating Artesania:', error);
       }
     },
     uploadFile(event) {
      this.file = event.target.files[0];
    },
    

    
   },
 };
 </script>
 