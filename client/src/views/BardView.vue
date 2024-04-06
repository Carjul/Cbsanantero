<template>
  <div class="container">
    <hr>
    <div class="row">
      <div class="col-md-4" v-for="bar in bares" :key="bar.id">
        <div class="card mb-4">
          <img @click="enviarIdGaleria(bar._id)" :src="bar.image" class="card-img-top" alt="Bar" />
          <div class="card-body">
            <h5 class="card-title">{{ bar.name }}</h5>
            <p class="card-text">{{ bar.description }}</p>
            <p class="card-text"><strong>Dirección:</strong> {{ bar.address }}</p>
            <div v-if="customerRol === 'Cliente'">
              <button class="btn btn-primary" data-toggle="modal" data-target="#solicitudModal"
                @click="enviarcid(bar.customer_id, bar._id)">Obtener Servicio</button>
            </div>

          </div>
        </div>
      </div>
    </div>

    <!-- Modal -->
    <div class="modal" id="solicitudModal" tabindex="-1" role="dialog" aria-labelledby="solicitudModalLabel"
      aria-hidden="true" >
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <!-- Modal Header -->
          <div class="modal-header">
            <h4 class="modal-title" id="solicitudModalLabel">Solicitar Servicio</h4>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close" ></button>
          </div>

          <!-- Modal Body -->
          <div class="modal-body">
            <b>
              <p>Creando servicio:</p>
            </b>

            <!-- Formulario con campos solicitados -->
            <form>
              <div class="form-group">
                <label for="nombre">Nombres:</label>
                <input type="text" class="form-control" id="nombre" placeholder="Describa nombre y apellido"
                  v-model="formData.nombre" readonly>
              </div>

              <div class="form-group">
                <label for="correo">Correo:</label>
                <input type="email" class="form-control" placeholder="Escriba su correo" id="apellidos"
                  v-model="formData.correo" readonly>
              </div>

              <div class="form-group">
                <label for="celular">Numero de celular:</label>
                <input type="text" class="form-control" id="celular" placeholder="eje +57 304 000 4445"
                  v-model="formData.celular" readonly>
              </div>

              <div class="form-group">
                <label for="description">Descripcion:</label>
                <textarea id="description" placeholder="Describa su solicitud" rows="3"
                  v-model="formData.description"></textarea>
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
      bares: [],
      customerRol: '',
      tipo: 'Bar',
      formData: {
        nombre: '',
        celular: '',
        correo: '',
        description: '',
        clienteId: null,
        customerId: null,
        negocioId: null
      },
    };
  },
  mounted() {
    this.customerRol = localStorage.getItem('usuarioRol');

    this.formData.clienteId = localStorage.getItem('customerId')
    this.formData.nombre = localStorage.getItem('nombreUsuario');
    this.formData.correo = localStorage.getItem('correoUsuario');
    this.formData.celular = localStorage.getItem('celularUsuario');

    this.fetchBares();
  },
  methods: {
    async fetchBares() {
      try {
        const response = await axios.get(`${process.env.API}/Bar`);
        this.bares = response.data;
      } catch (error) {
        console.error('Error al obtener los bares', error);
      }
    },
    enviarIdGaleria(id) {
      localStorage.setItem('negocioId', id);
      this.$router.push({ path: '/artegaleria' });
    },
    enviarcid(cid, id) {
      this.formData.customerId = cid;
      this.formData.negocioId = id;
    },


    async submitForm() {
      try {
        const response = await axios.post(`${process.env.API}/pedirServicio?tipo=${this.tipo}`, this.formData);
        console.log(response);
        this.formData.customerId = null;
        this.formData.negocioId = null;
        this.formData.description = '';
    
      } catch (error) {
        console.error('Error al enviar la solicitud', error);
      }



    },
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
  border-color: blue; /* Color del borde al pasar el cursor */
}

.card-img-top {
  width:auto; /* Ancho deseado */
  height:300px; /* Altura deseada */
  object-fit: cover; /* Para asegurar que la imagen se ajuste correctamente sin distorsión */
}
</style>
