import React, { useState } from 'react';
import { Card, Input, Button, Form, Row, Col, Typography } from 'antd';
import Display from './Display';

const { Title } = Typography;

function RouletteCard() {
    const [selectedNumbers, setSelectedNumbers] = useState([]);
    const [selectedColor, setSelectedColor] = useState(null);
    const [selectedBlack, setSelectedBlack] = useState(false);
    const [selectedRed, setSelectedRed] = useState(false);
    const [betAmount, setBetAmount] = useState('');
    const [evenBet, setEvenBet] = useState('');
    const [oddBet, setOddBet] = useState('');
    const [redBet, setRedBet] = useState('');
    const [blackBet, setBlackBet] = useState('');
    const [first12Bet, setFirst12Bet] = useState('');
    const [second12Bet, setSecond12Bet] = useState('');
    const [third12Bet, setThird12Bet] = useState('');
    const [oneToEighteenBet, setOneToEighteenBet] = useState('');
    const [nineteenToThirtySixBet, setNineteenToThirtySixBet] = useState('');
    const [first2To1Bet, setFirst2To1Bet] = useState('');
    const [second2To1Bet, setSecond2To1Bet] = useState('');
    const [third2To1Bet, setThird2To1Bet] = useState('');

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
        if (color === 'black') {
            setSelectedBlack(!selectedBlack);
            setSelectedRed(false);
        } else if (color === 'red') {
            setSelectedRed(!selectedRed);
            setSelectedBlack(false);
        }
    };

    const handleBetChange = (setter) => (e) => {
        setter(e.target.value);
    };

    const handleSubmit = async () => {
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
                },
                body: JSON.stringify({}),
            });

            if (!response.ok) {
                throw new Error('Network response was not ok');
            }

            const data = await response.json();
            console.log(data);
            setSpinResult(data.result); // предполагается, что в ответе есть поле "result" с результатом спина
            setIsSpinning(false);
        } catch (error) {
            console.error('Error:', error);
            setIsSpinning(false);
        }
    };

    const numbers = Array.from(Array(38).keys());

    return (
        <div style={styles.container}>
            <Card style={styles.card}>
                <Title level={2} style={styles.cardHeader}>Рулетка</Title>
                <div style={styles.displayContainer}>
                    <Display selectedNumbers={selectedNumbers} selectedColor={selectedColor} selectedBlack={selectedBlack} selectedRed={selectedRed} spinResult={spinResult} isSpinning={isSpinning} />
                </div>
                <div style={styles.scrollContainer}>
                    <Form layout="vertical" style={styles.form}>
                        <Form.Item label="Введите ставку на число">
                            <Input
                                placeholder="Ставка на число"
                                size="large"
                                prefix="₽"
                                value={betAmount}
                                onChange={handleBetChange(setBetAmount)}
                            />
                        </Form.Item>
                        <Row gutter={[16, 16]}>
                            <Col span={12}>
                                <Form.Item label="Ставка на четное">
                                    <Input
                                        placeholder="Ставка на четное"
                                        size="large"
                                        prefix="₽"
                                        value={evenBet}
                                        onChange={handleBetChange(setEvenBet)}
                                    />
                                </Form.Item>
                            </Col>
                            <Col span={12}>
                                <Form.Item label="Ставка на нечетное">
                                    <Input
                                        placeholder="Ставка на нечетное"
                                        size="large"
                                        prefix="₽"
                                        value={oddBet}
                                        onChange={handleBetChange(setOddBet)}
                                    />
                                </Form.Item>
                            </Col>
                        </Row>
                        <Row gutter={[16, 16]}>
                            <Col span={12}>
                                <Form.Item label="Ставка на красное">
                                    <Input
                                        placeholder="Ставка на красное"
                                        size="large"
                                        prefix="₽"
                                        value={redBet}
                                        onChange={handleBetChange(setRedBet)}
                                    />
                                </Form.Item>
                            </Col>
                            <Col span={12}>
                                <Form.Item label="Ставка на черное">
                                    <Input
                                        placeholder="Ставка на черное"
                                        size="large"
                                        prefix="₽"
                                        value={blackBet}
                                        onChange={handleBetChange(setBlackBet)}
                                    />
                                </Form.Item>
                            </Col>
                        </Row>
                        <Row gutter={[16, 16]}>
                            <Col span={8}>
                                <Form.Item label="Ставка на 1st 12">
                                    <Input
                                        placeholder="Ставка на 1st 12"
                                        size="large"
                                        prefix="₽"
                                        value={first12Bet}
                                        onChange={handleBetChange(setFirst12Bet)}
                                    />
                                </Form.Item>
                            </Col>
                            <Col span={8}>
                                <Form.Item label="Ставка на 2nd 12">
                                    <Input
                                        placeholder="Ставка на 2nd 12"
                                        size="large"
                                        prefix="₽"
                                        value={second12Bet}
                                        onChange={handleBetChange(setSecond12Bet)}
                                    />
                                </Form.Item>
                            </Col>
                            <Col span={8}>
                                <Form.Item label="Ставка на 3rd 12">
                                    <Input
                                        placeholder="Ставка на 3rd 12"
                                        size="large"
                                        prefix="₽"
                                        value={third12Bet}
                                        onChange={handleBetChange(setThird12Bet)}
                                    />
                                </Form.Item>
                            </Col>
                        </Row>
                        <Row gutter={[16, 16]}>
                            <Col span={12}>
                                <Form.Item label="Ставка на 1-18">
                                    <Input
                                        placeholder="Ставка на 1-18"
                                        size="large"
                                        prefix="₽"
                                        value={oneToEighteenBet}
                                        onChange={handleBetChange(setOneToEighteenBet)}
                                    />
                                </Form.Item>
                            </Col>
                            <Col span={12}>
                                <Form.Item label="Ставка на 19-36">
                                    <Input
                                        placeholder="Ставка на 19-36"
                                        size="large"
                                        prefix="₽"
                                        value={nineteenToThirtySixBet}
                                        onChange={handleBetChange(setNineteenToThirtySixBet)}
                                    />
                                </Form.Item>
                            </Col>
                        </Row>
                        <Row gutter={[16, 16]}>
                            <Col span={8}>
                                <Form.Item label="Ставка на First 2 to 1">
                                    <Input
                                        placeholder="Ставка на First 2 to 1"
                                        size="large"
                                        prefix="₽"
                                        value={first2To1Bet}
                                        onChange={handleBetChange(setFirst2To1Bet)}
                                    />
                                </Form.Item>
                            </Col>
                            <Col span={8}>
                                <Form.Item label="Ставка на Second 2 to 1">
                                    <Input
                                        placeholder="Ставка на Second 2 to 1"
                                        size="large"
                                        prefix="₽"
                                        value={second2To1Bet}
                                        onChange={handleBetChange(setSecond2To1Bet)}
                                    />
                                </Form.Item>
                            </Col>
                            <Col span={8}>
                                <Form.Item label="Ставка на Third 2 to 1">
                                    <Input
                                        placeholder="Ставка на Third 2 to 1"
                                        size="large"
                                        prefix="₽"
                                        value={third2To1Bet}
                                        onChange={handleBetChange(setThird2To1Bet)}
                                    />
                                </Form.Item>
                            </Col>
                        </Row>
                    </Form>
                </div>
                <div style={styles.submitButtonContainer}>
                    <Button type="primary" size="large" onClick={handleSubmit} style={styles.submitButton}>
                        Submit
                    </Button>
                </div>
                <div style={styles.cardBody}>
                    <div style={styles.numberGrid}>
                        {numbers.map((number) => (
                            <Button
                                key={number}
                                shape="circle"
                                size="large"
                                style={{
                                    ...styles.numberButton,
                                    backgroundColor: getColorForNumber(number),
                                    color: getTextColorForNumber(number),
                                }}
                                onClick={() => handleNumberClick(number)}
                            >
                                {number}
                                {selectedNumbers.includes(number) && (
                                    <span style={styles.coinIcon}>💰</span>
                                )}
                            </Button>
                        ))}
                    </div>
                </div>
                <div style={styles.colorButtons}>
                    <Button
                        style={{
                            ...styles.colorButton,
                            backgroundColor: 'black',
                            color: 'white',
                            opacity: selectedBlack ? '0.5' : '1',
                            width: '120px',
                            height: '50px',
                        }}
                        onClick={() => handleColorClick('black')}
                    >
                        Черное
                    </Button>
                    <Button
                        style={{
                            ...styles.colorButton,
                            backgroundColor: 'red',
                            color: 'white',
                            opacity: selectedRed ? '0.5' : '1',
                            width: '120px',
                            height: '50px',
                        }}
                        onClick={() => handleColorClick('red')}
                    >
                        Красное
                    </Button>
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
        width: '90%',
        maxWidth: '1200px',
        backgroundColor: 'white',
        padding: '24px',
        borderRadius: '8px',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
    },
    cardHeader: {
        textAlign: 'center',
        marginBottom: '24px',
    },
    form: {
        width: '100%',
    },
    inputContainer: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        marginBottom: '16px',
    },
    input: {
        width: '100%',
        maxWidth: '300px',
        marginBottom: '10px',
    },
    cardBody: {
        display: 'flex',
        justifyContent: 'center',
        marginTop: '20px',
    },
    numberGrid: {
        display: 'grid',
        gridTemplateColumns: 'repeat(10, 1fr)',
        gap: '10px',
        justifyContent: 'center',
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
        border: 'none',
        transition: 'transform 0.2s',
    },
    numberButtonSelected: {
        transform: 'scale(1.1)',
    },
    displayContainer: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '200px',
    },
    coinIcon: {
        marginLeft: '4px',
    },
    colorButtons: {
        display: 'flex',
        justifyContent: 'center',
        marginTop: '10px',
    },
    colorButton: {
        marginRight: '10px',
        padding: '10px 20px',
        fontSize: '16px',
        fontWeight: 'bold',
        borderRadius: '8px',
        border: 'none',
        transition: 'opacity 0.2s',
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
    scrollContainer: {
        maxHeight: '400px',
        overflowY: 'auto',
        paddingRight: '15px', // Compensate for scrollbar width
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

function getTextColorForNumber(number) {
    return 'white';
}

export default RouletteCard;
