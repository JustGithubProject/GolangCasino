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
    />
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
    position: 'absolute',
    top: '10px',
    right: '10px',
    background: 'transparent',
    border: 'none',
    color: '#ff4b2b',
    fontSize: '16px',
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
