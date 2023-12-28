import React, { useState } from 'react';

function AddPack() {
  const [size, setSize] = useState(null);
  const [res, setRes] = useState("");

  const handleSubmit = async (event) => {
    event.preventDefault();

    const response = await fetch('http://localhost:8082/pack', {
      method: 'POST',
      body: JSON.stringify({
        size: Number(size),
      })
    });

    const data = await response.json();

    setRes(data)
  }

  function err() {
    if (res.error) {
      return (
        <div style={{ backgroundColor: '#f2f2f2', padding: '20px', borderRadius: '5px', color: '#333', fontSize: '16px', fontWeight: 'bold' }}>
          <h1 style={{ color: '#ed1c24' }}>Error:</h1>
          <p style={{ margin: '10px 0' }}>{res.error.error_message}</p>
          <p style={{ margin: '10px 0' }}>Code: [{res.error.error_code.code}] {res.error.error_code.name}</p>
        </div>
      );
    }
  }


  function result() {
    if (res.message && res.response) {
      return (
        <div style={{ backgroundColor: '#f2f2f2', padding: '20px', borderRadius: '5px' }}>
          <h1 style={{ color: '#333', fontSize: '24px', fontWeight: 'bold' }}>{res.message}</h1>
          <h1 style={{ color: '#666', fontSize: '18px', marginTop: '10px' }}>response value is: [{res.response.size}]</h1>
        </div>
      );
    }
  }


  return (
    <div>
      <h1 style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>Add Pack</h1>
      <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', alignItems: 'center'}}>
        <div style={{ marginBottom: '10px', display: 'flex', alignItems: 'center' }}>
          <label style={{ marginRight: '5px', fontSize: '16px', fontWeight: 'bold' }}>size:</label>
          <input type="text" value={size} onChange={(event) => setSize(event.target.value)} style={{ fontSize: '16px', padding: '5px', border: '1px solid #9dd9f3', borderRadius: '5px' }} autoFocus />
        </div>
        <button type="submit" style={{ fontSize: '16px', padding: '5px 10px', border: 'none', borderRadius: '5px', backgroundColor: '#9dd9f3', color: '#fff' }}>Add Pack</button>
      </form>

      {result()}
      {err()}
    </div>
  );
}

export default AddPack;
