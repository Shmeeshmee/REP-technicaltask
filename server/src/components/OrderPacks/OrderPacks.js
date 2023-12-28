import React, { useState } from 'react';

function OrderPacks() {
  const [id, setID] = useState(null);
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

      <h1 style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>Order</h1>
      <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <div style={{ marginBottom: '10px', display: 'flex', alignItems: 'center' }}>
          <label htmlFor="ID" style={{ marginRight: '5px', fontSize: '16px', fontWeight: 'bold' }}>ID: </label>
          <input id="ID" type="text" value={id} onChange={(event) => setID(event.target.value)} style={{ fontSize: '16px', padding: '5px', border: '1px solid #9dd9f3', borderRadius: '5px' }} autoFocus />
        </div>
        <button type="submit" style={{ fontSize: '16px', padding: '5px 10px', border: 'none', borderRadius: '5px', backgroundColor: '#9dd9f3', color: '#fff' }}>Order!</button>
      </form>
      {result()}
    </div>
  );


  function result() {
    if (res.message) {
      return (
        <div style={{ backgroundColor: '#f2f2f2', padding: '20px', borderRadius: '5px' }}>
          <h1 style={{ color: '#333', fontSize: '24px', fontWeight: 'bold' }}>{res.message}</h1>
          <h1 style={{ color: '#666', fontSize: '18px', marginTop: '10px' }}>Result is: [{res.response.join(', ')}]</h1>
        </div>
      );
    }
  }

}


export default OrderPacks;
