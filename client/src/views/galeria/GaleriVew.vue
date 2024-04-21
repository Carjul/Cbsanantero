<template>
  <div>
    <hr>
    <div class="row gallery">
      <div class="col-md-11">

        <center>
          <div class="col-md-4" v-if="customerRol === 'Vendedor' && customerID === GaleriaEnv.customer_id">
            <form @submit.prevent="crearGal">
              <div class="input-group mb-3">

                <input type="file" class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFile'
                  multiple id="imagenBar" ref="file">
                <p style="color: transparent;">+</p>
                <input type="submit" value="enviar">

              </div>


            </form>
            <progress v-if="uploadProgress !== null" :value="uploadProgress" max="100"></progress>
          </div>
        </center>



        <hr>


        <div class="row">
          <div class="col-md-4" v-for="(obj, imgIndex) in Galeria.photos" :key="imgIndex">
            <div class="card mb-2" @click="openModal(obj.image)">
              <button  v-if="customerRol === 'Vendedor' && customerID === GaleriaEnv.customer_id" style="height: 50px; width: 50px;" type="button" class="btn btn-danger"
                @click.stop="EliminarPhoto(Galeria.photos[imgIndex])">X</button>
              <img :src="obj.image" class="card-img-top gallery-image" alt="imagen" loading="lazy" />
            </div>
          </div>
        </div>
      </div>

      <div v-if="modalOpen" class="modal-container">
      <span class="close-modal" @click="closeModal">&times;</span>
      <img :src="selectedImage" class="modal-img" alt="Imagen agrandada" />

      <!-- Botones de navegación -->
      <input type="radio" id="prev" name="nav" class="modal-nav-input" @click="preImage" />
      <label for="prev" class="modal-nav-btn prev">&#10094;</label>
      <input type="radio" id="next" name="nav" class="modal-nav-input" @click="nexImage" />
      <label for="next" class="modal-nav-btn next">&#10095;</label>
    </div>
    
    </div>
  </div>





</template>


<script>
import Swal from 'sweetalert2';
import axios from 'axios';

export default {
  data() {
    return {
      file: null,
      customerRol: '',
      customerID: '',
      Galeria: {
        photos: [],
        negocio_id: null,
        status: '',
        customer_id: null
      },
      GaleriaEnv: {
        photos: [],
        negocio_id: null,
        customer_id: null
      },
      modalOpen: false,
      selectedImage: null,
      uploadProgress: null,
      numSelectedImages: 0,
    };
  },
  methods: {
    uploadFile(event) {
      this.file = event.target.files;
      for (let key in this.file) {
        this.GaleriaEnv.photos[key] = this.file[key];
      }
      this.file = null;
    },
    openModal(image) {
      this.selectedImage = image;
      this.modalOpen = true;
    },
    closeModal() {
      this.modalOpen = false;
    },
    preImage() {
  const currentIndex = this.Galeria.photos.findIndex(photo => photo.image === this.selectedImage);
  const prevIndex = (currentIndex - 1 + this.Galeria.photos.length) % this.Galeria.photos.length;
  this.selectedImage = this.Galeria.photos[prevIndex].image;
},
nexImage() {
  const currentIndex = this.Galeria.photos.findIndex(photo => photo.image === this.selectedImage);
  const nextIndex = (currentIndex + 1) % this.Galeria.photos.length;
  this.selectedImage = this.Galeria.photos[nextIndex].image;
},

    async fetchGaleria() {
      try {
        const response = await axios.get(`${process.env.API}/galeria/${this.GaleriaEnv.negocio_id}`);
        let data = response.data;
        this.Galeria = {
          photos: data.photos,
          negocio_id: data.negocio_id,
          status: data.status,
          customer_id: data.customer_id
        };
      } catch (error) {
        console.error('Error al obtener las artesanías', error);
      }
    },
    async crearGal() {
      try {
        if (this.GaleriaEnv.customer_id !== null && this.GaleriaEnv.negocio_id !== null && this.GaleriaEnv.photos.length > 0) {
          const formData = new FormData();
          for (let i = 0; i < this.GaleriaEnv.photos.length; i++) { formData.append(`image${i}`, this.GaleriaEnv.photos[i]); }
          formData.append('logitud', this.GaleriaEnv.photos.length);
          formData.append('negocio_id', this.GaleriaEnv.negocio_id);
          formData.append('customer_id', this.GaleriaEnv.customer_id);
          const config = {
            onUploadProgress: (progressEvent) => {
              this.uploadProgress = Math.round((progressEvent.loaded * 100) / progressEvent.total);
            },
          };
          const response = await axios.post(`${process.env.API}/galeria`, formData, config);
          console.log(response);
          if (response.status === 202) {
            this.GaleriaEnv.photos = [];
            this.fetchGaleria();
          }
        } else {
          alert('No tienes las credenciales para subir fotos');
          console.log(this.GaleriaEnv);
        }
      } catch (error) {
        console.error('Error creating Galeria:', error);
      } finally {
        // Restaurar la variable de progreso después de completar la solicitud
        this.uploadProgress = null;
      }
    },
    async EliminarPhoto(photo) {
      const result = await Swal.fire({
        title: "¿Quieres eliminar esta imagen?",
        text: "No podrás revertir esto",
        icon: "warning",
        showCancelButton: true,
        confirmButtonColor: "#3085d6",
        cancelButtonColor: "#d33",
        confirmButtonText: "Sí, eliminar",
      });
      if (result.isConfirmed) {
        try {
          const response = await axios.put(`${process.env.API}/galeria/${photo.id}`, { photo: photo.image });
          console.log(response);
          this.fetchGaleria();
          Swal.fire({
            title: "Eliminada",
            text: "La imagen ha sido eliminada exitosamente.",
            icon: "success"
          });
        } catch (error) {
          console.log(error);
        }
      } else {
        console.log("Solicitud cancelada");
      }
    },
    openImage(image, index) {
      this.showModal = true;
      this.selectedImage = image;
      this.currentIndex = index;
    },
    closeImage() {
      this.showModal = false;
      this.selectedImage = "";
      this.currentIndex = 0;
    },
    prevImage() {
      this.currentIndex = (this.currentIndex - 1 + this.artesanias.length) % this.artesanias.length;
      this.selectedImage = this.artesanias[this.currentIndex].image;
    },
    nextImage() {
      this.currentIndex = (this.currentIndex + 1) % this.artesanias.length;
      this.selectedImage = this.artesanias[this.currentIndex].image;
    },
  },
  created() {
    this.GaleriaEnv.customer_id = localStorage.getItem('customerId');
    this.GaleriaEnv.negocio_id = localStorage.getItem('negocioId');
    this.customerRol = localStorage.getItem('usuarioRol');
    this.customerID = localStorage.getItem('customerNegocioId');
  },
  mounted() {
    this.fetchGaleria()
  },
};

