import React, { useState } from 'react';

function DeletePack() {
  const [id, setID] = useState(0);


  const handleSubmit = async (event) => {
    event.preventDefault();
    const url = `/api/pack/${id}`;

    const response = await fetch(url, {
      method: 'GET',
    });

    const data = await response.json();
  }


  return (
    <div>
      <h1>Delete Pack</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>id:</label>
          <input type="text" value={id} onChange={(event) => setID(event.target.value)} />
        </div>
        <button type="submit">Delete Pack</button>
      </form>
    </div>
  );
}

export default DeletePack;
