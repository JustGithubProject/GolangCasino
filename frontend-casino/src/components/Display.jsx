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
                currentNumber = (currentNumber + 1) % 37; // Ограничить числа до 36
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
            <Card style={styles.card}>
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
        height: '100%',
        padding: '15px',
    },
    card: {
        width: '100%',
        maxWidth: '700px',
        backgroundColor: 'white',
        borderRadius: '8px',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        padding: '20px',
    },
    numberText: {
        fontSize: '24px',
        fontWeight: 'bold',
        textAlign: 'center',
    },
    colorText: {
        fontSize: '20px',
        marginTop: '10px',
        textAlign: 'center',
    },
};

export default Display;