</script>


<style scoped>
.gallery {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
}

.gallery-item {
  width: 100%;
  height: 200px;
  /* Set your desired height here */
  object-fit: cover;
  /* Preserve aspect ratio while covering the specified height */
  margin-bottom: 20px;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
}


.gallery-item:hover {
  cursor: pointer;
  transform: scale(1.05);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
}

.modal-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.945);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-img {
  max-width: 50%;
  max-height: 50%;
}

.close-modal {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 30px;
  color: #fff;
  cursor: pointer;
}

.slider-container {
  display: flex;
  align-items: center;
}

.slider-arrow {
  font-size: 24px;
  color: #fff;
  margin: 0 10px;
  cursor: pointer;
}

.image-number {
  color: transparent;
  text-align: center;
}



/* ajustes de las impreciones como card 
 */
.gallery-image {
  width: 100%;
  height: auto;
}

.card {
  margin-bottom: 20px;
}

/* Agregué algunos estilos para el modal, ajusta según tus necesidades */
.gallery-image {
  width: 100%;
  height: auto;
  border: 1px solid #ddd;
  /* Añade un borde a las imágenes */
  border-radius: 5px;
  /* Redondea las esquinas */
}

.card {
  margin-bottom: 20px;
  border: 1px solid #ddd;
  border-radius: 5px;
  box-shadow: 0 10px 10px rgba(0, 0, 254, 0.1);
  /* Sombreado azul */

  /* Animación de semi zoom al pasar el cursor */
  transition: transform 0.3s ease-in-out;
}

.card:hover {
  transform: scale(1.05);
  /* Semi zoom al pasar el cursor */
}

/* Agregué algunos estilos para el modal, ajusta según tus necesidades */
.gallery-image {
  width: 100%;
  height: 300px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

.card {
  margin-bottom: 20px;
  border: 1px solid #ddd;
  border-radius: 5px;
  box-shadow: 0 10px 10px rgba(0, 0, 254, 0.1);
  /* Sombreado azul */
}



/*  modal que abre imagen  */


.modal-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  justify-content: center;
  align-items: center;
}

.close-modal {
  position: absolute;
  top: 10px;
  right: 10px;
  cursor: pointer;
  color: white;
  font-size: 24px;
}

.modal-img {
  max-width: 80%;
  max-height: 80%;
  border: 1px solid #ddd;
  border-radius: 5px;
  box-shadow: 0 10px 10px rgba(0, 0, 254, 0.1);
  /* Sombreado azul para la imagen modal */
}




/* Estilos para los botones de navegación del modal */
.modal-nav-input {
  display: none;
}

.modal-nav-btn {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  background: rgba(0, 0, 0, 0.5);
  color: white;
  border: none;
  padding: 10px;
  cursor: pointer;
}

.modal-nav-btn.prev {
  left: 10px;
}

.modal-nav-btn.next {
  right: 10px;
}
</style>