import React from 'react';
import { Typography } from 'antd';

const { Text } = Typography;

const BalanceDisplay = ({ username, balance }) => (
  <div style={styles.balanceContainer}>
    {username && <Text style={styles.balanceText}>Пользователь: {username}</Text>}
    {balance !== null && <Text style={styles.balanceText}>Баланс: ₽{balance}</Text>}
  </div>
);

const styles = {
  balanceContainer: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    marginBottom: '20px',
  },
  balanceText: {
    fontSize: '20px',
    fontWeight: 'bold',
    color: '#fff',
  },
};

export default BalanceDisplay;
