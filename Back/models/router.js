// models/router.js
const controller = require('./controller'); // Corrige la ruta del require
const express = require('express');
const bodyParser = require('body-parser');

const router = express.Router();
router.use(bodyParser.json());

// Ruta para obtener todos los clientes
router.get('/customers', controller.getAllCustomers);

// Ruta para crear un nuevo cliente
router.post('/customers', controller.createCustomer);

module.exports = router;
