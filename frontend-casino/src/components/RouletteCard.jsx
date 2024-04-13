import Display from "./Display"
import { Card, Input } from 'antd';

const gridStyle = {
    width: '10%',
    textAlign: 'center',
};

function RouletteCard() {
    // Создаем массив чисел от 0 до 37
    const numbers = Array.from(Array(38).keys());
    
    return (
        <div style={styles.container}>
            <Card style={styles.card}>
                <div style={styles.cardHeader}>
                    Рулетка
                </div>
                <div style={styles.displayContainer}>
                    <Display />
                </div>
                <div style={styles.inputContainer}>
                    <Input
                        style={styles.input}
                        placeholder="Введите ставку"
                        size="large"
                        prefix="₽"
                    />
                </div>
                <div style={styles.cardBody}>
                    {numbers.map((number) => (
                        <Card.Grid
                            key={number}
                            style={{
                                ...gridStyle,
                                backgroundColor: getColorForNumber(number),
                                color: getTextColorForNumber(number),
                            }}
                        >
                            {number}
                        </Card.Grid>
                    ))}
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
        height: '200px', // Example height, adjust as needed
    },
};

// Функция для определения цвета фона в зависимости от числа
function getColorForNumber(number) {
    if (number === 0) {
        return 'green'; // Зеленый цвет для числа 0
    } else if (number % 2 === 0) {
        return 'black'; // Черный цвет для четных чисел
    } else {
        return 'red'; // Красный цвет для нечетных чисел
    }
}

// Функция для определения цвета текста в зависимости от фона
function getTextColorForNumber(number) {
    if (number === 0) {
        return 'white'; // Белый цвет текста для числа 0
    } else {
        return 'white'; // Белый цвет текста для всех остальных чисел (на черном и красном фоне)
    }
}

export default RouletteCard;
