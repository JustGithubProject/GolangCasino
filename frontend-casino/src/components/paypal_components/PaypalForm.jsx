import React from 'react';
import styled from 'styled-components';
import { PayPalScriptProvider, PayPalButtons } from '@paypal/react-paypal-js';

const clientId = "AR472O-P4C1udEvKpzyCKdo7yxC4PuZqGYuVn43PR37ZbhlSPlToUN40nSlutsL4z4Sypnfkf6Pjazfz" 


const PayPalComponent = ({ amount, currency, returnURL, cancelURL }) => {
  const createOrder = (data, actions) => {
    return fetch('http://localhost:8081/paypal/paypal-payment/', { 
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        amount,
        currency,
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
  padding: 20px;
  background-color: #f7f7f7;
  border-radius: 10px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  max-width: 400px;
  margin: auto;
`;

const Header = styled.h1`
  font-size: 24px;
  color: #333;
  margin-bottom: 20px;
`;
