import React, { useState, useEffect } from "react";
import styles from './GetAllPacks.module.css';


GetAllPacks.defaultProps = {};

export default function GetAllPacks() {
  let [packs, setPacks] = useState([]);

  function displayArrValues() {
    const containerStyle = {
      // Add your container styles here
      backgroundColor: '#f0f0f0',
      padding: '10px',
      borderRadius: '5px',
      margin: '10px',
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'center', // Center items horizontally
      justifyContent: 'center', // Center items vertically
    };

    const itemStyle = {
      // Add your item styles here
      borderBottom: '1px solid #ccc',
      margin: '5px 0',
    };

    return (
      <div style={containerStyle}>
        {packs &&
          packs.map((pack, key) => (
            <div key={pack.key} style={itemStyle}>
              <h1>
                {pack.key} : {pack.value}
              </h1>
            </div>
          ))}
      </div>
    );
  }


  useEffect(() => {
    let url = "http://localhost:8082/packs";
    fetch(url)
      .then((response) => response.json())
      .then((data) => setPacks(data.response))
      .catch((error) => {
        console.error("Error fetching data: ", error);
      });
    console.log(packs);
  }, []);


  return (
    <div className={styles.GetAllPacks} data-testid="GetAllPacks">
      <h1 style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>Get All Packs</h1>

      {displayArrValues()}
    </div>
  );
}
