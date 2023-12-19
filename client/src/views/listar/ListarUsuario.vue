<template>
  <div id="app">
    <div class="container">
      <br>
      <div class="card">
  <h5 class="card-header">Panel de Clientes</h5>
  <div class="card-body">
    <h5 class="card-title">Creacion de Usuario</h5>
    <p class="card-text">En este espacio creamos nuevos usuarorios para COSTA BRISA</p>
    <button class="btn btn-success mb-4" @click="abrirModalCrear">Crear</button>
  </div>
</div>

      
     
     
      <hr>
      <strong>Todos los clientes</strong>
      <hr>
      <div class="row">
        <div v-for="customer in customers" :key="customer._id" class="col-md-4 mb-4">
          <div class="card">
            <img :src="customer.image" class="card-img-top" alt="Profile Image">
            <div class="card-body">
              <h5 class="card-title">{{ customer.name }}</h5>
              <p class="card-text">Email: {{ customer.email }}</p>
              <p class="card-text">Contraseña: {{ customer.password }}</p>
              <p class="card-text">Address: {{ customer.address }}</p>
              <p class="card-text">Phone: {{ customer.phone }}</p>
              <p class="card-text">Role: {{ customer.rol }}</p>
                <!-- Switch -->
                <!-- Lever Toggle Switch -->
                <div class="form-check form-switch">
                  <input type="checkbox" :checked="customer.status === 'Activo' ? true : false" class="form-check-input position-relative"  id="leverSwitch" @change="statusCustomer($event, customer)">
                  <label class="form-check-label position-absolute" for="leverSwitch" >
                    {{ customer.status === 'Activo'? 'Active' : 'Inactive' }}
                  </label>
                </div>

              <button class="btn btn-danger" @click="eliminarCliente(customer._id)">Eliminar</button>
              <button class="btn btn-primary" @click="abrirModalActualizar(customer)">Actualizar</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal Crear -->
    <div ref="modalCrear" class="modal fade" id="modalCrear" tabindex="" role="dialog" aria-labelledby="modalCrearLabel" aria-hidden="true">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="modalCrearLabel">Crear Cliente</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalCrear">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="crearCliente">
  <div class="row">
    <div class="col-md-6">
      <div class="form-group">
        <label for="newName">Nombre:</label>
        <input type="text" v-model="nuevoCliente.name" class="form-control" id="newName" required>
      </div>
      <div class="form-group">

        <input class="form-control"  accept="image/jpeg, image/jpg, image/png" @change='uploadFile()' id="file" ref="file" type="file" s
                tyle="display:none">
        
      </div>
      <div class="form-group">
        <label for="newEmail">Email:</label>
        <input type="email" v-model="nuevoCliente.email" class="form-control" id="newEmail" required>
      </div>
      <div class="form-group">
        <label for="newEmail">Contraseña:</label>
        <input type="text" v-model="nuevoCliente.password" class="form-control" id="newEmail" required>
      </div>
    </div>
    <div class="col-md-6">
      <div class="form-group">
        <label for="newAddress">Dirección:</label>
        <input type="text" v-model="nuevoCliente.address" class="form-control" id="newAddress" required>
      </div>
      <div class="form-group">
        <label for="newPhone">Teléfono:</label>
        <input type="text" v-model="nuevoCliente.phone" class="form-control" id="newPhone" required>
      </div>

      <div class="form-group">
        <label for="newRol">
          Rol:</label>
        <input type="text" v-model="nuevoCliente.rol" class="form-control" id="newRol" required>
      </div>

      

    </div>
  </div>
  <button type="submit" class="btn btn-primary">Crear Cliente</button>
