
<template>
    <!-- eslint-disable -->
  <nav>
    <!-- Topbar Start -->
    <div class="container-fluid bg-light pt-3 d-none d-lg-block">
      <div class="container">
        <div class="row">
          <div class="col-lg-6 text-center text-lg-left mb-2 mb-lg-0">
            <div class="d-inline-flex align-items-center">
              <p><i class="fa fa-envelope mr-2"></i>costabrisaweb@gmail.com</p>
              <p class="text-body px-3">|</p>
              <p><i class="fa fa-phone-alt mr-2"></i>+57 3117119510</p>
            </div>
          </div>
          <div class="col-lg-6 text-center text-lg-right">
            <div class="d-inline-flex align-items-center" id="res">
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
              <router-link class="text-primary pl-3" v-if="usuarioAutenticado && usuarioRol === 'Admin' || usuarioAutenticado && usuarioRol === 'Vendedor'" to="/admin">Dashboard</router-link>
              
              <router-link class="text-primary pl-3" v-if="!usuarioAutenticado"  to="/login">Iniciar</router-link>
              <div  class="nav-item nav-link">
               
                <i  v-if="usuarioAutenticado && usuarioRol !== ''" @click="cerrarSesion()" class="fas fa-sign-out-alt"></i> 
              
              </div>

              <div v-if="usuarioAutenticado" class="usuario-info">
                <img :src="imagenUsuario" alt="avatar" loading="lazy" class="avatar" >
                <p class="nombre-usuario">{{ nombreUsuario }}</p>
              </div>
                 <p style="color: transparent;">..</p>
                 <div class="notification-container">
                
              </div>



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
            
              <a href="" class="nav-item nav-link">
               
                <router-link to="/">
                  <i class="fas fa-home"></i> Inicio
                </router-link>
              </a>
              <router-link to="#" class="nav-item nav-link">
               Acerca de
              </router-link>
              <a href="service.html" class="nav-item nav-link">
               Mapa turístico
               </a>
               
               <router-link to="/tourn" class="nav-item nav-link">
                   Tours
                </router-link>

              
              <div class="nav-item dropdown">
                <a href="#" class="nav-link dropdown-toggle" data-toggle="dropdown">
                 Servicios
                </a>

                <div class="dropdown-menu border-0 rounded-0 m-0">
                  <router-link to="/artes" class="dropdown-item">Atesanias</router-link>
                  <router-link to="/tranport" class="dropdown-item">Transporte</router-link>
                  <router-link to="/hospe" class="dropdown-item">Hospedaje</router-link>
                  <router-link to="/hotels" class="dropdown-item">Hoteles</router-link>
                  <router-link to="/bard" class="dropdown-item">Pescadores</router-link>
                  <router-link to="/recre" class="dropdown-item">Juegos y Recreación</router-link>
                </div>
              </div>
              <!-- <a href="contact.html" class="nav-item nav-link">
                <i class="fas fa-envelope"></i> Contact
              </a> -->

            </div>
          </div>
        </nav>
      </div>
    </div>
    <!-- Navbar End -->

    <router-view></router-view>
  </nav>
  


  

    <!-- Footer Start -->
    <div class="container-fluid bg-dark text-white-50 py-5 px-sm-3 px-lg-5" style="margin-top: 90px;">
        <div class="row pt-6">
            <div class="col-lg-3 col-md-6 mb-5">
                <a href="" class="navbar-brand">
                    <h1 class="text-primary"><span class="text-white">COSTA</span>BRISA</h1>
                </a>
                <p>Sed ipsum clita tempor ipsum ipsum amet sit ipsum lorem amet labore rebum lorem ipsum dolor. No sed vero lorem dolor dolor</p>
                <h6 class="text-white text-uppercase mt-4 mb-3" style="letter-spacing: 5px;">Follow Us</h6>
                <div class="d-flex justify-content-start">
                    <a class="btn btn-outline-primary btn-square mr-2" href="#"><i class="fab fa-twitter"></i></a>
                    <a class="btn btn-outline-primary btn-square mr-2" href="#"><i class="fab fa-facebook-f"></i></a>
                    <a class="btn btn-outline-primary btn-square mr-2" href="#"><i class="fab fa-linkedin-in"></i></a>
                    <a class="btn btn-outline-primary btn-square" href="#"><i class="fab fa-instagram"></i></a>
                </div>
            </div>
            <div class="col-lg-3 col-md-6 mb-5">
                <h5 class="text-white text-uppercase mb-4" style="letter-spacing: 5px;">Our Services</h5>
                <div class="d-flex flex-column justify-content-start">
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>About</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Destination</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Services</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Packages</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Guides</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Testimonial</a>
                    <a class="text-white-50" href="#"><i class="fa fa-angle-right mr-2"></i>Blog</a>
                </div>
            </div>
            <div class="col-lg-3 col-md-6 mb-5">
                <h5 class="text-white text-uppercase mb-4" style="letter-spacing: 5px;">Usefull Links</h5>
                <div class="d-flex flex-column justify-content-start">
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>About</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Destination</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Services</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Packages</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Guides</a>
                    <a class="text-white-50 mb-2" href="#"><i class="fa fa-angle-right mr-2"></i>Testimonial</a>
                    <a class="text-white-50" href="#"><i class="fa fa-angle-right mr-2"></i>Blog</a>
                </div>
            </div>
            <div class="col-lg-3 col-md-6 mb-5">
                <h5 class="text-white text-uppercase mb-4" style="letter-spacing: 5px;">Contact Us</h5>
                <p><i class="fa fa-map-marker-alt mr-2"></i>123 Street, New York, USA</p>
                <p><i class="fa fa-phone-alt mr-2"></i>+012 345 67890</p>
                <p><i class="fa fa-envelope mr-2"></i>info@example.com</p>
                <h6 class="text-white text-uppercase mt-4 mb-3" style="letter-spacing: 5px;">Newsletter</h6>
                <div class="w-100">
                    <div class="input-group">
                        <input type="text" class="form-control border-light" style="padding: 25px;" placeholder="Your Email">
                        <div class="input-group-append">
                            <button class="btn btn-primary px-3">Sign Up</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
    <div class="container-fluid bg-dark text-white border-top py-4 px-sm-3 px-md-5" style="border-color: rgba(256, 256, 256, .1) !important;">
        <div class="row">
            <div class="col-lg-6 text-center text-md-left mb-3 mb-md-0">
               
            </div>
            <div class="col-lg-6 text-center text-md-right">
                <p class="m-0 text-white-50">Designed by <a href="#">CostaBrisa</a>
                </p>
            </div>
        </div>
    </div>
    <!-- Footer End -->
    

   <!-- eslint-enable -->
