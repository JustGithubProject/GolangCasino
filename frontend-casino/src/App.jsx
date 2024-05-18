import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import RouletteCard from './components/RouletteCard';

function App() {
  return (
      <Router>
          <Routes>
              <Route path="/roulette" element={<RouletteCard />} />
          </Routes>
      </Router>
  );
}

export default App;