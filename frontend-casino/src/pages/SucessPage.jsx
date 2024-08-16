import React from 'react';
import { Link } from 'react-router-dom';

const SuccessPage = () => {
  return (
    <div style={pageStyle}>
      <div style={contentStyle}>
        <h1 style={headerStyle}>Оплата успешна</h1>
        <p style={paragraphStyle}>Спасибо за вашу оплату. Ваш транзакция завершена успешно.</p>
        <p style={paragraphImportantStyle}><strong>Важно!</strong> Зайдите в "Мои платежи" и заберите свои средства</p>
        <Link to="/" style={linkStyle}>Вернуться на главную</Link>
      </div>
    </div>
  );
};

const pageStyle = {
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
  height: '100vh',
  backgroundColor: '#1c1c1c', // Темный фон
  color: '#e0e0e0', // Светлый текст
};

const contentStyle = {
  backgroundColor: '#2c2c2c', // Чуть светлее для выделения
  padding: '30px',
  borderRadius: '10px',
  textAlign: 'center',
  boxShadow: '0 4px 8px rgba(0, 0, 0, 0.5)', // Более мягкая тень
};

const headerStyle = {
  fontSize: '32px',
  marginBottom: '20px',
  color: '#ffffff', // Белый цвет заголовка
};

const paragraphStyle = {
  fontSize: '18px',
  marginBottom: '20px',
};

const paragraphImportantStyle = {
  fontSize: '18px',
  color: '#76ff03', // Зеленый для важного текста, чтобы он выделялся
  marginBottom: '20px',
};

const linkStyle = {
  textDecoration: 'none',
  color: '#81a1c1', // Спокойный синий цвет ссылки
  fontSize: '18px',
  fontWeight: 'bold',
  padding: '10px 20px',
  backgroundColor: '#3b4252', // Темно-синий фон для кнопки
  borderRadius: '5px',
  transition: 'background-color 0.3s ease',
};

linkStyle[':hover'] = {
  backgroundColor: '#4c566a', // Немного светлее при наведении
};

export default SuccessPage;
