import React, { useState } from 'react';
import { Card, Button, Typography, Spin, message } from 'antd';
import Display from './Display';
import BetForm from './BetForm';
import NumberGrid from './NumberGrid';

const { Title } = Typography;

function RouletteCard() {
    const [selectedNumbers, setSelectedNumbers] = useState([]);
    const [selectedColor, setSelectedColor] = useState(null);
    const [betValues, setBetValues] = useState({
        betAmount: '',
        evenBet: '',
        oddBet: '',
        redBet: '',
        blackBet: '',
        first12Bet: '',
        second12Bet: '',
        third12Bet: '',
        oneToEighteenBet: '',
        nineteenToThirtySixBet: '',
        first2To1Bet: '',
        second2To1Bet: '',
        third2To1Bet: '',
    });

    const [spinResult, setSpinResult] = useState(null);
    const [isSpinning, setIsSpinning] = useState(false);

    const handleNumberClick = (number) => {
        const index = selectedNumbers.indexOf(number);
        if (index === -1) {
            setSelectedNumbers([...selectedNumbers, number]);
        } else {
            const updatedNumbers = [...selectedNumbers];
            updatedNumbers.splice(index, 1);
            setSelectedNumbers(updatedNumbers);
        }
    };

    const handleColorClick = (color) => {
        setSelectedColor(color);
        setBetValues((prevValues) => ({
            ...prevValues,
            redBet: color === 'red' ? (prevValues.redBet ? '' : '100') : '',
            blackBet: color === 'black' ? (prevValues.blackBet ? '' : '100') : '',
        }));
    };

    const handleBetChange = (key) => (e) => {
        const value = e.target.value;
        setBetValues((prevValues) => ({
            ...prevValues,
            [key]: value,
        }));
    };

    const handleSectorClick = (sector) => {
        const fieldMap = {
            '1st 12': 'first12Bet',
            '2nd 12': 'second12Bet',
            '3rd 12': 'third12Bet',
            '1 to 18': 'oneToEighteenBet',
            'EVEN': 'evenBet',
            'ODD': 'oddBet',
            '19 to 36': 'nineteenToThirtySixBet',
            '2 to 1 (1)': 'first2To1Bet',
            '2 to 1 (2)': 'second2To1Bet',
            '2 to 1 (3)': 'third2To1Bet',
        };

        const field = fieldMap[sector];
        if (field) {
            setBetValues((prevValues) => ({
                ...prevValues,
                [field]: prevValues[field] === '100' ? '' : '100',
            }));
        }
    };

    const handleSubmit = async () => {
        const {
            betAmount, evenBet, oddBet, redBet, blackBet,
            first12Bet, second12Bet, third12Bet, oneToEighteenBet,
            nineteenToThirtySixBet, first2To1Bet, second2To1Bet, third2To1Bet,
        } = betValues;

        if (!betAmount && !evenBet && !oddBet && !redBet && !blackBet && !first12Bet &&
            !second12Bet && !third12Bet && !oneToEighteenBet && !nineteenToThirtySixBet &&
            !first2To1Bet && !second2To1Bet && !third2To1Bet) {
            message.error('Please place at least one bet.');
            return;
        }

        const params = new URLSearchParams();

        if (evenBet) params.append('even', evenBet);
        if (oddBet) params.append('odd', oddBet);
        if (redBet) params.append('red', redBet);
        if (blackBet) params.append('black', blackBet);
        if (first12Bet) params.append('1st12', first12Bet);
        if (second12Bet) params.append('2nd12', second12Bet);
        if (third12Bet) params.append('3rd12', third12Bet);
        if (betAmount) params.append('number', betAmount);
        if (selectedNumbers.length > 0) params.append('num', selectedNumbers.join(','));
        if (oneToEighteenBet) params.append('1To18', oneToEighteenBet);
        if (nineteenToThirtySixBet) params.append('19To36', nineteenToThirtySixBet);
        if (first2To1Bet) params.append('First2To1', first2To1Bet);
        if (second2To1Bet) params.append('Second2To1', second2To1Bet);
        if (third2To1Bet) params.append('Third2To1', third2To1Bet);

        const url = `http://localhost:8080/spin-roulette-v1/?${params.toString()}`;

        console.log('URL:', url);

        try {
            setIsSpinning(true);
            const response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTYwMTY2OTQsInVzZXJuYW1lIjoiS3JvcGl2YSJ9.6bzKBcZSUs938nN-sJmRSeC4Q9w29vGjOnACnDltezg',
                },
                body: JSON.stringify({}),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const data = await response.json();
            console.log(data);
            setSpinResult(data.dropped_number);
            setIsSpinning(false);
        } catch (error) {
            console.error('Error:', error);
            setIsSpinning(false);
        }
    };

    const numbers = [
        [3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36],
        [2, 5, 8, 11, 14, 17, 20, 23, 26, 29, 32, 35],
        [1, 4, 7, 10, 13, 16, 19, 22, 25, 28, 31, 34],
    ];

    return (
        <div style={styles.container}>
            <Card style={styles.card}>
                <Title level={2} style={styles.cardHeader}>Рулетка</Title>
                <div style={styles.displayContainer}>
                    <Display selectedNumbers={selectedNumbers} selectedColor={selectedColor} spinResult={spinResult} isSpinning={isSpinning} />
                </div>
                <div style={styles.scrollContainer}>
                    <BetForm betValues={betValues} handleBetChange={handleBetChange} />
                </div>
                <div style={styles.submitButtonContainer}>
                    <Button type="primary" size="large" onClick={handleSubmit} style={styles.submitButton} disabled={isSpinning}>
                        {isSpinning ? <Spin /> : 'Вращать'}
                    </Button>
                </div>
                <div style={styles.cardBody}>
                    <NumberGrid
                        numbers={numbers}
                        selectedNumbers={selectedNumbers}
                        handleNumberClick={handleNumberClick}
                        handleColorClick={handleColorClick}
                        handleSectorClick={handleSectorClick}
                    />
                </div>
            </Card>
        </div>
    );
}

const styles = {
    container: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        minHeight: '100vh',
        backgroundColor: '#f0f2f5',
        padding: '20px',
    },
    card: {
        width: '100%',
        maxWidth: '1200px',
        backgroundColor: 'white',
        padding: '24px',
        borderRadius: '8px',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
    },
    cardHeader: {
        textAlign: 'center',
        marginBottom: '24px',
        fontSize: '24px',
        fontWeight: 'bold',
        color: '#1890ff',
        textShadow: '1px 1px 2px rgba(0,0,0,0.1)',
    },
    displayContainer: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '200px',
    },
    scrollContainer: {
        maxHeight: '400px',
        overflowY: 'auto',
        paddingRight: '15px', // Compensate for scrollbar width
    },
    submitButtonContainer: {
        display: 'flex',
        justifyContent: 'center',
        position: 'sticky',
        bottom: '20px',
        backgroundColor: 'white',
        padding: '10px 0',
        boxShadow: '0 -2px 8px rgba(0, 0, 0, 0.1)',
        zIndex: 1,
    },
    submitButton: {
        padding: '10px 30px',
        fontSize: '16px',
        fontWeight: 'bold',
        borderRadius: '8px',
    },
    cardBody: {
        paddingTop: '24px',
    }
};

export default RouletteCard;
