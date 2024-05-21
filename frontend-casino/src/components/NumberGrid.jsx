import React from 'react';
import { Button } from 'antd';

function NumberGrid({ numbers, selectedNumbers, selectedColor, handleNumberClick, handleColorClick, handleSectorClick }) {
    const getButtonStyle = (number) => ({
        ...styles.numberButton,
        backgroundColor: getColorForNumber(number),
        color: 'white',
        opacity: selectedNumbers.includes(number) ? '0.7' : '1',
        transform: selectedNumbers.includes(number) ? 'scale(1.1)' : 'scale(1)',
    });

    const getColorButtonStyle = (color) => ({
        ...styles.diamondButton,
        backgroundColor: color,
        color: 'white',
        opacity: selectedColor === color ? '0.7' : '1',
        transform: selectedColor === color ? 'scale(1.1)' : 'scale(1)',
    });

    const getZeroButtonStyle = () => ({
        ...styles.diamondButton,
        backgroundColor: 'green',
        color: 'white',
        transform: 'rotate(45deg)',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    });

    return (
        <div style={styles.container}>
            <div style={styles.grid}>
                <div style={styles.zeroColumn}>
                    <Button
                        shape="circle"
                        size="large"
                        style={getZeroButtonStyle()}
                        onClick={() => handleNumberClick(0)}
                    >
                        <span style={styles.diamondText}>0</span>
                    </Button>
                </div>
                <div style={styles.numbers}>
                    {numbers.map((numberRow, rowIndex) => (
                        <div key={rowIndex} style={styles.row}>
                            {numberRow.map((number, numberIndex) => (
                                <Button
                                    key={numberIndex}
                                    style={getButtonStyle(number)}
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
                    <div style={styles.row}>
                        <Button style={styles.sectorButton} onClick={() => handleSectorClick('1st 12')}>1st 12</Button>
                        <Button style={styles.sectorButton} onClick={() => handleSectorClick('2nd 12')}>2nd 12</Button>
                        <Button style={styles.sectorButton} onClick={() => handleSectorClick('3rd 12')}>3rd 12</Button>
                    </div>
                    <div style={styles.row}>
                        <Button style={styles.sectorButton} onClick={() => handleSectorClick('1 to 18')}>1 to 18</Button>
                        <Button style={styles.sectorButton} onClick={() => handleSectorClick('EVEN')}>EVEN</Button>
                        <Button
                            style={getColorButtonStyle('red')}
                            onClick={() => handleColorClick('red')}
                        >
                            <span style={styles.diamondText}>Красное</span>
                        </Button>
                        <Button
                            style={getColorButtonStyle('black')}
                            onClick={() => handleColorClick('black')}
                        >
                            <span style={styles.diamondText}>Черное</span>
                        </Button>
                        <Button style={styles.sectorButton} onClick={() => handleSectorClick('ODD')}>ODD</Button>
                        <Button style={styles.sectorButton} onClick={() => handleSectorClick('19 to 36')}>19 to 36</Button>
                    </div>
                </div>
                <div style={styles.column}>
                    <Button style={styles.wideVerticalSectorButton} onClick={() => handleSectorClick('2 to 1 (1)')}>2 to 1</Button>
                    <Button style={styles.wideVerticalSectorButton} onClick={() => handleSectorClick('2 to 1 (2)')}>2 to 1</Button>
                    <Button style={styles.wideVerticalSectorButton} onClick={() => handleSectorClick('2 to 1 (3)')}>2 to 1</Button>
                </div>
            </div>
        </div>
    );
}

const styles = {
    container: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        padding: '20px',
        backgroundColor: '#006400', // Dark green background
    },
    grid: {
        display: 'flex',
        flexDirection: 'row', // Align columns in a row
        border: '2px solid white', // White border for the grid
        backgroundColor: '#006400', // Dark green background
        borderRadius: '10px', // Rounded corners for the grid
    },
    zeroColumn: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#006400', // Dark green background
        width: '60px', // Width of the zero block
        borderRight: '2px solid white', // White border for the zero block
        borderTopLeftRadius: '10px', // Rounded top left corner
        borderBottomLeftRadius: '10px', // Rounded bottom left corner
    },
    numbers: {
        display: 'flex',
        flexDirection: 'column',
        backgroundColor: '#006400', // Dark green background
    },
    row: {
        display: 'flex',
    },
    column: {
        display: 'flex',
        flexDirection: 'column',
    },
    numberButton: {
        width: '60px', // Increased button size
        height: '60px',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        fontSize: '18px', // Increased font size
        fontWeight: 'bold',
        borderRadius: '8px', // Rounded corners
        border: '2px solid white', // White border for each button
        transition: 'transform 0.2s, background-color 0.2s',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)', // Shadows for 3D effect
        margin: '2px', // Margin between buttons
    },
    selectedButton: {
        transform: 'scale(1.1)',
        boxShadow: '0 0 10px #1890ff',
    },
    coinIcon: {
        marginLeft: '4px',
    },
    sectorButton: {
        flex: 1,
        height: '60px',
        backgroundColor: '#006400', // Dark green background
        color: 'white',
        fontSize: '16px',
        fontWeight: 'bold',
        borderRadius: '8px',
        border: '2px solid white',
        cursor: 'pointer',
        textAlign: 'center',
        lineHeight: '60px',
        margin: '2px', // Margin between buttons
        transition: 'background-color 0.2s, transform 0.2s',
    },
    verticalSectorButton: {
        flex: 1,
        width: '60px',
        height: '60px',
        backgroundColor: '#006400', // Dark green background
        color: 'white',
        fontSize: '16px',
        fontWeight: 'bold',
        borderRadius: '8px',
        border: '2px solid white',
        cursor: 'pointer',
        textAlign: 'center',
        margin: '2px', // Margin between buttons
        transition: 'background-color 0.2s, transform 0.2s',
    },
    wideVerticalSectorButton: {
        flex: 1,
        width: '80px', // Increased width for wider appearance
        height: '60px',
        backgroundColor: '#006400', // Dark green background
        color: 'white',
        fontSize: '16px',
        fontWeight: 'bold',
        borderRadius: '8px',
        border: '2px solid white',
        cursor: 'pointer',
        textAlign: 'center',
        margin: '2px', // Margin between buttons
        transition: 'background-color 0.2s, transform 0.2s',
    },
    colorButton: {
        width: '80px', // Reduced button width
        height: '40px', // Reduced button height
        fontSize: '14px',
        fontWeight: 'bold',
        borderRadius: '8px',
        border: '2px solid white', // White border around the button
        cursor: 'pointer',
        transition: 'opacity 0.2s, transform 0.2s',
        textAlign: 'center',
        lineHeight: '1', // To center the text vertically
    },
    diamondButton: {
        width: '60px',
        height: '60px',
        backgroundColor: 'transparent',
        border: '2px solid white',
        transform: 'rotate(45deg)',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
    diamondText: {
        transform: 'rotate(-45deg)',
        textAlign: 'center',
    },
};

function getColorForNumber(number) {
    if (number === 0) {
        return 'green';
    } else if ([1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36].includes(number)) {
        return 'red';
    } else {
        return 'black';
    }
}

export default NumberGrid;