</template>
<script>
import { ref, watchEffect } from 'vue'

export default {
  setup() {
    // Declara las variables reactivas usando ref()
    const usuarioAutenticado = ref(localStorage.getItem('usuarioAutenticado'));
    const usuarioRol = ref(localStorage.getItem('usuarioRol'));
    const imagenUsuario = ref(localStorage.getItem('imagenUsuario'));
    const nombreUsuario = ref(localStorage.getItem('nombreUsuario'));
    // Observa los cambios en las variables reactivas
    watchEffect(() => {
      usuarioAutenticado.value = localStorage.getItem('usuarioAutenticado');
      usuarioRol.value = localStorage.getItem('usuarioRol');
      imagenUsuario.value = localStorage.getItem('imagenUsuario');
      nombreUsuario.value = localStorage.getItem('nombreUsuario');
    });

    return {
      usuarioAutenticado,
      usuarioRol,
      imagenUsuario,
      nombreUsuario
    };
  },
  created(){
   
  
  },
  
  
  methods: {
    cerrarSesion() {
      // Limpiar la información de autenticación y del usuario
      localStorage.clear();

      // Otra lógica de cierre de sesión si es necesario
      // Actualizar el estado local de autenticación y del usuario
      this.usuarioAutenticado = false;
      this.usuarioRol = '';
      this.imagenUsuario = '';
      this.nombreUsuario = '';
      this.logout()
    },
    logout() {
        this.$auth0.logout({ logoutParams: { returnTo: window.location.origin } });
      }
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
    background-color: rgb(255, 255, 255); /* Fondo blanco para el menú desplegable en dispositivos móviles */
  }
}
.usuario-info {
  display: flex;
  align-items: center;
}

.avatar {
  width: 50px; /* ajusta el tamaño según tus preferencias */
  height: 50px; /* ajusta el tamaño según tus preferencias */
  border-radius: 50%; /* para hacerlo circular, ajusta según tus preferencias */
  margin-right: 10px; /* ajusta el margen según tus preferencias */
}

.nombre-usuario {
  margin: 0; /* asegúrate de quitar cualquier margen predeterminado si es necesario */
}
@keyframes glowAnimation {
            0% {
                box-shadow: 0 0 5px rgba(255, 255, 255, 0.5);
            }
            50% {
                box-shadow: 0 0 20px rgba(255, 255, 255, 0.7);
            }
            100% {
                box-shadow: 0 0 5px rgba(255, 255, 255, 0.5);
            }
        }

        .glow-text {
            display: inline-block;
            animation: glowAnimation 2s infinite;
        }

        h1 {
            font-family: 'Arial', sans-serif;
            font-size: 2em;
            margin: 0;
        }

        .text-primary {
            color: #007bff;
        }

        .text-dark {
            color: #333;
        }



/* Estilos generales para el enlace */
.navbar-brand {
  text-decoration: none;
  color: #000; /* Color del texto sin pasar el cursor */
  position: relative;
  overflow: hidden;
  transition: color 0.3s ease;
}

/* Efecto de sombreado y transición al pasar el cursor */
.navbar-brand::before {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 100%;
  height: 2px;
  background-color: #007bff; /* Color del sombreado al pasar el cursor */
  transform: scaleX(0);
  transform-origin: bottom right;
  transition: transform 0.3s ease;
}

.navbar-brand:hover {
  color: #007bff; /* Color del texto al pasar el cursor */
}

.navbar-brand:hover::before {
  transform: scaleX(1);
  transform-origin: bottom left;
}




</style>
