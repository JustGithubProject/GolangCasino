import React from 'react';
import { Button } from 'antd';

function NumberGrid({ numbers, selectedNumbers, handleNumberClick, handleColorClick }) {
    return (
        <div style={styles.container}>
            <div style={styles.grid}>
                <div style={styles.zeroColumn}>
                    <Button
                        shape="circle"
                        size="large"
                        style={getButtonStyle(0, selectedNumbers)}
                        onClick={() => handleNumberClick(0)}
                    >
                        0
                        {selectedNumbers.includes(0) && (
                            <span style={styles.coinIcon}>💰</span>
                        )}
                    </Button>
                </div>
                <div style={styles.numbers}>
                    {numbers.map((numberRow, rowIndex) => (
                        <div key={rowIndex} style={styles.row}>
                            {numberRow.map((number, numberIndex) => (
                                <Button
                                    key={numberIndex}
                                    style={getButtonStyle(number, selectedNumbers)}
                                    onClick={() => handleNumberClick(number)}
                                >
                                    {number}
                                    {selectedNumbers.includes(number) && (
                                        <span style={styles.coinIcon}>💰</span>
                                    )}
                                </Button>
                            ))}
                        </div>
                    ))}
                    <div style={styles.labels}>
                        <div style={styles.label}>1st 12</div>
                        <div style={styles.label}>2nd 12</div>
                        <div style={styles.label}>3rd 12</div>
                    </div>
                    <div style={styles.labels}>
                        <div style={styles.label}>1 to 18</div>
                        <div style={styles.label}>EVEN</div>
                        <div style={styles.diamond}></div>
                        <div style={styles.square}></div>
                        <div style={styles.label}>ODD</div>
                        <div style={styles.label}>19 to 36</div>
                    </div>
                    <div style={styles.labels}>
                        <div style={styles.verticalLabel}>2 to 1</div>
                        <div style={styles.verticalLabel}>2 to 1</div>
                        <div style={styles.verticalLabel}>2 to 1</div>
                    </div>
                </div>
            </div>
        </div>
    );
}

const styles = {
    container: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        padding: '20px',
        backgroundColor: '#006400', // Темно-зеленый фон
    },
    grid: {
        display: 'flex',
        flexDirection: 'row', // Align columns in a row
        border: '2px solid white', // Белая рамка для сетки
        backgroundColor: '#006400', // Темно-зеленый фон
    },
    zeroColumn: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: 'green', // Темно-зеленый фон
        width: '60px', // Ширина блока 0
        borderRight: '2px solid white', // Белая рамка для нуля
    },
    hexagon: {
        position: 'relative',
        width: '60px',
        height: '60px', // Сделать высоту блока 0 равной остальным
        backgroundColor: '#006400', // Темно-зеленый фон
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
    numbers: {
        display: 'flex',
        flexDirection: 'column',
        backgroundColor: '#006400', // Темно-зеленый фон
    },
    row: {
        display: 'flex',
    },
    numberButton: {
        width: '60px', // Увеличенный размер кнопок
        height: '60px',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        fontSize: '18px', // Увеличенный размер шрифта
        fontWeight: 'bold',
        borderRadius: '8px', // Закругленные края
        border: '2px solid white', // Белая рамка для каждой кнопки
        transition: 'transform 0.2s, background-color 0.2s',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)', // Тени для 3D-эффекта
    },
    selectedButton: {
        transform: 'scale(1.1)',
        boxShadow: '0 0 10px #1890ff',
    },
    coinIcon: {
        marginLeft: '4px',
    },
    labels: {
        display: 'flex',
        justifyContent: 'space-between',
        width: '100%',
        padding: '10px 0',
        borderTop: '2px solid white', // Белая рамка сверху
        borderBottom: '2px solid white', // Белая рамка снизу
    },
    label: {
        color: 'white',
        fontSize: '16px',
        fontWeight: 'bold',
        textAlign: 'center',
        flex: 1,
        border: '2px solid white', // Белая рамка вокруг каждого лейбла
        padding: '10px',
    },
    verticalLabel: {
        color: 'white',
        fontSize: '16px',
        fontWeight: 'bold',
        textAlign: 'center',
        writingMode: 'vertical-rl',
        transform: 'rotate(180deg)',
        padding: '10px',
        flex: 1,
        border: '2px solid white', // Белая рамка вокруг каждого лейбла
    },
    diamond: {
        width: '16px',
        height: '16px',
        backgroundColor: 'red',
        transform: 'rotate(45deg)',
        border: '2px solid white', // Белая рамка вокруг ромба
    },
    square: {
        width: '16px',
        height: '16px',
        backgroundColor: 'black',
        border: '2px solid white', // Белая рамка вокруг квадрата
    }
};

function getButtonStyle(number, selectedNumbers) {
    const baseStyle = {
        ...styles.numberButton,
        backgroundColor: getColorForNumber(number),
        color: 'white',
    };

    if (selectedNumbers.includes(number)) {
        return { ...baseStyle, ...styles.selectedButton };
    }

    return baseStyle;
}

function getColorForNumber(number) {
    if (number === 0) {
        return 'green';
    } else if ([1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36].includes(number)) {
        return 'red';
    } else {
        return 'black';
    }
}

// Пример чисел для отображения в нужном порядке
const numbers = [
    [3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36],
    [2, 5, 8, 11, 14, 17, 20, 23, 26, 29, 32, 35],
    [1, 4, 7, 10, 13, 16, 19, 22, 25, 28, 31, 34]
];

export default NumberGrid;
