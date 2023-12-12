<template>
    <div id="app">
      <div class="sidebar">
        <a href="javascript:void(0)" @click="mostrarContenido('contenido1')">Sección 1</a>
        <a href="javascript:void(0)" @click="mostrarContenido('contenido2')">Sección 2</a>
        <a href="javascript:void(0)" @click="mostrarContenido('contenido3')">Sección 3</a>
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
          { label: 'contenido1', content: { title: 'Contenido de la Sección 1', text: 'Este es el contenido de la Sección 1.' } },
          { label: 'contenido2', content: { title: 'Contenido de la Sección 2', text: 'Este es el contenido de la Sección 2.' } },
          { label: 'contenido3', content: { title: 'Contenido de la Sección 3', text: 'Este es el contenido de la Sección 3.' } }
        ],
        contenido: { title: 'Contenido Principal', text: 'Selecciona una sección en la barra lateral para ver el contenido correspondiente.' }
      };
    },
    methods: {
      mostrarContenido(contentKey) {
        const selectedContent = this.sections.find(section => section.label === contentKey);
        if (selectedContent) {
          this.contenido = { ...selectedContent.content };
        }
      }
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
    background-color: #0680878c;
    padding: 2px;
  }
  
  .sidebar a {
    padding: 8px 8px 8px 32px;
    text-decoration: none;
    font-size: 18px;
    color: #818181;
    display: block;
  }
  
  .sidebar a:hover {
    color: #f1f1f1;
  }
  
  .content {
    padding: 16px;
  }
  </style>
  