import React from 'react';
import { GoogleOAuthProvider, GoogleLogin } from '@react-oauth/google';

const GoogleButton = () => {
  const handleLoginSuccess = (response) => {
    console.log('Google login success:', response);
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
