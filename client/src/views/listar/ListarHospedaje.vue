<template>
  <div id="app">
    <main>
      <nav class="main-menu">
        <strong>PANEL DE ADMINISTRACION</strong>
        <ul>  
        </ul>
      </nav>

      <section class="content">
        <!-- Panel de crear hospedaje -->
        <div class="left-content">
          <div class="card">
            <div class="card-body">
              <br>
              <br>
              <h5 class="card-title">Creación de Hospedaje</h5>
              <p class="card-text">En este espacio creamos nuevos hospedajes para COSTA BRISA</p>
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
                <h5 class="modal-title" id="modalCrearLabel">Crear Hospedaje</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalCrear">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <div class="modal-body">
                <form @submit.prevent="crearHospedaje" class="row">
                    <div class="form-group col-md-6">
                        <label for="nombreHospedaje">Nombre:</label>
                        <input type="text" v-model="nuevoHospedaje.name" class="form-control" id="nombreHospedaje" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="descriptionHospedaje">Descripcion:</label>
                        <input type="text" v-model="nuevoHospedaje.description" class="form-control" id="descriptionHospedaje" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="direccionHospedaje">Dirección:</label>
                        <input type="text" v-model="nuevoHospedaje.address" class="form-control" id="direccionHospedaje" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="imagenHospedaje">Imagen:</label>
                        <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFile' id="imagenHospedaje" ref="fileCrear" type="file">
                    </div>
                    <div class="form-group col-md-6">
                        <label for="telefonoHospedaje">Teléfono:</label>
                        <input type="text" v-model="nuevoHospedaje.phone" class="form-control" id="telefonoHospedaje" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="precioHospedaje">Precio:</label>
                        <input type="text" v-model="nuevoHospedaje.price" class="form-control" id="precioHospedaje" required>
                    </div>
                    <div class="form-group col-md-12">
                        <button type="submit" class="btn btn-primary">Crear Hospedaje</button>
                    </div>
                </form>
            </div>
        </div>
    </div>
</div>

        </div>

        <!-- Todos los hospedajes -->
        <div class="right-content">
          <strong>HOSPEDAJE ACTUAL</strong>
          <div class="row">
            <div v-for="hospedaje in hospedajes" :key="hospedaje._id" class="col-md-12 mb-6">
              <div class="card">
                <img :src="hospedaje.image" class="card-img-top" style="width:100%; height:200px;" alt="Hospedaje Image">
                <div class="card-body">
                  <h5 class="card-title">{{ hospedaje.name }}</h5>
                  <p class="card-text">Dirección: {{ hospedaje.address }}</p>
                  <p class="card-text">Descripción: {{ hospedaje.description }}</p>
                  <p>{{hospedaje.price}}</p>
                  <p class="card-text">Teléfono: {{ hospedaje.phone }}</p>

                  <button class="btn btn-danger" @click="eliminarHospedaje(hospedaje._id)">Eliminar</button>
                  <button class="btn btn-primary" @click="abrirModalActualizar(hospedaje)">Actualizar</button>
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
                <h5 class="modal-title" id="modalActualizarLabel">Actualizar Hospedaje</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalActualizar">
                    <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <div class="modal-body">
                <form @submit.prevent="actualizarHospedaje" class="row">
                    <div class="form-group col-md-6">
                        <label for="nombreHospedajeActualizar">Nombre:</label>
                        <input type="text" v-model="hospedajeActualizado.name" class="form-control" id="nombreHospedajeActualizar" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="descriptionHospedajeActualizar">Descripción:</label>
                        <input type="text" v-model="hospedajeActualizado.description" class="form-control" id="descriptionHospedajeActualizar" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="direccionHospedajeActualizar">Dirección:</label>
                        <input type="text" v-model="hospedajeActualizado.address" class="form-control" id="direccionHospedajeActualizar" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="imagenHospedajeActualizar">Imagen:</label>
                        <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar' id="imagenHospedajeActualizar" ref="fileActualizar" type="file">
                    </div>
                    <div class="form-group col-md-6">
                        <label for="telefonoHospedajeActualizar">Teléfono:</label>
                        <input type="text" v-model="hospedajeActualizado.phone" class="form-control" id="telefonoHospedajeActualizar" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="precioHospedajeActualizar">Precio:</label>
                        <input type="text" v-model="hospedajeActualizado.price" class="form-control" id="precioHospedajeActualizar" required>
                    </div>
                    <div class="form-group col-md-12">
                        <button type="submit" class="btn btn-primary">Guardar Cambios</button>
                    </div>
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
      hospedajes: [],
      hospedajeActualizado: {
        _id: null,
        name: '',
        description: '',
        address: '',
        image: '',
        phone: '',
        price:'',
    
      },
      file: null,
      fileActualizar: null,
      nuevoHospedaje: {
        name: '',
        description:'',
        address: '',
        image: '',
        phone: '',
        price:'',
        customer_id: '',
      },
    };
  },
 async mounted() {
    await this.fetchHospedajes();
   
  },
  methods: {
    async fetchHospedajes() {
      try {
        const response = await axios.get(`${process.env.API}/Hospedaje`);
        let cid =  localStorage.getItem('customerId');
        response.data?  this.hospedajes = response.data.filter((e)=> e.customer_id === cid):null
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
          formData.append('description',this.nuevoHospedaje.description);
          formData.append('address', this.nuevoHospedaje.address);
          formData.append('phone', this.nuevoHospedaje.phone);
          formData.append('price', this.nuevoHospedaje.price);
          formData.append('customer_id', this.nuevoHospedaje.customer_id);
          const response = await axios.post(`${process.env.API}/Hospedaje`, formData);
        
        console.log(response);

        // Reset the form and close the modal
        this.nuevoHospedaje = {
          name: '',
          description:'',
          address: '',
          image: '',
          phone: '',
          price: '',
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
        forData.append('description', this.hospedajeActualizado.description);
        formData.append('address', this.hospedajeActualizado.address);
        formData.append('phone', this.hospedajeActualizado.phone);
        formData.append('price',this.hospedajeActualizado.price);
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
          console.log(this.actualizarHospedaje)
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
    uploadFile(event) {
      this.file = event.target.files[0];
    },
    uploadFileActualizar(event) {
      this.fileActualizar = event.target.files[0];
    },
  
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
