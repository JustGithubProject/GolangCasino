import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { FaEye, FaEyeSlash } from 'react-icons/fa';

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const navigate = useNavigate();

  const handleLogin = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');
    try {
      const response = await axios.post(
        'http://127.0.0.1:8081/login/user/',
        { username, password },
        {
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json',
          },
        }
      );

      const data = response.data;
      const token = data.token;
      localStorage.setItem('token', token);

      setSuccess('Login successful!');
      navigate('/roulette');
    } catch (err) {
      setError('Login failed. Please try again.');
    }
  };

  const toggleShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const styles = {
    form: {
      display: 'flex',
      flexDirection: 'column',
      background: '#fff',
      padding: '30px',
      borderRadius: '15px',
      boxShadow: '0 8px 20px rgba(0, 0, 0, 0.2)',
      color: '#333',
      width: '350px',
      transition: 'transform 0.3s ease, box-shadow 0.3s ease',
    },
    formGroup: {
      marginBottom: '20px',
      position: 'relative', // Add this to position the toggle button correctly
    },
    label: {
      marginBottom: '8px',
      fontSize: '14px',
      fontWeight: 'bold',
    },
    input: {
      padding: '12px 40px 12px 12px', // Add right padding to make space for the icon
      borderRadius: '20px',
      border: '1px solid #ddd',
      width: '100%',
      fontSize: '14px',
      outline: 'none',
      transition: 'border-color 0.3s ease, box-shadow 0.3s ease',
    },
    button: {
      padding: '12px',
      borderRadius: '20px',
      border: 'none',
      background: 'linear-gradient(to right, #ff416c, #ff4b2b)',
      color: '#fff',
      cursor: 'pointer',
      fontSize: '16px',
      transition: 'background 0.3s ease, transform 0.3s ease',
    },
    buttonHover: {
      background: 'linear-gradient(to right, #ff4b2b, #ff416c)',
      transform: 'scale(1.05)',
    },
    error: {
      color: '#ff4b2b',
      marginBottom: '10px',
    },
    success: {
      color: '#28a745',
      marginBottom: '10px',
    },
    toggleButton: {
      position: 'absolute',
      right: '10px',
      top: '50%',
      transform: 'translateY(-50%)',
      background: 'none',
      border: 'none',
      cursor: 'pointer',
      fontSize: '18px',
      color: '#aaa',
      padding: '0', // Ensure the button has no padding for correct positioning
    },
    icon: {
      verticalAlign: 'middle',
    }
  };

  return (
    <form onSubmit={handleLogin} style={styles.form}>
      <div style={styles.formGroup}>
        <label style={styles.label}>Username:</label>
        <input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
          style={username ? { ...styles.input, borderColor: '#185a9d', boxShadow: '0 0 8px rgba(24, 90, 157, 0.2)' } : styles.input}
        />
      </div>
      <div style={styles.formGroup}>
        <label style={styles.label}>Password:</label>
        <input
          type={showPassword ? 'text' : 'password'}
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
          style={password ? { ...styles.input, borderColor: '#185a9d', boxShadow: '0 0 8px rgba(24, 90, 157, 0.2)' } : styles.input}
        />
        <button type="button" onClick={toggleShowPassword} style={styles.toggleButton}>
          {showPassword ? <FaEyeSlash style={styles.icon} /> : <FaEye style={styles.icon} />}
        </button>
      </div>
      <button
        type="submit"
        style={styles.button}
        onMouseOver={(e) => {
          e.currentTarget.style.background = styles.buttonHover.background;
          e.currentTarget.style.transform = styles.buttonHover.transform;
        }}
        onMouseOut={(e) => {
          e.currentTarget.style.background = styles.button.background;
          e.currentTarget.style.transform = 'scale(1)';
        }}
      >
        Login
      </button>
      {error && <div style={styles.error}>{error}</div>}
      {success && <div style={styles.success}>{success}</div>}
    </form>
  );
};

export default Login;
