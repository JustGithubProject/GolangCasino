import React from 'react';
import { GoogleOAuthProvider, GoogleLogin } from '@react-oauth/google';
import axios from 'axios';

const GoogleButton = () => {
  const handleLoginSuccess = async (response) => {
    console.log('Google login success:', response);
    const token = response.credential;

    // Отправляем токен на сервер для обмена на данные о пользователе
    try {
      const result = await axios.post('http://127.0.0.1:8081/google/auth/callback', {
        token,
      });
      console.log(result.data);
    } catch (error) {
      console.error('Error during authentication:', error);
    }
  };

  const handleLoginFailure = (error) => {
    console.error('Google login failed:', error);
  };

  return (
    <GoogleOAuthProvider clientId="609582150163-ejcmmse6ut85n5iv2sm6s7k4nauirlk8.apps.googleusercontent.com">
      <GoogleLogin
        onSuccess={handleLoginSuccess}
        onError={handleLoginFailure}
        useOneTap
      />
    </GoogleOAuthProvider>
  );
};

export default GoogleButton;
