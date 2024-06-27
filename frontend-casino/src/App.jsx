import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import RouletteCard from './components/RouletteCard';
import RouletteCardV2 from './components/RouletteCard_v2'
import RouletteCardV3 from './components/RouletteCard_v3'
import Login from './components/auth_components/Login'
import Register from './components/auth_components/Register'
import Home from './components/Home'


function App() {
  return (
      <Router>
          <Routes>
              <Route path="/room/roulette/1" element={<RouletteCard />} />
              <Route path="/room/roulette/2" element={<RouletteCardV2/>} />
              <Route path="/room/roulette/3" element={<RouletteCardV3/>} />
              <Route path="/room/roulette/4" element={<RouletteCardV3/>} />
              <Route path="/room/roulette/5" element={<RouletteCardV3/>} />
              <Route path="/room/roulette/6" element={<RouletteCardV3/>} />
              <Route path="/login" element={<Login/>} />
              <Route path="/register" element={<Register/>} />
              <Route path="/" element={<Home/>} />
          </Routes>
      </Router>
  );
}

export default App;