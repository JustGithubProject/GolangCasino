import React from 'react';
import { Button, Spin } from 'antd';

const BetButton = ({ isSpinning, handleSubmit, handleReset }) => (
  <div style={styles.submitButtonContainer}>
    <Button type="primary" size="large" onClick={handleSubmit} style={styles.submitButton} disabled={isSpinning}>
      {isSpinning ? <Spin /> : 'Вращать'}
    </Button>
    <Button type="default" size="large" onClick={handleReset} style={styles.resetButton} disabled={isSpinning}>
      Очистить
    </Button>
  </div>
);

const styles = {
  submitButtonContainer: {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    marginTop: '32px',
  },
  submitButton: {
    padding: '12px 60px',
    fontSize: '16px',
    fontWeight: 'bold',
    borderRadius: '20px',
    marginLeft: '15px',
    background: 'linear-gradient(to right, #ff416c, #ff4b2b)',
    border: 'none',
    color: 'white',
    transition: 'background 0.3s ease, transform 0.3s ease',
    boxShadow: '0 4px 10px rgba(0, 0, 0, 0.2)',
  },
  resetButton: {
    padding: '12px 40px',
    fontSize: '16px',
    fontWeight: 'bold',
    borderRadius: '20px',
    marginLeft: '15px',
    background: 'linear-gradient(to right, #6a11cb, #2575fc)',
    border: 'none',
    color: 'white',
    transition: 'background 0.3s ease, transform 0.3s ease',
    boxShadow: '0 4px 10px rgba(0, 0, 0, 0.2)',
  },
};

export default BetButton;
