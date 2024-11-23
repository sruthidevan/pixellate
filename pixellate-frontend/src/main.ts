import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import AOS from "aos";
import "aos/dist/aos.css";

(AOS as any).init();
import './index.css'; // Import Tailwind styles

const app = createApp(App);
app.use(router);
app.mount('#app');
