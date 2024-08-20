import React, { useState } from 'react';
import axios from 'axios';

const WithdrawFundsPage = () => {
  const [amount, setAmount] = useState('');
  const [currency, setCurrency] = useState('USD');
  const [email, setEmail] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
        const url = "http://127.0.0.1:8081/paypal/withdraw/funds/"
        const token = localStorage.getItem("token");
        const response = await axios.post(url, {
                total: amount,
                currency: currency,
                receiver_email: email
            },
            {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${token}`,
            },
        });

    } catch (error) {
        console.log("Failed to execute request: ", error);
    }
    
  };

  return (
    <div style={styles.container}>
      <h2 style={styles.heading}>Снять средства</h2>
      <form className="withdraw-form" onSubmit={handleSubmit}>
        <div style={styles.formGroup}>
          <label style={styles.label} htmlFor="amount">Сумма</label>
          <input
            type="number"
            id="amount"
            value={amount}
            onChange={(e) => setAmount(e.target.value)}
            placeholder="Enter amount"
            required
            style={styles.input}
          />
        </div>

        <div style={styles.formGroup}>
          <label style={styles.label} htmlFor="currency">Валюта</label>
          <select
            id="currency"
            value={currency}
            onChange={(e) => setCurrency(e.target.value)}
            required
            style={styles.input}
          >
            <option value="USD">USD</option>
            <option value="EUR">EUR</option>
            <option value="GBP">GBP</option>
            <option value="JPY">JPY</option>
          </select>
        </div>

        <div style={styles.formGroup}>
          <label style={styles.label} htmlFor="email">Электронная почта получателя PayPal</label>
          <input
            type="email"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="example@paypal.com"
            required
            style={styles.input}
          />
        </div>

        <button
          type="submit"
          style={{ ...styles.submitButton, ...styles.submitButtonHover }}
        >
          Снять средства
        </button>
      </form>
    </div>
  );
};

const styles = {
    container: {
      maxWidth: '400px',
      margin: '50px auto',
      padding: '30px',
      backgroundColor: '#ffffff',
      boxShadow: '0 8px 20px rgba(0, 0, 0, 0.1)',
      borderRadius: '8px',
      textAlign: 'center',
      fontFamily: "'Segoe UI', Tahoma, Geneva, Verdana, sans-serif",
    },
    heading: {
      marginBottom: '20px',
      color: '#333',
    },
    formGroup: {
      marginBottom: '15px',
      textAlign: 'left',
    },
    label: {
      display: 'block',
      marginBottom: '5px',
      color: '#555',
      fontWeight: 600,
    },
    input: {
      width: '100%',
      padding: '10px',
      border: '1px solid #ddd',
      borderRadius: '5px',
      fontSize: '16px',
      outline: 'none',
      transition: 'border-color 0.3s',
    },
    inputFocus: {
      borderColor: '#007bff',
    },
    submitButton: {
      width: '100%',
      padding: '12px',
      backgroundColor: '#007bff',
      border: 'none',
      borderRadius: '5px',
      color: '#ffffff',
      fontSize: '18px',
      cursor: 'pointer',
      transition: 'background-color 0.3s',
    },
    submitButtonHover: {
      backgroundColor: '#0056b3',
    },
  };

export default WithdrawFundsPage;
