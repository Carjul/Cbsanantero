<template>
  <div>
    <MostrarVista :infoUser="infoReset"></MostrarVista>
    <hr>

    <!-- Agrega los botones que mostrarán la vista de usuario o artesanías al hacer clic -->
    <button v-if="usuarioAutenticado && usuarioRol === 'Admin'" @click="mostrarUsuario">Usuario</button>
    <button @click="mostrarArtesanias">Artesanías</button>
    <button @click="mostrarBar">Bares</button>

    <!-- Muestra el componente "Us" solo si la variable mostrarUsuarioComponent es verdadera -->
    <div v-if="mostrarUsuarioComponent">
      <Us></Us>
    </div>

    <!-- Muestra el componente "Ar" solo si la variable mostrarArtesaniasComponent es verdadera -->
    <div v-if="mostrarArtesaniasComponent">
      <Ar></Ar>
    </div>

    <!-- Muestra el componente "Ba" solo si la variable mostrarBarComponet es verdadera -->
    <div v-if="mostrarBarComponet">
      <Ba></Ba>
    </div>
  </div>
</template>

<script>
import Us from "@/views/listar/ListarUsuario.vue";
import Ar from "@/views/listar/ListarArtesanias.vue";
import Ba from "@/views/listar/ListarBar.vue";

export default {
  components: {
    Us,
    Ar,
    Ba,
  },
  data() {
    return {
      mostrarUsuarioComponent: false,
      mostrarArtesaniasComponent: false,
      mostrarBarComponet: false,
      usuarioAutenticado: localStorage.getItem('usuarioAutenticado') === 'true',
      usuarioRol: localStorage.getItem('usuarioRol') || '',
    };
  },
  methods: {
    mostrarUsuario() {
      // Muestra el componente "Us" y oculta los componentes "Ar" y "Ba"
      this.mostrarUsuarioComponent = true;
      this.mostrarArtesaniasComponent = false;
      this.mostrarBarComponet = false;
    },
    mostrarArtesanias() {
      // Muestra el componente "Ar" y oculta los componentes "Us" y "Ba"
      this.mostrarArtesaniasComponent = true;
      this.mostrarUsuarioComponent = false;
      this.mostrarBarComponet = false;
    },
    mostrarBar() {
      // Muestra el componente "Ba" y oculta los componentes "Us" y "Ar"
      this.mostrarUsuarioComponent = false;
      this.mostrarArtesaniasComponent = false;
      this.mostrarBarComponet = true;
    }
  },
};
</script>
