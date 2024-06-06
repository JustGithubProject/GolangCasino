import React from 'react';
import { Typography } from 'antd';
import { CSSTransition } from 'react-transition-group';

const { Text } = Typography;

const ResultOverlay = ({ showResult, resultMessage, spinResult }) => (
  <CSSTransition
    in={showResult}
    timeout={300}
    classNames="fade"
    unmountOnExit
  >
    <div style={styles.resultOverlay}>
      <Text style={styles.resultText}>{resultMessage}<p></p></Text>
      {spinResult !== null && <Text style={styles.spinResultText}>Выпавшее число: {spinResult}</Text>}
    </div>
  </CSSTransition>
);

const styles = {
  resultOverlay: {
    position: 'fixed',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    backgroundColor: 'rgba(0, 0, 0, 0.85)',
    padding: '30px',
    borderRadius: '20px',
    boxShadow: '0 6px 20px rgba(0, 0, 0, 0.5)',
    textAlign: 'center',
    zIndex: 1000,
  },
  resultText: {
    fontSize: '26px',
    fontWeight: 'bold',
    color: '#fff',
    marginBottom: '15px',
  },
  spinResultText: {
    fontSize: '22px',
    fontWeight: 'bold',
    color: '#fff',
  },
};

export default ResultOverlay;
