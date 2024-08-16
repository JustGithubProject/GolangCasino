import React, { useState, useEffect } from 'react';
import axios from 'axios';
import styled from 'styled-components';

const Container = styled.div`
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
    background-color: #f9f9f9;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
    border-radius: 8px;
`;

const Header = styled.h1`
    text-align: center;
    color: #333;
    margin-bottom: 20px;
`;

const PaymentList = styled.ul`
    list-style: none;
    padding: 0;
`;

const PaymentItem = styled.li`
    background-color: #fff;
    margin-bottom: 15px;
    padding: 15px;
    border-radius: 8px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
    gap: 10px;
`;

const PaymentInfo = styled.div`
    font-size: 16px;
    color: #555;
`;

const Button = styled.button`
    padding: 10px 20px;
    font-size: 14px;
    color: #fff;
    background-color: #007bff;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.3s ease;

    &:hover {
        background-color: #0056b3;
    }
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

    const updateUserBalance = async () => {
        
    }

    useEffect(() => {
        getListOrdersOrPayments();
    }, []);

    if (loading) {
        return <div>Loading...</div>;
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
                        <Button>Забрать деньги</Button>
                    </PaymentItem>
                ))}
            </PaymentList>
        </Container>
    );
};

export default PaymentHistoryPage;
