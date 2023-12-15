<template>
  <div class="container">
    <hr>
    <div class="row">
      <div
        class="col-md-4"
        v-for="transporte in transportes"
        :key="transporte.id"
      >
        <div class="card mb-4">
          <img
            :src="transporte.image"
            class="card-img-top"
            alt="Vehículo de Transporte"
          />
          <div class="card-body">
            <h5 class="card-title">{{ transporte.tipo }}</h5>
            <p class="card-text"><strong>Placa:</strong> {{ transporte.placa }}</p>
            <p class="card-text"><strong>Conductor:</strong> {{ transporte.conductor }}</p>
            <p class="card-text"><strong>Celular:</strong> {{ transporte.celular }}</p>
          
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
      transportes: [],
    };
  },
  mounted() {
    this.fetchTransportes();
  },
  methods: {
    async fetchTransportes() {
      try {
        const response = await axios.get('http://localhost:3000/Trasporte');
        this.transportes = response.data;
      } catch (error) {
        console.error('Error al obtener la información de transporte', error);
      }
    },
  },
};
</script>

