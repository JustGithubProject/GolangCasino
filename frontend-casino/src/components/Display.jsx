import { Card, Typography, Space } from 'antd';
import { useEffect, useState } from 'react';

const { Text } = Typography;

function Display({ selectedNumbers, selectedColor, spinResult, isSpinning, betValues }) {
    const [displayedNumber, setDisplayedNumber] = useState(null);

    useEffect(() => {
        if (isSpinning) {
            let timeout;
            let currentNumber = 0;
            const spinInterval = setInterval(() => {
                setDisplayedNumber(currentNumber);
                currentNumber = (currentNumber + 1) % 37;
            }, 100);

            timeout = setTimeout(() => {
                clearInterval(spinInterval);
                setDisplayedNumber(spinResult);
            }, 5000);

            return () => {
                clearInterval(spinInterval);
                clearTimeout(timeout);
            };
        }
    }, [isSpinning, spinResult]);

    const getSelectedBetsDisplay = () => {
        const bets = [];

        if (selectedNumbers.length > 0) {
            bets.push(`Числа: ${selectedNumbers.join(', ')}`);
        }
        if (selectedColor) {
            bets.push(`Цвет: ${selectedColor === 'red' ? 'Красное' : 'Черное'}`);
        }
        Object.keys(betValues).forEach((key) => {
            if (betValues[key] && betValues[key] !== '') {
                const value = parseInt(betValues[key], 10);
                if (value > 0) {
                    bets.push(`${key}: ₽${value}`);
                }
            }
        });

        return bets.length > 0 ? bets.join(' | ') : 'Нет ставок';
    };

    return (
        <div style={styles.displayContainer}>
            <Card style={{ ...styles.card, ...(isSpinning && styles.cardSpinning) }}>
                <Space direction="vertical" align="center">
                    {isSpinning ? (
                        <Text strong style={styles.numberText}>
                            {displayedNumber !== null && <span>{displayedNumber}</span>}
                        </Text>
                    ) : (
                        <>
                            <Text strong style={styles.numberText}>
                                {spinResult !== null ? `Выпавшее число: ${spinResult}` : 'Выберите ставку'}
                            </Text>
                            <Text style={styles.betText}>{getSelectedBetsDisplay()}</Text>
                        </>
                    )}
                </Space>
            </Card>
        </div>
    );
}

const styles = {
    displayContainer: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '200px',
        marginBottom: '20px',
        background: 'linear-gradient(135deg, #ececec 0%, #f5f5f5 100%)',
    },
    card: {
        width: '100%',
        maxWidth: '900px',
        backgroundColor: '#ffffff',
        borderRadius: '12px',
        boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        padding: '30px',
        transition: 'all 0.5s ease',
        textAlign: 'center',
        transform: 'scale(1)',
    },
    cardSpinning: {
        maxWidth: '1100px',
        boxShadow: '0 0 20px #1890ff',
        animation: 'spinEffect 5s cubic-bezier(0.68, -0.55, 0.27, 1.55)',
        transform: 'scale(1.05)',
    },
    numberText: {
        fontSize: '40px',
        fontWeight: 'bold',
        color: '#1890ff',
    },
    betText: {
        fontSize: '20px',
        color: '#595959',
        marginTop: '10px',
    },
    '@keyframes spinEffect': {
        '0%': { transform: 'rotate(0deg)' },
        '80%': { transform: 'rotate(2880deg)' },
        '90%': { transform: 'rotate(2950deg)' },
        '100%': { transform: 'rotate(2880deg)' },
    },
};

export default Display;
