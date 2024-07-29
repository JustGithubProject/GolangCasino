import React, { useState } from 'react';
import styled from 'styled-components';
import { PayPalScriptProvider, PayPalButtons } from '@paypal/react-paypal-js';

const clientId = "AR472O-P4C1udEvKpzyCKdo7yxC4PuZqGYuVn43PR37ZbhlSPlToUN40nSlutsL4z4Sypnfkf6Pjazfz";

const PayPalComponent = ({ returnURL, cancelURL }) => {
  const [amount, setAmount] = useState('');
  const [currency, setCurrency] = useState('USD');

  const createOrder = (data, actions) => {
    return fetch('http://localhost:8081/paypal/paypal-payment/', { 
      method: 'POST',
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      body: new URLSearchParams({
        Total: amount,
        Currency: currency,
        return_url: returnURL,
        cancel_url: cancelURL,
      }),
    })
      .then(response => response.json())
      .then(payment => {
        return payment.id;
      });
  };

  const onApprove = (data, actions) => {
    return actions.order.capture().then(details => {
      console.log('Transaction completed by', details.payer.name.given_name);
    });
  };

  return (
    <Container>
      <Header>Завершите оплату</Header>
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
          createOrder={createOrder}
          onApprove={onApprove}
          style={{ layout: 'vertical', color: 'blue', shape: 'pill', label: 'paypal' }}
        />
      </PayPalScriptProvider>
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
