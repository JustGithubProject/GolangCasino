import React, { useState, useEffect } from 'react';
import axios from 'axios';
import styled, { keyframes } from 'styled-components';

// Анимация загрузки
const spin = keyframes`
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
`;

// Анимация появления элемента
const fadeIn = keyframes`
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
`;

// Общий контейнер страницы
const Container = styled.div`
  max-width: 900px;
  margin: 40px auto;
  padding: 30px;
  background-color: #f4f7fa;
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.15);
  border-radius: 20px;
  overflow: hidden;
  font-family: 'Roboto', sans-serif;
`;

// Заголовок страницы
const Header = styled.h1`
  text-align: center;
  color: #333;
  font-size: 2.8rem;
  margin-bottom: 40px;
  font-weight: bold;
  text-transform: uppercase;
  letter-spacing: 1.5px;
`;

// Список платежей
const PaymentList = styled.ul`
  list-style: none;
  padding: 0;
  margin: 0;
`;

// Элемент списка платежей
const PaymentItem = styled.li`
  background-color: #ffffff;
  margin-bottom: 24px;
  padding: 24px;
  border-radius: 15px;
  box-shadow: 0 6px 15px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  gap: 16px;
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  animation: ${fadeIn} 0.6s ease-out;

  &:hover {
    transform: translateY(-6px);
    box-shadow: 0 12px 25px rgba(0, 0, 0, 0.15);
  }
`;

// Информация о платеже
const PaymentInfo = styled.div`
  font-size: 18px;
  color: #555;
  line-height: 1.6;
  font-weight: 500;

  & strong {
    color: #333;
    font-weight: 700;
  }
`;

// Кнопка для получения денег
const Button = styled.button`
  padding: 16px 32px;
  font-size: 18px;
  color: #fff;
  background-color: #28a745;
  border: none;
  border-radius: 50px;
  cursor: pointer;
  align-self: flex-end;
  transition: background-color 0.3s ease, transform 0.2s ease;
  box-shadow: 0 4px 12px rgba(40, 167, 69, 0.4);

  &:hover {
    background-color: #218838;
    box-shadow: 0 6px 18px rgba(40, 167, 69, 0.6);
  }
  
  &:active {
    transform: scale(0.95);
  }
`;

// Контейнер для загрузки
const Loader = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
`;

// Анимация спиннера
const Spinner = styled.div`
  border: 10px solid #f3f3f3;
  border-top: 10px solid #007bff;
  border-radius: 50%;
  width: 80px;
  height: 80px;
  animation: ${spin} 1s linear infinite;
`;

const PaymentHistoryPage = () => {
    const [payments, setPayments] = useState([]);
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

            for (const payment of response.data) {
                if (payment.Status == "APPROVED") {
                }
                await checkOrderIDPayment(payment.OrderID);

            }

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
                            <strong>Date:</strong> {new Date(payment.CreatedAt).toLocaleDateString()}
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
