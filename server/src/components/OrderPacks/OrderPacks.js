import React, { useState } from 'react';

function OrderPacks() {
  const [id, setID] = useState(0);
  const [res, setRes] = useState({});


  const handleSubmit = async (event) => {
    event.preventDefault();
    const url = `http://localhost:8082/calculate/${id}`;

    const response = await fetch(url, {
      method: 'GET',
    });

    const data = await response.json();

    setRes(data)
  }


  return (
    <div className='order'>
      <h1>Order</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>id:</label>
          <input type="text" value={id} onChange={(event) => setID(event.target.value)} />
        </div>
        <button type="submit">Order</button>
      </form>
{result()}
    </div>
  );


  function result(){
    if (res.message) return (
        <div>
          <h1> {res.message } </h1>
          <h1> Result is: [ {res.response.join(', ')} ] </h1>
        </div>
      )
  }
}


export default OrderPacks;
