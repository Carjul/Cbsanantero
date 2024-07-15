<template>
  <div id="app">
    <main>
      <nav class="main-menu">
        <strong>PANEL DE ADMINISTRACION</strong>
        <ul>  
        </ul>
      </nav>

      <section class="content">
        <!-- Panel de crear restaurante -->
        <div class="left-content">
          <div class="card">
            <div class="card-body">
              <br>
              <br>
              <h5 class="card-title">Creación de Restaurante</h5>
              <p class="card-text">En este espacio creamos nuevos restaurantes para COSTA BRISA</p>
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
                  <h5 class="modal-title" id="modalCrearLabel">Crear Restaurante</h5>
                  <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalCrear">
                    <span aria-hidden="true">&times;</span>
                  </button>
                </div>
                <div class="modal-body">
                  <form @submit.prevent="crearRestaurante" class="row">
                    <div class="form-group col-md-6">
                      <label for="nombreRestaurante">Nombre:</label>
                      <input type="text" v-model="nuevoRestaurante.name" class="form-control" id="nombreRestaurante" required>
                    </div>
                    <div class="form-group col-md-6">
                      <label for="descripcionRestaurante">Descripción:</label>
                      <input type="text" v-model="nuevoRestaurante.description" class="form-control" id="descripcionRestaurante" required>
                    </div>
                    <div class="form-group col-md-6">
                      <label for="direccionRestaurante">Dirección:</label>
                      <input type="text" v-model="nuevoRestaurante.address" class="form-control" id="direccionRestaurante" required>
                    </div>
                    <div class="form-group col-md-6">
                      <label for="imagenRestaurante">Imagen:</label>
                      <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFile' id="imagenRestaurante" ref="fileCrear" type="file">
                    </div>
                    <div class="form-group col-md-6">
                      <label for="telefonoRestaurante">Teléfono:</label>
                      <input type="text" v-model="nuevoRestaurante.phone" class="form-control" id="telefonoRestaurante" required>
                    </div>
                    <div class="form-group col-md-6">
                      <label for="precioRestaurante">Precio:</label>
                      <input type="text" v-model="nuevoRestaurante.price" class="form-control" id="precioRestaurante" required>
                    </div>
                    <div class="form-group col-md-12">
                      <button type="submit" class="btn btn-primary">Crear Restaurante</button>
                    </div>
                  </form>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Todos los restaurantes -->
        <div class="right-content">
          <strong>RESTAURANTES ACTUALES</strong>
          <div class="row">
            <div v-for="restaurante in restaurantes" :key="restaurante._id" class="col-md-12 mb-6">
              <div class="card">
                <img :src="restaurante.image" class="card-img-top" style="width:100%; height:200px;" alt="Restaurante Image">
                <div class="card-body">
                  <h5 class="card-title">{{ restaurante.name }}</h5>
                  <p class="card-text">Dirección: {{ restaurante.address }}</p>
                  <p class="card-text">Descripción: {{ restaurante.description }}</p>
                  <p>{{ restaurante.price }}</p>
                  <p class="card-text">Teléfono: {{ restaurante.phone }}</p>

                  <button class="btn btn-danger" @click="eliminarRestaurante(restaurante._id)">Eliminar</button>
                  <button class="btn btn-primary" @click="abrirModalActualizar(restaurante)">Actualizar</button>
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
                <h5 class="modal-title" id="modalActualizarLabel">Actualizar Restaurante</h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalActualizar">
                  <span aria-hidden="true">&times;</span>
                </button>
              </div>
              <div class="modal-body">
                <form @submit.prevent="actualizarRestaurante" class="row">
                  <div class="form-group col-md-6">
                    <label for="nombreRestauranteActualizar">Nombre:</label>
                    <input type="text" v-model="restauranteActualizado.name" class="form-control" id="nombreRestauranteActualizar" required>
                  </div>
                  <div class="form-group col-md-6">
                    <label for="descripcionRestauranteActualizar">Descripción:</label>
                    <input type="text" v-model="restauranteActualizado.description" class="form-control" id="descripcionRestauranteActualizar" required>
                  </div>
                  <div class="form-group col-md-6">
                    <label for="direccionRestauranteActualizar">Dirección:</label>
                    <input type="text" v-model="restauranteActualizado.address" class="form-control" id="direccionRestauranteActualizar" required>
                  </div>
                  <div class="form-group col-md-6">
                    <label for="imagenRestauranteActualizar">Imagen:</label>
                    <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFileActualizar' id="imagenRestauranteActualizar" ref="fileActualizar" type="file">
                  </div>
                  <div class="form-group col-md-6">
                    <label for="telefonoRestauranteActualizar">Teléfono:</label>
                    <input type="text" v-model="restauranteActualizado.phone" class="form-control" id="telefonoRestauranteActualizar" required>
                  </div>
                  <div class="form-group col-md-6">
                    <label for="precioRestauranteActualizar">Precio:</label>
                    <input type="text" v-model="restauranteActualizado.price" class="form-control" id="precioRestauranteActualizar" required>
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
      restaurantes: [],
      restauranteActualizado: {
        _id: null,
        name: '',
        description: '',
        address: '',
        image: '',
        phone: '',
        price: '',
      },
      file: null,
      fileActualizar: null,
      nuevoRestaurante: {
        name: '',
        description: '',
        address: '',
        image: '',
        phone: '',
        price: '',
        customer_id: '',
      },
    };
  },
  async mounted() {
    await this.fetchRestaurantes();
  },
  methods: {
    async fetchRestaurantes() {
      let id = localStorage.getItem('customerId');
      try {
        const response = await axios.get(`${process.env.API}/Restaurante`);
        response.data ? this.restaurantes = response.data.filter((e) => e.customer_id === id) : [];
      } catch (error) {
        console.error('Error fetching restaurantes:', error);
      }
    },
    async eliminarRestaurante(id) {
      try {
        await axios.delete(`${process.env.API}/Restaurante/${id}`);
        this.fetchRestaurantes();
      } catch (error) {
        console.error('Error deleting restaurante:', error);
      }
    },
    abrirModalActualizar(restaurante) {
      this.restauranteActualizado = { ...restaurante };
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
    async crearRestaurante() {
      let id = localStorage.getItem('customerId');
      try {
        if (id) {
          this.nuevoRestaurante.customer_id = id;
          const formData = new FormData();
          formData.append('image', this.file);
          formData.append('name', this.nuevoRestaurante.name);
          formData.append('description', this.nuevoRestaurante.description);
          formData.append('address', this.nuevoRestaurante.address);
          formData.append('phone', this.nuevoRestaurante.phone);
          formData.append('price', this.nuevoRestaurante.price);
          formData.append('customer_id', this.nuevoRestaurante.customer_id);
          const response = await axios.post(`${process.env.API}/Restaurante`, formData);
        
          console.log(response);

          // Reset the form and close the modal
          this.nuevoRestaurante = {
            name: '',
            description: '',
            address: '',
            image: '',
            phone: '',
            price: '',
            customer_id: '',
          };

          // Close the modal
          this.cerrarModalCrear();
          this.fetchRestaurantes();
        }
      } catch (error) {
        console.error('Error creating Restaurante:', error);
      }
    },
    async actualizarRestaurante() {
      if (this.fileActualizar) {
        const formData = new FormData();
        formData.append('name', this.restauranteActualizado.name);
        formData.append('description', this.restauranteActualizado.description);
        formData.append('address', this.restauranteActualizado.address);
        formData.append('phone', this.restauranteActualizado.phone);
        formData.append('price', this.restauranteActualizado.price);
        formData.append('customer_id', this.restauranteActualizado.customer_id);
        formData.append('imagen', this.fileActualizar);
        formData.append('image', "");
        try {
          const response = await axios.put(`${process.env.API}/Restaurante/${this.restauranteActualizado._id}`, formData);
          console.log(response);
          this.fetchRestaurantes();
          this.cerrarModalActualizar();
          this.fileActualizar = null;
          
        } catch (error) {
          console.error('Error updating Restaurante:', error);
        }
      } else {
        try {
          const response = await axios.put(`${process.env.API}/Restaurante/${this.restauranteActualizado._id}`, this.restauranteActualizado);
          console.log(response);
          this.fetchRestaurantes();
          // Close the modal
          this.cerrarModalActualizar();
        } catch (error) {
          console.error('Error updating Restaurante:', error);
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
