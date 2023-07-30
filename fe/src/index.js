import React from 'react';
import ReactDOM from 'react-dom/client';
import './style.css';

class Submit extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            orderMsg: "",
        };
    }
    onClick() {
        const xhr = new XMLHttpRequest();
        xhr.open('POST', 'http://35.229.224.119:1234/order');
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.onload = function() {
            console.log(xhr.responseText)
            this.setState({orderMsg: JSON.parse(xhr.responseText).reason});
        }.bind(this);
        xhr.onerror = function() {
            this.setState({orderMsg: "request error"});
        }.bind(this);
        xhr.send(JSON.stringify({ "name": this.props.name, "password": this.props.password}));
    }
    render() {
        return (
            <div>
                <button onClick={() => this.onClick()}>
                    Đăng ký
                </button>
                <label>
                    {this.state.orderMsg}
                </label>
            </div>
        )
    }
}

class Game extends React.Component {
    constructor(props) {
        super(props);
        const xhr = new XMLHttpRequest();
        xhr.open('GET', 'http://35.229.224.119:1234/orders');
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.onload = function() {
            console.log(xhr.responseText)
            let data = JSON.parse(xhr.responseText).data
            if (data == null) {
                return
            }
            this.setState({ordered: JSON.parse(xhr.responseText).data})
        }.bind(this);
        xhr.onerror = function() {
            console.log("error")
        }.bind(this);
        xhr.send();
        this.state = {
            name : "",
            password :"",
            ordered: [
                {id: "1", name: "chưa ai đặt cơm"},
            ]
        }
    }
    handleNameChange(s) {
        this.setState({name : s.target.value})
    }
    handlePasswordChange(s) {
        this.setState({password : s.target.value})
    }
    render() {
        return (
            <div className='game'>
                <h1>Đặt cơm</h1>
                <div> 
                    <label>
                        Tên: 
                    </label>
                    <input onChange={(event) => this.handleNameChange(event)}>
                    </input>
                    <br></br>
                    <br></br>
                    <label>
                        Mật khẩu:
                    </label>
                    <input type='password' onChange={(event) => this.handlePasswordChange(event)}>
                    </input>
                    <br></br>
                    <br></br>
                    <Submit name={this.state.name} password={this.state.password}>

                    </Submit>
                </div>
                <div>
                    <h2>Danh sách đặt cơm hôm nay</h2>
                    <ol>
                        {this.state.ordered.map(data => (
                        <li key={data.id}> {data.name} </li>
                        ))}
                    </ol>
                </div>
            </div>
        )
    }
}

const root = ReactDOM.createRoot(document.getElementById("root"))
root.render(<Game/>)
