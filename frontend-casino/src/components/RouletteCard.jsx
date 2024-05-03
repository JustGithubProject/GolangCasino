import React, { useState } from 'react';
import { Card, Input, Button } from 'antd';
import Display from './Display';
import SimpleButton from './SimpleButton';

function RouletteCard() {
    const [selectedNumbers, setSelectedNumbers] = useState([]);

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

    const numbers = Array.from(Array(38).keys());

    return (
        <div style={styles.container}>
            <Card style={styles.card}>
                <div style={styles.cardHeader}>Рулетка</div>
                <div style={styles.displayContainer}>
                    <Display selectedNumbers={selectedNumbers} />
                </div>
                <div style={styles.inputContainer}>
                    <Input
                        style={styles.input}
                        placeholder="Введите ставку"
                        size="large"
                        prefix="₽"
                    />
                    <SimpleButton />
                </div>
                <div style={styles.cardBody}>
                    {numbers.map((number) => (
                        <Button
                            key={number}
                            style={{
                                ...gridStyle,
                                backgroundColor: getColorForNumber(number),
                                color: getTextColorForNumber(number),
                                marginBottom: '10px',
                            }}
                            onClick={() => handleNumberClick(number)}
                        >
                            {number}
                            {selectedNumbers.includes(number) && ( // Проверяем, выбрано ли число
                                <span style={styles.coinIcon}>💰</span>
                            )}
                        </Button>
                    ))}
                </div>
            </Card>
        </div>
    );
}

const gridStyle = {
    width: '15%',
    textAlign: 'center',
};

const styles = {
    container: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        minHeight: '100vh',
    },
    card: {
        width: '80%',
        backgroundColor: 'white',
        padding: '24px',
        borderRadius: '8px',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
    },
    cardHeader: {
        fontSize: '24px',
        fontWeight: 'bold',
        textAlign: 'center',
        marginBottom: '16px',
    },
    inputContainer: {
        display: 'flex',
        justifyContent: 'center',
        marginBottom: '16px',
    },
    input: {
        width: '100%',
        maxWidth: '300px',
    },
    cardBody: {
        display: 'flex',
        flexWrap: 'wrap',
        justifyContent: 'center',
    },
    displayContainer: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '200px',
    },
    coinIcon: {
        marginLeft: '4px', // Отступ для разделения числа и монетки
    },
};

function getColorForNumber(number) {
    if (number === 0) {
        return 'green';
    } else if (
        [1, 3, 5, 7, 9, 12, 14, 16, 18, 19, 21, 23, 25, 27, 30, 32, 34, 36].includes(number)
    ) {
        return 'red';
    } else {
        return 'black';
    }
}

function getTextColorForNumber(number) {
    if (number === 0) {
        return 'white';
    } else {
        return 'white';
    }
}

export default RouletteCard;
