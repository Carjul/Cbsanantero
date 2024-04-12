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
              <p class="card-text">Peticiones: {{ 0}}</p>


              <p v-if="customer.rol === 'Vendedor'" class="card-text">Tipo Negocio: {{ customer.tipo_negocio }}</p>
              <!-- Switch -->
              <!-- Lever Toggle Switch -->
              <div class="form-check form-switch">
                <input type="checkbox" :checked="customer.status === 'Activo' ? true : false"
                  class="form-check-input position-relative" id="leverSwitch"
                  @change="statusCustomer($event, customer)">
                <label class="form-check-label position-absolute" for="leverSwitch">
                  {{ customer.status === 'Activo' ? 'Active' : 'Inactive' }}
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
    <div ref="modalCrear" class="modal fade" id="modalCrear" tabindex="" role="dialog" aria-labelledby="modalCrearLabel"
      aria-hidden="true">
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

                    <input class="form-control" accept="image/jpeg, image/jpg, image/png" @change='uploadFile' id="file"
                      ref="file" type="file" s tyle="display:none">

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
                      <span>Rol:</span>
                    </label>
                    <select v-model="nuevoCliente.rol" class="form-control" id="newRol" required>
                      <option value="Admin">Admin</option>
                      <option value="Cliente">Cliente</option>
                      <option value="Vendedor">Vendedor</option>
                    
                    </select>
                  </div>
                  <div class="form-group" v-if="nuevoCliente.rol === 'Vendedor'">
                    <label for="newRol">
                      <span>Tipo Negocio</span>
                    </label>
                    <select v-model="nuevoCliente.tipo_negocio" class="form-control" id="newRol" required>
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

                  


                </div>
              </div>
              <button type="submit" class="btn btn-primary">Crear Cliente</button>
            </form>

          </div>
        </div>
      </div>
    </div>

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
                <form @submit.prevent="actualizarCliente" class="row">
                    <div class="form-group col-md-6">
                        <label for="name">Nombre:</label>
                        <input type="text" v-model="clienteActualizado.name" class="form-control" id="name" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="email">Email:</label>
                        <input type="email" v-model="clienteActualizado.email" class="form-control" id="email" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="password">Contraseña:</label>
                        <input type="password" v-model="clienteActualizado.password" class="form-control" id="password" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="address">Address:</label>
                        <input type="text" v-model="clienteActualizado.address" class="form-control" id="address" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="phone">Phone:</label>
                        <input type="text" v-model="clienteActualizado.phone" class="form-control" id="phone" required>
                    </div>
                    <div class="form-group col-md-6">
                        <label for="newRol">Rol:</label>
                        <select v-model="clienteActualizado.rol" class="form-control" id="newRol" required>
                            <option value="Admin">Admin</option>
                            <option value="Cliente">Cliente</option>
                            <option value="Vendedor">Vendedor</option>
                        </select>
                    </div>
                    <div class="form-group col-md-6" v-if="clienteActualizado.rol === 'Vendedor'">
                        <label for="tipo_negocio">Tipo Negocio</label>
                        <select v-model="clienteActualizado.tipo_negocio" class="form-control" id="tipo_negocio" required>
                            <option value="Transporte">Transporte</option>
                            <option value="Hospedaje">Hospedaje</option>
                            <option value="Tour">Tour</option>
                            <option value="Hoteles">Hoteles</option>
                            <option value="Artesanías">Artesanías</option>
                            <option value="Bares">Bares y discotecas</option>
                            <option value="Restaurantes">Restaurantes</option>
                            <option value="Juegos">Recreación</option>
                        </select>
                    </div>
                    <div class="form-group col-12">
                        <button type="submit" @click="ActualizarCustomer(clienteActualizado, clienteActualizado._id)" class="btn btn-primary">Guardar Cambios</button>
                    </div>
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
        tipo_negocio:''
        
      },
      file: null,
      nuevoCliente: {
        name: '',
        image: '',
        email: '',
        password: '',
        address: '',
        phone: '',
        rol: null,
        tipo_negocio:null
      },
    };
  },
  mounted() {
    this.fetchServicio();
    this.fetchCustomers();
  },
  methods: {
    async fetchServicio(cid) {
      
      try {
        const response = await axios.get(`${process.env.API}/pedirServicioc/${cid}`);
      
        console.log(response.data);
       return response.data.length

  
      } catch (error) {
        console.error('Error fetching customers:', error);
      }
    },
    async fetchCustomers() {
      let id = localStorage.getItem('customerId');
      try {
        const response = await axios.get(`${process.env.API}/Customer`);
        const Filtro = response.data.filter((e) => e._id !== id);
        console.log(response.data);
        console.log(Filtro);
        this.customers = Filtro;
      } catch (error) {
        console.error('Error fetching customers:', error);
      }
    },
    async statusCustomer(e, c) {
      try {
        const res = await axios.put(`${process.env.API}/customerStatus/${c._id}`, {
          status: e.target.checked,
        });
        console.log(res.data);
      } catch (error) {
        if (error.response) {
          console.error('Error de respuesta:', error.response.data);
          console.error('Código de estado HTTP:', error.response.status);
        } else if (error.request) {
          console.error('No se recibió respuesta del servidor:', error.request);
        } else {
          console.error('Error al configurar la solicitud:', error.message);
        }
      }
      this.fetchCustomers();
    },
    async eliminarCliente(id) {
      try {
        await axios.delete(`${process.env.API}/Customer/${id}`);
        this.fetchCustomers();
      } catch (error) {
        console.error('Error deleting customer:', error);
      }
    },
    abrirModalActualizar(customer) {
      this.clienteActualizado = { ...customer };
      const modalElement = this.$refs.modalActualizar;
      if (modalElement) {
        modalElement.classList.add('show');
        modalElement.style.display = 'block';
      }
    },
    async ActualizarCustomer(customer, id) {
      try {
       
        const response = await axios.put(`${process.env.API}/Customer/${id}`, customer);
        console.log(response);
        this.fetchCustomers();
        this.cerrarModalActualizar();
      } catch (error) {
        console.error('Error updating customer:', error);
      }
    },
    uploadFile(event) {
      this.file = event.target.files[0];
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
        /*  console.log('Nuevo Cliente antes de enviar:', this.nuevoCliente); */
        const formData = new FormData();
        formData.append('name', this.nuevoCliente.name);
        formData.append('image', this.file);
        formData.append('email', this.nuevoCliente.email);
        formData.append('password', this.nuevoCliente.password);
        formData.append('address', this.nuevoCliente.address);
        formData.append('phone', this.nuevoCliente.phone);
        formData.append('rol', this.nuevoCliente.rol);
        formData.append('tipo_negocio', this.nuevoCliente.tipo_negocio);
        console.log('Nuevo Cliente:',this.nuevoCliente);
        const response = await axios.post(`${process.env.API}/Customer`, formData);

        console.log(response.data);
        this.fetchCustomers();

        // Resetear el formulario y cerrar el modal
        this.nuevoCliente = {
          name: '',
          image: '',
          email: '',
          password: '',
          address: '',
          phone: '',
          rol: '',
        };

        this.cerrarModalCrear();
      } catch (error) {
        console.error('Error creando cliente:', error);
      }
    },
  },
};

</script>