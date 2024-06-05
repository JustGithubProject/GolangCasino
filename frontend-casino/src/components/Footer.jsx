import React from 'react';

const footerStyle = {
  backgroundColor: '#000',
  color: '#fff',
  textAlign: 'center',
  padding: '20px 0',
  position: 'absolute',
  bottom: 0,
  width: '100%',
};

const Footer = () => {
  return (
    <div style={footerStyle}>
      <p>© 2024 Casino Royale. All Rights Reserved.</p>
    </div>
  );
};

export default Footer;
