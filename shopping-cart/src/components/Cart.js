
import React from 'react';

class Cart extends React.Component {

   //Cada vez que se presiona el boton de agregar al carrito, hacer un post al carrito del back, que grabe en la bd
   //y que tambien reste stock de ese producto. luego, en el path /cart, mostrar los productos enviados a la bd
   //haciendo un get, probar hacer funcion de total en el back y devolver eso tambien
   //si se elimina algun elemento del carrito, eliminarlo de la bd y agregar stock
   //cuando se haga el checkout, pasar los datos de la tabla a una funcion orden en el back, que devuelva un json con 
   //datos de usuario, articulos comprados, fecha y total
   //hacer un get desde el front y mostrar en un path /order. 
   //boton volver al menu, eliminar tabla carrito.

   constructor(props) {
      super(props);

      this.state = {
          postId: null
      };
  }

	// ComponentDidMount is used to
	// execute the code
	componentDidMount() {
      const requestOptions = {
         method: 'POST',
         headers: { 'Content-Type': 'application/json' },
         body: JSON.stringify({ title: 'Sinsajo', author: 'Suzanne Collins' })
     };
     fetch('https://localhost:8090/cart', requestOptions)
         .then(response => response.json())
         .then(data => this.setState({ postId: data.id }));
	}
	
   render() {
      const { postId } = this.state;
      return (
          <div className="card text-center m-3">
              <h5 className="card-header">Simple POST Request</h5>
              <div className="card-body">
                  Returned Id: {postId}
              </div>
          </div>
      );
  }
}




export default Cart;