import React, { useState } from 'react';
import { Elements, CardElement, useStripe, useElements } from '@stripe/react-stripe-js';
import { loadStripe } from '@stripe/stripe-js';
import styled from 'styled-components';
import axios from 'axios';

// Ваш публикуемый ключ Stripe
const stripePromise = loadStripe('your_publishable_key');

const Wrapper = styled.div`
  max-width: 500px;
  margin: 0 auto;
  padding: 20px;
  background: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
`;

const Form = styled.form`
  display: flex;
  flex-direction: column;
`;

const CardInput = styled.div`
  margin-bottom: 20px;
`;

const SubmitButton = styled.button`
  padding: 10px;
  background: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
  transition: background 0.2s ease;

  &:hover {
    background: #0056b3;
  }
`;

const ErrorMessage = styled.div`
  color: red;
  margin-top: 10px;
`;

const StripeForm = () => {
  const stripe = useStripe();
  const elements = useElements();
  const [error, setError] = useState(null);
  const [isProcessing, setIsProcessing] = useState(false);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setError(null);
    setIsProcessing(true);

    if (!stripe || !elements) {
      return;
    }

    const { error, paymentIntent } = await axios.post('/create-payment-intent', {
      amount: 5000 // сумма в центах
    }).then(response => response.data);

    if (error) {
      setError(error.message);
      setIsProcessing(false);
      return;
    }

    const { error: stripeError } = await stripe.confirmCardPayment(paymentIntent.client_secret, {
      payment_method: {
        card: elements.getElement(CardElement),
        billing_details: {
          name: 'Customer Name',
        },
      },
    });

    if (stripeError) {
      setError(stripeError.message);
    } else {
      // Платеж успешен
      alert('Payment successful!');
    }

    setIsProcessing(false);
  };

  return (
    <Wrapper>
      <Form onSubmit={handleSubmit}>
        <CardInput>
          <CardElement />
        </CardInput>
        {error && <ErrorMessage>{error}</ErrorMessage>}
        <SubmitButton type="submit" disabled={isProcessing}>
          {isProcessing ? 'Processing...' : 'Pay'}
        </SubmitButton>
      </Form>
    </Wrapper>
  );
};

const App = () => (
  <Elements stripe={stripePromise}>
    <StripeForm />
  </Elements>
);

export default App;
