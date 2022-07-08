import React from 'react';
import Cookies from 'universal-cookie';
 
import "./Checkout.css"
 
const cookies = new Cookies();
function gopath(path){
	window.location = window.location.origin + path
	}
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

async function postAddress (){
  
  return fetch("http://localhost:8090/address", {
    method:"POST",
    mode: 'no-cors',
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({
    country:document.getElementById("inputCountry").value ,
    city: document.getElementById("inputCity").value,
    street: document.getElementById("inputStreet").value ,
    number: Number(document.getElementById("inputNumber").value),
    cod_postal:Number(document.getElementById("inputZip").value),
    user_id: Number(cookies.get("user_id"))
    })
  })
}
async function shop (){
  crearOrden()
  postAddress()
  cookies.remove("cart")
  alert(`Felicitaciones, ${cookies.get("username")}! Tu compra fue realizada con éxito.`)
  cookies.remove("total")
 
  gopath("/home")
}



function Checkout(){
    



            return(
            <div className="checkout-message">
            <div id="sub">Está por realizar una compra de ${cookies.get("total")}</div>
            <div id="msg">Ingrese la dirección de envío</div>
            <div class="form-group">
              <form>
  <input type="text" class="form-control" id="inputStreet" placeholder="Calle" required/>
  <input type="number" class="form-control" id="inputNumber" placeholder="Numeración"/>
  <input type="text" class="form-control" id="inputCity" placeholder="Ciudad" required />
  <input type="number" class="form-control" id="inputZip" placeholder="CP" required/>
  <input type="text"  class="form-control"  id="inputCountry"  placeholder="País" required/>
  </form>
</div>
            <button onClick={shop} id="shopbutton" >Confirmar</button>
           
            </div>
            


)
}
export default Checkout;