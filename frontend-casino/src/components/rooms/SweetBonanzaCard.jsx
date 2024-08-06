import React, { useState, useEffect } from 'react';
import styled, { keyframes } from 'styled-components';
import backgroundImage from '../../images/backgroundCasinoNew.jpg';
import Header from '../header/Header';
import bananaImage from '../../images/sweetbonanza-items/banana.png';
import grapesImage from '../../images/sweetbonanza-items/grapes.png'
import watermelonImage from '../../images/sweetbonanza-items/watermelon.png'
import plumImage from '../../images/sweetbonanza-items/plum.avif'
import appleImage from '../../images/sweetbonanza-items/apple.jpg'
import bluecandyImage from '../../images/sweetbonanza-items/bluecandy.png'
import greencandyImage from '../../images/sweetbonanza-items/greencandy.png'
import purpleImage from '../../images/sweetbonanza-items/purplecandy.png'
import scatter from '../../images/sweetbonanza-items/scatter.webp'
import redCandy from '../../images/sweetbonanza-items/redcandy.png'
import bomb100xImage from '../../images/sweetbonanza-items/bomb100x.png'
import { fetchWithAuth } from '../auth_components/fetchWrapper';
import * as jwtDecodeModule from 'jwt-decode';

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  background: linear-gradient(to bottom, rgba(255, 64, 108, 0.8), rgba(255, 75, 43, 0.8));
  background-image: url(${backgroundImage});
  background-size: cover;
  background-position: center;
  background-attachment: fixed;
  height: 100vh;
  width: 100vw;
  box-sizing: border-box;
  padding: 60px 20px;
  border: 2px solid #333;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.7);
  overflow: hidden;
`;

const spinAnimation = keyframes`
  0% { transform: rotate(0); }
  50% { transform: rotate(5deg); }
  100% { transform: rotate(0); }
`;

const GameBoard = styled.div`
  display: grid;
  grid-template-columns: repeat(6, 1fr);
  grid-gap: 15px;
  background: rgba(255, 255, 255, 0.15);
  padding: 20px;
  border-radius: 15px;
  margin-top: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.4);
  border: 2px solid #444;
  animation: ${props => props.isSpinning ? spinAnimation : 'none'} 0.6s ease-in-out;
`;

const bounce = keyframes`
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-8px);
  }
`;

const Symbol = styled.div`
  width: 70px;
  height: 70px;
  display: flex;
  justify-content: center;
  align-items: center;
  background: ${props => props.color || '#666'};
  border-radius: 12px;
  font-size: 28px;
  color: #fff;
  animation: ${bounce} 1s infinite;
  transition: transform 0.3s ease, box-shadow 0.3s ease;

  &:hover {
    transform: scale(1.1);
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.6);
  }
`;

const Button = styled.button`
  margin-top: 20px;
  padding: 12px 24px;
  background: #28a745;
  color: #fff;
  border: 2px solid #444;
  border-radius: 8px;
  cursor: pointer;
  font-size: 20px;
  font-weight: bold;
  transition: background 0.3s ease, transform 0.3s ease, box-shadow 0.3s ease;

  &:hover {
    background: #218838;
    transform: scale(1.05);
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.4);
  }
`;

const BalanceText = styled.p`
  margin-top: 15px;
  font-size: 22px;
  color: #ffdf00;
  background: rgba(0, 0, 0, 0.8);
  padding: 12px 24px;
  border-radius: 8px;
  border: 2px solid #555;
  text-shadow: 2px 2px 5px rgba(0, 0, 0, 0.8);
`;

const Title = styled.h1`
  color: #fff;
  margin-top: 20px;
  font-size: 40px;
  text-shadow: 3px 3px 5px rgba(0, 0, 0, 0.8), 0 0 15px #ff4500;
  border: 2px solid #444;
  padding: 15px;
  border-radius: 8px;
  background: rgba(0, 0, 0, 0.7);
`;

const InnerWrapper = styled.div`
  border: 2px solid #444;
  border-radius: 15px;
  padding: 20px;
  background: rgba(255, 255, 255, 0.3);
  backdrop-filter: blur(6px);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.5);
`;

const BetInput = styled.input`
  margin-top: 15px;
  padding: 12px;
  font-size: 20px;
  border: 2px solid #555;
  border-radius: 8px;
  background: rgba(0, 0, 0, 0.8);
  color: #fff;
  width: 90px;
  text-align: center;
