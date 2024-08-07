import { Link, NavLink } from "@remix-run/react";
import styles from "./Nav.module.css";

function Nav() {
  return (
    <div className={styles.wrapper}>
      <div className={styles.logo}>
        <Link to="/">Chatty</Link>
      </div>
      <nav>
        <ul className={styles.navlinks}>
          <li className={styles.navlink}>
            <NavLink to="/login">Login</NavLink>
          </li>
        </ul>
      </nav>
    </div>
  );
}

export default Nav;
