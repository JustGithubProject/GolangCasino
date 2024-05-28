import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import RouletteCard from './components/RouletteCard';
import Login from './components/auth_components/Login'
import Register from './components/auth_components/Register'

function App() {
  return (
      <Router>
          <Routes>
              <Route path="/roulette" element={<RouletteCard />} />
              <Route path="/login" element={<Login/>} />
              <Route path="/register" element={<Register/>} />
          </Routes>
      </Router>
  );
}

export default App;