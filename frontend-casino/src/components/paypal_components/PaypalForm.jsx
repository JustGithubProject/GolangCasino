import React, { useEffect, useRef, useState } from 'react';
import styled from 'styled-components';
import axios from 'axios';

const PayPalComponent = () => {
  const paypalRef = useRef();
  const [amount, setAmount] = useState('');
  const [currency, setCurrency] = useState('USD');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    const loadPayPalScript = () => {
      const script = document.createElement('script');
      script.src = "https://www.paypal.com/sdk/js?client-id=AX4y6W2OQIABZYyC4OoPfLY1pxRTlQsg80XghZG2Jxs0Fd-OJkNR6gQXTesLUUeGuoxfFO7MOw9Ql-2G&components=buttons&enable-funding=venmo,paylater";
      script.addEventListener('load', () => {
        window.paypal.Buttons({
          style: {
            shape: 'rect',
            layout: 'vertical',
          },
          createOrder: async () => {
            try {
              const orderId = await createPayPalOrder(amount, currency, setLoading, setError);
              return orderId;
            } catch (error) {
              console.error('Error creating PayPal order:', error);
            }
          },
          onApprove: async (data, actions) => {
            try {
              await actions.order.capture();
              console.log('Transaction completed by', data.payer.name.given_name);
            } catch (error) {
              console.error('Error capturing PayPal order:', error);
            }
          },
        }).render(paypalRef.current);
      });
      document.body.appendChild(script);
    };

    loadPayPalScript();
  }, [amount, currency]);

  const createPayPalOrder = async (amountP, currencyP, setLoading, setError) => {
    setLoading(true);
    setError('');
    try {
      const token = localStorage.getItem('token');
      if (!token) {
        throw new Error('No authentication token found');
      }

      const url = "http://127.0.0.1:8081/paypal/create/order/";
      const response = await axios.post(url, {
        currency_code: currencyP,
        value: amountP,
      }, {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
      });
      window.location.href = response.data.links[1].href
      return response.data.id; 
    } catch (error) {
      setError(error.message);
      console.error('Error creating PayPal order:', error);
    } finally {
      setLoading(false);
    }
  };

  

  return (
    <Container>
      <Header>Завершите оплату</Header>
      {error && <ErrorMessage>{error}</ErrorMessage>}
      <Form>
        <Label>
          Сумма:
          <Input 
            type="number"
            value={amount} 
            onChange={(e) => setAmount(e.target.value)} 
            placeholder="Введите сумму" 
          />
        </Label>
        <Label>
          Валюта:
          <Select 
            value={currency} 
            onChange={(e) => setCurrency(e.target.value)}
          >
            <option value="USD">USD</option>
            <option value="EUR">EUR</option>
            <option value="RUB">RUB</option>
            <option value="GBP">GBP</option>
          </Select>
        </Label>
      </Form>
      <div ref={paypalRef}></div>

    </Container>
  );
};

export default PayPalComponent;

const Container = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 30px;
  background-color: #ffffff;
  border-radius: 12px;
  box-shadow: 0 6px 10px rgba(0, 0, 0, 0.1);
  max-width: 450px;
  margin: 50px auto;
`;

const Header = styled.h1`
  font-size: 28px;
  color: #333;
  margin-bottom: 20px;
  text-align: center;
`;

const Form = styled.div`
  display: flex;
  flex-direction: column;
  width: 100%;
  margin-bottom: 20px;
`;

const Label = styled.label`
  font-size: 18px;
  color: #333;
  margin-bottom: 15px;
  display: flex;
  flex-direction: column;
`;

const Input = styled.input`
  padding: 10px;
  font-size: 16px;
  border-radius: 8px;
  border: 1px solid #ccc;
  margin-top: 5px;
  transition: border-color 0.3s;
  
  &:focus {
    border-color: #3498db;
    outline: none;
  }
`;

const Select = styled.select`
  padding: 10px;
  font-size: 16px;
  border-radius: 8px;
  border: 1px solid #ccc;
  margin-top: 5px;
  transition: border-color 0.3s;
  
  &:focus {
    border-color: #3498db;
    outline: none;
  }
`;

const CreditCardButton = styled.button`
  padding: 10px 20px;
  font-size: 16px;
  color: #fff;
  background-color: #007bff;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  margin-top: 20px;
  transition: background-color 0.3s;
  
  &:hover {
    background-color: #0056b3;
  }
  
  &:disabled {
    background-color: #cccccc;
    cursor: not-allowed;
  }
`;

const ErrorMessage = styled.div`
  color: red;
  margin-bottom: 20px;
  font-size: 16px;
`;
