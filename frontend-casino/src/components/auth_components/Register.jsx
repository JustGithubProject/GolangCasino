import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { FaEye, FaEyeSlash } from 'react-icons/fa';

const Register = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');
  const navigate = useNavigate();

  const handleRegister = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');
    try {
      const response = await axios.post(
        'http://127.0.0.1:8081/register/user/',
        { username, password, email },
        {
          withCredentials: true,
          headers: {
            'Content-Type': 'application/json',
          },
        }
      );

      if (response.status !== 201) {
        throw new Error('Registration failed');
      }

      setSuccess('Registration successful!');
      navigate('/');
      window.location.reload(); 
    } catch (err) {
      setError('Registration failed. Please try again.');
    }
  };

  const toggleShowPassword = () => {
    setShowPassword(!showPassword);
  };

  const styles = {
    form: {
      display: 'flex',
      flexDirection: 'column',
      background: '#2c3e50', // Dark background color
      padding: '30px',
      borderRadius: '15px',
      boxShadow: '0 8px 20px rgba(0, 0, 0, 0.2)',
      color: '#ECF0F1',
      width: '350px',
      transition: 'transform 0.3s ease, box-shadow 0.3s ease',
    },
    formGroup: {
      marginBottom: '20px',
      position: 'relative',
    },
    label: {
      marginBottom: '8px',
      fontSize: '14px',
      fontWeight: 'bold',
      color: '#ECF0F1', // Light text color
    },
    input: {
      padding: '12px 40px 12px 12px',
      borderRadius: '20px',
      border: '1px solid #34495e',
      background: '#34495e', // Darker input background
      color: '#ECF0F1', // Light text color
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
      padding: '0',
    },
    icon: {
      verticalAlign: 'middle',
    },
  };

  return (
    <form onSubmit={handleRegister} style={styles.form}>
      <div style={styles.formGroup}>
        <label style={styles.label}>Логин:</label>
        <input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          required
          style={username ? { ...styles.input, borderColor: '#185a9d', boxShadow: '0 0 8px rgba(24, 90, 157, 0.2)' } : styles.input}
        />
      </div>
      <div style={styles.formGroup}>
        <label style={styles.label}>Пароль:</label>
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
      <div style={styles.formGroup}>
        <label style={styles.label}>Email:</label>
        <input
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
          style={email ? { ...styles.input, borderColor: '#185a9d', boxShadow: '0 0 8px rgba(24, 90, 157, 0.2)' } : styles.input}
        />
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
        Регистрация
      </button>
      {error && <div style={styles.error}>{error}</div>}
      {success && <div style={styles.success}>{success}</div>}
    </form>
  );
};

export default Register;
