<template>
  <div id="app">
    <main>
      <nav class="main-menu">
        <strong>PANEL DE ADMINISTRACION</strong>
        <ul>  
        </ul>
      </nav>

      <section class="content">
        <!-- Panel de crear negocio -->
        <div class="left-content">
          <div class="card">
            <div class="card-body">
              <br>
              <br>
              <h5 class="card-title">Creación de Hotel</h5>
              <p class="card-text">En este espacio creamos nuevos hoteles para COSTA BRISA</p>
              <button class="btn btn-success mb-4" @click="abrirModalCrear">Crear</button>
            </div>
          </div>
          
          <div class="card">
            <div class="card-body">
              <br>
              <br>
              <h5 class="card-title">¡Crea tu Galeria!</h5>
              <p class="card-text">En este espacio podrás publicar tu álbum de fotografías</p>
              <button class="btn btn-success mb-4">Crear Galería</button>
            </div>
          </div>

          <!-- Modal Crear -->
          <div ref="modalCrear" class="modal fade" id="modalCrear" tabindex="" role="dialog" aria-labelledby="modalCrearLabel" aria-hidden="true">
            <div class="modal-dialog modal-dialog-centered" role="document">
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
                      <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFile' id="imagenHotel" ref="fileCrear" type="file">
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
        </div>

        <!-- Todos los negocios -->
        <div class="right-content">
          <strong>NEGOCIO ACTUAL</strong>
          <div class="row">
            <div v-for="hotel in hoteles" :key="hotel._id" class="col-md-12 mb-6">
              <div class="card">
                <img :src="hotel.image" class="card-img-top" style="width:100%; height:200px;" alt="Hotel Image">
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

        <!-- Modal Actualizar -->
       <!-- Modal Actualizar -->
<div ref="modalActualizar" class="modal fade" id="modalActualizar" tabindex="-1" role="dialog" aria-labelledby="modalActualizarLabel" aria-hidden="true">
  <div class="modal-dialog modal-dialog-centered" role="document">
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
            <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar' id="imagenHotelActualizar" ref="fileActualizar" type="file">
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
      </section>
    </main>
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
    file:null,
    fileA:null,
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
      const response = await axios.get(`${process.env.API}/Hotel`);
      this.hoteles = response.data;
    } catch (error) {
      console.error('Error fetching hoteles:', error);
    }
  },
  async eliminarHotel(id) {
    try {
      await axios.delete(`${process.env.API}/Hotel/${id}`);
      this.fetchHoteles();
    } catch (error) {
      console.error('Error deleting hotel:', error);
    }
  },
  
  enviarIdGaleria(id, cid) {
localStorage.setItem('negocioId', id);
localStorage.setItem('dueñoId', cid);
this.$router.push({ path: '/artegaleria' });
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
      var formData = new FormData();
      formData.append('image', this.file);
      formData.append('name', this.nuevoHotel.name);
      formData.append('address', this.nuevoHotel.address);
      formData.append('phone', this.nuevoHotel.phone);
      formData.append('customer_id', this.nuevoHotel.customer_id);
      const response = await axios.post(`${process.env.API}/Hotel`, formData);
      
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
    }
    } catch (error) {
      console.error('Error creating Hotel:', error);
    }
  },
  async actualizarHotel() {
    if (this.fileA){
      var formData = new FormData();
      formData.append('name', this.hotelActualizado.name);
      formData.append('address', this.hotelActualizado.address);
      formData.append('phone', this.hotelActualizado.phone);
      formData.append('customer_id', this.hotelActualizado.customer_id);
      formData.append('imagen', this.fileA);
      formData.append('image', "");
      const response = await axios.put(`${process.env.API}/Hotel/${this.hotelActualizado._id}`, formData);
      console.log(response);
      // Si la solicitud se realizó con éxito, actualizamos la lista de hoteles
      this.fetchHoteles();
      this.cerrarModalActualizar();
      this.fileA = null;
    }else{

    
  try {
    const response = await axios.put(`${process.env.API}/Hotel/${this.hotelActualizado._id}`, this.hotelActualizado);
    console.log(response);
    // Si la solicitud se realizó con éxito, actualizamos la lista de hoteles
    this.fetchHoteles();
    // Cerramos el modal de actualización
    this.cerrarModalActualizar();
  } catch (error) {
    console.error('Error updating Hotel:', error);
    // Aquí puedes agregar cualquier lógica adicional para manejar el error, como mostrar un mensaje de error al usuario.
  }
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


.right-content h1 {
  font-size: 1.3rem;
  font-weight: 700;
  margin-bottom: 15px;
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
