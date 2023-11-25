<template>
  <nav>
   
      <!-- Navbar de Bootstrap -->
  <nav class="navbar navbar-expand-lg navbar-dark fixed-top">
    <a class="navbar-brand" href="#">
      CB
    </a>
    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav ml-auto">
        <li class="nav-item active">
          <a class="nav-link" href="#"> <router-link to="/">Home</router-link>  </a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">  <router-link to="/about">About</router-link> </a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">Servicios</a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="#">Contacto</a>
        </li>
      </ul>
    </div>
  </nav>



    
   
  </nav>
  <router-view/>
</template>

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
      padding-top: 56px; /* Ajusta el cuerpo de la página para evitar que el navbar oculte el contenido */
    }

    .navbar {
      background-color: #3498db; /* Color de fondo del navbar */
      height: 90px;
      width: 100%;
    }

    .navbar-brand, .navbar-nav .nav-link {
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
