import React, { useState } from 'react';

function DeletePack() {
  const [id, setID] = useState(0);


  const handleSubmit = async (event) => {
    event.preventDefault();
    const url = `http://api:8082/pack/${id}`;

    const response = await fetch(url, {
      method: 'DELETE',
    });

    const data = await response.json();
  }


  return (
    <div className='deletePack'>
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
