import React from 'react';
import { Button } from 'antd';
import { CloseOutlined, PlusOutlined } from '@ant-design/icons';

const BetFormToggle = ({ showBetForm, toggleBetForm }) => (
  showBetForm ? (
    <Button
      type="text"
      icon={<CloseOutlined />}
      onClick={toggleBetForm}
      style={styles.closeButton}
    >
      Hide Bet Form
    </Button>
  ) : (
    <Button
      type="default"
      icon={<PlusOutlined />}
      onClick={toggleBetForm}
      style={styles.openButton}
    >
      Show Bet Form
    </Button>
  )
);

const styles = {
  closeButton: {
    marginTop: '20px',
    fontSize: '16px',
    fontWeight: 'bold',
    borderRadius: '12px',
    background: 'linear-gradient(to right, #ff4b2b, #ff416c)',
    border: 'none',
    color: 'white',
    transition: 'background 0.3s ease, transform 0.3s ease',
    boxShadow: '0 4px 10px rgba(0, 0, 0, 0.2)',
  },
  openButton: {
    marginTop: '20px',
    fontSize: '16px',
    fontWeight: 'bold',
    borderRadius: '12px',
    background: 'linear-gradient(to right, #6a11cb, #2575fc)',
    border: 'none',
    color: 'white',
    transition: 'background 0.3s ease, transform 0.3s ease',
    boxShadow: '0 4px 10px rgba(0, 0, 0, 0.2)',
  },
};

export default BetFormToggle;
