<template>
  <div class="container mt-5">
    <div class="row justify-content-center">   
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h3 class="text-center">Inicio de Sesión</h3>
          </div>
          <div class="card-body">
            <!-- Formulario de inicio de sesión -->
            <form @submit.prevent="submitForm">
              <div class="form-group">
                <label for="usuario">Usuario:</label>
                <input v-model="usuario" type="text" class="form-control" id="usuario" placeholder="Ingrese su usuario">
              </div>
              <div class="form-group">
                <label for="contrasena">Contraseña:</label>
                <input v-model="contrasena" type="password" class="form-control" id="contrasena"
                  placeholder="Ingrese su contraseña">
              </div>
              <button type="submit" class="btn btn-primary btn-block" @click="submitForm" :disabled="iniciandoSesion">
                Iniciar Sesión
              </button>
              <div>
                <hr>
                <label> Inicio de sececion con Google</label>
                <br>
                <svg style="width:100px;color: brown;" @click="login" xmlns="http://www.w3.org/2000/svg" width="16"
                  height="16" fill="currentColor" class="bi bi-google" viewBox="0 0 16 16">

                  <path
                    d="M15.545 6.558a9.4 9.4 0 0 1 .139 1.626c0 2.434-.87 4.492-2.384 5.885h.002C11.978 15.292 10.158 16 8 16A8 8 0 1 1 8 0a7.7 7.7 0 0 1 5.352 2.082l-2.284 2.284A4.35 4.35 0 0 0 8 3.166c-2.087 0-3.86 1.408-4.492 3.304a4.8 4.8 0 0 0 0 3.063h.003c.635 1.893 2.405 3.301 4.492 3.301 1.078 0 2.004-.276 2.722-.764h-.003a3.7 3.7 0 0 0 1.599-2.431H8v-3.08z" />
                </svg>



              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  data() {
    return {
      usuario: '',
      contrasena: '',
      iniciandoSesion: false,
    };
  },
  methods: {
    login() {
      this.$auth0.loginWithRedirect();
    },
   
    submitForm() {
      if (this.iniciandoSesion) {
        return;
      } else {
        this.iniciandoSesion = true;
        fetch('http://localhost:3000/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ email: this.usuario, password: this.contrasena }),
        })
          .then((response) => response.json())
          .then((data) => {

            if (data.token) {
              localStorage.setItem('usuarioAutenticado', true);
              localStorage.setItem('usuarioRol', data.user.rol);
              localStorage.setItem('customerId', data.user._id);
              localStorage.setItem('nombreUsuario', data.user.name);
              localStorage.setItem('correoUsuario', data.user.email);
              localStorage.setItem('celularUsuario', data.user.phone);
              localStorage.setItem('imagenUsuario', data.user.image);
              localStorage.setItem('Tipo_negocio', data.user.tipo_negocio);
              

              if (localStorage.getItem('usuarioRol', data.user.rol) === 'Admin' || localStorage.getItem('usuarioRol', data.user.rol) === 'Cliente' || localStorage.getItem('usuarioRol', data.user.rol) === 'Vendedor'){
                location.href = '/';
                console.log("sesion iniciada")
              } else {
                this.$router.push({ path: '/login' });
              }
            }
          })
          .catch((error) => {
            console.error('Error en la solicitud:', error);
          })
          .finally(() => {
            this.iniciandoSesion = false;
          });
      }

    },
  },
  mounted() {
    let c = localStorage.getItem('usuarioAutenticado')
  
    if (c === 'true') {
      this.$router.push({ path: '/' });
    }
  

    
  },
  beforeRouteLeave(to, from, next) {
    // Restablece los valores de usuario y contrasena al abandonar la ruta
    this.usuario = '';
    this.contrasena = '';
    next();
  },
};
</script>