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
          <form>
            <div class="form-group">
              <label for="usuario">Usuario:</label>
              <input type="text" class="form-control" id="usuario" placeholder="Ingrese su usuario">
            </div>
            <div class="form-group">
              <label for="contrasena">Contraseña:</label>
              <input type="password" class="form-control" id="contrasena" placeholder="Ingrese su contraseña">
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
document.getElementById('loginForm').addEventListener('submit', function (event) {
  event.preventDefault();

  var usuario = document.getElementById('usuario').value;
  var contrasena = document.getElementById('contrasena').value;

  
  fetch('http://localhost:3000/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ usuario: usuario, contrasena: contrasena })
  })
  .then(response => response.json())
  .then(data => {
    // Verificar la respuesta del servidor
    if (data.autenticado) {
      alert('Inicio de sesión exitoso');
    } else {
      alert('Inicio de sesión fallido. Verifica tus credenciales.');
    }
  })
  .catch(error => {
    console.error('Error en la solicitud:', error);
  });
});
</script>

