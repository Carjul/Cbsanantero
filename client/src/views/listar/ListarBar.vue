<template>
  <div id="app">
    <main>
      <nav class="main-menu">
        <strong>PANEL DE ADMINISTRACION</strong>
        <ul>  
        </ul>
      </nav>

      <section class="content">
        <!-- Panel de crear bar -->
        <div class="left-content">
          <div class="card">
            <div class="card-body">
              <br>
              <br>
              <h5 class="card-title">Creación de Bar</h5>
              <p class="card-text">En este espacio creamos nuevos bares para COSTA BRISA</p>
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
                      <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFile' id="imagenBar" ref="fileCrear" type="file">
                    </div>
                    <div class="form-group">
                      <label for="telefonoBar">Teléfono:</label>
                      <input type="text" v-model="nuevoBar.phone" class="form-control" id="telefonoBar" required>
                    </div>

                    <button type="submit" class="btn btn-primary">Crear Bar</button>
                  </form>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Todos los bares -->
        <div class="right-content">
          <strong>BARES ACTUALES</strong>
          <div class="row">
            <div v-for="bar in bares" :key="bar._id" class="col-md-12 mb-6">
              <div class="card">
                <img :src="bar.image" class="card-img-top" style="width:100%; height:200px;" alt="Bar Image">
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

        <!-- Modal Actualizar -->
        <div ref="modalActualizar" class="modal fade" id="modalActualizar" tabindex="-1" role="dialog" aria-labelledby="modalActualizarLabel" aria-hidden="true">
          <div class="modal-dialog modal-dialog-centered" role="document">
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
                    <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar' id="imagenBarActualizar" ref="fileActualizar" type="file">
                  </div>
                  <div class="form-group">
                    <label for="telefonoBarActualizar">Teléfono:</label>
                    <input type="text" v-model="barActualizado.phone" class="form-control" id="telefonoBarActualizar" required>
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
