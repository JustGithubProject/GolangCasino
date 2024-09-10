import React, { useState } from 'react';
import { Elements, CardElement, useStripe, useElements } from '@stripe/react-stripe-js';
import { loadStripe } from '@stripe/stripe-js';
import styled from 'styled-components';
import axios from 'axios';

const stripePromise = loadStripe('your_publishable_key');

const Wrapper = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: #f0f2f5;
`;

const FormContainer = styled.div`
  max-width: 400px;
  width: 100%;
  padding: 20px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
`;

const Form = styled.form`
  display: flex;
  flex-direction: column;
`;

const FormField = styled.div`
  margin-bottom: 20px;
`;

const Label = styled.label`
  margin-bottom: 8px;
  font-weight: bold;
`;

const Input = styled.input`
  padding: 8px;
  font-size: 16px;
  border-radius: 4px;
  border: 1px solid #ccc;
  width: 100%;
`;

const SubmitButton = styled.button`
  padding: 12px;
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

  &:disabled {
    background: #cccccc;
    cursor: not-allowed;
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
  const [amount, setAmount] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();
    setError(null);
    setIsProcessing(true);

    if (!stripe || !elements) {
      return;
    }

    const amountInCents = Math.round(parseFloat(amount) * 100);

    const { error, paymentIntent } = await axios.post('http://127.0.0.1:8081/stripe/create/payment/intent', {
      amount: amountInCents // сумма в центах
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
    <Elements stripe={stripePromise}>
      <Wrapper>
        <FormContainer>
          <Form onSubmit={handleSubmit}>
            <FormField>
              <Label htmlFor="amount">Сумма</Label>
              <Input 
                id="amount" 
                type="number" 
                value={amount} 
                onChange={(e) => setAmount(e.target.value)} 
                placeholder="Сумма"
                required 
              />
            </FormField>
            <FormField>
              <Label>Детали карты</Label>
              <CardElement />
            </FormField>
            {error && <ErrorMessage>{error}</ErrorMessage>}
            <SubmitButton type="submit" disabled={isProcessing || !amount}>
              {isProcessing ? 'Processing...' : 'Pay'}
            </SubmitButton>
          </Form>
        </FormContainer>
      </Wrapper>
    </Elements>
  );
};

export default StripeForm;
