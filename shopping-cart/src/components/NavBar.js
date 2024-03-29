import React, {useState} from 'react'
import {Link} from 'react-router-dom'
import './NavBar.css';
import image from  "./shopping-cart.png"

import Cookies from 'universal-cookie';



const cookies = new Cookies();

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
var username=  cookies.get("name");

window.addEventListener('resize', showButton);
var username=  cookies.get("name");

        return (
        <>
        <nav className="NavBar" >
        <div className="navbar-container">
        <Link to="/home" className="navbar-logo">
        LIBRERIA
        </Link>
        <div className='menu-icon' onClick={handleClick}>
        <i className={click ? 'fas fa-times': 'fas fa-bars'}/>
        </div>
        <ul id="lista" className={click ? 'nav-menu active' : 'nav-menu'}> 

        <li className='nav-item'>
        <Link to='/cart' className='nav-links' onClick={closeMobileMenu}>
        <img src={image} className="imagen"/>
        </Link>
        </li>
        <li className='nav-item'>
        <Link to='/account' className='nav-links' onClick={closeMobileMenu}>
        Mi Cuenta
        </Link>
        </li>
        <li className='nav-item'>
        <Link to='/' className='nav-links' onClick={closeMobileMenu}>
        Cerrar Sesión
        </Link>
        </li>
       
        </ul>
        <div className="logInfo" >

        Bienvenid@, {username}!


        </div>

        </div>
        </nav>
        </>
        )
}

export default NavBar
