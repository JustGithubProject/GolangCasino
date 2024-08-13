import React from 'react';
import axios from 'axios';

const PaymentHistoryPage = () => {

    const getListOrdersOrPayments = async () => {
        try{
            const token = localStorage.getItem("token");
            if (!token){
                // TODO...
            }
        }catch(error){

        }
    }

    return (
        <div>
            <h1 align="center">История платежей</h1>
        </div>
    );
};

export default PaymentHistoryPage;