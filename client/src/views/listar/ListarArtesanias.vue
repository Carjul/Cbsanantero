<template>
  <div id="app">
    <main>
      <nav class="main-menu">
        <strong>PANEL DE ADMINISTRACION</strong>
        <ul>  
        </ul>
      </nav>

      <section class="content">
        <!-- Panel de crear artesanía -->
        <div class="left-content">
          <div class="card">
            <div class="card-body">
              <br>
              <br>
              <h5 class="card-title">Creación de Artesanía</h5>
              <p class="card-text">En este espacio creamos nuevas artesanías para COSTA BRISA</p>
              <button class="btn btn-success mb-4" @click="abrirModalCrear">Crear</button>
            </div>
          </div>
          
          <!-- Modal Crear -->
           <!-- Modal Crear -->
    <div ref="modalCrear" class="modal fade" id="modalCrear" tabindex="" role="dialog" aria-labelledby="modalCrearLabel" aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="modalCrearLabel">Crear Artesanía</h5>
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
              <button type="submit" class="btn btn-primary">Crear Artesanía</button>
            </form>
          </div>
        </div>
      </div>
    </div>
        </div>

        <!-- Todas las artesanías -->
        <div class="right-content">
          <strong>ARTESANÍA ACTUAL</strong>
          <div class="row">
            <div v-for="artesania in artesanias" :key="artesania._id" class="col-md-12 mb-6">
              <div class="card">
                <img :src="artesania.image" class="card-img-top" style="width:100%; height:200px;" alt="Artesanía Image">
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

        <!-- Modal Actualizar -->
        <div ref="modalActualizar" class="modal fade" id="modalActualizar" tabindex="-1" role="dialog" aria-labelledby="modalActualizarLabel" aria-hidden="true">
          <div class="modal-dialog modal-dialog-centered" role="document">
            <div class="modal-content">
              <div class="modal-header">
                <h5 class="modal-title" id="modalActualizarLabel">Actualizar Artesanía</h5>
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
                    <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar()' id="imagenArtesaniaActualizar" ref="fileActualizar" type="file">
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
      </section>
    </main>
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
    },
    uploadFile(event) {
      this.file = event.target.files[0];
    },
    uploadFileActualizar(event) {
      this.fileActualizar = event.target.files[0];
    },
    
 };
 </script>
 
<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Nunito:wght@200;300;400;500;600;700;800;900;1000&family=Roboto:wght@300;400;500;700&display=swap");

*,
*::before,
*::after {
  box-sizing: border-box;
  padding: 0;
  margin: 0;
}

nav {
  user-select: none;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  -o-user-select: none;
}

nav ul,
nav ul li {
  outline: 0;
}

nav ul li a {
  text-decoration: none;
}

/* MAIN MENU */

main-menu
.nav-item a {
  position: relative;
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  color: #fff;
  font-size: 1rem;
  padding: 15px 0;
  margin-left: 10px;
  border-top-left-radius: 20px;
  border-bottom-left-radius: 20px;
}

.nav-item b:nth-child(1),
.nav-item b:nth-child(2) {
  position: absolute;
  height: 15px;
  width: 100%;
  background: #fff;
  display: none;
}

.nav-item b:nth-child(1)::before,
.nav-item b:nth-child(2)::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: #68bee7;
}

.nav-item.active b:nth-child(1),
.nav-item.active b:nth-child(2) {
  display: block;
}

.nav-item.active a {
  text-decoration: none;
  color: #000;
  background: rgb(254, 254, 254);
}

/* CONTENT */

.content {
  display: grid;
  grid-template-columns: 75% 25%;
}

/* LEFT CONTENT */

.left-content {
  display: grid;
  grid-template-rows: 50% 50%;
  background: #f6f7fb;
  margin: 15px;
  padding: 20px;
  border-radius: 15px;
}

/* ACTIVITIES */

.activities h1 {
  margin: 0 0 20px;
  font-size: 1.4rem;
  font-weight: 700;
}

.activity-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  grid-auto-rows: 150px;
  gap: 10px;
}

.image-container {
  position: relative;
  margin-bottom: 10px;
  box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 3px;
  border-radius: 10px;
}

.image-container img {
  width: 100%;
  height: 100%;
  border-radius: 10px;
  object-fit: cover;
}

.overlay {
  position: absolute;
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  justify-content: flex-end;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    180deg,
    transparent,
    transparent,
    rgba(3, 3, 185, 0.5)
  );
  border-radius: 10px;
  transition: all 0.6s linear;
}

.image-container:hover .overlay {
  opacity: 0;
}

.overlay h3 {
  margin-bottom: 8px;
  margin-right: 10px;
  color: #fff;
}

