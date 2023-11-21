const mongoose = require('mongoose');

const customerSchema = new mongoose.Schema({
  username: String,
  name: String,
  address: String,
  birthdate: Date,
  email: String,
  active: Boolean,
  accounts: [String],
  tier_and_details: Object,
});

const Customer = mongoose.model('Customer', customerSchema);

module.exports = Customer;
