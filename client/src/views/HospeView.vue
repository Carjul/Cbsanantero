<template>
  <div class="container">
   <hr>
    <div class="row">
      <div
        class="col-md-4"
        v-for="hospedaje in hospedajes"
        :key="hospedaje.id"
      >
        <div class="card mb-4">
          <img
            :src="hospedaje.image"
            class="card-img-top"
            alt="Hospedaje"
          />
          <div class="card-body">
            <h5 class="card-title">{{ hospedaje.name }}</h5>
            <p class="card-text">{{ hospedaje.address }}</p>
            <p class="card-text"><strong>Teléfono:</strong> {{ hospedaje.phone }}</p>
            <p class="card-text"><strong>Precio:</strong> {{ hospedaje.price }}</p>
           
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


<script>
import axios from 'axios';

export default {
  data() {
    return {
      hospedajes: [],
    };
  },
  mounted() {
    this.fetchHospedajes();
  },
  methods: {
    async fetchHospedajes() {
      try {
        const response = await axios.get('http://localhost:3000/Hospedaje');
        this.hospedajes = response.data;
      } catch (error) {
        console.error('Error al obtener los hospedajes', error);
      }
    },
  },
};
</script>

