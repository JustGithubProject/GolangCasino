import React, { useState, useEffect } from 'react';
import axios from 'axios';

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

            console.log("Response:", response.status);
            console.log("Response data", response.data);

            const orderDetailsDict = {};
            for (const payment of response.data) {
                const data = await checkOrderIDPayment(payment.OrderID);
                console.log("DATA=", data);
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

            console.log("Response:", response.status);
            console.log("Response data", response.data);

            return response.data;

        } catch (error) {
            console.error("Error fetching payments", error);
            return null;
        }
    };

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
        <div>
            <h1 align="center">История платежей</h1>
            <ul>
                {payments.map((payment) => (
                    <li key={payment.ID}>
                        <h1>--------------------</h1>
                        <div>
                            <strong>ID:</strong> {payment.ID}
                        </div>
                        <div>
                            <strong>OrderID:</strong> {payment.OrderID}
                        </div>
                        <div>
                            <strong>Status:</strong> {dictOrder[payment.OrderID]?.status || "Неизвестен"}
                        </div>
                        <div>
                            <strong>Amount:</strong> {payment.amount}
                        </div>
                        <div>
                            <strong>Date: </strong> {payment.CreatedAt}
                        </div>
                        <div>
                            <button>Забрать деньги</button>
                        </div>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default PaymentHistoryPage;
