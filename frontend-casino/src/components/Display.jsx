import { Card } from 'antd';


function Display({ selectedNumbers }) {
    return (
        <div style={styles.displayContainer}>
            <Card style={styles.card}>
                {selectedNumbers.length > 0 ? (
                    <p style={styles.numberText}>Выбранные числа: {selectedNumbers.join(', ')}</p>
                ) : (
                    <p style={styles.numberText}>Выберите числа от 0 до 37</p>
                )}
            </Card>
        </div>
    );
}

const styles = {
    displayContainer: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100%', // Use 100% height to fill the parent container
        padding: '15px', // Add padding for spacing
    },
    card: {
        width: 700,
        height: 150,
        backgroundColor: 'white', // Example background color
        borderRadius: '8px',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)', // Example shadow
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
    numberText: {
        fontSize: '24px',
        fontWeight: 'bold',
    },
};

export default Display;