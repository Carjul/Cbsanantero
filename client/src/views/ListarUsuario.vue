<template>
    <div id="app">
      <div class="container">
        <h1>Listado de Clientes</h1>
        <div class="row">
          <div v-for="customer in customers" :key="customer._id" class="col-md-4 mb-4">
            <div class="card">
              <img :src="customer.image" class="card-img-top" alt="Profile Image">
              <div class="card-body">
                <h5 class="card-title">{{ customer.name }}</h5>
                <p class="card-text">Email: {{ customer.email }}</p>
                <p class="card-text">Address: {{ customer.address }}</p>
                <p class="card-text">Phone: {{ customer.phone }}</p>
                <p class="card-text">Role: {{ customer.rol }}</p>
                <button class="btn btn-danger" @click="eliminarCliente(customer._id)">Eliminar</button>
                <button class="btn btn-primary" @click="abrirModalActualizar(customer)">Actualizar</button>
              </div>
            </div>
          </div>
        </div>
      </div>
  
      <!-- Modal Actualizar -->
      <div class="modal fade" id="modalActualizar" tabindex="-1" role="dialog" aria-labelledby="modalActualizarLabel" aria-hidden="true">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="modalActualizarLabel">Actualizar Cliente</h5>
              <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="modal-body">
              <!-- Formulario para actualizar los datos del cliente -->
              <form @submit.prevent="actualizarCliente">
                <div class="form-group">
                  <label for="name">Nombre:</label>
                  <input type="text" v-model="clienteActualizado.name" class="form-control" id="name" required>
                </div>
                <div class="form-group">
                  <label for="email">Email:</label>
                  <input type="email" v-model="clienteActualizado.email" class="form-control" id="email" required>
                </div>
                <!-- Agrega más campos según tus necesidades -->
  
                <button type="submit" class="btn btn-primary">Guardar Cambios</button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
 /* eslint-disable */
import axios from 'axios';
const $ = window.jQuery; // Asigna jQuery a la variable $ si estás utilizando jQuery

export default {
  data() {
    return {
      customers: [],
      clienteActualizado: {
        _id: null,
        name: '',
        email: '',
        // Agrega más propiedades según tus necesidades
      },
    };
  },
  mounted() {
    this.fetchCustomers();
  },
  methods: {
    async fetchCustomers() {
      try {
        const response = await axios.get('http://localhost:3000/Customer');
        this.customers = response.data;
        console.log('Customers:', this.customers);
      } catch (error) {
        console.error('Error fetching customers:', error);
      }
    },
    async eliminarCliente(id) {
      // Lógica para eliminar el cliente con el ID proporcionado
    },
    abrirModalActualizar(customer) {
      // Asigna los datos del cliente al objeto clienteActualizado
      this.clienteActualizado = { ...customer };
      // Abre el modal
      if ($) $('#modalActualizar').modal('show');
    },
    async actualizarCliente() {
      // Lógica para actualizar los datos del cliente
      // Utiliza this.clienteActualizado para acceder a los datos actualizados
      // Cierra el modal después de la actualización
      if ($) $('#modalActualizar').modal('hide');
    },
  },
};

  </script>
  
  <style scoped>
  /* Puedes agregar estilos específicos aquí si es necesario */
  </style>
  