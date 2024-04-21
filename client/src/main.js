import { createApp } from 'vue'
import { createAuth0 } from '@auth0/auth0-vue';
import App from './App.vue'
import router from './router'
import Swal from 'sweetalert2';
const app = createApp(App);

app.use(
    createAuth0({
      domain: "dev-d1d2gby4uy86jijx.us.auth0.com",
      clientId: "wShFwTaMGgMnOdCH5kNHPmxSHIteYfcp",
      authorizationParams: {
        redirect_uri: window.location.origin
      }
    })
  );
app.use(router)
app.mount('#app')
/* addEventListener('keypress', function(e) {
  if(e.key == "Enter") {
    localStorage.clear();
    location.reload();
  }
}); */

addEventListener('DOMContentLoaded', function() {
  var x = this.localStorage.getItem('usuarioAutenticado');
  setTimeout(()=>{
    if(x == null) return;
    Swal.fire({
      title: "La sesión ha expirado",
      text: "Por favor iniciar sesión nuevamente",
      icon: "warning",
    }).then((result) => {
      if (result.isConfirmed) {
        localStorage.clear();
        location.reload();
      }
    });

  },600000);
  
  })
  

