import { createApp } from 'vue'
import { createAuth0 } from '@auth0/auth0-vue';
import App from './App.vue'
import router from './router'

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