import React, { useState, useEffect } from 'react';
import axios from 'axios';
import styled from 'styled-components';
import backgroundImage from '../../images/casinoImage_2.png';
import Header from '../Header'; // Импортируем Header

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
  padding-top: 80px; /* Учитываем высоту хедера */
`;

const GameBoard = styled.div`
  display: grid;
  grid-template-columns: repeat(6, 1fr);
  grid-gap: 10px;
  background: rgba(255, 255, 255, 0.8); /* полупрозрачный белый фон */
  padding: 10px;
  border-radius: 10px;
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
`;

const Button = styled.button`
  margin-top: 20px;
  padding: 10px 20px;
  background: #28a745;
  color: #fff;
  border: none;
  border-radius: 5px;
  cursor: pointer;

  &:hover {
    background: #218838;
  }
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
        <p>Balance: ${balance}</p>
      </Wrapper>
    </>
  );
};

export default SweetBonanzaCard;
