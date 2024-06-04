import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import Login from './auth_components/Login';
import Register from './auth_components/Register';

const Home = () => {
  const [isLogin, setIsLogin] = useState(true);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleLogout = async () => {
    try {
      await axios.post('http://127.0.0.1:8081/logout');

      localStorage.removeItem('token');

      window.location.href = '/roulette';
    } catch (error) {
      console.error('Failed to logout:', error);
    }
  };

  const containerStyle = {
    display: 'flex',
    flexDirection: 'row',
    height: '100vh',
    fontFamily: 'Arial, sans-serif',
    background: 'linear-gradient(to right, #43cea2, #185a9d)',
    color: '#fff',
  };

  const navStyle = {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'flex-start',
    width: '200px',
    padding: '20px',
    backgroundColor: '#333',
    color: '#fff',
    boxShadow: '2px 0 5px rgba(0, 0, 0, 0.1)',
  };

  const ulStyle = {
    listStyleType: 'none',
    padding: '0',
    margin: '0',
    width: '100%',
  };

  const liStyle = {
    margin: '15px 0',
  };

  const linkStyle = {
    textDecoration: 'none',
    color: '#fff',
    fontWeight: 'bold',
    fontSize: '16px',
    display: 'block',
    padding: '10px',
    borderRadius: '4px',
    transition: 'background-color 0.3s, transform 0.3s',
  };

  const linkHoverStyle = {
    backgroundColor: '#007BFF',
    transform: 'scale(1.05)',
  };

  const buttonStyle = {
    backgroundColor: '#007BFF',
    color: '#fff',
    border: 'none',
    padding: '10px',
    fontSize: '16px',
    fontWeight: 'bold',
    borderRadius: '4px',
    cursor: 'pointer',
    transition: 'background-color 0.3s, transform 0.3s',
  };

  const mainStyle = {
    flex: 1,
    padding: '40px',
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
  };

  const tabStyle = {
    display: 'flex',
    justifyContent: 'center',
    marginBottom: '20px',
  };

  const tabButtonStyle = {
    padding: '10px 20px',
    cursor: 'pointer',
    border: 'none',
    background: 'none',
    fontSize: '18px',
    fontWeight: 'bold',
    borderBottom: isLogin ? '2px solid #007BFF' : 'none',
  };

  const inactiveTabButtonStyle = {
    padding: '10px 20px',
    cursor: 'pointer',
    border: 'none',
    background: 'none',
    fontSize: '18px',
    fontWeight: 'bold',
    borderBottom: !isLogin ? '2px solid #007BFF' : 'none',
  };

  return (
    <div style={containerStyle}>
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
                e.currentTarget.style.backgroundColor = '#0056b3';
                e.currentTarget.style.transform = 'scale(1.05)';
              }}
              onMouseLeave={(e) => {
                e.currentTarget.style.backgroundColor = '#007BFF';
                e.currentTarget.style.transform = 'scale(1)';
              }}
            >
              Logout
            </button>
          </li>
        </ul>
      </nav>
      <div style={mainStyle}>
        <div style={{
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          background: '#fff',
          color: '#333',
          padding: '30px',
          borderRadius: '15px',
          boxShadow: '0 8px 20px rgba(0, 0, 0, 0.2)',
          transition: 'transform 0.3s ease, box-shadow 0.3s ease',
          width: '350px',
        }}>
          <div style={tabStyle}>
            <button
              style={tabButtonStyle}
              onClick={() => setIsLogin(true)}
            >
              Логин
            </button>
            <button
              style={inactiveTabButtonStyle}
              onClick={() => setIsLogin(false)}
            >
              Регистрация
            </button>
          </div>
          {error && <p style={{ color: '#ff4b2b' }}>{error}</p>}
          {success && <p style={{ color: '#28a745' }}>{success}</p>}
          {isLogin ? (
            <Login />
          ) : (
            <Register />
          )}
        </div>
      </div>
    </div>
  );
};

export default Home;