`;

const symbols = [
  { id: 1, name: 'Banana', image: bananaImage },
  { id: 2, name: 'Grapes', image: grapesImage },
  { id: 3, name: 'Watermelon', image: watermelonImage },
  { id: 4, name: 'Plum', image: plumImage },
  { id: 5, name: 'Apple', image: appleImage },
  { id: 6, name: 'BlueCandy', image: bluecandyImage },
  { id: 7, name: 'GreenCandy', image: greencandyImage },
  { id: 8, name: 'PurpleCandy', image: purpleImage },
  { id: 9, name: 'RedCandy', image: redCandy },
  { id: 10, name: 'Scatter', image: scatter },
  { id: 11, name: 'Bomb2X', image: '../../images/sweetbonanza-items/banana.png' },
  { id: 12, name: 'Bomb3X', image: '../../images/sweetbonanza-items/banana.png' },
  { id: 13, name: 'Bomb5X', image: '../../images/sweetbonanza-items/banana.png' },
  { id: 14, name: 'Bomb10X', image: '../../images/sweetbonanza-items/banana.png' },
  { id: 15, name: 'Bomb25X', image: '../../images/sweetbonanza-items/banana.png' },
  { id: 16, name: 'Bomb50X', image: '../../images/sweetbonanza-items/banana.png' },
  { id: 17, name: 'Bomb100X', image: bomb100xImage }
];

// Функция генерации случайного игрового поля
const generateRandomGameBoard = () => {
  const gameBoard = [];
  for (let i = 0; i < 5; i++) {
    const row = [];
    for (let j = 0; j < 6; j++) {
      const randomSymbol = symbols[Math.floor(Math.random() * symbols.length)].id;
      row.push(randomSymbol);
    }
    gameBoard.push(row);
  }
  return gameBoard;
};

const SweetBonanzaCard = () => {
  const [gameBoard, setGameBoard] = useState(generateRandomGameBoard());
  const [username, setUsername] = useState(null);
  const [balance, setBalance] = useState(null);
  const [isSpinning, setIsSpinning] = useState(false);
  const [bet, setBet] = useState(10);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      const decodedToken = jwtDecodeModule.jwtDecode(token);
      const username = decodedToken.username;
      setUsername(username);
      fetchUserBalance(username);
    }
  }, []);

  const fetchUserBalance = async (username) => {
    try {
      const response = await fetchWithAuth(`http://localhost:8081/user/name/${username}`);
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      const data = await response.json();
      setBalance(data.Balance);
    } catch (error) {
      console.error('Error fetching user balance:', error);
    }
  };

  const handleSpin = async () => {
    if (bet <= 0) {
      alert('Ставка должна быть больше нуля');
      return;
    }

    setIsSpinning(true);
    setGameBoard(generateRandomGameBoard());

    setTimeout(async () => {
      try {
        const url = `http://localhost:8081/spin-slot-v1/?spinBet=${bet}`;
        const response = await fetchWithAuth(url, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({}),
        });

        if (!response.ok) {
          throw new Error('Network response was not ok');
        }

        const { playingField, balance } = await response.json();
        setGameBoard(playingField);
        setBalance(balance);
      } catch (error) {
        console.error('Error fetching spin data:', error);
      } finally {
        setIsSpinning(false);
      }
    }, 500);
  };

  return (
    <Wrapper>
      <Header username={username} balance={balance} handleLogout={() => {
          localStorage.removeItem('token');
          window.location.reload();
        }} />
      <InnerWrapper>
        <Title>Sweet Bonanza</Title>
        <BalanceText>Ваш баланс: {balance} UAH</BalanceText>
        <BetInput
          type="number"
          value={bet}
          onChange={e => setBet(Number(e.target.value))}
        />
        <Button onClick={handleSpin} disabled={isSpinning}>
          {isSpinning ? 'Спините...' : 'Спин'}
        </Button>
        <GameBoard isSpinning={isSpinning}>
          {gameBoard.flat().map((symbolId, index) => {
            const symbol = symbols.find(s => s.id === symbolId);
            return (
              <Symbol key={index} color={symbol.color}>
                <img src={symbol.image} alt={symbol.name} style={{ width: '100%', height: '100%' }} />
              </Symbol>
            );
          })}
        </GameBoard>
      </InnerWrapper>
    </Wrapper>
  );
};

export default SweetBonanzaCard;
