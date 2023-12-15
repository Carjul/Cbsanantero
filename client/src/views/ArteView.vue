
<template>
    <div class="container">
      <hr>
      <div class="row">
        <div
          class="col-md-4"
          v-for="artesania in artesanias"
          :key="artesania.id"
        >
          <div class="card mb-4">
            <img
              :src="artesania.image"
              class="card-img-top"
              alt="Artesanía"
            />
            <div class="card-body">
              <h5 class="card-title">{{ artesania.name }}</h5>
              <p class="card-text">{{ artesania.description }}</p>
              <p class="card-text"><strong>Dirección:</strong> {{ artesania.address }}</p>
              <p class="card-text"><strong>Teléfono:</strong> {{ artesania.phone }}</p>
             
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <style scoped>
  .card {
    transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
  }
  
  .card:hover {
    transform: scale(1.05);
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  }
  </style>
  
 // En tu componente.vue
 <script>
 import axios from 'axios';
 
 export default {
   data() {
     return {
       artesanias: [],
     };
   },
   mounted() {
     this.fetchArtesanias();
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
   },
 };
 </script>
 

 