
<template>
  <div class="container">
    <div class="row">
      <div
        class="col-md-4"
        v-for="bar in bares"
        :key="bar.id"
      >
        <div class="card mb-4">
          <img
            :src="bar.image"
            class="card-img-top"
            alt="Bar"
          />
          <div class="card-body">
            <h5 class="card-title">{{ bar.name }}</h5>
            <p class="card-text">{{ bar.description }}</p>
            <p class="card-text"><strong>Dirección:</strong> {{ bar.address }}</p>
          
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
      bares: [],
    };
  },
  mounted() {
    this.fetchBares();
  },
  methods: {
    async fetchBares() {
      try {
        const response = await axios.get('http://localhost:3000/Bar');
        this.bares = response.data;
      } catch (error) {
        console.error('Error al obtener los bares', error);
      }
    },
  },
};
</script>


