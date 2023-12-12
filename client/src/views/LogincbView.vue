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
                  <input v-model="contrasena" type="password" class="form-control" id="contrasena" placeholder="Ingrese su contraseña">
                </div>
                <button type="submit" class="btn btn-primary btn-block" @click="submitForm" :disabled="iniciandoSesion">
                  Iniciar Sesión
                </button>
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
      submitForm() {
        if (this.iniciandoSesion) {
          return;
        }
  
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
            console.log(data);
            if (data.autenticado) {
              localStorage.setItem('usuarioAutenticado', true);
              localStorage.setItem('usuarioRol', data.rol);
  
              if (data.rol === 'admin') {
                // Redirige a la vista principal si el usuario es administrador
                this.$router.push({ name: 'home' });
              } else {
                // Redirige a otra vista si no es administrador
                // Puedes personalizar esta lógica según tus necesidades
                // Por ahora, redirigimos a la página de inicio
                this.$router.push({ path: '/' });
              }
            }
          })
          .catch((error) => {
            console.error('Error en la solicitud:', error);
          })
          .finally(() => {
            this.iniciandoSesion = false;
          });
      },
    },
    beforeRouteLeave(to, from, next) {
      // Restablece los valores de usuario y contrasena al abandonar la ruta
      this.usuario = '';
      this.contrasena = '';
      next();
    },
  };
  </script>
  