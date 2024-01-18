import { createApp } from 'vue'
import App from './App.vue'
import BaseCard from './components/ui/BaseCard.vue'
import './assets/index.css'

const app = createApp(App);
app.component('base-card', BaseCard);

app.mount('#app');