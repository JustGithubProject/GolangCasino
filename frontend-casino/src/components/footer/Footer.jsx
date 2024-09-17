import React from 'react';

const footerStyle = {
  background: 'linear-gradient(to right, rgba(19, 19, 19, 0.8), rgba(30, 30, 30, 0.8))',
  color: '#eee',
  textAlign: 'center',
  padding: '1px 0',
  position: 'fixed',
  bottom: 0,
  width: '100%',
  fontFamily: 'Roboto, sans-serif',
  boxShadow: '0 -4px 12px rgba(0, 0, 0, 0.7)',
  zIndex: 1000,
};

const footerTextStyle = {
  margin: '10px 0',
  fontSize: '1.1em',
  fontWeight: '300',
  letterSpacing: '1px',
  color: '#ddd',
};

const linkContainerStyle = {
  margin: '10px 0',
};

const linkStyle = {
  color: '#faf7f5',
  textDecoration: 'none',
  margin: '0 15px',
  fontSize: '1.2em',
  transition: 'color 0.3s ease',
};

const linkHoverStyle = {
  color: '#f2f6f7',
};

const iconStyle = {
  margin: '0 12px',
  color: '#faf7f5',
  fontSize: '1.6em',
  transition: 'color 0.3s ease',
};

const iconHoverStyle = {
  color: '#e76f51',
};

const Footer = () => {
  return (
    <div style={footerStyle}>
      <div style={linkContainerStyle}>
        <a href="https://twitter.com" style={linkStyle} 
           onMouseOver={(e) => e.currentTarget.style.color = linkHoverStyle.color}
           onMouseOut={(e) => e.currentTarget.style.color = linkStyle.color}>
          <i className="fab fa-twitter" style={iconStyle}
             onMouseOver={(e) => e.currentTarget.style.color = iconHoverStyle.color}
             onMouseOut={(e) => e.currentTarget.style.color = iconStyle.color}></i>
        </a>
        <a href="https://facebook.com" style={linkStyle}
           onMouseOver={(e) => e.currentTarget.style.color = linkHoverStyle.color}
           onMouseOut={(e) => e.currentTarget.style.color = linkStyle.color}>
          <i className="fab fa-facebook-f" style={iconStyle}
             onMouseOver={(e) => e.currentTarget.style.color = iconHoverStyle.color}
             onMouseOut={(e) => e.currentTarget.style.color = iconStyle.color}></i>
        </a>
        <a href="https://instagram.com" style={linkStyle}
           onMouseOver={(e) => e.currentTarget.style.color = linkHoverStyle.color}
           onMouseOut={(e) => e.currentTarget.style.color = linkStyle.color}>
          <i className="fab fa-instagram" style={iconStyle}
             onMouseOver={(e) => e.currentTarget.style.color = iconHoverStyle.color}
             onMouseOut={(e) => e.currentTarget.style.color = iconStyle.color}></i>
        </a>
      </div>
      <p style={footerTextStyle}>© 2024 Casino Royale. All Rights Reserved.</p>
    </div>
  );
};

export default Footer;
