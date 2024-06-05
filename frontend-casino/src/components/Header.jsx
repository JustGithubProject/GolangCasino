import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEye, faEyeSlash } from '@fortawesome/free-solid-svg-icons';

const Header = ({ username, balance, handleLogout }) => {
  const [showBalance, setShowBalance] = useState(false);

  const toggleBalance = () => {
    setShowBalance(!showBalance);
  };

  const headerStyle = {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: '20px 40px',
    backgroundColor: 'rgba(44, 62, 80, 0.8)', // More transparent background
    color: '#ECF0F1',
    boxShadow: '0 2px 10px rgba(0, 0, 0, 0.15)',
    backdropFilter: 'blur(10px)', // Adds a blur effect to the background
  };

  const userInfoStyle = {
    display: 'flex',
    alignItems: 'center',
  };

  const avatarStyle = {
    width: '50px',
    height: '50px',
    borderRadius: '50%',
    backgroundColor: 'rgba(52, 152, 219, 0.8)', // More transparent avatar background
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
  };

  const balanceStyle = {
    fontSize: '16px',
    color: '#BDC3C7',
    marginLeft: '10px',
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
    backgroundColor: 'rgba(52, 152, 219, 0.8)', // More transparent hover background
    transform: 'scale(1.1)',
  };

  const buttonStyle = {
    backgroundColor: 'rgba(231, 76, 60, 0.8)', // More transparent button background
    color: '#fff',
    border: 'none',
    padding: '10px 20px',
    fontSize: '16px',
    fontWeight: 'bold',
    borderRadius: '30px',
    cursor: 'pointer',
    transition: 'background-color 0.3s, transform 0.3s',
  };

  const iconStyle = {
    cursor: 'pointer',
    marginLeft: '10px',
    fontSize: '20px',
  };

  const formatBalance = (balance) => {
    return new Intl.NumberFormat('ru-RU', { style: 'currency', currency: 'RUB' }).format(balance);
  };

  return (
    <header style={headerStyle}>
      <div style={userInfoStyle}>
        <div style={avatarStyle}>{username.charAt(0).toUpperCase()}</div>
        <div>
          <div style={usernameStyle}>{username}</div>
          {showBalance && <div style={balanceStyle}>Баланс: {formatBalance(balance)}</div>}
        </div>
        <FontAwesomeIcon
          icon={showBalance ? faEyeSlash : faEye}
          style={iconStyle}
          onClick={toggleBalance}
        />
      </div>
      <nav style={navStyle}>
        <ul style={ulStyle}>
          <li style={liStyle}>
            <Link
              to="/roulette"
              style={linkStyle}
              onMouseEnter={(e) => {
                e.currentTarget.style.backgroundColor = linkHoverStyle.backgroundColor;
                e.currentTarget.style.transform = linkHoverStyle.transform;
              }}
              onMouseLeave={(e) => {
                e.currentTarget.style.backgroundColor = 'transparent';
                e.currentTarget.style.transform = 'scale(1)';
              }}
            >
              Roulette
            </Link>
          </li>
          <li style={liStyle}>
            <button
              onClick={handleLogout}
              style={buttonStyle}
              onMouseEnter={(e) => {
                e.currentTarget.style.backgroundColor = 'rgba(192, 57, 43, 0.8)'; // More transparent hover background
                e.currentTarget.style.transform = 'scale(1.1)';
              }}
              onMouseLeave={(e) => {
                e.currentTarget.style.backgroundColor = 'rgba(231, 76, 60, 0.8)'; // More transparent button background
                e.currentTarget.style.transform = 'scale(1)';
              }}
            >
              Logout
            </button>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
