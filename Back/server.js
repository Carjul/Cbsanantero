//mongodb+srv://hotelsapp:EgB0AQDudz9NP8eP@cluster0.zzgdpky.mongodb.net/
// server.js
const express = require('express');
const mongoose = require('mongoose');
const customerRouter = require('./models/router');

const app = express();
const PORT = process.env.PORT || 3000;

// Conéctate a MongoDB Atlas
mongoose.connect('mongodb+srv://hotelsapp:EgB0AQDudz9NP8eP@cluster0.zzgdpky.mongodb.net/', {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

// Verifica la conexión
const db = mongoose.connection;
db.on('error', console.error.bind(console, 'Error de conexión a MongoDB:'));
db.once('open', () => {
  console.log('Conexión exitosa a MongoDB');
});

// Usa el router para gestionar las rutas relacionadas con los clientes
app.use('/api', customerRouter);

// Resto de la configuración de tu servidor...

app.listen(PORT, () => {
  console.log(`Servidor Express escuchando en el puerto ${PORT}`);
});
