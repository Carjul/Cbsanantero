<template>
    <div>
    <hr>
      <div class="row gallery">
        <div
          class="col-md-4"
          v-for="(artesania, index) in artesanias"
          :key="artesania.id"
        >
          <img
            :src="artesania.image"
            class="img-fluid gallery-item"
            alt="Artesanía"
            @click="openImage(artesania.image, index)"
          />
        </div>
  
        <!-- Modal para la imagen agrandada con slider -->
        <div v-if="showModal" class="modal-container">
          <span class="close-modal" @click="closeImage">&times;</span>
          <div class="slider-container">
            <div class="slider-arrow" @click="prevImage">&#8249;</div>
            <img :src="selectedImage" class="modal-img" alt="Imagen agrandada" />
            <div class="slider-arrow" @click="nextImage">&#8250;</div>
          </div>
          <p class="image-number"> {{ currentIndex  }} {{ artesanias.length }}</p>
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
        showModal: false,
        selectedImage: "",
        currentIndex: 0,
      };
    },
    methods: {
      async fetchArtesanias() {
        try {
          const response = await axios.get('http://localhost:3000/Artesania');
          this.artesanias = response.data;
        } catch (error) {
          console.error('Error al obtener las artesanías', error);
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
      this.fetchArtesanias();
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
  height: 200px; /* Set your desired height here */
  object-fit: cover; /* Preserve aspect ratio while covering the specified height */
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
    max-width: 90%;
    max-height: 90%;
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
  color:transparent;
  text-align: center;
}
  </style>
  