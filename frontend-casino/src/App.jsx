import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import RouletteCard from './components/rooms/RouletteCard';
import RouletteCardV2 from './components/rooms/RouletteCard_v2'
import RouletteCardV3 from './components/rooms/RouletteCard_v3'
import SweetBonanza from "./components/rooms/SweetBonanzaCard"
import Login from './components/auth_components/Login'
import Register from './components/auth_components/Register'
import Home from './components/home/Home'
import SuccessPage from './pages/SucessPage';
import CancelPage from './pages/CancelPage';
import TopUpPage from './components/payments/TopUpComponent';
import PaymentHistoryPage from './components/payments/PaymentsHistory';
import WithdrawFundsPage from './components/payments/WithdrawFundsPaypal';

function App() {
  return (
      <Router>
          <Routes>
              <Route path="/room/roulette/1" element={<RouletteCardV3/>} />
              <Route path="/room/roulette/2" element={<RouletteCardV3/>} />
              <Route path="/room/roulette/3" element={<RouletteCardV3/>} />
              <Route path="/room/roulette/4" element={<RouletteCardV3/>} />
              <Route path="/room/slot/sweetbonanza" element={<SweetBonanza/>} />
              <Route path="/room/roulette/6" element={<RouletteCardV3/>} />
              <Route path="/login" element={<Login/>} />
              <Route path="/register" element={<Register/>} />
              <Route path="/" element={<Home/>} />
              <Route path="/sucess-payment" element={<SuccessPage/>} />
              <Route path="/cancel-payment" element={<CancelPage/>} />
              <Route path="/top-up-balance" element={<TopUpPage/>} />
              <Route path="/payment-history" element={<PaymentHistoryPage/>} />
              <Route path="/withdraw-funds" element={<WithdrawFundsPage/>} />
          </Routes>
      </Router>
  );
}

export default App;