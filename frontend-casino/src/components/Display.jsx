import { Card } from 'antd';

function Display() {
    return (
        <div style={styles.displayContainer}>
            <Card style={styles.card}>
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
        width: 500,
        height: 150,
        backgroundColor: 'white', // Example background color
        borderRadius: '8px',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)', // Example shadow
    },
};

export default Display;