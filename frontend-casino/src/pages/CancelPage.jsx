import React from 'react';
import { Link } from 'react-router-dom';

const CancelPage = () => {
  return (
    <div style={pageStyle}>
      <h1 style={headerStyle}>Оплата отменена</h1>
      <p style={paragraphStyle}>Ваша оплата была отменена. Пожалуйста, попробуйте снова.</p>
      <Link to="/" style={linkStyle}>Вернуться на главную</Link>
    </div>
  );
};

const pageStyle = {
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  justifyContent: 'center',
  height: '100vh',
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

const linkStyle = {
  textDecoration: 'none',
  color: '#3498db',
  fontSize: '18px',
  fontWeight: 'bold',
};

export default CancelPage;
