<template>
  <div>
    <hr>
    <div class="row gallery">

      <div class="col-md-4">
        <form @submit.prevent="crearGal">
          <div class="form-group">
            <label for="imagenBar">Imagen:</label>
            <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFile' multiple
              id="imagenBar" ref="file" type="file" style="display:none">
          </div>
          <input type="submit" value="enviar">
        </form>
      </div>
      <div class="row">

        <div class="card mb-4">
          <div class="" v-for="(obj, imgIndex) in Galeria.photos" :key="imgIndex">
            <buttom @click="EliminarPhoto(Galeria.photos[imgIndex])">X</buttom>
            <img :src="obj.image" class="card-img-top" alt="imagen" loading="lazy" />
          </div>
        </div>

      </div>



      <!-- Modal para la imagen agrandada con slider -->
      <div v-if="showModal" class="modal-container">
        <span class="close-modal" @click="closeImage">&times;</span>
        <div class="slider-container">
          <div class="slider-arrow" @click="prevImage">&#8249;</div>
          <img :src="selectedImage" class="modal-img" alt="Imagen agrandada" />
          <div class="slider-arrow" @click="nextImage">&#8250;</div>
        </div>
        <p class="image-number"> {{ currentIndex }} {{ artesanias.length }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      file: null,
      Galeria: {
        photos: [],
        negocio_id: null,
        customer_id: null
      },
      showModal: false,
      selectedImage: "",
      currentIndex: 0,
    };
  },
  methods: {
    uploadFile(event) {
      const arr = [];
      this.file = event.target.files;
      for (let key in this.file) {
        arr[key] = this.file[key]
      }
      this.file = arr;

    },
    async fetchGaleria() {
      try {
        const response = await axios.get(`http://localhost:3000/galeria/${this.Galeria.negocio_id}`);
        let data = response.data;

        if (data.length !== 0) {
          for (let i = 0; i < data.length; i++) {
            const element = data[i];
            for (let j = 0; j < element.photos.length; j++) {
              const photo = element.photos[j];
              const objetoEncontrado = this.Galeria.photos.find(obj => obj.image === photo);

              if (!objetoEncontrado) {
                this.Galeria.photos.unshift({ id: element._id, image: photo });
              }

            }
            this.Galeria.negocio_id = element.negocio_id;
            this.Galeria.customer_id = element.customer_id;
          }
        }



      } catch (error) {
        console.error('Error al obtener las artesanías', error);
      }
    },
    async crearGal() {
      try {
        var formData = new FormData();

        if (this.Galeria.customer_id != '') {

          for (let i = 0; i < this.file.length; i++) {
            const element = this.file[i];

            let cadena = i.toString();
            let name = 'image' + cadena;

            formData.append(name, element)
          }
          formData.append('logitud', this.file.length)
          formData.append('negocio_id', this.Galeria.negocio_id)
          formData.append('customer_id', this.Galeria.customer_id)

          const response = await axios.post('http://localhost:3000/galeria', formData);
          console.log(response);
          this.file = null;

          this.fetchGaleria()

        } else {
          alert('No tienes permisos para crear una artesanía')
        }
      } catch (error) {
        console.error('Error creating Artesania:', error);
      }
    },
    async EliminarPhoto(photo) {
      const respuesta = window.confirm("¿Quieres Eliminar esta imagen?");

      if (respuesta) {
       
      try {
        const response = await axios.put(`http://localhost:3000/galeria/${photo.id}`, { photo: photo.image });
        console.log(response)
        this.Galeria.photos = this.Galeria.photos.filter(obj => obj.image !== photo.image);
        this.fetchGaleria()
      } catch (error) {
        console.log(error)
      }
      } else {
        console.log("solicitud cancelad");
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
  mounted() {
    this.Galeria.customer_id = localStorage.getItem('customerId');
    this.Galeria.negocio_id = localStorage.getItem('negocioId');
    localStorage.removeItem('idtemp')
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
  background: rgba(0, 0, 0, 0.8);
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
</style>