import React from 'react';
class List extends React.Component {
    constructor(props) {
        super(props);
    }
    state = {  }
    render() { 
        return ( <table class="table">
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Nombre</th>
                    
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td scope="row">1</td>
                    <td>Prod1</td>
                    
                </tr>
                <tr>
                    <td scope="row">2</td>
                    <td>Prod2</td>
            
                </tr>
            </tbody>
        </table> );
    }
}
 
export default List;