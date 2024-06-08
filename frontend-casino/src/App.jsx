import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import RouletteCard from './components/RouletteCard';
import RouletteCardV2 from './components/RouletteCard_v2'
import Login from './components/auth_components/Login'
import Register from './components/auth_components/Register'
import Home from './components/Home'


function App() {
  return (
      <Router>
          <Routes>
              <Route path="/room/roulette/1" element={<RouletteCard />} />
              <Route path="/room/roulette/2" element={<RouletteCardV2/>} />
              <Route path="/login" element={<Login/>} />
              <Route path="/register" element={<Register/>} />
              <Route path="/" element={<Home/>} />
          </Routes>
      </Router>
  );
}

export default App;