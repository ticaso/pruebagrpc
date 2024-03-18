import { h, Fragment } from 'preact';
import styles from './obteneraspirante.module.css';

import { useState, useEffect } from 'preact/hooks';
const ObtenerAspirante = ({ id }) => {
  const [aspirante, setAspirante] = useState(null);
  const [error, setError] = useState('');

  useEffect(() => {
    fetch(`http://localhost:8080/aspirante/${id}`)
      .then(response => {
        if (!response.ok) {
          throw new Error('No se pudo obtener la información del aspirante');
        }
        return response.json();
      })
      .then(data => {
        console.log("Datos recibidos para establecer en el estado:", data);
        setAspirante(data.aspirante);
      })
      .catch(error => {
        console.error('There has been a problem with your fetch operation:', error);
        setError(error.toString());
      });
  }, [id]);

  return (
    <Fragment>
      {error && <div className="alert alert-danger" role="alert">{error}</div>}
      {!aspirante && !error && <div className="spinner-border text-primary" role="status"><span className="visually-hidden">Cargando...</span></div>}
      {aspirante && (
        <div className="container mt-4 d-flex justify-content-center">

        <div className={styles.card}>
            <img src={aspirante.imagen_principal} className={styles.perfil} alt={`Foto de perfil de ${aspirante.nombre}`} />
            <div className={styles.container}>
              <h5 className={styles.title}>{aspirante.nombre} {aspirante.apellido}</h5>
              <p className={styles.carrera}>{aspirante.carrera}</p>
              <p className={styles.lenguajes}>
                <small className="text-muted">He trabajado en proyectos desarrollados en: {aspirante.lenguajes_programacion.join(', ')}</small>
              </p>
              <p className={styles.favorito}>Lenguaje de programación favorito:</p>
              <img 
                src={aspirante.imagen_secundaria} 
                className={styles.imagenfavorito} 
                alt={`Lenguaje de programación favorito de ${aspirante.nombre}`} 
                
              /> 
            </div>
          </div>
        </div>
      )}
    </Fragment>
  );
  
};

export default ObtenerAspirante;
