
<template>
    <!-- eslint-disable -->
  <nav>
    <!-- Topbar Start -->
    <div class="container-fluid bg-light pt-3 d-none d-lg-block">
      <div class="container">
        <div class="row">
          <div class="col-lg-6 text-center text-lg-left mb-2 mb-lg-0">
            <div class="d-inline-flex align-items-center">
              <p><i class="fa fa-envelope mr-2"></i>costabrisa@gmail.com</p>
              <p class="text-body px-3">|</p>
              <p><i class="fa fa-phone-alt mr-2"></i>+57 3117119510</p>
            </div>
          </div>
          <div class="col-lg-6 text-center text-lg-right">
            <div class="d-inline-flex align-items-center">
              <a class="text-primary px-3" href="">
                <i class="fab fa-facebook-f"></i>
              </a>
              <a class="text-primary px-3" href="">
                <i class="fab fa-twitter"></i>
              </a>
              <a class="text-primary px-3" href="">
                <i class="fab fa-linkedin-in"></i>
              </a>
              <a class="text-primary px-3" href="">
                <i class="fab fa-instagram"></i>
              </a>
              <a class="text-primary pl-3" href="">
                <i class="fab fa-youtube"></i>
              </a>

              <router-link class="text-primary pl-3" to="/login">Iniciar</router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- Topbar End -->

    <!-- Navbar Start -->
    <div class="container-fluid position-relative nav-bar p-0">
      <div class="container-lg position-relative p-0 px-lg-3" style="z-index: 9;">
        <nav class="navbar navbar-expand-lg bg-light navbar-light shadow-lg py-3 py-lg-0 pl-3 pl-lg-5">
          <a href="" class="navbar-brand">
            <h1 class="m-0 text-primary"><span class="text-dark">COSTA</span>BRISA</h1>
          </a>
          <button type="button" class="navbar-toggler" data-toggle="collapse" data-target="#navbarCollapse">
            <span class="navbar-toggler-icon"></span>
          </button>
          <div class="collapse navbar-collapse justify-content-between px-3" id="navbarCollapse">
            <div class="navbar-nav ml-auto py-0">
              <!--
                  <router-link v-if="usuarioAutenticado && usuarioRol === 'admin'" to="/admin">Admin</router-link>
              <router-link v-if="usuarioAutenticado && usuarioRol === 'admin'" to="/admin">Publicar</router-link>
              -->
            
              <a href="" class="nav-item nav-link"> <router-link  to="/">Home</router-link></a>
              <a href="" class="nav-item nav-link"><router-link to="/about">About</router-link></a>
              <a href="service.html" class="nav-item nav-link">Services</a>
              <a href="package.html" class="nav-item nav-link">Tour Packages</a>
              <div class="nav-item dropdown">
                <a href="#" class="nav-link dropdown-toggle" data-toggle="dropdown">Servicios</a>
                <div class="dropdown-menu border-0 rounded-0 m-0">
                  <router-link to="/artes" class="dropdown-item">Atesanias</router-link>
                  <router-link to="/bard" class="dropdown-item">Bares y Discotecas</router-link>
                  <router-link to="/hospe" class="dropdown-item">Hospedaje</router-link>
                  <router-link to="/hotels" class="dropdown-item">Hoteles</router-link>
                  <router-link to="/recre" class="dropdown-item">Juegos y Recreación</router-link>
                </div>
              </div>
              <a href="contact.html" class="nav-item nav-link">Contact</a>
            </div>
          </div>
        </nav>
      </div>
    </div>
    <!-- Navbar End -->

    <router-view></router-view>
  </nav>
   <!-- eslint-enable -->
</template>

<script>
export default {
  data() {
    return {
      usuarioAutenticado: localStorage.getItem('usuarioAutenticado') === 'true',
      usuarioRol: localStorage.getItem('usuarioRol') || '',
    };
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 5px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}

/* Estilos personalizados para el navbar */
body {
  padding-top: 10px; /* Ajusta el cuerpo de la página para evitar que el navbar oculte el contenido */
}

.navbar {
  background-color: #3498db; /* Color de fondo del navbar */
  height: 90px;
  width: 100%;
}

.navbar-brand,
.navbar-nav .nav-link {
  color: #fff; /* Color del texto del navbar */
}

.navbar-toggler-icon {
  background-color: rgba(255, 255, 255, 0.142)ccc; /* Color del ícono del botón de alternar */
}

/* Animación personalizada para el navbar */
@keyframes slideInLeftToRight {
  from {
    opacity: 0;
    transform: translate3d(-100%, 0, 0);
  }
  to {
    opacity: 1;
    transform: none;
  }
}

/* Estilos para dispositivos móviles */
@media (max-width: 767px) {
  .navbar-collapse {
    animation: slideInLeftToRight 0.5s ease;
    background-color: rgba(255, 255, 255, 0.413); /* Fondo blanco para el menú desplegable en dispositivos móviles */
  }
}
</style>
