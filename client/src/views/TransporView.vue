<template>
  <div class="container">
    <hr>
    <div class="row">
      <div
        class="col-md-4"
        v-for="trasporte in transportes"
        :key="trasporte.id"
      >
        <div class="card mb-4">
          <img
            @click="enviarIdGaleria(trasporte.customer_id, trasporte._id)"
            :src="trasporte.image"
            class="card-img-top"
            alt="Trasporte"
          />
          <div class="card-body">
            <hr>
            <p class="h3 text-decoration-none">{{ trasporte.name }}</p>
            <hr>
            <!-- <p class="h5 text-decoration-none">{{ trasporte.description }}</p> -->
            <p class="card-text"></p>
            <div class="d-flex justify-content-between">
              <h5 class="m-0"><i class="fa fa-map-marker-alt text-primary mr-2"></i>{{ trasporte.address }}</h5>
              <h5 class="m-0"><i class="fas fa-coins text-primary mr-2"></i>{{ trasporte.price }}</h5>
            </div>
            <hr>
            <div v-if="customerRol === null">
              <router-link to="/register">
                <button class="btn btn-primary" data-toggle="modal" data-target="#solicitudModal">Solicitar servicio</button>
              </router-link>
            </div>
            <div v-if="customerRol === 'Cliente'">
              <button class="btn btn-primary" data-toggle="modal" data-target="#solicitudModal" @click="enviarcid(trasporte.customer_id, trasporte._id)">Obtener Servicio</button>
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
            <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="closeModal"></button>
          </div>
          <div class="modal-body">
            <b>
              <p>Creando servicio:</p>
            </b>
            <form>
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
              <button type="button" class="btn btn-success" @click="submitForm">Solicitar</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      transportes: [],
      customerRol: null,
      tipo: 'Trasporte',
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
    this.fetchTransportes();
  },
  methods: {
    async fetchTransportes() {
      try {
        const response = await axios.get(`${process.env.API}/Trasporte`);
        this.transportes = response.data;
      } catch (error) {
        console.error('Error al obtener los trasportes', error);
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
    closeModal() {
      this.formData = {
        _id: null,
        nombre: '',
        celular: '',
        correo: '',
        description: '',
        person: '',
        entrada: '',
        salida: '',
        customerId: null,
      };
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
