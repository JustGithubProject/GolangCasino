import React, { useState, useEffect } from 'react';
import { fetchWithAuth } from './auth_components/fetchWrapper';
import Login from './auth_components/Login';
import Register from './auth_components/Register';
import * as jwtDecodeModule from 'jwt-decode';
import { Link } from 'react-router-dom';
import Header from './Header';
import Footer from './Footer';
import { Carousel } from 'react-responsive-carousel';
import 'react-responsive-carousel/lib/styles/carousel.min.css'; // Import carousel styles
import backgroundImage from '../images/casinoImage_2.png';
import cardBackgroundImage from '../images/card.png'; // Import the card background image

import image1 from '../images/cas_image_99.png';
import image2 from '../images/cas_image_9.png';
import image3 from '../images/cas_image_11.png';

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
      window.location.href = '/';
    } else {
      console.error('Failed to logout:', response.status);
    }
  };

  const containerStyle = {
    display: 'flex',
    flexDirection: 'column',
    height: '100vh',
    fontFamily: '"Roboto", sans-serif',
    background: `url(${backgroundImage}) no-repeat center center fixed`,
    backgroundSize: 'cover',
    color: '#fff',
    position: 'relative',
    paddingBottom: '60px', // Space for the footer
    overflow: 'hidden', // Prevent scrolling
  };

  const mainStyle = {
    flex: 1,
    padding: '20px',
    display: 'flex',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
  };

  const tabStyle = {
    display: 'flex',
    justifyContent: 'center',
    marginBottom: '10px',
  };

  const tabButtonStyle = (active) => ({
    padding: '5px 10px',
    cursor: 'pointer',
    border: 'none',
    background: 'none',
    fontSize: '16px',
    fontWeight: 'bold',
    borderBottom: active ? '2px solid #28a745' : 'none',
    color: active ? '#28a745' : '#fff', // Dark theme color
  });

  const formContainerStyle = {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    background: '#2c3e50', // Dark background color
    color: '#ECF0F1', // Light text color
    padding: '20px',
    borderRadius: '10px',
    boxShadow: '0 8px 20px rgba(0, 0, 0, 0.2)',
    transition: 'transform 0.3s ease, box-shadow 0.3s ease',
    width: '350px', // Increased width
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
    padding: '30px',
    borderRadius: '10px',
    boxShadow: '0 8px 16px rgba(0, 0, 0, 0.3)',
    width: '200px',
    height: '160px',
    transition: 'transform 0.3s ease, box-shadow 0.3s ease',
    animation: 'fadeIn 1s',
    border: '2px solid rgba(255, 255, 255, 0.8)',
    marginBottom: '5px',
  };

  const cardButtonStyle = {
    backgroundColor: '#4CAF50',
    color: '#fff',
    border: 'none',
    padding: '8px 16px',
    borderRadius: '5px',
    cursor: 'pointer',
    fontWeight: 'bold',
    textDecoration: 'none',
    display: 'inline-block',
    transition: 'background-color 0.3s ease, transform 0.3s ease',
  };

  const handleMouseEnter = (event) => {
    event.target.style.backgroundColor = '#388E3C';
    event.target.style.transform = 'scale(1.05)';
  };

  const handleMouseLeave = (event) => {
    event.target.style.backgroundColor = '#4CAF50';
    event.target.style.transform = 'scale(1)';
  };

  const cardGridStyle = {
    display: 'grid',
    gridTemplateColumns: 'repeat(auto-fill, minmax(250px, 1fr))',
    gap: '10px',
    padding: '10px',
    gridTemplateColumns: 'repeat(3, 1fr)',
  };

  const carouselImageStyle = {
    maxHeight: '250px',
    objectFit: 'cover',
  };

  const cardLabelStyle = {
    backgroundColor: '#000',
    color: '#fff',
    fontSize: '14px',
    margin: 0,
    padding: '5px 0',
    borderRadius: '0 0 10px 10px',
  };

  return (
    <div style={containerStyle}>
      <Header
        username={username}
        balance={balance}
        handleLogout={handleLogout}
        style={{ backgroundColor: '#007BFF', padding: '10px', color: '#fff' }}
      />
      <Carousel
        showThumbs={false}
        autoPlay
        infiniteLoop
        interval={1200}
        dynamicHeight={false}
        showStatus={false}
        showIndicators={false}
      >
        <div>
          <img src={image1} alt="Image 1" style={carouselImageStyle} />
        </div>
        <div>
          <img src={image2} alt="Image 2" style={carouselImageStyle} />
        </div>
        <div>
          <img src={image3} alt="Image 3" style={carouselImageStyle} />
        </div>
      </Carousel>
      <div style={mainStyle}>
        {isAuthenticated ? (
          <div style={cardGridStyle}>
            {[...Array(6)].map((_, index) => (
              <div key={index} style={{ textAlign: 'center' }}>
                <div style={cardStyle}>
                  <Link
                    to="/roulette"
                    style={cardButtonStyle}
                    onMouseEnter={handleMouseEnter}
                    onMouseLeave={handleMouseLeave}
                  >
                    Играть
                  </Link>
                </div>
                <p style={cardLabelStyle}>{`Комната ${index + 1}`}</p>
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
                Войти
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
      <Footer />
    </div>
  );
};

export default Home;
