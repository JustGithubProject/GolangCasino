import { Card, Typography, Space } from 'antd';
import { CheckCircleTwoTone, CloseCircleTwoTone } from '@ant-design/icons';
import { useEffect, useState } from 'react';

const { Text } = Typography;

function Display({ selectedNumbers, selectedColor, spinResult, isSpinning }) {
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
                            {selectedNumbers.length > 0 ? (
                                <Text strong style={styles.numberText}>
                                    <CheckCircleTwoTone twoToneColor="#52c41a" /> Выбранные числа: {selectedNumbers.join(', ')}
                                </Text>
                            ) : (
                                <Text strong style={styles.numberText}>
                                    <CloseCircleTwoTone twoToneColor="#eb2f96" /> Выберите числа от 0 до 36
                                </Text>
                            )}
                            {selectedColor && (
                                <Text style={styles.colorText}>
                                    <span style={{ color: selectedColor === 'red' ? '#ff4d4f' : '#595959' }}>
                                        {selectedColor === 'red' ? '🔴 Красный' : '⚫ Черный'}
                                    </span>
                                </Text>
                            )}
                            {spinResult !== null && (
                                <Text strong style={styles.resultText}>
                                    Выпавшее число: {spinResult}
                                </Text>
                            )}
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
    },
    card: {
        width: '100%',
        maxWidth: '900px',
        backgroundColor: 'white',
        borderRadius: '12px',
        boxShadow: '0 8px 16px rgba(0, 0, 0, 0.2)',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        padding: '30px',
        transition: 'all 0.5s ease',
        textAlign: 'center',
    },
    cardSpinning: {
        maxWidth: '1100px',
        boxShadow: '0 0 20px #1890ff',
        animation: 'spinEffect 5s cubic-bezier(0.68, -0.55, 0.27, 1.55)',
    },
    numberText: {
        fontSize: '32px',
        fontWeight: 'bold',
        color: '#1890ff',
    },
    colorText: {
        fontSize: '24px',
        marginTop: '15px',
    },
    resultText: {
        fontSize: '32px',
        fontWeight: 'bold',
        color: '#ff4d4f',
        marginTop: '20px',
    },
    '@keyframes spinEffect': {
        '0%': { transform: 'rotate(0deg)' },
        '80%': { transform: 'rotate(2880deg)' },
        '90%': { transform: 'rotate(2950deg)' },
        '100%': { transform: 'rotate(2880deg)' },
    },
};

export default Display;
