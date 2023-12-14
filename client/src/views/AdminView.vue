<template>
  <div id="app">
    <div class="sidebar">
      <div v-for="(section, index) in sections" :key="index">
        <a href="javascript:void(0)" @click="mostrarContenido(section.label)">
          {{ section.label }}
        </a>
        <div v-if="section.expanded">
          <ul>
            <li><a href="javascript:void(0)" @click="mostrarSubseccion(section.label, 'crear')">Crear</a></li>
            <li><a href="javascript:void(0)" @click="mostrarSubseccion(section.label, 'listar')">Listar</a></li>
          </ul>
        </div>
      </div>
    </div>

    <!-- Contenido principal -->
    <div class="content" :style="{ 'margin-left': sidebarWidth + 'px' }">
      <h2>{{ contenido.title }}</h2>
      <p v-html="contenido.text"></p>
    </div>
  </div>
</template>

<script>

export default {
  data() {
    return {
      sidebarWidth: 260,
      sections: [
        {
          label: 'Transporte',
          content: { title: 'Contenido de Transporte', text: 'Este es el contenido de la sección Transporte.' },
          expanded: false
        },
        {
          label: 'Usuarios',
          content: { title: 'Contenido de Usuarios', text: 'Este es el contenido de la sección Usuarios.' },
          expanded: false
        },
        {
          label: 'Hospedaje',
          content: { title: 'Contenido de Hospedaje', text: 'Este es el contenido de la sección Hospedaje.' },
          expanded: false
        },
        {
          label: 'Tour',
          content: { title: 'Contenido de Tour', text: 'Este es el contenido de la sección Tour.' },
          expanded: false
        },
        {
          label: 'Hoteles',
          content: { title: 'Contenido de Hoteles', text: 'Este es el contenido de la sección Hoteles.' },
          expanded: false
        },
        {
          label: 'Restaurantes',
          content: { title: 'Contenido de Restaurantes', text: 'Este es el contenido de la sección Restaurantes.' },
          expanded: false
        },
        {
          label: 'Recreación',
          content: { title: 'Contenido de Recreación', text: 'Este es el contenido de la sección Recreación.' },
          expanded: false
        },
        {
          label: 'Bares y discotecas',
          content: { title: 'Contenido de Bares y discotecas', text: 'Este es el contenido de la sección Bares y discotecas.' },
          expanded: false
        },
        {
          label: 'Artesanías',
          content: { title: 'Contenido de Artesanías', text: 'Este es el contenido de la sección Artesanías.' },
          expanded: false
        }
      ],
      contenido: { title: 'Contenido Principal', text: 'Selecciona una sección en la barra lateral para ver el contenido correspondiente.' }
    };
  },
  methods: {
    mostrarContenido(sectionKey) {
      const selectedSection = this.sections.find(section => section.label === sectionKey);

      if (selectedSection) {
        if (sectionKey === 'Usuarios') {
          // Si la sección es 'Usuarios', muestra la vista ListarUsuario.vue
          this.$router.push({ path: '/usuario' });
        } else {
          // Para otras secciones, actualiza el contenido normalmente
          this.contenido = { ...selectedSection.content };
        }

        // Cerrar las demás secciones
        this.sections.forEach(section => (section.expanded = section.label === sectionKey));
      }
    },

    // ... (otros métodos)
  }
};
</script>

<style>
/* Ajustes de estilo para la barra lateral */
.sidebar {
  height: 70vh;
  width: 250px;
  position: fixed;
  z-index: 4;
  top: 4;
  left: 0;
  background-color: #111;
  padding: 2px;
}

.sidebar a {
  padding: 8px 8px 8px 32px;
  text-decoration: none;
  font-size: 18px;
  color: #818181;
  display: block;
  margin-bottom: 5px; /* Ajusta la distancia entre las etiquetas */
}

.sidebar a:hover {
  color: #f1f1f1;
}

.sidebar div {
  margin-bottom: -15px;
}

.sidebar ul {
  list-style-type: none;
  padding-left: 15px;
  margin: 0;
}

.sidebar li {
  margin-bottom: 5px;
}

.content {
  padding: 16px;
}
</style>
