import React, { useState, useEffect } from 'react';
import axios from 'axios';
import styled, { keyframes } from 'styled-components';
import backgroundImage from '../../images/casinoImage_2.png';
import Header from '../Header';

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  background: linear-gradient(to bottom, #ff416c, #ff4b2b);
  padding: 20px;
  border-radius: 10px;
  background-image: url(${backgroundImage});
  background-size: cover;
  background-position: center;
  height: 100vh;
  width: 100vw;
  box-sizing: border-box;
  padding-top: 80px;
`;

const fadeIn = keyframes`
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
`;

const GameBoard = styled.div`
  display: grid;
  grid-template-columns: repeat(6, 1fr);
  grid-gap: 10px;
  background: rgba(255, 255, 255, 0.8);
  padding: 20px;
  border-radius: 10px;
  animation: ${fadeIn} 0.5s ease-in;
  margin-top: 20px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
`;

const bounce = keyframes`
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
`;

const Symbol = styled.div`
  width: 60px;
  height: 60px;
  display: flex;
  justify-content: center;
  align-items: center;
  background: ${props => props.color || '#ccc'};
  border-radius: 10px;
  font-size: 24px;
  color: #fff;
  animation: ${bounce} 1s infinite;
  transition: transform 0.3s ease-in-out;

  &:hover {
    transform: scale(1.1);
  }
`;

const Button = styled.button`
  margin-top: 20px;
  padding: 10px 20px;
  background: #28a745;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 18px;
  transition: background 0.3s ease, transform 0.3s ease;

  &:hover {
    background: #218838;
    transform: scale(1.05);
  }
`;

const BalanceText = styled.p`
  margin-top: 10px;
  font-size: 18px;
  color: #fff;
  background: rgba(0, 0, 0, 0.5);
  padding: 10px 20px;
  border-radius: 5px;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
  animation: ${fadeIn} 0.5s ease-in;
`;

const Title = styled.h1`
  color: #fff;
  margin-top: 20px;
  font-size: 36px;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);
`;

const symbols = [
  { id: 1, name: 'Banana', emoji: '🍌', color: '#FFE135' },
  { id: 2, name: 'Grapes', emoji: '🍇', color: '#6F2DA8' },
  { id: 3, name: 'Watermelon', emoji: '🍉', color: '#FC6C85' },
  { id: 4, name: 'Plum', emoji: '🍑', color: '#DDA0DD' },
  { id: 5, name: 'Apple', emoji: '🍎', color: '#FF0800' },
  { id: 6, name: 'BlueCandy', emoji: '🍬', color: '#4169E1' },
  { id: 7, name: 'GreenCandy', emoji: '🍬', color: '#32CD32' },
  { id: 8, name: 'PurpleCandy', emoji: '🍬', color: '#9370DB' },
  { id: 9, name: 'RedCandy', emoji: '🍬', color: '#FF4500' },
  { id: 10, name: 'Scatter', emoji: '⭐', color: '#FFD700' },
  { id: 11, name: 'Bomb2X', emoji: '💣', color: '#FFA500' },
  { id: 12, name: 'Bomb3X', emoji: '💣', color: '#FF4500' },
  { id: 13, name: 'Bomb5X', emoji: '💣', color: '#FF6347' },
  { id: 14, name: 'Bomb10X', emoji: '💣', color: '#FF0000' },
  { id: 15, name: 'Bomb25X', emoji: '💣', color: '#8B0000' },
  { id: 16, name: 'Bomb50X', emoji: '💣', color: '#800000' },
  { id: 17, name: 'Bomb100X', emoji: '💣', color: '#4B0082' }
];

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
  const [balance, setBalance] = useState(0);

  useEffect(() => {
    // Initialize game board with random symbols
    setGameBoard(generateRandomGameBoard());
  }, []);

  const handleSpin = async () => {
    try {
      const response = await axios.post('http://localhost:8081/spin-slot-v1/?spinBet=10');
      const { playingField, balance } = response.data;
      setGameBoard(playingField);
      setBalance(balance);
    } catch (error) {
      console.error('Error fetching spin data:', error);
    }
  };

  return (
    <>
      <Header username="User" balance={balance} handleLogout={() => {}} />
      <Wrapper>
        <Title>Sweet Bonanza</Title>
        <GameBoard>
          {gameBoard.flat().map((symbolId, index) => {
            const symbol = symbols.find(s => s.id === symbolId);
            return (
              <Symbol key={index} color={symbol.color}>
                {symbol.emoji}
              </Symbol>
            );
          })}
        </GameBoard>
        <Button onClick={handleSpin}>Spin</Button>
        <BalanceText>Balance: ${balance}</BalanceText>
      </Wrapper>
    </>
  );
};

export default SweetBonanzaCard;
