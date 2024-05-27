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

    const getColorForNumber = (number) => {
        if (number === 0) return 'green';
        const redNumbers = [1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36];
        return redNumbers.includes(number) ? 'red' : 'black';
    };

    const spinResultColor = getColorForNumber(spinResult);

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
                            <Text strong style={styles.resultText}>
                                Выпавшее число: <span style={{ color: spinResultColor }}>{spinResult !== null ? spinResult : 'Выберите ставку'}</span>
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
        background: 'linear-gradient(135deg, #e0e0e0 0%, #f0f0f0 100%)',
        borderRadius: '12px',
        padding: '10px',
    },
    card: {
        width: '100%',
        maxWidth: '900px',
        backgroundColor: '#ffffff',
        borderRadius: '16px',
        boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        padding: '30px',
        transition: 'all 0.5s ease',
        textAlign: 'center',
        transform: 'scale(1)',
        border: '1px solid #d9d9d9',
    },
    cardSpinning: {
        maxWidth: '1100px',
        boxShadow: '0 0 20px #1890ff',
        animation: 'spinEffect 5s cubic-bezier(0.68, -0.55, 0.27, 1.55)',
        transform: 'scale(1.1)',
        border: '1px solid #1890ff',
    },
    numberText: {
        fontSize: '48px',
        fontWeight: 'bold',
        color: '#1890ff',
        textShadow: '2px 2px 4px rgba(0, 0, 0, 0.2)',
        transition: 'color 0.5s ease',
    },
    resultText: {
        fontSize: '48px',
        fontWeight: 'bold',
        color: '#000000',
        textShadow: '2px 2px 4px rgba(0, 0, 0, 0.2)',
    },
    betText: {
        fontSize: '20px',
        color: '#595959',
        marginTop: '10px',
        fontStyle: 'italic',
    },
    '@keyframes spinEffect': {
        '0%': { transform: 'rotate(0deg)' },
        '80%': { transform: 'rotate(2880deg)' },
        '90%': { transform: 'rotate(2950deg)' },
        '100%': { transform: 'rotate(2880deg)' },
    },
};

export default Display;
