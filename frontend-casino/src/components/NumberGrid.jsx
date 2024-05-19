import React from 'react';
import { Button } from 'antd';

function NumberGrid({ numbers, selectedNumbers, handleNumberClick }) {
    return (
        <div style={styles.container}>
            <div style={styles.grid}>
                {numbers.map((numberRow, rowIndex) => (
                    <div key={rowIndex} style={styles.row}>
                        {numberRow.map((number, numberIndex) => (
                            <Button
                                key={numberIndex}
                                shape="circle"
                                size="large"
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
                <div style={styles.bottomRow}>
                    <div style={styles.label}>1st 12</div>
                    <div style={styles.label}>2nd 12</div>
                    <div style={styles.label}>3rd 12</div>
                </div>
                <div style={styles.bottomRow}>
                    <div style={styles.label}>1 to 18</div>
                    <div style={styles.label}>EVEN</div>
                    <div style={styles.label}>&#x25A0;</div>
                    <div style={styles.label}>&#x25A1;</div>
                    <div style={styles.label}>ODD</div>
                    <div style={styles.label}>19 to 36</div>
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
        backgroundColor: '#006400', // Dark green background
    },
    grid: {
        display: 'grid',
        gap: '2px',
        border: '2px solid white', // White border for the grid
        backgroundColor: '#006400', // Dark green background
    },
    row: {
        display: 'flex',
    },
    numberButton: {
        width: '50px',
        height: '50px',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        fontSize: '16px',
        fontWeight: 'bold',
        borderRadius: '50%',
        border: '2px solid white', // White border for each button
        transition: 'transform 0.2s, background-color 0.2s',
    },
    selectedButton: {
        transform: 'scale(1.1)',
        boxShadow: '0 0 10px #1890ff',
    },
    coinIcon: {
        marginLeft: '4px',
    },
    bottomRow: {
        display: 'flex',
        justifyContent: 'space-between',
        width: '100%',
        padding: '10px 0',
    },
    label: {
        color: 'white',
        fontSize: '16px',
        fontWeight: 'bold',
        textAlign: 'center',
        flex: 1,
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

// Sample numbers to render
const numbers = [
    [0, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33],
    [36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25],
    [24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13],
    [12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1]
];

export default function App() {
    const [selectedNumbers, setSelectedNumbers] = React.useState([]);

    const handleNumberClick = (number) => {
        setSelectedNumbers((prev) => (
            prev.includes(number) ? prev.filter(n => n !== number) : [...prev, number]
        ));
    };

    return (
        <NumberGrid 
            numbers={numbers} 
            selectedNumbers={selectedNumbers} 
            handleNumberClick={handleNumberClick} 
        />
    );
}
