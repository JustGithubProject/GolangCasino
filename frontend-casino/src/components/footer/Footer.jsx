import React from 'react';

const footerStyle = {
  background: 'linear-gradient(to right, rgba(19, 19, 19, 0.8), rgba(30, 30, 30, 0.8))', // Gradient with transparency
  color: '#fff',
  textAlign: 'center',
  padding: '20px 0',
  position: 'absolute',
  bottom: 0,
  width: '100%',
  fontFamily: 'Arial, sans-serif',
  boxShadow: '0 -2px 10px rgba(0, 0, 0, 0.5)',
  borderTop: '1px solid #555', // Border to add some separation
};

const footerTextStyle = {
  margin: 0,
  fontSize: '1.2em',
  fontWeight: 'bold',
  letterSpacing: '1px',
  color: '#fff', // Accent color to match the overall theme
};

const Footer = () => {
  return (
    <div style={footerStyle}>
      <p style={footerTextStyle}>© 2024 Casino Royale. All Rights Reserved.</p>
    </div>
  );
};

export default Footer;
