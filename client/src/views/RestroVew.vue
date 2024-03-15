<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <h2>Registro de Usuario</h2>
        <form @submit.prevent="submitForm">
          <div class="form-group">
            <label for="name">Nombres</label>
            <input v-model="formData.name" type="text" class="form-control" id="name" placeholder="Ingrese sus nombres" required>
          </div>
          <div class="form-group">
            <label for="file">Imagen</label>
            <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change="uploadFile($event)" id="file"
                      ref="file" type="file" style="display:none">
          </div>
          <div class="form-group">
            <label for="email">Correo Electrónico</label>
            <input v-model="formData.email" type="email" class="form-control" id="email" placeholder="Ingrese su correo electrónico" required>
          </div>
          <div class="form-group">
            <label for="password">Contraseña</label>
            <input v-model="formData.password" type="password" class="form-control" id="password" placeholder="Ingrese su contraseña" required>
          </div>
          <div class="form-group">
            <label for="address">Dirección</label>
            <input v-model="formData.address" type="text" class="form-control" id="address" placeholder="Ingrese su dirección" required>
          </div>
          <div class="form-group">
            <label for="phone">Celular</label>
            <input v-model="formData.phone" type="text" class="form-control" id="phone" placeholder="Ingrese su número de celular" required>
          </div>
          <!-- Agregar control para el campo de rol si es necesario -->
          
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
      file: null,
      formData: {
        name: '',
        image: '', // Make sure this matches the key used for file in uploadFile method
        email: '',
        password: '',
        address: '',
        phone: '',
        
      
      }
    };
  },
  methods: {
    uploadFile(event) {
      if (event.target.files.length > 0) {
        this.formData.image = event.target.files[0]; // Assign to formData.image
      }
    },
    submitForm() {
      const formData = new FormData();
      formData.append('name', this.formData.name);
      formData.append('image', this.formData.image);
      formData.append('email', this.formData.email);
      formData.append('password', this.formData.password);
      formData.append('address', this.formData.address);
      formData.append('phone', this.formData.phone);
    
      
      axios.post('http://localhost:3000/Customer', formData)
  .then(response => {
    if (response && response.data) {
      console.log('Respuesta del servidor:', response.data);
      // Aquí puedes manejar la respuesta del servidor si lo deseas
    } else {
      console.error('Error: No se recibió una respuesta válida del servidor.');
    }
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
