import React from 'react';
import { Form, Input, Row, Col, Card, Button } from 'antd';

function BetForm({ betValues, handleBetChange, reset }) {
    const {
        betAmount, evenBet, oddBet, redBet, blackBet, first12Bet, second12Bet, third12Bet,
        oneToEighteenBet, nineteenToThirtySixBet, first2To1Bet, second2To1Bet, third2To1Bet
    } = betValues;

    return (
        <Card style={styles.card}>
            <Form layout="vertical" style={styles.form}>
                <Form.Item label="Введите ставку на число" style={styles.formItem}>
                    <Input
                        placeholder="Ставка на число"
                        size="large"
                        prefix="₽"
                        value={betAmount}
                        onChange={handleBetChange('betAmount')}
                        style={styles.input}
                    />
                </Form.Item>
                <Row gutter={[16, 16]}>
                    <Col span={12}>
                        <Form.Item label="Ставка на четное" style={styles.formItem}>
                            <Input
                                placeholder="Ставка на четное"
                                size="large"
                                prefix="₽"
                                value={evenBet}
                                onChange={handleBetChange('evenBet')}
                                style={styles.input}
                            />
                        </Form.Item>
                    </Col>
                    <Col span={12}>
                        <Form.Item label="Ставка на нечетное" style={styles.formItem}>
                            <Input
                                placeholder="Ставка на нечетное"
                                size="large"
                                prefix="₽"
                                value={oddBet}
                                onChange={handleBetChange('oddBet')}
                                style={styles.input}
                            />
                        </Form.Item>
                    </Col>
                </Row>
                <Row gutter={[16, 16]}>
                    <Col span={12}>
                        <Form.Item label="Ставка на красное" style={styles.formItem}>
                            <Input
                                placeholder="Ставка на красное"
                                size="large"
                                prefix="₽"
                                value={redBet}
                                onChange={handleBetChange('redBet')}
                                style={styles.input}
                            />
                        </Form.Item>
                    </Col>
                    <Col span={12}>
                        <Form.Item label="Ставка на черное" style={styles.formItem}>
                            <Input
                                placeholder="Ставка на черное"
                                size="large"
                                prefix="₽"
                                value={blackBet}
                                onChange={handleBetChange('blackBet')}
                                style={styles.input}
                            />
                        </Form.Item>
                    </Col>
                </Row>
                <Row gutter={[16, 16]}>
                    <Col span={8}>
                        <Form.Item label="Ставка на 1st 12" style={styles.formItem}>
                            <Input
                                placeholder="Ставка на 1st 12"
                                size="large"
                                prefix="₽"
                                value={first12Bet}
                                onChange={handleBetChange('first12Bet')}
                                style={styles.input}
                            />
                        </Form.Item>
                    </Col>
                    <Col span={8}>
                        <Form.Item label="Ставка на 2nd 12" style={styles.formItem}>
                            <Input
                                placeholder="Ставка на 2nd 12"
                                size="large"
                                prefix="₽"
                                value={second12Bet}
                                onChange={handleBetChange('second12Bet')}
                                style={styles.input}
                            />
                        </Form.Item>
                    </Col>
                    <Col span={8}>
                        <Form.Item label="Ставка на 3rd 12" style={styles.formItem}>
                            <Input
                                placeholder="Ставка на 3rd 12"
                                size="large"
                                prefix="₽"
                                value={third12Bet}
                                onChange={handleBetChange('third12Bet')}
                                style={styles.input}
                            />
                        </Form.Item>
                    </Col>
                </Row>
                <Row gutter={[16, 16]}>
                    <Col span={12}>
                        <Form.Item label="Ставка на 1-18" style={styles.formItem}>
                            <Input
                                placeholder="Ставка на 1-18"
                                size="large"
                                prefix="₽"
                                value={oneToEighteenBet}
                                onChange={handleBetChange('oneToEighteenBet')}
                                style={styles.input}
                            />
                        </Form.Item>
                    </Col>
                    <Col span={12}>
                        <Form.Item label="Ставка на 19-36" style={styles.formItem}>
                            <Input
                                placeholder="Ставка на 19-36"
                                size="large"
                                prefix="₽"
                                value={nineteenToThirtySixBet}
                                onChange={handleBetChange('nineteenToThirtySixBet')}
                                style={styles.input}
                            />
                        </Form.Item>
                    </Col>
                </Row>
                <Row gutter={[16, 16]}>
                    <Col span={8}>
                        <Form.Item label="Ставка на First 2 to 1" style={styles.formItem}>
                            <Input
                                placeholder="Ставка на First 2 to 1"
                                size="large"
                                prefix="₽"
                                value={first2To1Bet}
                                onChange={handleBetChange('first2To1Bet')}
                                style={styles.input}
                            />
                        </Form.Item>
                    </Col>
                    <Col span={8}>
                        <Form.Item label="Ставка на Second 2 to 1" style={styles.formItem}>
                            <Input
                                placeholder="Ставка на Second 2 to 1"
                                size="large"
                                prefix="₽"
                                value={second2To1Bet}
                                onChange={handleBetChange('second2To1Bet')}
                                style={styles.input}
                            />
                        </Form.Item>
                    </Col>
                    <Col span={8}>
                        <Form.Item label="Ставка на Third 2 to 1" style={styles.formItem}>
                            <Input
                                placeholder="Ставка на Third 2 to 1"
                                size="large"
                                prefix="₽"
                                value={third2To1Bet}
                                onChange={handleBetChange('third2To1Bet')}
                                style={styles.input}
                            />
                        </Form.Item>
                    </Col>
                </Row>  
            </Form>
        </Card>
    );
}

const styles = {
    card: {
        padding: '32px',
        borderRadius: '16px',
        boxShadow: '0 10px 20px rgba(0, 0, 0, 0.15)',
        background: 'linear-gradient(135deg, #ff416c 0%, #ff4b2b 100%)',
        marginBottom: '24px',
        border: '1px solid #fff',
    },
    form: {
        width: '100%',
        backgroundColor: 'transparent',
    },
    formItem: {
        marginBottom: '20px',
        backgroundColor: 'transparent',
    },
    input: {
        width: '100%',
        borderRadius: '12px',
        border: '1px solid #e0e0e0',
        padding: '12px 18px',
        fontSize: '16px',
        backgroundColor: '#fff',
        color: '#333',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
        transition: 'border-color 0.3s ease',
    },
    inputFocus: {
        borderColor: '#ff416c',
    }
};


export default BetForm;
