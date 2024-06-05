import React, { useState, useEffect } from 'react';
import { fetchWithAuth } from './auth_components/fetchWrapper';
import Login from './auth_components/Login';
import Register from './auth_components/Register';
import * as jwtDecodeModule from 'jwt-decode';
import { Link } from 'react-router-dom';
import Header from './Header';
import backgroundImage from '../images/casinoImage_2.png';
import cardBackgroundImage from '../images/card.png'; // Import the card background image

const Home = () => {
  const [isLogin, setIsLogin] = useState(true);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [username, setUsername] = useState('');
  const [balance, setBalance] = useState(0);
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      const decodedToken = jwtDecodeModule.jwtDecode(token);
      const username = decodedToken.username;
      setUsername(username);
      fetchUserBalance(username);
      setIsAuthenticated(true);
    }
  }, []);

  const fetchUserBalance = async (username) => {
    try {
      const response = await fetchWithAuth(`http://127.0.0.1:8081/user/name/${username}`);
      if (response.ok) {
        const data = await response.json();
        setBalance(data.Balance);
      } else {
        console.error('Failed to fetch user balance:', response.status);
      }
    } catch (error) {
      console.error('Error fetching user balance:', error);
    }
  };

  const handleLogout = async () => {
    const response = await fetchWithAuth('http://127.0.0.1:8081/logout', { method: 'POST' });
    if (response.ok) {
      localStorage.removeItem('token');
      window.location.href = '/roulette';
    } else {
      console.error('Failed to logout:', response.status);
    }
  };

  const containerStyle = {
    display: 'flex',
    flexDirection: 'column',
    height: '100vh',
    fontFamily: 'Arial, sans-serif',
    background: `url(${backgroundImage}) no-repeat center center fixed`,
    backgroundSize: 'cover',
    color: '#fff',
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

  const tabButtonStyle = (active) => ({
    padding: '10px 20px',
    cursor: 'pointer',
    border: 'none',
    background: 'none',
    fontSize: '18px',
    fontWeight: 'bold',
    borderBottom: active ? '2px solid #28a745' : 'none',
    color: active ? '#28a745' : '#333',
  });

  const formContainerStyle = {
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
  };

  const errorStyle = {
    color: '#ff4b2b',
  };

  const successStyle = {
    color: '#28a745',
  };

  const cardStyle = {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    background: `url(${cardBackgroundImage}) no-repeat center center`,
    backgroundSize: 'cover',
    color: '#fff',
    padding: '40px', // Increase padding for larger card
    borderRadius: '10px',
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.3)',
    width: '206px', // Increase width for larger card
    height: '155px', // Increase height for larger card
    transition: 'transform 0.2s ease, box-shadow 0.2s ease', // Faster animation
    animation: 'fadeIn 1s',
    border: '2px solid #fff', // Add border
  };

  const cardButtonStyle = {
    backgroundColor: '#4CAF50', // Match the green color from the image
    color: '#fff',
    border: 'none',
    padding: '10px 20px',
    borderRadius: '5px',
    cursor: 'pointer',
    fontWeight: 'bold',
    textDecoration: 'none',
    display: 'inline-block',
    transition: 'background-color 0.2s ease, transform 0.2s ease', // Faster animation
  };

  const handleMouseEnter = (event) => {
    event.target.style.backgroundColor = '#43A047'; // Darker green for hover
    event.target.style.transform = 'scale(1.05)';
  };

  const handleMouseLeave = (event) => {
    event.target.style.backgroundColor = '#4CAF50';
    event.target.style.transform = 'scale(1)';
  };

  const cardGridStyle = {
    display: 'grid',
    gridTemplateColumns: 'repeat(auto-fill, minmax(300px, 1fr))',
    gap: '20px',
    padding: '20px',
    gridTemplateColumns: 'repeat(3, 1fr)', // Three columns per row
  };

  return (
    <div style={containerStyle}>
      <Header
        username={username}
        balance={balance}
        handleLogout={handleLogout}
        style={{ backgroundColor: '#007BFF', padding: '10px', color: '#fff' }}
      />
      <div style={mainStyle}>
        {isAuthenticated ? (
          <div style={cardGridStyle}>
            {/* Duplicate the card multiple times for demonstration */}
            {[...Array(6)].map((_, index) => (
              <div key={index} style={cardStyle}>
                <Link
                  to="/roulette"
                  style={cardButtonStyle}
                  onMouseEnter={handleMouseEnter}
                  onMouseLeave={handleMouseLeave}
                >
                  Играть
                </Link>
              </div>
            ))}
          </div>
        ) : (
          <div style={formContainerStyle}>
            <div style={tabStyle}>
              <button
                style={tabButtonStyle(isLogin)}
                onClick={() => setIsLogin(true)}
              >
                Логин
              </button>
              <button
                style={tabButtonStyle(!isLogin)}
                onClick={() => setIsLogin(false)}
              >
                Регистрация
              </button>
            </div>
            {error && <p style={errorStyle}>{error}</p>}
            {success && <p style={successStyle}>{success}</p>}
            {isLogin ? (
              <Login />
            ) : (
              <Register />
            )}
          </div>
        )}
      </div>
    </div>
  );
};

export default Home;
