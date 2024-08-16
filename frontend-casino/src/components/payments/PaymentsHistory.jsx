import React, { useState, useEffect } from 'react';
import axios from 'axios';
import styled, { keyframes } from 'styled-components';

// Загрузка анимации
const spin = keyframes`
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
`;

const Container = styled.div`
    max-width: 900px;
    margin: 40px auto;
    padding: 30px;
    background-color: #ffffff;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    border-radius: 12px;
    overflow: hidden;
`;

const Header = styled.h1`
    text-align: center;
    color: #333;
    font-size: 2rem;
    margin-bottom: 30px;
`;

const PaymentList = styled.ul`
    list-style: none;
    padding: 0;
`;

const PaymentItem = styled.li`
    background-color: #fefefe;
    margin-bottom: 20px;
    padding: 20px;
    border-radius: 10px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
    gap: 12px;
`;

const PaymentInfo = styled.div`
    font-size: 16px;
    color: #555;
`;

const Button = styled.button`
    padding: 12px 24px;
    font-size: 16px;
    color: #fff;
    background-color: #007bff;
    border: none;
    border-radius: 8px;
    cursor: pointer;
    transition: background-color 0.3s ease, transform 0.2s ease;
    
    &:hover {
        background-color: #0056b3;
    }
    
    &:active {
        transform: scale(0.98);
    }
`;

const Loader = styled.div`
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
`;

const Spinner = styled.div`
    border: 8px solid #f3f3f3;
    border-top: 8px solid #007bff;
    border-radius: 50%;
    width: 60px;
    height: 60px;
    animation: ${spin} 1s linear infinite;
`;

const PaymentHistoryPage = () => {
    const [payments, setPayments] = useState([]);
    const [dictOrder, setDictOrder] = useState({});
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    const getListOrdersOrPayments = async () => {
        try {
            const token = localStorage.getItem("token");
            if (!token) {
                throw new Error("No authentication token found");
            }
            const url = "http://127.0.0.1:8081/paypal/payments/history/";

            const response = await axios.get(url, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`,
                },
            });

            const orderDetailsDict = {};
            for (const payment of response.data) {
                const data = await checkOrderIDPayment(payment.OrderID);
                orderDetailsDict[payment.OrderID] = data;
            }

            setDictOrder(orderDetailsDict);
            setPayments(response.data);
            setLoading(false);

        } catch (error) {
            console.error("Error fetching payments", error);
            setError("Error fetching payments");
            setLoading(false);
        }
    };

    const checkOrderIDPayment = async (orderID) => {
        try {
            const token = localStorage.getItem("token");
            if (!token) {
                throw new Error("No authentication token found");
            }
            const url = `http://127.0.0.1:8081/paypal/info/order/?id=${orderID}`;

            const response = await axios.get(url, {
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`,
                },
            });

            return response.data;

        } catch (error) {
            console.error("Error fetching payments", error);
            return null;
        }
    };

    const updateUserBalance = async (amountOfMoney, orderID, currentStatus) => {
        const url = "http://127.0.0.1:8081/paypal/update/pickup/money";
        const token = localStorage.getItem("token");
        
        const data = {
            total: amountOfMoney,
            order_id: orderID,
            status: currentStatus
        };

        const headers = {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`,  
        };
        try {
            const response = await axios.post(url, data, { headers });
            console.log("Response", response.data);
            window.location.href = "/"
        } catch(error) {
            console.log("Error", error);
        }
    };

    useEffect(() => {
        getListOrdersOrPayments();
    }, []);

    if (loading) {
        return (
            <Loader>
                <Spinner />
            </Loader>
        );
    }

    if (error) {
        return <div>{error}</div>;
    }

    return (
        <Container>
            <Header>История платежей</Header>
            <PaymentList>
                {payments.map((payment) => (
                    <PaymentItem key={payment.ID}>
                        <PaymentInfo>
                            <strong>OrderID:</strong> {payment.OrderID}
                        </PaymentInfo>
                        <PaymentInfo>
                            <strong>Status:</strong> {payment.Status}
                        </PaymentInfo>
                        <PaymentInfo>
                            <strong>Amount:</strong> {payment.Amount} USD
                        </PaymentInfo>
                        <PaymentInfo>
                            <strong>Date:</strong> {payment.CreatedAt}
                        </PaymentInfo>
                        {payment.Status === 'APPROVED' && (
                            <Button onClick={() => updateUserBalance(payment.Amount, payment.OrderID, payment.Status)}>
                                Забрать деньги
                            </Button>
                        )}
                    </PaymentItem>
                ))}
            </PaymentList>
        </Container>
    );
};

export default PaymentHistoryPage;
