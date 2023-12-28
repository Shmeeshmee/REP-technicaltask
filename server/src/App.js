import React from 'react';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import GetAllPacks from './components/GetAllPacks/GetAllPacks';
import GetPackById from './components/GetPackById/GetPackById';
import DeletePack from './components/DeletePack/DeletePack';
import AddPack from './components/AddPack/AddPack';
import OrderPacks from './components/OrderPacks/OrderPacks';
import NavBar from "./components/navbar/NavBar";

import "./App.css";

function App() {
  return (
    <BrowserRouter>
      <NavBar />
      <Routes>
        <Route path="/">
          <Route index element={<GetAllPacks />} />
          <Route path="get_all_packs" element={<GetAllPacks />} />
          <Route path="get_pack_by_id" element={<GetPackById />} />
          <Route path="delete_pack" element={<DeletePack />} />
          <Route path="add_pack" element={<AddPack />} />
          <Route path="order_packs" element={<OrderPacks />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
