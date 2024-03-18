import { h, render } from 'preact';
import Navbar from './components/navbar/navbar';
import Footer from './components/footer/footer';
import ObtenerAspirante from './components/obteneraspirante/obteneraspirante'; 
import './style.css';
import './components/obteneraspirante/obteneraspirante.module.css';
import 'bootstrap/dist/css/bootstrap.min.css';


const App = () => (
  <div id="app">
    <Navbar />
    
    <ObtenerAspirante id="1" />
    <Footer />
  </div>
);

render(<App />, document.getElementById('app'));
