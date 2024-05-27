import React, { useState } from 'react';
import { Card, Button, Typography, Spin, message, Space } from 'antd';
import BetForm from './BetForm';
import NumberGrid from './NumberGrid';
import { CSSTransition } from 'react-transition-group';
import './styles.css';

const { Title, Text } = Typography;

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
    const [selectedCoin, setSelectedCoin] = useState(1);

    const [spinResult, setSpinResult] = useState(null);
    const [isSpinning, setIsSpinning] = useState(false);
    const [resultMessage, setResultMessage] = useState(null);
    const [showResult, setShowResult] = useState(false);

    const handleNumberClick = (number, amount) => {
        setBetValues((prevValues) => ({
            ...prevValues,
            betAmount: (parseInt(prevValues.betAmount) || 0) + amount,
        }));
        const index = selectedNumbers.indexOf(number);
        if (index === -1) {
            setSelectedNumbers([...selectedNumbers, number]);
        }
    };

    const handleColorClick = (color) => {
        setSelectedColor(color);
        setBetValues((prevValues) => ({
            ...prevValues,
            redBet: color === 'red' ? (parseInt(prevValues.redBet) || 0) + selectedCoin : prevValues.redBet,
            blackBet: color === 'black' ? (parseInt(prevValues.blackBet) || 0) + selectedCoin : prevValues.blackBet,
        }));
    };

    const handleBetChange = (key) => (e) => {
        const value = e.target.value;
        setBetValues((prevValues) => ({
            ...prevValues,
            [key]: value,
        }));
    };

    const handleSectorClick = (sector, amount) => {
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
                [field]: (parseInt(prevValues[field]) || 0) + amount,
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

        const url = `http://localhost:8081/spin-roulette-v1/?${params.toString()}`;

        console.log('URL:', url);

        try {
            setIsSpinning(true);
            const response = await fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY3OTg5NjYsInVzZXJuYW1lIjoiS3JvcGl2YSJ9.Qg0GIkgQAZD_AwMzeSGYQP7A_RLnD6Tk90qRj7uJcrM',
                },
                body: JSON.stringify({}),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const data = await response.json();
            console.log(data);
            setSpinResult(data.dropped_number);
            calculateWinnings(data.dropped_number);
            setIsSpinning(false);
            setShowResult(true);
            setTimeout(() => {
                setShowResult(false);
                handleReset();  // Очистка формы после игры
            }, 5000); // Show the result for 5 seconds
        } catch (error) {
            console.error('Error:', error);
            setIsSpinning(false);
        }
    };

    const calculateWinnings = (droppedNumber) => {
        let winnings = 0;
        if (selectedNumbers.includes(droppedNumber)) {
            winnings += (parseInt(betValues.betAmount) || 0) * 36;
        }
        if ((droppedNumber % 2 === 0 && betValues.evenBet) || (droppedNumber % 2 !== 0 && betValues.oddBet)) {
            winnings += (parseInt(betValues.evenBet) || 0) * 2;
            winnings += (parseInt(betValues.oddBet) || 0) * 2;
        }
        if ((droppedNumber <= 18 && betValues.oneToEighteenBet) || (droppedNumber >= 19 && betValues.nineteenToThirtySixBet)) {
            winnings += (parseInt(betValues.oneToEighteenBet) || 0) * 2;
            winnings += (parseInt(betValues.nineteenToThirtySixBet) || 0) * 2;
        }
        if ((droppedNumber >= 1 && droppedNumber <= 12 && betValues.first12Bet) ||
            (droppedNumber >= 13 && droppedNumber <= 24 && betValues.second12Bet) ||
            (droppedNumber >= 25 && droppedNumber <= 36 && betValues.third12Bet)) {
            winnings += (parseInt(betValues.first12Bet) || 0) * 3;
            winnings += (parseInt(betValues.second12Bet) || 0) * 3;
            winnings += (parseInt(betValues.third12Bet) || 0) * 3;
        }
        if ((droppedNumber % 3 === 1 && betValues.first2To1Bet) ||
            (droppedNumber % 3 === 2 && betValues.second2To1Bet) ||
            (droppedNumber % 3 === 0 && betValues.third2To1Bet)) {
            winnings += (parseInt(betValues.first2To1Bet) || 0) * 3;
            winnings += (parseInt(betValues.second2To1Bet) || 0) * 3;
            winnings += (parseInt(betValues.third2To1Bet) || 0) * 3;
        }
        if (selectedColor && (selectedColor === 'red' || selectedColor === 'black')) {
            const redNumbers = [1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36];
            const isRed = redNumbers.includes(droppedNumber);
            if ((selectedColor === 'red' && isRed) || (selectedColor === 'black' && !isRed)) {
                winnings += (parseInt(betValues.redBet) || 0) * 2;
                winnings += (parseInt(betValues.blackBet) || 0) * 2;
            }
        }

        setResultMessage(winnings > 0 ? `Вы выиграли: ₽${winnings}` : 'Вы не выиграли');
    };

    const handleReset = () => {
        setSelectedNumbers([]);
        setSelectedColor(null);
        setBetValues({
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
        setSpinResult(null);
        setResultMessage(null);
    };

    const toggleFullscreen = () => {
        if (!document.fullscreenElement) {
            document.documentElement.requestFullscreen().catch((err) => {
                alert(`Error attempting to enable full-screen mode: ${err.message} (${err.name})`);
            });
        } else {
            document.exitFullscreen();
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
                {resultMessage && (
                    <CSSTransition
                        in={showResult}
                        timeout={300}
                        classNames="fade"
                        unmountOnExit
                    >
                        <div style={styles.resultOverlay}>
                            <Text style={styles.resultText}>{resultMessage}<p></p></Text>
                            {spinResult !== null && <Text style={styles.spinResultText}>Выпавшее число: {spinResult}</Text>}
                        </div>
                    </CSSTransition>
                )}
                <div style={styles.scrollContainer}>
                    <NumberGrid
                        numbers={numbers}
                        selectedNumbers={selectedNumbers}
                        selectedColor={selectedColor}
                        handleNumberClick={handleNumberClick}
                        handleColorClick={handleColorClick}
                        handleSectorClick={handleSectorClick}
                        reset={handleReset}
                        setSelectedCoin={setSelectedCoin}
                    />
                </div>
                <div style={styles.submitButtonContainer}>
                    <Button type="primary" size="large" onClick={handleSubmit} style={styles.submitButton} disabled={isSpinning}>
                        {isSpinning ? <Spin /> : 'Вращать'}
                    </Button>
                    <Button type="default" size="large" onClick={handleReset} style={styles.resetButton} disabled={isSpinning}>
                        Сбросить ставки
                    </Button>
                    <Button type="default" size="large" onClick={toggleFullscreen} style={styles.fullscreenButton}>
                        Полноэкранный режим
                    </Button>
                </div>
                <div style={styles.scrollContainer}>
                    <BetForm betValues={betValues} handleBetChange={handleBetChange} reset={handleReset} />
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
        background: 'linear-gradient(to right, #43cea2, #185a9d)',
        padding: '20px',
    },
    card: {
        width: '100%',
        maxWidth: '1200px',
        backgroundColor: '#f8f9fa',
        padding: '24px',
        borderRadius: '16px',
        boxShadow: '0 4px 12px rgba(0, 0, 0, 0.2)',
    },
    cardHeader: {
        textAlign: 'center',
        marginBottom: '24px',
        color: '#fff',
        fontSize: '32px',
        fontWeight: 'bold',
        background: 'linear-gradient(to right, #ff416c, #ff4b2b)',
        padding: '16px',
        borderRadius: '8px',
    },
    resultOverlay: {
        position: 'fixed',
        top: '50%',
        left: '50%',
        transform: 'translate(-50%, -50%)',
        backgroundColor: 'rgba(0, 0, 0, 0.8)',
        padding: '20px',
        borderRadius: '16px',
        boxShadow: '0 4px 12px rgba(0, 0, 0, 0.5)',
        textAlign: 'center',
        zIndex: 1000,
    },
    resultText: {
        fontSize: '24px',
        fontWeight: 'bold',
        color: '#fff',
        marginBottom: '10px',
    },
    spinResultText: {
        fontSize: '20px',
        fontWeight: 'bold',
        color: '#fff',
    },
    scrollContainer: {
        maxHeight: '400px',
        overflowY: 'auto',
        paddingRight: '15px',
    },
    submitButtonContainer: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        marginTop: '24px',
    },
    submitButton: {
        padding: '10px 30px',
        fontSize: '16px',
        fontWeight: 'bold',
        borderRadius: '8px',
        marginLeft: '10px',
        background: 'linear-gradient(to right, #ff416c, #ff4b2b)',
        border: 'none',
        color: 'white',
    },
    resetButton: {
        padding: '10px 10px',
        fontSize: '16px',
        fontWeight: 'bold',
        borderRadius: '8px',
        marginLeft: '10px',
        background: 'linear-gradient(to right, #6a11cb, #2575fc)',
        border: 'none',
        color: 'white',
    },
    fullscreenButton: {
        padding: '10px 30px',
        fontSize: '16px',
        fontWeight: 'bold',
        borderRadius: '8px',
        marginLeft: '10px',
        background: 'linear-gradient(to right, #43cea2, #185a9d)',
        border: 'none',
        color: 'white',
        
    },
};

export default RouletteCard;
