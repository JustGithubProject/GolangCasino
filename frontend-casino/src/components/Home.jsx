import React, { useState, useEffect } from 'react';
import { fetchWithAuth } from './auth_components/fetchWrapper';
import Login from './auth_components/Login';
import Register from './auth_components/Register';
import * as jwtDecodeModule from 'jwt-decode';
import Header from './Header';

const Home = () => {
  const [isLogin, setIsLogin] = useState(true);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const [username, setUsername] = useState('');
  const [balance, setBalance] = useState(0);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      const decodedToken = jwtDecodeModule.jwtDecode(token);
      const username = decodedToken.username;
      setUsername(username);
      fetchUserBalance(username);
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
    background: 'linear-gradient(to right, #43cea2, #185a9d)',
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
    borderBottom: active ? '2px solid #007BFF' : 'none',
    color: active ? '#007BFF' : '#333',
  });

  return (
    <div style={containerStyle}>
      <Header username={username} balance={balance} handleLogout={handleLogout} />
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
