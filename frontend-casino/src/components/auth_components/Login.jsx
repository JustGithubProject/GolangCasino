import React, { useState } from 'react';

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleLogin = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');
    try {
      const response = await fetch('http://localhost:8081/login/user', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
      });

      if (!response.ok) {
        throw new Error('Login failed');
      }

      const data = await response.json();
      const token = data.token; // Assuming your backend returns the token in the "token" field
      localStorage.setItem('token', token); // Store the token in localStorage

      setSuccess('Login successful!');
    } catch (err) {
      setError('Login failed. Please try again.');
    }
  };

  const styles = {
    authContainer: {
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'center',
      justifyContent: 'center',
      height: '100vh',
      background: 'linear-gradient(to right, #0f2027, #203a43, #2c5364)',
      color: '#fff',
      fontFamily: 'Arial, sans-serif',
      padding: '20px',
    },
    form: {
      display: 'flex',
      flexDirection: 'column',
      background: '#fff',
      padding: '30px',
      borderRadius: '10px',
      boxShadow: '0 4px 15px rgba(0, 0, 0, 0.2)',
      color: '#333',
      width: '300px',
    },
    formGroup: {
      marginBottom: '20px',
    },
    label: {
      marginBottom: '8px',
      fontSize: '14px',
    },
    input: {
      padding: '12px',
      borderRadius: '20px',
      border: '1px solid #ddd',
      width: '100%',
      fontSize: '14px',
    },
    button: {
      padding: '12px',
      borderRadius: '20px',
      border: 'none',
      background: 'linear-gradient(to right, #ff416c, #ff4b2b)',
      color: '#fff',
      cursor: 'pointer',
      fontSize: '16px',
      transition: 'background 0.3s ease',
    },
    error: {
      color: '#ff4b2b',
      marginBottom: '10px',
    },
    success: {
      color: '#28a745',
      marginBottom: '10px',
    },
  };

  return (
    <div style={styles.authContainer}>
      <h2>Login</h2>
      {error && <p style={styles.error}>{error}</p>}
      {success && <p style={styles.success}>{success}</p>}
      <form onSubmit={handleLogin} style={styles.form}>
        <div style={styles.formGroup}>
          <label style={styles.label}>Username:</label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
            style={styles.input}
          />
        </div>
        <div style={styles.formGroup}>
          <label style={styles.label}>Password:</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            style={styles.input}
          />
        </div>
        <button
          type="submit"
          style={styles.button}
          onMouseOver={(e) => e.target.style.background = 'linear-gradient(to right, #ff4b2b, #ff416c)'}
          onMouseOut={(e) => e.target.style.background = 'linear-gradient(to right, #ff416c, #ff4b2b)'}
        >
          Login
        </button>
      </form>
    </div>
  );
};

export default Login;
