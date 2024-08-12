import React from 'react';
import PayPalComponent from '../paypal_components/PaypalForm';

const TopUpPage = () => {
  return (
    <div style={pageStyle}>
      <div style={headerContainerStyle}>
        <h1 style={headerStyle}>Пополнить баланс</h1>
        <p style={paragraphStyle}>Выберите способ пополнения вашего баланса:</p>
      </div>
      <div style={paymentMethodContainerStyle}>
        <div style={paymentMethodStyle}>
          <h2 style={methodHeaderStyle}>PayPal</h2>
          <PayPalComponent />
        </div>
      </div>
    </div>
  );
};

const pageStyle = {
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  justifyContent: 'center',
  minHeight: '100vh',
  padding: '20px',
  background: 'linear-gradient(145deg, #f9f9f9, #e0e0e0)',
  fontFamily: "'Poppins', sans-serif",
  boxSizing: 'border-box',
};

const headerContainerStyle = {
  textAlign: 'center',
  marginBottom: '40px',
  padding: '20px',
  backgroundColor: '#ffffff',
  borderRadius: '15px',
  boxShadow: '0 8px 30px rgba(0, 0, 0, 0.1)',
  width: '100%',
  maxWidth: '800px',
};

const headerStyle = {
  fontSize: '48px',
  fontWeight: '800',
  color: '#333',
  margin: '0',
  background: 'linear-gradient(90deg, #4a90e2, #50e3c2)',
  backgroundClip: 'text',
  WebkitBackgroundClip: 'text',
  color: 'transparent',
};

const paragraphStyle = {
  fontSize: '22px',
  color: '#666',
  margin: '10px 0 0',
};

const paymentMethodContainerStyle = {
  display: 'flex',
  justifyContent: 'center',
  width: '100%',
  maxWidth: '800px',
};

const paymentMethodStyle = {
  width: '100%',
  padding: '40px',
  backgroundColor: '#ffffff',
  borderRadius: '15px',
  boxShadow: '0 12px 50px rgba(0, 0, 0, 0.2)',
  transition: 'transform 0.3s ease, box-shadow 0.3s ease',
  position: 'relative',
  zIndex: 1,
  overflow: 'hidden',
  transform: 'scale(1)',
};

const methodHeaderStyle = {
  fontSize: '32px',
  fontWeight: '700',
  color: '#333',
  marginBottom: '20px',
  borderBottom: '4px solid #4a90e2',
  paddingBottom: '10px',
  transition: 'color 0.3s ease',
};

const paymentMethodHoverStyle = {
  transform: 'scale(1.05)',
  boxShadow: '0 15px 60px rgba(0, 0, 0, 0.3)',
};

export default TopUpPage;