</form>

          </div>
        </div>
      </div>
    </div>


    <!-- Modal Actualizar -->
    <div ref="modalActualizar" class="modal fade" id="modalActualizar" tabindex="-1" role="dialog" aria-labelledby="modalActualizarLabel" aria-hidden="true">
      <div class="modal-dialog" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <h5 class="modal-title" id="modalActualizarLabel">Actualizar Cliente</h5>
            <button type="button" class="close" data-dismiss="modal" aria-label="Close" @click="cerrarModalActualizar">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <form @submit.prevent="actualizarCliente">
              <div class="form-group">
                <label for="name">Nombre:</label>
                <input type="text" v-model="clienteActualizado.name" class="form-control" id="name" required>
              </div>
              <div class="form-group">
                <label for="email">Email:</label>
                <input type="email" v-model="clienteActualizado.email" class="form-control" id="email" required>
              </div>
              <div class="form-group">
                <label for="email">Contraseña:</label>
                <input type="text" v-model="clienteActualizado.password" class="form-control" id="email" required>
              </div>
              <div class="form-group">
                <label for="address">Address:</label>
                <input type="text" v-model="clienteActualizado.address" class="form-control" id="address" required>
              </div>
              <div class="form-group">
                <label for="phone">Phone:</label>
                <input type="text" v-model="clienteActualizado.phone" class="form-control" id="phone" required>
              </div>
              <div class="form-group">
                <label for="role">Role:</label>
                <input type="text" v-model="clienteActualizado.rol" class="form-control" id="role" required>
              </div>
              <button type="submit" @click="ActualizarCustomer(clienteActualizado,clienteActualizado._id)"  class="btn btn-primary">Guardar Cambios</button>
            </form>
          </div>
          <div class="modal-footer">
            <p v-if="successMessage" class="text-success">{{ successMessage }}</p>
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
      customers: [],
      clienteActualizado: {
        _id: null,
        name: '',
        email: '',
        password: '',
        address: '',
        phone: '',
        rol: '',
      },
      file:null,
      nuevoCliente: {
        name: '',
        image: '',
        email: '',
        password:'',
        address: '',
        phone: '',
        rol: null,
      },
    };
  },
  mounted() {
    this.fetchCustomers();

  },
  methods: {
    async fetchCustomers() {
      let id = localStorage.getItem('customerId')
      try {
        const response = await axios.get('http://localhost:3000/Customer');
        const Filtro = response.data.filter((e)=>e._id!==id)
        console.log(response.data)
        console.log(Filtro)
        this.customers = Filtro;
      } catch (error) {
        console.error('Error fetching customers:', error);
      }
    },
    async statusCustomer(e,c) {
  

      try {
  const res = await axios.put(`http://localhost:3000/customerStatus/${c._id}`, {
    status: e.target.checked
  });
  console.log(res.data);
} catch (error) {
  if (error.response) {
    // El servidor respondió con un código de estado fuera del rango 2xx
    console.error('Error de respuesta:', error.response.data);
    console.error('Código de estado HTTP:', error.response.status);
  } else if (error.request) {
    // La solicitud fue realizada pero no se recibió respuesta
    console.error('No se recibió respuesta del servidor:', error.request);
  } else {
    // Se produjo un error al configurar la solicitud
    console.error('Error al configurar la solicitud:', error.message);
  }
}
      this.fetchCustomers(); 
    },
    async eliminarCliente(id) {
      try {
        await axios.delete(`http://localhost:3000/Customer/${id}`);
        this.fetchCustomers();
      } catch (error) {
        console.error('Error deleting customer:', error);
      }
    },
    abrirModalActualizar(customer) {
  this.clienteActualizado =  { ...customer };
  const modalElement = this.$refs.modalActualizar;
  if (modalElement) {
    modalElement.classList.add('show');
    modalElement.style.display = 'block';
  }

    },

   async ActualizarCustomer(customer,id){
  const response= await axios.put(`http://localhost:3000/Customer/${id}`,customer);
  console.log(response)
  this.fetchCustomers();
   },
   uploadFile() {
      this.file = this.$refs.file.files[0];
      var sizeByte = this.$refs.file.files[0].size;
      var siezekiloByte = parseInt(sizeByte / 1024)
      if (siezekiloByte > 2048) {
       console.log("la imagen es muy grande")
      }
      else{
        this.createImage(this.file)
      }
        
},
  createImage(file) {

var reader = new FileReader();
reader.onload = (e) => {
  this.nuevoCliente.image = e.target.result;
};
reader.readAsDataURL(file);
},

    cerrarModalActualizar() {
      const modalElement = this.$refs.modalActualizar;
      if (modalElement) {
        modalElement.classList.remove('show');
        modalElement.style.display = 'none';
      }
    },
    abrirModalCrear() {
      const modalElement = this.$refs.modalCrear;
      if (modalElement) {
        modalElement.classList.add('show');
        modalElement.style.display = 'block';
      }
    },
    cerrarModalCrear() {
      const modalElement = this.$refs.modalCrear;
      if (modalElement) {
        modalElement.classList.remove('show');
        modalElement.style.display = 'none';
      }
    },
    async crearCliente() {
      try {
        const response = await axios.post('http://localhost:3000/Customer', this.nuevoCliente);
        console.log(response.data);
        this.fetchCustomers();

        // Reset the form and close the modal
        this.nuevoCliente = {
          name: '',
          image: '',
          email: '',
          password:'',
          address: '',
          phone: '',
          rol: '',
        };

        // Close the modal
        this.cerrarModalCrear();
      } catch (error) {
        console.error('Error creating customer:', error);
      }
    },
  },
};
</script>
