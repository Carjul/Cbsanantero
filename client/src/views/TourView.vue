<!--- <div class=" container">
   <hr>
    <div class="row">
      <div
        class="col-md-4"
        v-for="tour in tours"
        :key="tour.id"
      >
        <div class="card mb-4">
          <img
            :src="tour.image"
            class="card-img-top"
            alt="Transporte para Tours"
          />
          <div class="card-body">
            <h5 class="card-title">{{ tour.name }}</h5>
            <p class="card-text">{{ tour.address }}</p>
            <p class="card-text"><strong>Teléfono:</strong> {{ tour.phone }}</p>
            <p class="card-text"><strong>Precio:</strong> {{ tour.price }}</p>
            <p class="card-text"><strong>Transporte:</strong> {{ tour.transporte }}</p>
          
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
      tours: [],
    };
  },
  mounted() {
    this.fetchTours();
  },
  methods: {
    async fetchTours() {
      try {
        const response = await axios.get(`${process.env.API}/Tour`);
        this.tours = response.data;
      } catch (error) {
        console.error('Error al obtener los tours y transporte', error);
      }
    },
  },
};
</script>---->

<template>
  <!-- eslint-disable vue/no-unused-components -->
  <div class="container">
    <hr>
    <vue-slick-carousel :options="slickOptions">
      <div v-for="tour in tours" :key="tour.id">
        <div class="card mb-4">
          <img
            :src="tour.image"
            class="card-img-top"
            alt="Transporte para Tours"
          />
          <div class="card-body">
            <h5 class="card-title">{{ tour.name }}</h5>
            <p class="card-text">{{ tour.address }}</p>
            <p class="card-text"><strong>Teléfono:</strong> {{ tour.phone }}</p>
            <p class="card-text"><strong>Precio:</strong> {{ tour.price }}</p>
            <p class="card-text"><strong>Transporte:</strong> {{ tour.transporte }}</p>
          </div>
        </div>
      </div>
    </vue-slick-carousel>
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
import 'vue-slick-carousel/dist/vue-slick-carousel.css';
import 'vue-slick-carousel/dist/vue-slick-carousel-theme.css';
import { Slick } from 'vue-slick-carousel';

export default {
  components: {
    'vue-slick-carousel': Slick,
  },
  data() {
    return {
      tours: [],
      slickOptions: {
        vertical: true,
        verticalSwiping: true,
        slidesToShow: 3,
        slidesToScroll: 1,
        dots: true,
        responsive: [
          {
            breakpoint: 768,
            settings: {
              vertical: false,
              verticalSwiping: false,
              slidesToShow: 1,
            },
          },
        ],
      },
    };
  },
  mounted() {
    this.fetchTours();
  },
  methods: {
    async fetchTours() {
      try {
        const response = await axios.get(`${process.env.API}/Tour`);
        this.tours = response.data;
      } catch (error) {
        console.error('Error al obtener los tours y transporte', error);
      }
    },
  },
};
</script>

<style scoped>
.card {
  transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
}

.card:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  border-color: blue; /* Color del borde al pasar el cursor */
}

.card-img-top {
  width:auto; /* Ancho deseado */
  height:300px; /* Altura deseada */
  object-fit: cover; /* Para asegurar que la imagen se ajuste correctamente sin distorsión */
}
</style>
