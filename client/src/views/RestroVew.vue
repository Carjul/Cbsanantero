<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <h2>Registro de Usuario</h2>
        <form @submit.prevent="submitForm">
          <div class="form-group">
            <label for="nombres">Nombres</label>
            <input v-model="formData.nombres" type="text" class="form-control" id="nombres" placeholder="Ingrese sus nombres">
          </div>
          <div class="form-group">
            <label for="imagen">Imagen</label>
            <input type="file" class="form-control-file" id="imagen">
          </div>
          <div class="form-group">
            <label for="correo">Correo Electrónico</label>
            <input v-model="formData.correo" type="email" class="form-control" id="correo" placeholder="Ingrese su correo electrónico">
          </div>
          <div class="form-group">
            <label for="contraseña">Contraseña</label>
            <input v-model="formData.contraseña" type="password" class="form-control" id="contraseña" placeholder="Ingrese su contraseña">
          </div>
          <div class="form-group">
            <label for="direccion">Dirección</label>
            <input v-model="formData.direccion" type="text" class="form-control" id="direccion" placeholder="Ingrese su dirección">
          </div>
          <div class="form-group">
            <label for="celular">Celular</label>
            <input v-model="formData.celular" type="text" class="form-control" id="celular" placeholder="Ingrese su número de celular">
          </div>
          <div class="form-group">
            <label for="rol">Rol</label>
            <select v-model="formData.rol" class="form-control" id="rol">
              <option value="cliente">Cliente</option>
              <option value="vendedor">Vendedor</option>
            </select>
          </div>
          <!-- Mostrar el campo "Tipo de negocio" solo cuando el rol seleccionado sea "Vendedor" -->
          <div v-if="formData.rol === 'vendedor'" class="form-group">
            <label for="tipoNegocio">Tipo de Negocio</label>
            <select v-model="formData.tipoNegocio" class="form-control" id="tipoNegocio">
              <option value="Transporte">Transporte</option>
              <option value="Hospedaje">Hospedaje</option>
              <option value="Tour">Tour</option>
              <option value="Hoteles">Hoteles</option>
              <option value="Restaurantes">Restaurantes</option>
              <option value="Recreación">Recreación</option>
              <option value="Bares">Bares y discotecas</option>
              <option value="Artesanías">Artesanías</option>
            </select>
          </div>
          <button type="submit" class="btn btn-primary">Registrarse</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      formData: {
        nombres: '',
        correo: '',
        contraseña: '',
        direccion: '',
        celular: '',
        rol: '',
        tipoNegocio: ''
      }
    };
  },
  methods: {
    submitForm() {
      axios.post('http://localhost:3000/Customer', this.formData)
        .then(response => {
          console.log('Respuesta del servidor:', response.data);
          // Aquí puedes manejar la respuesta del servidor si lo deseas
        })
        .catch(error => {
          console.error('Error al enviar datos:', error);
          // Aquí puedes manejar el error si ocurre alguno
        });
    }
  }
}
</script>

<style scoped>
/* Estilos específicos si los necesitas */
</style>