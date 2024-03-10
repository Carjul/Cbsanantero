import { createApp } from 'vue'
import { createAuth0 } from '@auth0/auth0-vue';
import App from './App.vue'
import router from './router'
import Swal from 'sweetalert2';
const app = createApp(App);

app.use(
    createAuth0({
      domain: "dev-35dy5wz22m0uql2f.us.auth0.com",
      clientId: "iiFhow7XesbA4pFa9u6Rqiu8VKz1mgE1",
      authorizationParams: {
        redirect_uri: window.location.origin
      }
    })
  );
app.use(router)
app.mount('#app')
addEventListener('keypress', function(e) {
  if(e.key == "Enter") {
    localStorage.clear();
    location.reload();
  }
});

addEventListener('DOMContentLoaded', function() {
  var x = this.localStorage.getItem('usuarioAutenticado');

  for (let i = 0; i <= 599; i++) {
    setTimeout(() => {
  
      if (i === 599 && x) {
        Swal.fire({
          title: "La sesión ha expirado",
          text: "Favor iniciar sesión nuevamente",
          icon: "warning",
        });
  
        setTimeout(() => {
          localStorage.clear();
          location.reload();
        }, 1000);
      }
    }, 60000);
  }
  


})