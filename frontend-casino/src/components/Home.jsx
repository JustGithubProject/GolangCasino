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
    background: 'linear-gradient(to right, #43cea2, #185a9d)',
    color: '#fff',
  };

  const navStyle = {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'flex-start',
    width: '250px',
    padding: '20px',
    backgroundColor: '#333',
    color: '#fff',
    boxShadow: '2px 0 5px rgba(0, 0, 0, 0.1)',
    borderTopLeftRadius: '10px',
    borderBottomLeftRadius: '10px',
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
    fontSize: '18px',
    width: '100%',
    display: 'block',
    padding: '10px 15px',
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
    padding: '10px 15px',
    fontSize: '18px',
    fontWeight: 'bold',
    width: '100%',
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

  const headingStyle = {
    fontSize: '48px',
    marginBottom: '20px',
    textAlign: 'center',
    textShadow: '2px 2px 4px rgba(0, 0, 0, 0.2)',
  };

  const paragraphStyle = {
    fontSize: '24px',
    textAlign: 'center',
    maxWidth: '600px',
    lineHeight: '1.5',
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
            <Link
              to="/login"
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
              Login
            </Link>
          </li>
          <li style={liStyle}>
            <Link
              to="/register"
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
              Register
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
        <h1 style={headingStyle}>Welcome to Our Site</h1>
        <p style={paragraphStyle}>
          Select an option from the navigation menu to get started.
        </p>
      </div>
    </div>
  );
};

export default Home;
