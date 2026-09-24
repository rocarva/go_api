import React, { Component } from 'react';
import axios from 'axios';
import '../components/Personalidades.css';

const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8000';

export default class Personalidades extends Component {
    state = {
        personalidades: []
    }

    componentDidMount() {
        axios.get(`${API_URL}/api/personalidades`)
            .then(res => {
                const personalidades = res.data;
                this.setState({ personalidades })
            })
    }

    render() {
        return (
            <div>
                {this.state.personalidades.map((p, id )=>
                    <div className="CardPersonalidades" key={id}>
                        <h3>{p.nome}</h3>
                        <p>{p.historia}</p>
                    </div>)}
            </div>
        );
    }
}
