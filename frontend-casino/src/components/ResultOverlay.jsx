import React from 'react';
import { Typography } from 'antd';
import { CSSTransition } from 'react-transition-group';

const { Text } = Typography;

const ResultOverlay = ({ showResult, resultMessage, spinResult }) => (
  <CSSTransition
    in={showResult}
    timeout={150}
    classNames="fade"
    unmountOnExit
  >
    <div style={styles.resultOverlay}>
      <Text style={styles.resultText}>
        {resultMessage}
        {spinResult !== null && (
          <>
            <br />
            <Text style={styles.spinResultText}>Выпавшее число: {spinResult}</Text>
          </>
        )}
      </Text>
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
    padding: '40px 50px',
    borderRadius: '20px',
    boxShadow: '0 6px 20px rgba(0, 0, 0, 0.5)',
    textAlign: 'center',
    zIndex: 1000,
    width: '80%', // Added width to constrain the text
    maxWidth: '600px', // Added max-width to limit the container size
    whiteSpace: 'normal', // Allow text to wrap
  },
  resultText: {
    fontSize: '28px',
    fontWeight: 'bold',
    color: '#ffffff',
    marginBottom: '20px',
    lineHeight: '1.5',
    textShadow: '2px 2px 4px rgba(0, 0, 0, 0.5)',
  },
  spinResultText: {
    fontSize: '24px',
    fontWeight: 'bold',
    color: '#ffdd57',
    textShadow: '1px 1px 3px rgba(0, 0, 0, 0.5)',
  },
};

export default ResultOverlay;
