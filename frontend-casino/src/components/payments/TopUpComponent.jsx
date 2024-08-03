import React from 'react';
import PayPalComponent from '../paypal_components/PaypalForm';

const TopUpPage = () => {
  return (
    <div style={pageStyle}>
      <h1 style={headerStyle}>Пополнить баланс</h1>
      <p style={paragraphStyle}>Выберите способ пополнения вашего баланса:</p>
      <div style={paymentMethodStyle}>
        <h2 style={methodHeaderStyle}>PayPal</h2>
        <PayPalComponent />
      </div>
    </div>
  );
};

const pageStyle = {
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  justifyContent: 'center',
  padding: '20px',
  backgroundColor: '#f7f7f7',
  color: '#333',
};

const headerStyle = {
  fontSize: '32px',
  marginBottom: '20px',
};

const paragraphStyle = {
  fontSize: '18px',
  marginBottom: '20px',
};

const paymentMethodStyle = {
  width: '100%',
  maxWidth: '600px',
  margin: '20px 0',
  padding: '20px',
  backgroundColor: '#fff',
  borderRadius: '8px',
  boxShadow: '0 2px 10px rgba(0, 0, 0, 0.1)',
};

const methodHeaderStyle = {
  fontSize: '24px',
  marginBottom: '10px',
};

export default TopUpPage;
