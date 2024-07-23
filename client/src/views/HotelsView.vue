<template>
  <div class="container">
    <hr>
    <div class="row">
      <div
        class="col-md-4"
        v-for="hotel in hoteles"
        :key="hotel.id"
      >
        <div class="card mb-4">
          <img
            @click="enviarIdGaleria(hotel.customer_id, hotel._id)"
            :src="hotel.image"
            class="card-img-top"
            alt="Hotel"
          />
          <div class="card-body">
            <hr>
            <p class="h3 text-decoration-none">{{ hotel.name }}</p>
            <hr>
            <p class="card-text"></p>
            <div class="d-flex justify-content-between">
              <h5 class="m-0"><i class="fa fa-map-marker-alt text-primary mr-2"></i>{{ hotel.address }}</h5>
              <h5 class="m-0"><i class="fas fa-coins text-primary mr-2"></i>{{ hotel.price }}</h5>
            </div>
            <hr>
            <div v-if="customerRol === null">
              <button class="btn btn-primary" @click="confirmarRegistro">Solicitar servicio</button>
            </div>
            <div v-if="customerRol === 'Cliente'">
              <button class="btn btn-primary" data-toggle="modal" data-target="#solicitudModal" @click="enviarcid(hotel.customer_id, hotel._id)">Obtener Servicio</button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="modal" id="solicitudModal" tabindex="-1" role="dialog" aria-labelledby="solicitudModalLabel" aria-hidden="true">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h4 class="modal-title" id="solicitudModalLabel">Solicitar Servicio</h4>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModal">X</button>
          </div>
          <div class="modal-body">
           
            <form @submit.prevent="submitForm">
              <div class="form-group">
                <label for="nombre">Nombres:</label>
                <input type="text" class="form-control" id="nombre" placeholder="Describa nombre y apellido" v-model="formData.nombre">
              </div>
              <div class="form-group">
                <label for="correo">Correo:</label>
                <input type="email" class="form-control" placeholder="Escriba su correo" id="apellidos" v-model="formData.correo">
              </div>
              <div class="form-group">
                <label for="celular">Numero de celular:</label>
                <input type="text" class="form-control" id="celular" placeholder="eje +57 304 000 4445" v-model="formData.celular">
              </div>
              <div class="form-group">
                <label for="description">Descripcion:</label>
                <textarea id="description" placeholder="Describa su solicitud" rows="3" v-model="formData.description"></textarea>
              </div>
              <div class="form-group">
                <label for="person">Personal:</label>
                <input type="text" class="form-control" id="person" placeholder="Número de personas" v-model="formData.person">
              </div>
              <div class="form-group">
                <label for="entrada">Entrada:</label>
                <input type="date" class="form-control" id="entrada" v-model="formData.entrada" required>
              </div>
              <div class="form-group">
                <label for="salida">Salida:</label>
                <input type="date" class="form-control" id="salida" v-model="formData.salida" required>
              </div>
              <button type="submit" class="btn btn-success">Solicitar</button>
              <button type="button" class="btn btn-secondary" @click="cerrarModal">Cerrar</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import Swal from 'sweetalert2';

export default {
  data() {
    return {
      hoteles: [],
      customerRol: null,
      tipo: 'Hotel',
      formData: {
        nombre: '',
        celular: '',
        correo: '',
        description: '',
        person: '',
        entrada: '',
        salida: '',
        clienteId: null,
        negocioId: null,
        customer_id: null,
      },
    };
  },
  mounted() {
    this.customerRol = localStorage.getItem('usuarioRol');
    this.formData.clienteId = localStorage.getItem('customerId');
    this.formData.nombre = localStorage.getItem('nombreUsuario');
    this.formData.correo = localStorage.getItem('correoUsuario');
    this.formData.celular = localStorage.getItem('celularUsuario');
    this.formData.person = localStorage.getItem('person');
    this.formData.entrada = localStorage.getItem('entrada');
    this.formData.salida = localStorage.getItem('salida');
    this.fetchHoteles();
  },
  methods: {
    async fetchHoteles() {
      try {
        const response = await axios.get(`${process.env.API}/Hotel`);
        this.hoteles = response.data;
      } catch (error) {
        console.error('Error al obtener los hoteles', error);
      }
    },
    async submitForm() {
      try {
        const response = await axios.post(`${process.env.API}/pedirServicio?tipo=${this.tipo}`, this.formData);
        console.log(response);
        this.formData.customer_id = null;
        this.formData.negocioId = null;
        this.formData.description = '';
        this.formData.person = '';
        this.formData.entrada = '';
        this.formData.salida = '';
        
        // Mostrar la alerta de éxito
        await Swal.fire({
        title: "¡Bien hecho!",
        text: "¡Has enviado la solicitud con éxito!",
        icon: "success"
      });
       
         // Recargar la página después de enviar con éxito
      window.location.reload();
        
      } catch (error) {
        console.error('Error al enviar la solicitud', error);
      }
    },
    enviarIdGaleria(cid, id) {
      localStorage.setItem('customerNegocioId', cid);
      localStorage.setItem('negocioId', id);
      this.$router.push({ path: '/artegaleria' });
    },
    enviarcid(cid, id) {
      this.formData.customer_id = cid;
      this.formData.negocioId = id;
    },
    confirmarRegistro() {
      Swal.fire({
        icon: "error",
        title: "Oops...",
        text: "Si desea solicitar un servicio, debe registrarse.",
        footer: '<a href="https://api.whatsapp.com/send?phone=573125637781&text=Hola%2C%20me%20gustar%C3%ADa%20tomar%20un%20servicio%20desde%20la%20plataforma%20COSTA%20BRISA.">Why do I have this issue?</a>'
      }).then(() => {
        this.$router.push({
          path: '/resgister',
          name: 'Res',
          component: () => import('../views/RestroVew.vue') // Ajusta la ruta y componente según tu estructura
        });
      });
    },
    cerrarModal() {
      // Limpiar datos del formulario si es necesario
      this.formData = {
        nombre: '',
        celular: '',
        correo: '',
        description: '',
        person: '',
        entrada: '',
        salida: '',
        clienteId: null,
        negocioId: null,
        customer_id: null,
      };
      // Cerrar el modal de Bootstrap
      const modal = document.getElementById('solicitudModal');
      if (modal) {
        modal.classList.remove('show');
        modal.style.display = 'none';
        const modalBackdrop = document.getElementsByClassName('modal-backdrop');
        document.body.removeChild(modalBackdrop[0]);
      }
    }
  },
};
</script>






<style scoped>
.card {
  transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
}

.card:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}
</style>

<style scoped>
.card {
  transition: transform 0.3s ease-in-out, box-shadow 0.3s ease-in-out;
}

.card:hover {
  transform: scale(1.05);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  border-color: blue; /* Color del borde al pasar el cursor */
}

.card-img-top {
  width:auto; /* Ancho deseado */
  height:300px; /* Altura deseada */
  object-fit: cover; /* Para asegurar que la imagen se ajuste correctamente sin distorsión */
}
</style>
