import React from 'react';
import { Button } from 'antd';

function ColorButtons({ selectedColor, handleColorClick }) {
    const getButtonStyle = (color) => ({
        ...styles.colorButton,
        backgroundColor: color,
        color: 'white',
        opacity: selectedColor === color ? '0.7' : '1',
        transform: selectedColor === color ? 'scale(1.1)' : 'scale(1)',
    });

    return (
        <div style={styles.colorButtons}>
            <Button
                style={getButtonStyle('black')}
                onClick={() => handleColorClick('black')}
            >
                Черное
            </Button>
            <Button
                style={getButtonStyle('red')}
                onClick={() => handleColorClick('red')}
            >
                Красное
            </Button>
        </div>
    );
}

const styles = {
    colorButtons: {
        display: 'flex',
        justifyContent: 'center',
        marginTop: '20px',
    },
    colorButton: {
        width: '120px',
        height: '50px',
        marginRight: '10px',
        padding: '10px 20px',
        fontSize: '16px',
        fontWeight: 'bold',
        borderRadius: '8px',
        border: 'none',
        cursor: 'pointer',
        transition: 'opacity 0.2s, transform 0.2s',
    },
};

export default ColorButtons;
