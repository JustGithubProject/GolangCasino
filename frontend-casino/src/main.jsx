import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.jsx'
import './index.css'

import { GoogleOAuthProvider } from '@react-oauth/google';

ReactDOM.createRoot(document.getElementById('root')).render(
    <GoogleOAuthProvider clientId="609582150163-ejcmmse6ut85n5iv2sm6s7k4nauirlk8.apps.googleusercontent.com">
        <App />
    </GoogleOAuthProvider>
)
