import React, { useState, useEffect } from "react";
import styles from './GetAllPacks.module.css';


GetAllPacks.defaultProps = {};

export default function GetAllPacks() {
  let [packs, setPacks] = useState([]);

  function displayArrValues() {
    if (packs){
      return (
  <div>
      {packs.map((pack, key) => (
        <div key={pack.key}>
         <h1>
            {pack.key} : {pack.value}
          </h1> 
        </div>
      ))}
    </div>
  )
    }
    return 
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
      {displayArrValues()}
    </div>
  );
}
