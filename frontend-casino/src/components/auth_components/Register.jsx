import React, { useState } from 'react';
import { fetchWithAuth } from './fetchWrapper'; // Use named import

const Register = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');
  const [balance, setBalance] = useState('');
  const [error, setError] = useState('');
  const [success, setSuccess] = useState('');

  const handleRegister = async (e) => {
    e.preventDefault();
    setError('');
    setSuccess('');
    try {
      const response = fetchWithAuth('/register/user/', {
        method: 'POST',
        body: JSON.stringify({ username, password, email, balance: parseFloat(balance) }),
      });

      if (!response.ok) {
        throw new Error('Registration failed');
      }

      const data = await response.json();
      setSuccess('Registration successful!');
    } catch (err) {
      setError('Registration failed. Please try again.');
    }
  };

  const styles = {
    authContainer: {
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'center',
      justifyContent: 'center',
      height: '100vh',
      background: 'linear-gradient(to right, #333, #555)',
      color: '#fff',
      fontFamily: 'Arial, sans-serif',
    },
    form: {
      display: 'flex',
      flexDirection: 'column',
      background: '#fff',
      padding: '20px',
      borderRadius: '8px',
      boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
      color: '#333',
    },
    formGroup: {
      marginBottom: '15px',
    },
    label: {
      marginBottom: '5px',
    },
    input: {
      padding: '10px',
      borderRadius: '20px',
      border: '1px solid #ddd',
    },
    button: {
      padding: '10px',
      borderRadius: '20px',
      border: 'none',
      background: '#ff416c',
      color: '#fff',
      cursor: 'pointer',
      fontSize: '16px',
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
      <h2>Register</h2>
      {error && <p style={styles.error}>{error}</p>}
      {success && <p style={styles.success}>{success}</p>}
      <form onSubmit={handleRegister} style={styles.form}>
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
        <div style={styles.formGroup}>
          <label style={styles.label}>Email:</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
            style={styles.input}
          />
        </div>
        <div style={styles.formGroup}>
          <label style={styles.label}>Balance:</label>
          <input
            type="number"
            step="0.01"
            value={balance}
            onChange={(e) => setBalance(e.target.value)}
            required
            style={styles.input}
          />
        </div>
        <button type="submit" style={styles.button}>Register</button>
      </form>
    </div>
  );
};

export default Register;
