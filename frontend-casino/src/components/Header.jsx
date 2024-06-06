import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEye, faEyeSlash, faHome } from '@fortawesome/free-solid-svg-icons';

const Header = ({ username, balance, handleLogout }) => {
  const [showBalance, setShowBalance] = useState(false);
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      setIsAuthenticated(true);
    }
  }, []);

  const toggleBalance = () => {
    setShowBalance(!showBalance);
  };

  const headerStyle = {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: '20px 40px',
    background: 'linear-gradient(to right, rgba(19, 19, 19, 0.8), rgba(30, 30, 30, 0.8))',
    color: '#ECF0F1',
    boxShadow: '0 2px 10px rgba(0, 0, 0, 0.15)',
    backdropFilter: 'blur(10px)',
    width: '100%',
    position: 'fixed',
    top: '0',
    zIndex: '1000',
  };

  const rightSectionStyle = {
    display: 'flex',
    alignItems: 'center',
  };

  const userInfoStyle = {
    display: 'flex',
    alignItems: 'center',
    marginRight: '20px',
  };

  const avatarStyle = {
    width: '50px',
    height: '50px',
    borderRadius: '50%',
    backgroundColor: 'rgba(52, 152, 219, 0.8)',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    color: '#fff',
    fontSize: '24px',
    fontWeight: 'bold',
    marginRight: '15px',
  };

  const usernameStyle = {
    fontSize: '18px',
    fontWeight: 'bold',
    display: 'flex',
    alignItems: 'center',
  };

  const balanceStyle = {
    fontSize: '16px',
    color: '#fff',
    marginLeft: '20px',
  };

  const navStyle = {
    display: 'flex',
    alignItems: 'center',
  };

  const ulStyle = {
    listStyleType: 'none',
    padding: '0',
    margin: '0',
    display: 'flex',
  };

  const liStyle = {
    margin: '0 15px',
  };

  const linkStyle = {
    textDecoration: 'none',
    color: '#ECF0F1',
    fontWeight: 'bold',
    fontSize: '16px',
    display: 'block',
    padding: '10px 20px',
    borderRadius: '30px',
    transition: 'background-color 0.3s, transform 0.3s',
  };

  const linkHoverStyle = {
    backgroundColor: 'rgba(52, 152, 219, 0.8)',
    transform: 'scale(1.1)',
  };

  const buttonStyle = {
    backgroundColor: 'rgba(231, 76, 60, 0.8)',
    color: '#fff',
    border: 'none',
    padding: '10px 20px',
    fontSize: '16px',
    fontWeight: 'bold',
    borderRadius: '30px',
    cursor: 'pointer',
    marginRight: '30px',
    transition: 'background-color 0.3s, transform 0.3s',
    marginLeft: 'auto',
  };

  const iconStyle = {
    cursor: 'pointer',
    marginLeft: '10px',
    fontSize: '20px',
  };

  const homeIconStyle = {
    cursor: 'pointer',
    fontSize: '24px',
    color: '#ECF0F1',
  };

  const formatBalance = (balance) => {
    return new Intl.NumberFormat('ru-RU', { style: 'currency', currency: 'RUB' }).format(balance);
  };

  return (
    <header style={headerStyle}>
      <Link to="/" style={{ textDecoration: 'none' }}>
        <FontAwesomeIcon icon={faHome} style={homeIconStyle} />
      </Link>
      {isAuthenticated && (
        <button
          onClick={handleLogout}
          style={buttonStyle}
          onMouseEnter={(e) => {
            e.currentTarget.style.backgroundColor = 'rgba(192, 57, 43, 0.8)';
            e.currentTarget.style.transform = 'scale(1.1)';
          }}
          onMouseLeave={(e) => {
            e.currentTarget.style.backgroundColor = 'rgba(231, 76, 60, 0.8)';
            e.currentTarget.style.transform = 'scale(1)';
          }}
        >
          Выйти
        </button>
      )}
      <nav style={navStyle}>
        <ul style={ulStyle}>
          {/* Additional navigation links can be added here */}
        </ul>
      </nav>
      {isAuthenticated && (
        <div style={rightSectionStyle}>
          <div style={userInfoStyle}>
            <div style={avatarStyle}>{username.charAt(0).toUpperCase()}</div>
            <div style={usernameStyle}>
              <div>{username}</div>
              {showBalance && <div style={balanceStyle}>Баланс: {formatBalance(balance)}</div>}
              <FontAwesomeIcon
                icon={showBalance ? faEyeSlash : faEye}
                style={iconStyle}
                onClick={toggleBalance}
              />
            </div>
          </div>
        </div>
      )}
    </header>
  );
};

export default Header;
