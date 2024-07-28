import React from 'react';
import { PayPalScriptProvider, PayPalButtons } from '@paypal/react-paypal-js';

const clientId = 'AR472O-P4C1udEvKpzyCKdo7yxC4PuZqGYuVn43PR37ZbhlSPlToUN40nSlutsL4z4Sypnfkf6Pjazfz';

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
    <PayPalScriptProvider options={{ "client-id": clientId }}>
      <PayPalButtons
        createOrder={createOrder}
        onApprove={onApprove}
      />
    </PayPalScriptProvider>
  );
};

export default PayPalComponent;
