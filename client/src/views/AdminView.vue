<template>
  <div>
    <MostrarVista :infoUser="infoReset"></MostrarVista>
     <hr>
    <!-- Agrega los botones que mostrarán la vista de usuario o artesanías al hacer clic -->
    <button v-if="usuarioAutenticado && usuarioRol === 'Admin'" @click="mostrarUsuario">Usuario</button>
    <button @click="mostrarArtesanias">Artesanías</button>

    <!-- Muestra el componente "Us" solo si la variable mostrarUsuario es verdadera -->
    <div v-if="mostrarUsuarioComponent">
      <Us></Us>
    </div>

    <!-- Muestra el componente "Ar" solo si la variable mostrarArtesanias es verdadera -->
    <div v-if="mostrarArtesaniasComponent">
      <Ar></Ar>
    </div>
  </div>
</template>

<script>
import Us from "@/views/listar/ListarUsuario.vue";
import Ar from "@/views/listar/ListarArtesanias.vue";

export default {
  components: {
    Us,
    Ar,
  },
  data() {
    return {
      mostrarUsuarioComponent: false,
      mostrarArtesaniasComponent: false,
      usuarioAutenticado: localStorage.getItem('usuarioAutenticado') === 'true',
      usuarioRol: localStorage.getItem('usuarioRol') || '',
    };
  },
  methods: {
    mostrarUsuario() {
      // Muestra el componente "Us" y oculta el componente "Ar"
      this.mostrarUsuarioComponent = true;
      this.mostrarArtesaniasComponent = false;
      localStorage.removeItem('usuarioAutenticado');
      localStorage.removeItem('usuarioRol');
    },
    mostrarArtesanias() {
      // Muestra el componente "Ar" y oculta el componente "Us"
      this.mostrarArtesaniasComponent = true;
      this.mostrarUsuarioComponent = false;
    },
  },
};
</script>
