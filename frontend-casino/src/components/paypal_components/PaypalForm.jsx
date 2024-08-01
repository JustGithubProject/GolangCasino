import React, { useState } from 'react';
import styled from 'styled-components';
import axios from 'axios';
import { PayPalScriptProvider, PayPalButtons } from '@paypal/react-paypal-js';

const clientId = "AR472O-P4C1udEvKpzyCKdo7yxC4PuZqGYuVn43PR37ZbhlSPlToUN40nSlutsL4z4Sypnfkf6Pjazfz";

const PayPalComponent = ({ returnURL, cancelURL }) => {
  const [amount, setAmount] = useState('');
  const [currency, setCurrency] = useState('USD');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');


  console.log("Total", amount)
  console.log("Currency", currency)
  
  const createPayPalOrder = async (data, actions) => {
    setLoading(true);
    setError('');

    try {
      const response = await axios.post('http://localhost:8081/paypal/paypal-payment/', {
        Total: amount,
        Currency: currency,
        return_url: returnURL,
        cancel_url: cancelURL,
      });

      if (response.status !== 200) {
        throw new Error(response.data.message || 'Error creating PayPal order');
      }

      return response.data.id;  
    } catch (error) {
      setError(error.message);
      console.error('Error creating PayPal order:', error);
    } finally {
      setLoading(false);
    }
  };
  
  


  const onPayPalApprove = async (data, actions) => {
    try {
      await actions.order.capture();
      console.log('Transaction completed by', data.payer.name.given_name);
    } catch (error) {
      console.error('Error capturing PayPal order:', error);
    }
  };

  const createCreditCardPayment = async () => {
    setLoading(true);
    setError('');

    try {
      const response = await fetch('http://localhost:8081/paypal/creditCard-payment/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
        },
        body: new URLSearchParams({
          Total: amount,
          Currency: currency,
        }),
      });

      const payment = await response.json();

      if (!response.ok) {
        throw new Error(payment.message || 'Error initiating credit card payment');
      }

      console.log('Credit card payment initiated:', payment);
    } catch (error) {
      setError(error.message);
      console.error('Error initiating credit card payment:', error);
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
      <PayPalScriptProvider options={{ "client-id": clientId }}>
        <PayPalButtons
          createOrder={createPayPalOrder}
          onApprove={onPayPalApprove}
          style={{ layout: 'vertical', color: 'blue', shape: 'pill', label: 'paypal' }}
        />
      </PayPalScriptProvider>
      <CreditCardButton onClick={createCreditCardPayment} disabled={loading}>
        {loading ? 'Обработка...' : 'Оплатить кредитной картой'}
      </CreditCardButton>
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