import React from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';

const Home = () => {
  const handleLogout = async () => {
    try {
      await axios.post('http://127.0.0.1:8081/logout');

      localStorage.removeItem('authToken');

      window.location.href = '/login';
    } catch (error) {
      console.error('Failed to logout:', error);
    }
  };

  const containerStyle = {
    display: 'flex',
    flexDirection: 'row',
    height: '100vh',
    fontFamily: 'Arial, sans-serif',
    backgroundColor: '#f0f0f0'
  };

  const navStyle = {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'flex-start',
    width: '250px',
    padding: '20px',
    backgroundColor: '#333',
    color: '#fff',
    boxShadow: '2px 0 5px rgba(0, 0, 0, 0.1)'
  };

  const ulStyle = {
    listStyleType: 'none',
    padding: '0',
    margin: '0',
    width: '100%'
  };

  const liStyle = {
    margin: '15px 0'
  };

  const linkStyle = {
    textDecoration: 'none',
    color: '#fff',
    fontWeight: 'bold',
    fontSize: '18px',
    width: '100%',
    display: 'block',
    padding: '10px 15px',
    borderRadius: '4px',
    transition: 'background-color 0.3s'
  };

  const linkHoverStyle = {
    backgroundColor: '#007BFF'
  };

  const buttonStyle = {
    backgroundColor: '#007BFF',
    color: '#fff',
    border: 'none',
    padding: '10px 15px',
    fontSize: '18px',
    fontWeight: 'bold',
    width: '100%',
    borderRadius: '4px',
    cursor: 'pointer',
    transition: 'background-color 0.3s'
  };

  return (
    <div style={containerStyle}>
      <nav style={navStyle}>
        <ul style={ulStyle}>
          <li style={liStyle}>
            <Link 
              to="/roulette" 
              style={linkStyle} 
              onMouseEnter={(e) => e.currentTarget.style.backgroundColor = linkHoverStyle.backgroundColor} 
              onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
            >
              Roulette
            </Link>
          </li>
          <li style={liStyle}>
            <Link 
              to="/login" 
              style={linkStyle} 
              onMouseEnter={(e) => e.currentTarget.style.backgroundColor = linkHoverStyle.backgroundColor} 
              onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
            >
              Login
            </Link>
          </li>
          <li style={liStyle}>
            <Link 
              to="/register" 
              style={linkStyle} 
              onMouseEnter={(e) => e.currentTarget.style.backgroundColor = linkHoverStyle.backgroundColor} 
              onMouseLeave={(e) => e.currentTarget.style.backgroundColor = 'transparent'}
            >
              Register
            </Link>
          </li>
          <li style={liStyle}>
            <button 
              onClick={handleLogout} 
              style={buttonStyle} 
              onMouseEnter={(e) => e.currentTarget.style.backgroundColor = '#0056b3'}
              onMouseLeave={(e) => e.currentTarget.style.backgroundColor = '#007BFF'}
            >
              Logout
            </button>
          </li>
        </ul>
      </nav>
      <div style={{ flex: 1, padding: '40px' }}>
        <h1>Welcome to Our Site</h1>
        <p>Select an option from the navigation menu to get started.</p>
      </div>
    </div>
  );
};

export default Home;
