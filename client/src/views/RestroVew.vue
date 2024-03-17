<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <h2>Registro de Usuario</h2>
        <form @submit.prevent="submitForm" class="registration-form">
          <div class="form-group animated">
            <label for="name">Nombres</label>
            <input v-model="formData.name" type="text" class="form-control falling-animation" id="name" placeholder="Ingrese sus nombres" required>
          </div>
          <div class="form-group animated">
            <label for="file">Imagen</label>
            <input class="form-control falling-animation" accept="image/jpeg, image/jpg, image/png" @change="uploadFile($event)" id="file" ref="file" type="file">
          </div>
          <div class="form-group animated">
            <label for="email">Correo Electrónico</label>
            <input v-model="formData.email" type="email" class="form-control falling-animation" id="email" placeholder="Ingrese su correo electrónico" required>
          </div>
          <div class="form-group animated">
            <label for="password">Contraseña</label>
            <input v-model="formData.password" type="password" class="form-control falling-animation" id="password" placeholder="Ingrese su contraseña" required>
          </div>
          <div class="form-group animated">
            <label for="address">Dirección</label>
            <input v-model="formData.address" type="text" class="form-control falling-animation" id="address" placeholder="Ingrese su dirección" required>
          </div>
          <div class="form-group animated">
            <label for="phone">Celular</label>
            <input v-model="formData.phone" type="text" class="form-control falling-animation" id="phone" placeholder="Ingrese su número de celular" required>
          </div>
          <!-- Agregar control para el campo de rol si es necesario -->
          <button type="submit" class="btn btn-primary animated">Registrarse</button>
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
      file:null,
      formData: {
        name: '',
        email: '',
        password: '',
        address: '',
        phone: '',
        rol: 'Cliente',
        tipo_negocio:''
      }
    };
  },
  methods: {
    uploadFile(event) {
      this.file = event.target.files[0];
    },
   
    submitForm() {
      const Form = new FormData();
        Form.append('name', this.formData.name);
        Form.append('image', this.file);
        Form.append('email', this.formData.email);
        Form.append('password', this.formData.password);
        Form.append('address', this.formData.address);
        Form.append('phone', this.formData.phone);
        Form.append('rol', this.formData.rol);
        Form.append('tipo_negocio', this.formData.tipo_negocio);

      axios.post('http://localhost:3000/Customer', Form)
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
.container {
  background-color: #f9f9f9;
  padding: 20px;
}

.registration-form {
  background-color: #fff;
  padding: 20px;
  border-radius: 10px;
  box-shadow: 1px 10px 10px rgb(8, 168, 221);
}

.registration-form h2 {
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 20px;
}

label {
  font-weight: bold;
}

.btn-primary {
  background-color: #007bff;
  border-color: #007bff;
}

.btn-primary:hover {
  background-color: #0056b3;
  border-color: #0056b3;
}

.form-control {
  border-radius: 5px;
}

/* Animaciones */
.animated {
  animation-duration: 0.5s;
  animation-fill-mode: both;
}
/* 
@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.fadeIn {
  animation-name: fadeIn;
}

/* Animación de caída 
@keyframes fallIn {
  0% {
    opacity: 0;
    transform: translateY(-50px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}

.falling-animation {
  animation-name: fallIn;
} */
</style>
