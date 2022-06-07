import React, {useState} from 'react'
import {Link} from 'react-router-dom'
import './NavBar.css';
import image from  "./img/shopping-cart.png"

function NavBar() {
    const [click, setClick]=useState(false);
    const [setButton]=useState(true);
    const handleClick=()=>setClick(!click);
    const closeMobileMenu=()=>setClick(false);
    const showButton =()=> {
        if(window.innerWidth <=960){
            setButton(false);
        } else {
            setButton(true);
        }
    };
    window.addEventListener('resize', showButton);
  return (
    <>
    <nav className="NavBar">
<div className="navbar-container">
    <Link to="/home" className="navbar-logo">
        LIBRERIA
    </Link>
    <div className='menu-icon' onClick={handleClick}>
        <i className={click ? 'fas fa-times': 'fas fa-bars'}/>
    </div>
    <ul className={click ? 'nav-menu active' : 'nav-menu'}> 
         <li className='nav-item'>
             <Link to='/categories' className='nav-links' onClick={closeMobileMenu}>
                 Categorías
             </Link>
         </li>
         <li className='nav-item'>
             <Link to='/cart' className='nav-links' onClick={closeMobileMenu}>
                <img src={image} className="imagen"/>
             </Link>
         </li>
         
    </ul>
    
</div>
    </nav>
    </>
  )
}

export default NavBar
