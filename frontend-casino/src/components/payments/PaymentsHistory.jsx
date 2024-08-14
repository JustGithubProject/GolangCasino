import React, { useState, useEffect } from 'react';
import axios from 'axios';

const PaymentHistoryPage = () => {
    const [payments, setPayments] = useState([]);
    const [orderDetails, setOrderDetails] = useState();
    const [dictOrder, setDictOrder] = useState({});
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    // Function to fetch payments
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

          
        
            //... TODO
            const paymentOrderDetails = response.data.map((payment) => {
                const data = await checkOrderIDPayment(payment.OrderID) 
                setDictOrder(prevDict => ({
                    ...prevDict,
                    [payment.OrderID]: value
                  }));
            });

            setPayments(response.data);
            setLoading(false); // Set loading to false once data is fetched

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
        }
    };

    // Fetch payments when component mounts
    useEffect(() => {
        getListOrdersOrPayments();
    }, []);

    if (loading) {
        return <div>Loading...</div>; // Show a loading message or spinner
    }

    if (error) {
        return <div>{error}</div>; // Show error message
    }

    return (
        <div>
            <h1 align="center">История платежей</h1>
            <ul>
                {payments.map((payment) => (
                    <li key={payment.ID}>
                        <div>
                            <strong>ID:</strong> {payment.ID}
                        </div>
                        <div>
                            <strong>OrderID:</strong> {payment.OrderID}
                        </div>
                        <div>
                            <strong>Status:</strong> {payment.Status}
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
