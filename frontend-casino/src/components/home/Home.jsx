import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { Carousel } from 'react-responsive-carousel';
import * as jwtDecodeModule from 'jwt-decode';

import Header from '../header/Header';
import Footer from '../footer/Footer';
import Login from '../auth_components/Login';
import Register from '../auth_components/Register';
import { fetchWithAuth } from '../auth_components/fetchWrapper';

import 'react-responsive-carousel/lib/styles/carousel.min.css'; 

import backgroundImage from '../../images/backgroundCasinoNew.jpg';
import cardBackgroundImage from '../../images/card.png'; 
import sweetbonanzaImage from '../../images/sweet-bonanza.png'; 
import doghouseImage from '../../images/dog-house.png'; 
import image1 from '../../images/cas_image_99.png';
import image2 from '../../images/cas_image_9.png';
import image3 from '../../images/cas_image_11.png';


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

    const interval = setInterval(() => {
      localStorage.removeItem('token');
      setIsAuthenticated(false);
      setUsername('');
      setBalance(0);
    }, 30 * 60 * 1000); // в миллисекундах

    return () => clearInterval(interval);
  }, []);

  useEffect(() => {
    const checkTokenExpiration = () => {
      const token = localStorage.getItem('authToken');

      if (token) {
        try {
          const decodedToken = jwtDecodeModule.jwtDecode(token);
          const expirationTime = decodedToken.exp;
          const currentTime = Math.floor(Date.now() / 1000);

          if (expirationTime < currentTime) {
            localStorage.removeItem('authToken');
            console.log('Token has been deleted');
          } else {
            console.log('Correct Token');
          }
        } catch (error) {
          console.error('Failed to decode token', error);
        }
      } else {
        console.log('Token not found');
      }
    };

    checkTokenExpiration();
  }, []);
  

  useEffect(() => {
    const queryParams = new URLSearchParams(location.search);
    const token = localStorage.getItem("token");
    const orderID = queryParams.get('token'); 
    if (token && orderID) { 
      fetch(`http://127.0.0.1:8081/paypal/update/approved/order/?token=${orderID}`, { 
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`, 
        },
      })
      .then(response => response.json())
      .then(data => {
        console.log('Success:', data);
        window.location.href = "/sucess-payment"
      })
      .catch(error => {
        console.error('Error:', error);
      });
    }
  }, [location.search]);

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

  const cardStyle = (index) => {
    let backgroundImage;
  
    if (index === 4) {
      backgroundImage = `url(${sweetbonanzaImage}) no-repeat center center`;
    } else if (index === 5) {
      backgroundImage = `url(${doghouseImage}) no-repeat center center`;
    } else {
      backgroundImage = `url(${cardBackgroundImage}) no-repeat center center`;
    }
  
    return {
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'center',
      justifyContent: 'center',
      background: backgroundImage,
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

  const carouselContainerStyle = {
    marginTop: '80px', // Space below the header
    width: '100%',
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
      />
      <div style={carouselContainerStyle}>
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
      </div>
      <div style={mainStyle}>
        {isAuthenticated ? (
          <div style={cardGridStyle}>
          {[...Array(6)].map((_, index) => (
          <div key={index} style={{ textAlign: 'center' }}>
            <div style={cardStyle(index)}>
              <Link
                to={index === 4 ? '/room/slot/sweetbonanza' : (index === 5 ? '/room/slot/doghouse' : `/room/roulette/${index + 1}`)}
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
