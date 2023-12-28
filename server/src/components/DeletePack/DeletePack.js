import React, { useState } from 'react';

function DeletePack() {
  const [id, setID] = useState(null);
  const [res, setRes] = useState("");


  const handleSubmit = async (event) => {
    event.preventDefault();
    console.log(`http://localhost:8082/pack/${id}`);

    const response = await fetch(`http://localhost:8082/delete_pack/${id}`, {
      method: 'GET',
    })
    .then((response) => response.json())
      .then((data) => setRes(data))
      .catch((error) => {
        console.error("Error fetching data: ", error);
      });;
  }

function result(){
  if (res.message) return (
    <div>
      <h1>
        {res.message}
      </h1>
    </div>
  )
}


  function err() {
    if (res.error) {
      return (
        <div style={{ color: 'ed', fontSize: '16px', fontWeight: 'bold' }}>
          <p>{res.error.error_message}</p>
        </div>
      );
    }
  }


  return (
    <div className='deletePack'>
      <h1 style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>Delete Pack</h1>
      <p style={{ display: 'flex', flexDirection: 'column', alignItems: 'center', fontSize: '0.9em', color: 'lightblue'}}>
        delete values by using the ID value [GET ALL PACKS > [pack.ID]]
      </p>
      <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <div style={{ marginBottom: '10px', display: 'flex', alignItems: 'center' }}>
          <label style={{ marginRight: '5px', fontSize: '16px', fontWeight: 'bold' }}>id:</label>
          <input type="text" value={id} onChange={(event) => setID(event.target.value)} style={{ fontSize: '16px', padding: '5px', border: '1px solid #9dd9f3', borderRadius: '5px' }} autoFocus />
        </div>
        <button type="submit" style={{ fontSize: '16px', padding: '5px 10px', border: 'none', borderRadius: '5px', backgroundColor: '#9dd9f3', color: '#fff' }}>Add Pack</button>
      </form>

      {result()}
      {err()}
    </div>
  );
}

export default DeletePack;
