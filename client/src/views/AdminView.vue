<template>
  <div>
    <MostrarVista :infoUser="infoReset"></MostrarVista>
    <hr>

    <!-- Agrega los botones que mostrarán la vista de usuario o artesanías al hacer clic -->
<!--     <button v-if="usuarioAutenticado && usuarioRol === 'Admin'" @click="mostrarUsuario">Usuario</button>
 -->    <button v-if="usuarioAutenticado && tipo_negocio === 'Admin'" @click="mostrarArtesanias">Artesanías</button>
    <button v-if="usuarioAutenticado && tipo_negocio === 'Hospedaje'" @click="mostrarBar">Bares</button>
    <button v-if="usuarioAutenticado && tipo_negocio === 'Hospedaje'" @click="mostrarHospedaje">Hospedaje</button>
    <button v-if="usuarioAutenticado && tipo_negocio === 'Hotel'" @click="mostrarHotel">Hotel</button>
    <button v-if="usuarioAutenticado && tipo_negocio === 'Trasporte'" @click="mostrarTrasporte">Trasporte</button>
    <button v-if="usuarioAutenticado && tipo_negocio === 'Juegos'" @click="mostrarJuegos">Juegos</button>

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

    <!-- Muestra el componente "Ho" solo si la variable mostrarHospedajeComponent es verdadera -->
    <div v-if="mostrarHospedajeComponent">
      <Ho></Ho>
    </div>

    <!-- Muestra el componente "Te" solo si la variable mostrarHotelComponent es verdadera -->
    <div v-if="mostrarHotelComponent">
      <Te></Te>
    </div>
    
 <!-- Muestra el componente "Te" solo si la variable mostrarHotelComponent es verdadera -->
 <div v-if="mostrarTrasporteComponent">
      <Tra></Tra>
    </div>

     <!-- Muestra el componente "Jue" solo si la variable mostrarHotelComponent es verdadera -->
 <div v-if="mostrarJuegosComponent">
      <Jue></Jue>
    </div>

  </div>

</template>

<script>
import Us from "@/views/listar/ListarUsuario.vue";
import Ar from "@/views/listar/ListarArtesanias.vue";
import Ba from "@/views/listar/ListarBar.vue";
import Ho from "@/views/listar/ListarHospedaje.vue";
import Te from "@/views/listar/ListarHotel.vue";
import Tra from "@/views/listar/ListarTrasporte.vue"
import Jue from "@/views/listar/ListarJuegos.vue"
export default {
  components: {
    Us,
    Ar,
    Ba,
    Ho,
    Te, 
    Tra,
    Jue,
  },
  data() {
    return {
      mostrarUsuarioComponent: false,
      mostrarArtesaniasComponent: false,
      mostrarBarComponet: false,
      mostrarHospedajeComponent: false,
      mostrarHotelComponent: false, 
      mostrarTrasporteComponent: false,
      mostrarJuegosComponent: false, 
      usuarioAutenticado: localStorage.getItem('usuarioAutenticado') === 'true',
      usuarioRol: localStorage.getItem('usuarioRol') || '',
      tipo_negocio: localStorage.getItem('Tipo_negocio') || '',
    };
  },
  mounted(){
    let c = localStorage.getItem('usuarioAutenticado')
    let d = localStorage.getItem('usuarioRol')
    
  
  if (c !== 'true' || (c === 'true' && d === 'Cliente') ) {
    this.$router.push('/')
  }
  },
  
  methods: {
    mostrarUsuario() {
      // Muestra el componente "Us" y oculta los componentes "Ar", "Ba", "Ho" y "Te"
      this.mostrarUsuarioComponent = true;
      this.mostrarArtesaniasComponent = false;
      this.mostrarBarComponet = false;
      this.mostrarHospedajeComponent = false;
      this.mostrarHotelComponent = false;
      this.mostrarTrasporteComponent = false;
      this.mostrarJuegosComponent = false; 
    },
    mostrarArtesanias() {
      // Muestra el componente "Ar" y oculta los componentes "Us", "Ba", "Ho" y "Te"
      this.mostrarArtesaniasComponent = true;
      this.mostrarUsuarioComponent = false;
      this.mostrarBarComponet = false;
      this.mostrarHospedajeComponent = false;
      this.mostrarHotelComponent = false;
      this.mostrarTrasporteComponent = false;
      this.mostrarJuegosComponent = false; 
    },
    mostrarBar() {
      // Muestra el componente "Ba" y oculta los componentes "Us", "Ar", "Ho" y "Te"
      this.mostrarUsuarioComponent = false;
      this.mostrarArtesaniasComponent = false;
      this.mostrarBarComponet = true;
      this.mostrarHospedajeComponent = false;
      this.mostrarHotelComponent = false;
      this.mostrarTrasporteComponent = false;
      this.mostrarJuegosComponent = false; 
    },
    mostrarHospedaje() {
      // Muestra el componente "Ho" y oculta los componentes "Us", "Ar", "Ba" y "Te"
      this.mostrarUsuarioComponent = false;
      this.mostrarArtesaniasComponent = false;
      this.mostrarBarComponet = false;
      this.mostrarHospedajeComponent = true;
      this.mostrarHotelComponent = false;
      this.mostrarTrasporteComponent = false;
      this.mostrarJuegosComponent = false; 
    },
    mostrarHotel() {
      // Muestra el componente "Te" y oculta los componentes "Us", "Ar", "Ba" y "Ho"
      this.mostrarUsuarioComponent = false;
      this.mostrarArtesaniasComponent = false;
      this.mostrarBarComponet = false;
      this.mostrarHospedajeComponent = false;
      this.mostrarHotelComponent = true;
      this.mostrarTrasporteComponent = false;
      this.mostrarJuegosComponent = false; 
    },

    mostrarTrasporte(){
      this.mostrarUsuarioComponent = false;
      this.mostrarArtesaniasComponent = false;
      this.mostrarBarComponet = false;
      this.mostrarHospedajeComponent = false;
      this.mostrarHotelComponent = false;
      this.mostrarTrasporteComponent = true;
      this.mostrarJuegosComponent = false; 
    },
    mostrarJuegos(){
      this.mostrarUsuarioComponent = false;
      this.mostrarArtesaniasComponent = false;
      this.mostrarBarComponet = false;
      this.mostrarHospedajeComponent = false;
      this.mostrarHotelComponent = false;
      this.mostrarTrasporteComponent = false;
      this.mostrarJuegosComponent = true;  
    }
  },
};
</script>
