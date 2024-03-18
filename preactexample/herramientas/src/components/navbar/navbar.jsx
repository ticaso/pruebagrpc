import { h } from 'preact';
import styles from './navbar.module.css';

const Navbar = () => {
  return (
    <nav className={`navbar navbar-expand-lg ${styles.navbar}`}>
      <div className="container-fluid">
        <p className={styles.navLink}>Portafolio de carrera</p>
      </div>
    </nav>
  );
};

export default Navbar;
