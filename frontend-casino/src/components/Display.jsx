import { Card, Typography, Space } from 'antd';
import { CheckCircleTwoTone, CloseCircleTwoTone } from '@ant-design/icons';

const { Text } = Typography;

function Display({ selectedNumbers, selectedColor }) {
    return (
        <div style={styles.displayContainer}>
            <Card style={styles.card}>
                <Space direction="vertical" align="center">
                    {selectedNumbers.length > 0 ? (
                        <Text strong style={styles.numberText}>
                            <CheckCircleTwoTone twoToneColor="#52c41a" /> Выбранные числа: {selectedNumbers.join(', ')}
                        </Text>
                    ) : (
                        <Text strong style={styles.numberText}>
                            <CloseCircleTwoTone twoToneColor="#eb2f96" /> Выберите числа от 0 до 37
                        </Text>
                    )}
                    {selectedColor && (
                        <Text style={styles.colorText}>
                            <span style={{ color: selectedColor === 'red' ? '#ff4d4f' : '#595959' }}>
                                {selectedColor === 'red' ? '🔴 Красный' : '⚫ Черный'}
                            </span>
                        </Text>
                    )}
                </Space>
            </Card>
        </div>
    );
}

const styles = {
    displayContainer: {
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        height: '100%',
        padding: '15px',
    },
    card: {
        width: '100%',
        maxWidth: '700px',
        backgroundColor: 'white',
        borderRadius: '8px',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        padding: '20px',
    },
    numberText: {
        fontSize: '24px',
        fontWeight: 'bold',
        textAlign: 'center',
    },
    colorText: {
        fontSize: '20px',
        marginTop: '10px',
        textAlign: 'center',
    },
};

export default Display;
