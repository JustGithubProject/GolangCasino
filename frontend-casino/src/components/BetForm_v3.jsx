import React from 'react';
import { Form, Input, Row, Col, Card, Button } from 'antd';
import { CloseOutlined } from '@ant-design/icons';

function BetFormV3({ betValues, handleBetChange, reset, closeForm }) {
    const {
        betAmount, zeroBet, evenBet, oddBet, redBet, blackBet, first12Bet, second12Bet, third12Bet,
        oneToEighteenBet, nineteenToThirtySixBet, first2To1Bet, second2To1Bet, third2To1Bet
    } = betValues;

    return (
        <Card style={styles.card} hoverable>
            <div style={styles.header}>
                <Button
                    type="text"
                    icon={<CloseOutlined />}
                    onClick={closeForm}
                    style={styles.closeButton}
                />
            </div>
            <Form layout="vertical" style={styles.form}>
                <Form.Item label={<span style={styles.label}>Введите ставку на число</span>} style={styles.formItem}>
                    <Input
                        placeholder="Ставка на число"
                        size="large"
                        prefix="₽"
                        value={betAmount}
                        onChange={handleBetChange('betAmount')}
                        style={styles.input}
                    />
                </Form.Item>
                <Form.Item label={<span style={styles.label}>Ставка на зеро (0)</span>} style={styles.formItem}>
                    <Input
                        placeholder="Ставка на зеро"
                        size="large"
                        prefix="₽"
                        value={zeroBet}
                        onChange={handleBetChange('zeroBet')}
                        style={styles.input}
                    />
                </Form.Item>
                <Row gutter={[16, 16]}>
                    <Col span={12}>
                        <Form.Item label={<span style={styles.label}>Ставка на четное</span>} style={styles.formItem}>
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
                        <Form.Item label={<span style={styles.label}>Ставка на нечетное</span>} style={styles.formItem}>
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
                        <Form.Item label={<span style={styles.label}>Ставка на красное</span>} style={styles.formItem}>
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
                        <Form.Item label={<span style={styles.label}>Ставка на черное</span>} style={styles.formItem}>
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
                        <Form.Item label={<span style={styles.label}>Ставка на 1st 12</span>} style={styles.formItem}>
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
                        <Form.Item label={<span style={styles.label}>Ставка на 2nd 12</span>} style={styles.formItem}>
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
                        <Form.Item label={<span style={styles.label}>Ставка на 3rd 12</span>} style={styles.formItem}>
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
                        <Form.Item label={<span style={styles.label}>Ставка на 1-18</span>} style={styles.formItem}>
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
                        <Form.Item label={<span style={styles.label}>Ставка на 19-36</span>} style={styles.formItem}>
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
                        <Form.Item label={<span style={styles.label}>Ставка на First 2 to 1</span>} style={styles.formItem}>
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
                        <Form.Item label={<span style={styles.label}>Ставка на Second 2 to 1</span>} style={styles.formItem}>
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
                        <Form.Item label={<span style={styles.label}>Ставка на Third 2 to 1</span>} style={styles.formItem}>
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
                <div style={styles.footer}>
                    <Button type="primary" onClick={reset} style={styles.resetButton}>
                        Сбросить
                    </Button>
                </div>
            </Form>
        </Card>
    );
}

const styles = {
    card: {
        padding: '32px',
        borderRadius: '16px',
        boxShadow: '0 10px 20px rgba(0, 0, 0, 0.15)',
        background: 'linear-gradient(135deg, #1a1a1a 0%, #000 100%)',
        marginBottom: '24px',
        border: '1px solid #fff',
        transition: 'transform 0.3s ease',
    },
    cardHover: {
        transform: 'scale(1.05)',
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
        border: '1px solid #444',
        padding: '12px 18px',
        fontSize: '16px',
        backgroundColor: '#fff',
        color: '#000',
        boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
        transition: 'border-color 0.3s ease, transform 0.3s ease',
    },
    inputFocus: {
        borderColor: '#ff416c',
    },
    header: {
        display: 'flex',
        justifyContent: 'flex-end',
    },
    closeButton: {
        fontSize: '18px',
        color: '#fff',
        backgroundColor: 'transparent',
        border: 'none',
        cursor: 'pointer',
    },
    label: {
        color: '#fff',
        fontWeight: 'bold',
    },
    footer: {
        display: 'flex',
        justifyContent: 'flex-end',
        marginTop: '20px',
    },
    resetButton: {
        backgroundColor: '#ff416c',
        borderColor: '#ff416c',
    }
};

export default BetFormV3;
