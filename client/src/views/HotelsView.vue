<template>
  <div class="container">
    <hr>
    <div class="row">
      <div
        class="col-md-4"
        v-for="hotel in hoteles"
        :key="hotel.id"
      >
        <div class="card mb-4">
          <img
            :src="hotel.image"
            class="card-img-top"
            alt="Hotel"
          />
          <div class="card-body">
            <h5 class="card-title">{{ hotel.name }}</h5>
            <p class="card-text">{{ hotel.address }}</p>
            <p class="card-text"><strong>Email:</strong> {{ hotel.email }}</p>
            <p class="card-text"><strong>Teléfono:</strong> {{ hotel.phone }}</p>
            <p class="card-text"><strong>Precio:</strong> {{ hotel.price }}</p>
         
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
      hoteles: [],
    };
  },
  mounted() {
    this.fetchHoteles();
  },
  methods: {
    async fetchHoteles() {
      try {
        const response = await axios.get('http://localhost:3000/Hotel');
        this.hoteles = response.data;
      } catch (error) {
        console.error('Error al obtener los hoteles', error);
      }
    },
  },
};
</script>

