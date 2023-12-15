<!-- En tu componenteRecreacion.vue -->
<template>
  <div class="container">
    <hr>
    <div class="row">
      <div
        class="col-md-4"
        v-for="recreacion in recreaciones"
        :key="recreacion.id"
      >
        <div class="card mb-4">
          <img
            :src="recreacion.image"
            class="card-img-top"
            alt="Centro de Recreación"
          />
          <div class="card-body">
            <h5 class="card-title">{{ recreacion.name }}</h5>
            <p class="card-text">{{ recreacion.address }}</p>
            <p class="card-text"><strong>Teléfono:</strong> {{ recreacion.phone }}</p>
            <p class="card-text"><strong>Servicios:</strong> {{ recreacion.services }}</p>
            <p class="card-text"><strong>Precio:</strong> {{ recreacion.price }}</p>
            <!-- Agrega aquí otros campos que quieras mostrar -->
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
      recreaciones: [],
    };
  },
  mounted() {
    this.fetchRecreaciones();
  },
  methods: {
    async fetchRecreaciones() {
      try {
        const response = await axios.get('http://localhost:3000/Recreacion');
        this.recreaciones = response.data;
      } catch (error) {
        console.error('Error al obtener los centros de recreación', error);
      }
    },
  },
};
</script>

<style>
/* Agrega estilos según tus preferencias o utiliza estilos de Bootstrap aquí */
</style>