/* LEFT BOTTOM */

.left-bottom {
  display: grid;
  grid-template-columns: 55% 40%;
  gap: 40px;
}

/* WEEKLY SCHEDULE */

.weekly-schedule {
  display: flex;
  flex-direction: column;
}

.weekly-schedule h1 {
  margin-top: 20px;
  font-size: 1.3rem;
  font-weight: 700;
}

.calendar {
  margin-top: 10px;
}

.day-and-activity {
  display: grid;
  grid-template-columns: 15% 60% 25%;
  align-items: center;
  border-radius: 14px;
  margin-bottom: 5px;
  color: #484d53;
  box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 3px;
}

.day {
  display: flex;
  flex-direction: column;
  align-items: center;
  transform: translateY(-10px);
}

.day h1 {
  font-size: 1.6rem;
  font-weight: 600;
}

.day p {
  text-transform: uppercase;
  font-size: 0.9rem;
  font-weight: 600;
  transform: translateY(-3px);
}

.activity {
  border-left: 3px solid #484d53;
}

.participants {
  display: flex;
  margin-left: 20px;
}

.participants img {
  width: 28px;
  height: 28px;
  border-radius: 50%;
  z-index: calc(2 * var(--i));
  margin-left: -10px;
  box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 3px;
}

.activity h2 {
  margin-left: 10px;
  font-size: 1.25rem;
  font-weight: 600;
  border-radius: 12px;
}

.btn {
  display: block;
  padding: 8px 24px;
  margin: 10px auto;
  border-radius: 25px;
  color: #fff;
  background: #68bee7;
  font-size: 1rem;
  font-weight: 600;
  text-transform: uppercase;
  transition: all 0.2s ease-in-out;
}

.btn:hover {
  background: #59a6d5;
}

/* RIGHT CONTENT */

.right-content {
  background: #f6f7fb;
  margin: 15px;
  padding: 20px;
  border-radius: 15px;
}

/* CALENDAR */

.calendar-container {
  position: relative;
  width: 100%;
}

.calendar-img {
  width: 100%;
  border-radius: 15px;
  box-shadow: rgba(0, 0, 0, 0.16) 0px 1px 3px;
}

.calendar-date {
  position: absolute;
  top: 10px;
  left: 10px;
  padding: 5px 10px;
  border-radius: 10px;
  background: rgba(104, 190, 231, 0.6);
  color: #fff;
  font-weight: 600;
}

.right-content h1 {
  font-size: 1.3rem;
  font-weight: 700;
  margin-bottom: 15px;
}

.calendar-link {
  display: block;
  padding: 8px 24px;
  border-radius: 25px;
  color: #fff;
  background: #68bee7;
  font-size: 1rem;
  font-weight: 600;
  text-transform: uppercase;
  text-align: center;
  transition: all 0.2s ease-in-out;
}

.calendar-link:hover {
  background: #59a6d5;
}

/* MEDIA QUERIES */

@media only screen and (max-width: 1024px) {
  .main-menu h1 {
    font-size: 1.2rem;
  }
  .logo {
    display: block;
    margin: 0 auto;
  }
  .nav-item b:nth-child(1),
  .nav-item b:nth-child(2) {
    display: block;
  }
}

@media only screen and (max-width: 768px) {
  .main {
    grid-template-columns: 100%;
  }
  .left-content,
  .right-content {
    margin: 15px 0;
  }
  .content {
    display: block;
  }
  .left-bottom {
    grid-template-columns: 100%;
  }
  .weekly-schedule h1 {
    margin-top: 15px;
  }
  .calendar-date {
    font-size: 0.9rem;
  }
  .calendar-link {
    font-size: 0.9rem;
  }
}

@media only screen and (max-width: 480px) {
  .main-menu {
    padding-top: 0;
  }
  .main-menu h1 {
    font-size: 1rem;
    margin: 10px 0 20px;
  }
  .nav-item a {
    padding: 12px 0;
    margin: 0;
  }
  .main-menu h1 {
    margin: 10px 0 20px;
  }
  .content {
    grid-template-columns: 100%;
  }
  .activities h1 {
    font-size: 1.2rem;
  }
  .activity-container {
    grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
  }
  .overlay h3 {
    font-size: 0.9rem;
    margin-bottom: 6px;
  }
  .left-bottom {
    grid-template-columns: 100%;
  }
  .weekly-schedule h1 {
    font-size: 1.1rem;
    margin-top: 15px;
  }
  .calendar-date {
    font-size: 0.8rem;
    padding: 4px 8px;
  }
  .calendar-link {
    font-size: 0.8rem;
    padding: 6px 16px;
  }
}
</style>
