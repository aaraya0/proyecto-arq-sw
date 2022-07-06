import React from 'react';
import Cookies from 'universal-cookie';
 
import "./Checkout.css"
 
const cookies = new Cookies();
async function getProductById(id){
  return fetch("http://localhost:8090/product/" + id).then(response => response.json())
  }
async function crearOrden(){
  let details=[]
  let current= cookies.get('cart')
  let items=[]
  let a=current.split(',')
  for (let i = 0; i < a.length; i++){
  let item = a[i];
  let prod= item.split('=')
  let prodID=prod[0]
  let cant= prod[1]
  if (prodID!=0 && prodID!="undefined" && cant!="undefined" && cant>0){
  let product = {product_id:0, quantity:0}
  product.quantity = cant;
  product.product_id=prodID;
items.push(product)
}
  }
items.forEach((item2) => {
  let detail = {
    product_id: Number(item2.product_id),
    quantity: Number(item2.quantity)
  }
  details.push(detail)
});



return fetch("http://localhost:8090/order", {
  method:"POST",
  mode: 'no-cors',
  headers: {
    "Content-Type": "application/json"
  },
  body: JSON.stringify({
    user_id: Number(cookies.get("user_id")),
    orderDetail: details
  })
})
}

async function postOrden (){
    let user_Id = parseInt(cookies.get("user_id"))
   
    var url = 'http://localhost:8090/order'
    let detail= await(crearOrden())
var data = {"user_id":user_Id, "orderDetail": detail}

fetch(url, {
  method: 'POST',
  mode: 'no-cors',
  body: JSON.stringify(data), 
  headers:{
    'Content-Type': 'application/json',
  }
}).then(res => res.json())
.catch(error => console.error('Error:', error))
.then(response => console.log('Success:', response));
}



function Checkout(){
    
crearOrden()


            return(
            <div className="checkout-message">
            <div id="sub">Estas por realizar una compra de ${cookies.get("total")}</div>
            <div id="msg">Ingresá tu dirección, pronto recibirás tu pedido.</div>
            <button>Confirmar</button>
            </div>



)
}
export default Checkout;