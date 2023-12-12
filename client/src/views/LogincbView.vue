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
              <label  for="contrasena">Contraseña:</label>
              <input v-model="contrasena" type="password" class="form-control" id="contrasena" placeholder="Ingrese su contraseña">
            </div>
            <button type="submit" class="btn btn-primary btn-block">Iniciar Sesión</button>
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
      contrasena: ''
    };
  },
  methods: {
    submitForm() {
      fetch('http://localhost:3000/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email: this.usuario, password: this.contrasena })
      })
      .then(response => response.json())
      .then(data => {
       console.log(data)
      })
      .catch(error => {
        console.error('Error en la solicitud:', error);
      });
    }
  }
};
</script>