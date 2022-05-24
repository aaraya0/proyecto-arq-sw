import React from 'react';

class Create extends React.Component {
    constructor(props) {
        super(props);
    }
    state = {  }
    render() { 
        return (
            <div class="card">
                <div class="card-header">
                   Productos
                </div>
                <div class="card-body">
                    <h4 class="card-title">Title</h4>
                    <p class="card-text">Text</p>
                </div>
                <div class="card-footer text-muted">
                    Footer
                </div>
            </div>
          );
    }
}
 
export default Create;