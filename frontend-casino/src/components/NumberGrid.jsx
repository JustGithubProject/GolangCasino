import React, { useState } from 'react';
import { Button } from 'antd';

function NumberGrid({ numbers, selectedNumbers, selectedColor, handleNumberClick, handleColorClick, handleSectorClick }) {
    const [selectedCoin, setSelectedCoin] = useState(1);
    const [bets, setBets] = useState({});

    const handleNumberBet = (number) => {
        handleNumberClick(number, selectedCoin);
        setBets((prevBets) => ({
            ...prevBets,
            [number]: (prevBets[number] || 0) + selectedCoin,
        }));
    };

    const handleSectorButtonClick = (sector) => {
        handleSectorClick(sector, selectedCoin);
    };

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
        ...styles.zeroButton,
        backgroundColor: 'green',
        color: 'white',
    });

    const getSectorButtonStyle = (sector) => ({
        ...styles.sectorButton,
    });

    const handleCoinClick = (coin) => {
        setSelectedCoin(coin);
    };

    return (
        <div style={styles.container}>
            <div style={styles.leftPanel}>
                {[1, 5, 10, 50].map((coin) => (
                    <Button
                        key={coin}
                        style={styles.coinButton}
                        onClick={() => handleCoinClick(coin)}
                    >
                        {coin}
                    </Button>
                ))}
            </div>
            <div style={styles.grid}>
                <div style={styles.zeroColumn}>
                    <Button
                        shape="circle"
                        size="large"
                        style={getZeroButtonStyle()}
                        onClick={() => handleNumberBet(0)}
                    >
                        0
                    </Button>
                </div>
                <div style={styles.numbers}>
                    {numbers.map((numberRow, rowIndex) => (
                        <div key={rowIndex} style={styles.row}>
                            {numberRow.map((number, numberIndex) => (
                                <Button
                                    key={numberIndex}
                                    style={getButtonStyle(number)}
                                    onClick={() => handleNumberBet(number)}
                                >
                                    {number}
                                    {bets[number] && (
                                        <span style={styles.betAmount}>{bets[number]}</span>
                                    )}
                                </Button>
                            ))}
                        </div>
                    ))}
                    <div style={styles.row}>
                        <Button style={getSectorButtonStyle('1st 12')} onClick={() => handleSectorButtonClick('1st 12')}>1st 12</Button>
                        <Button style={getSectorButtonStyle('2nd 12')} onClick={() => handleSectorButtonClick('2nd 12')}>2nd 12</Button>
                        <Button style={getSectorButtonStyle('3rd 12')} onClick={() => handleSectorButtonClick('3rd 12')}>3rd 12</Button>
                    </div>
                    <div style={styles.row}>
                        <Button style={getSectorButtonStyle('1 to 18')} onClick={() => handleSectorButtonClick('1 to 18')}>1 to 18</Button>
                        <Button style={getSectorButtonStyle('EVEN')} onClick={() => handleSectorButtonClick('EVEN')}>EVEN</Button>
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
                        <Button style={getSectorButtonStyle('ODD')} onClick={() => handleSectorButtonClick('ODD')}>ODD</Button>
                        <Button style={getSectorButtonStyle('19 to 36')} onClick={() => handleSectorButtonClick('19 to 36')}>19 to 36</Button>
                    </div>
                </div>
                <div style={styles.column}>
                    <Button style={getSectorButtonStyle('2 to 1 (1)')} onClick={() => handleSectorButtonClick('2 to 1 (1)')}>2 to 1</Button>
                    <Button style={getSectorButtonStyle('2 to 1 (2)')} onClick={() => handleSectorButtonClick('2 to 1 (2)')}>2 to 1</Button>
                    <Button style={getSectorButtonStyle('2 to 1 (3)')} onClick={() => handleSectorButtonClick('2 to 1 (3)')}>2 to 1</Button>
                </div>
            </div>
        </div>
    );
}

const styles = {
    container: {
        display: 'flex',
        flexDirection: 'row',
        justifyContent: 'center',
        alignItems: 'center',
        padding: '20px',
        backgroundColor: '#006400',
    },
    leftPanel: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        padding: '10px',
        backgroundColor: '#006400',
    },
    grid: {
        display: 'flex',
        flexDirection: 'row',
        border: '2px solid white',
        backgroundColor: '#006400',
        borderRadius: '10px',
    },
    zeroColumn: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        backgroundColor: '#006400',
        width: '60px',
        borderRight: '2px solid white',
        borderTopLeftRadius: '10px',
        borderBottomLeftRadius: '10px',
    },
    zeroButton: {
        width: '100%',
        height: '100%',
        fontSize: '24px',
        fontWeight: 'bold',
        borderRadius: '0px',
        border: '2px solid white',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
    numbers: {
        display: 'flex',
        flexDirection: 'column',
        backgroundColor: '#006400',
    },
    row: {
        display: 'flex',
    },
    column: {
        display: 'flex',
        flexDirection: 'column',
    },
    numberButton: {
        width: '60px',
        height: '60px',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        fontSize: '18px',
        fontWeight: 'bold',
        borderRadius: '8px',
        border: '2px solid white',
        transition: 'transform 0.2s, background-color 0.2s',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.2)',
        margin: '2px',
    },
    betAmount: {
        marginLeft: '4px',
        fontSize: '14px',
        fontWeight: 'bold',
        color: 'yellow',
    },
    sectorButton: {
        flex: 1,
        height: '60px',
        backgroundColor: '#006400',
        color: 'white',
        fontSize: '16px',
        fontWeight: 'bold',
        borderRadius: '8px',
        border: '2px solid white',
        cursor: 'pointer',
        textAlign: 'center',
        lineHeight: '60px',
        margin: '2px',
        transition: 'background-color 0.2s, transform 0.2s',
    },
    colorButton: {
        width: '80px',
        height: '40px',
        fontSize: '14px',
        fontWeight: 'bold',
        borderRadius: '8px',
        border: '2px solid white',
        cursor: 'pointer',
        transition: 'opacity 0.2s, transform 0.2s',
        textAlign: 'center',
        lineHeight: '1',
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
    coinButton: {
        width: '50px',
        height: '50px',
        margin: '5px',
        fontSize: '18px',
        fontWeight: 'bold',
        borderRadius: '50%',
        backgroundColor: '#f0ad4e',
        color: 'white',
        border: '2px solid white',
        transition: 'transform 0.2s',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
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
