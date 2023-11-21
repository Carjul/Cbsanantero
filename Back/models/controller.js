// models/controller.js
const Customer = require('./customer');

const getAllCustomers = async (req, res) => {
  try {
    const customers = await Customer.find();
    res.json(customers);
  } catch (error) {
    console.error(error);
    res.status(500).send('Error del servidor');
  }
};

const createCustomer = async (req, res) => {
  try {
    const newCustomer = new Customer(req.body);
    await newCustomer.save();
    res.json(newCustomer);
  } catch (error) {
    console.error(error);
    res.status(500).send('Error del servidor al crear un nuevo cliente');
  }
};

module.exports = {
  getAllCustomers,
  createCustomer,
};
