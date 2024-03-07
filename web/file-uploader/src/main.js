import {createApp} from 'vue'
import App from './App.vue'
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";


let app = createApp(App);
app.component('font-awesome-icon', FontAwesomeIcon);

app.mount('#app')
