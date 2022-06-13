import React from 'react';
import Cookies from 'universal-cookie';
 
import "./Checkout.css"
 
const cookies = new Cookies();
function Checkout(){
    

            return(
            <div className="checkout-message">

            <div id="head">¡Felicidades, {cookies.get("username")}!</div>
            <div id="sub">Tu compra de ${cookies.get("total")} se realizó con éxito</div>
            <div id="msg">Pronto recibirás tu pedido</div>

            </div>



)
}
export default Checkout;