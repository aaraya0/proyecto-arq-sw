
import React from 'react';
import Cookies from 'universal-cookie'
const cookies = new Cookies();

class Cart extends React.Component{
    constructor() {
        super();
        this.state = {producto:[]};

    }
componentDidMount(){
    
        let current= cookies.get('cart')
        let items=current.split(',')
        items.forEach(item =>{
            
            let prod= item.split('=')
            let prodID=prod[0]
            let cant= prod[1]
            if (prodID!=0 && cant!="undefined" ){
                fetch( "http://localhost:8090/product/"+prodID)
                .then(res => res.json())
                .then((json) => {
                    this.setState({
                        producto: json 
                    });
                })
            }
            })
        
        }
 render() {
  
    const { producto } = this.state;
   
            return (
            <div className="product">
           <ol key = { producto.id } className="item">
            
            <div id="titulo">{ producto.title}</div>
            <div id="autor"> { producto.author }</div>	
             <div id="precio">{producto.base_price }</div> 
        
                </ol>
             </div>
           
            )
}
}
export default Cart;

/*fetch( "http://localhost:8090/product/"+prodID)
                .then(res => res.json())
                .then((json) => {
                    this.setState({
                        producto: json 
                      
                    });
                })*/