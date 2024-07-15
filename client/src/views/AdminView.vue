<template>
  <div>
    <MostrarVista :infoUser="infoReset"></MostrarVista>
    <hr>

    <!-- Agrega los botones que mostrarán la vista de usuario o artesanías al hacer clic -->
 
    <!-- Muestra el componente correspondiente al iniciar -->
    <div v-if="mostrarInicial">
      <component :is="componenteInicial"></component>
    </div>

  </div>
</template>

<script>
import Us from "@/views/listar/ListarUsuario.vue";
import Ar from "@/views/listar/ListarArtesanias.vue";
import Ba from "@/views/listar/ListarBar.vue";
import Ho from "@/views/listar/ListarHospedaje.vue";
import Te from "@/views/listar/ListarHotel.vue";
import Ta from "@/views/listar/ListarTrasporte.vue";
import Jue from "@/views/listar/ListarJuegos.vue";
import Tures from "@/views/listar/ListarTures.vue";
import Rest from "@/views/listar/ListarRestaurant.vue";

export default {
  components: {
    Us,
    Ar,
    Ba,
    Ho,
    Te, 
    Ta, 
    Jue,
    Tures,
    Rest
  },
  data() {
    return {
      usuarioAutenticado: localStorage.getItem('usuarioAutenticado') === 'true',
      usuarioRol: localStorage.getItem('usuarioRol') || '',
      tipo_negocio: localStorage.getItem('Tipo_negocio') || '',
    };
  },
  computed: {
    componenteInicial() {
      if (this.usuarioAutenticado && this.tipo_negocio === 'Artesanías') return 'Ar';
      else if (this.usuarioAutenticado && this.tipo_negocio === 'Bares') return 'Ba';
      else if (this.usuarioAutenticado && this.tipo_negocio === 'Hospedaje') return 'Ho';
      else if (this.usuarioAutenticado && this.tipo_negocio === 'Hoteles') return 'Te';
      else if (this.usuarioAutenticado && this.tipo_negocio === 'Juegos') return 'Jue';
      else if (this.usuarioAutenticado && this.tipo_negocio === 'Restaurantes') return 'Rest';
      else if (this.usuarioAutenticado && this.tipo_negocio === 'Tour') return 'Tures';
      else if (this.usuarioAutenticado && this.tipo_negocio === 'Transporte') return 'Ta'; // Actualizamos el tipo de negocio
      else if (this.usuarioAutenticado && this.usuarioRol === 'Admin') return 'Us';
      else return null;
    },
    mostrarInicial() {
      return !!this.componenteInicial;
    }
  },
};
</script>
