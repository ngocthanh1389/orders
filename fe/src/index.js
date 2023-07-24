import React from 'react';
import ReactDOM from 'react-dom/client';
import './style.css';

class Submit extends React.Component {
    onClick() {
        const xhr = new XMLHttpRequest();
        xhr.open('POST', 'http://localhost:1234/order');
        xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
        xhr.onload = function() {
        //   if (xhr.status === 200) {
            console.log(xhr.responseText)
            // setData(JSON.parse(xhr.responseText));
        //   }
        };
        xhr.send(JSON.stringify({ "name": "abc", "password": "abc" }));
    }
    render() {
        return (
            <div>
                <button onClick={() => this.onClick()}>
                    Đăng ký
                </button>
                <label>
                
                </label>
            </div>
        )
    }
}

class Game extends React.Component {
    render() {
        return (
            <div className='game'>
                <h1>Đặt cơm</h1>
                <div> 
                    <label>
                        Tên: 
                    </label>
                    <input>

                    </input>
                    <br></br>
                    <br></br>
                    <label>
                        Mật khẩu:
                    </label>
                    <input type='password'>
                    </input>
                    <br></br>
                    <br></br>
                    <Submit>

                    </Submit>
                </div>
                <div>
                    <h2>Danh sách đặt cơm hôm nay</h2>
                </div>

            </div>
        )
    }
}

const root = ReactDOM.createRoot(document.getElementById("root"))
root.render(<Game/>)
