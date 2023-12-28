import React, { useState } from 'react';

function AddPack() {
  const [size, setSize] = useState(0);
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

    setRes(data.message)
  }

  return (
    <div>
      <h1>Add Pack</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>size:</label>
          <input type="text" value={size} onChange={(event) => setSize(event.target.value)} />
        </div>
        <button type="submit">Add Pack</button>
      </form>

      <h1> {res} </h1>
    </div>
  );
}

export default AddPack;
