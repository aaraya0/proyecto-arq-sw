import React from 'react';
import Cookies from 'universal-cookie';
 
import "./Checkout.css"
 
const cookies = new Cookies();

async function crearOrden(){
    
    let orderDetails=[]
    let orderDetail={'product_id':0, 'quantity':0}

    let current= cookies.get('cart')
    let items=[]
    let a=current.split(',')
    for (let i = 0; i < a.length; i++){
    let item = a[i];
    let prod= item.split('=')
    
    let prodID=prod[0]
    let cant= prod[1]
    if (prodID!=0 && cant!="undefined" && cant>0){
    orderDetail ={'product_id':prodID,'quantity':cant}
    orderDetails.push(orderDetail)
   // let product = await getProductById(prodID)
    //product.quantity = cant;
//items.push(product)
}
}
alert (orderDetails)
return orderDetails
}

async function postOrden (){
    let user_id = cookies.get("user_id")
    let orderDetails = crearOrden()
    var url = 'http://localhost:8090/order';
var data = {'user_id':user_id, 'orderDetail':orderDetails};

fetch(url, {
  method: 'POST', // or 'PUT'
  body: JSON.stringify(data), // data can be `string` or {object}!
  headers:{
    'Content-Type': 'application/json'
  }
}).then(res => res.json())
.catch(error => console.error('Error:', error))
.then(response => console.log('Success:', response));
}

function Checkout(){
    
postOrden()


            return(
            <div className="checkout-message">

            <div id="head">¡Felicidades, {cookies.get("username")}!</div>
            <div id="sub">Tu compra de ${cookies.get("total")} se realizó con éxito</div>
            <div id="msg">Pronto recibirás tu pedido</div>

            </div>



)
}
export default Checkout;