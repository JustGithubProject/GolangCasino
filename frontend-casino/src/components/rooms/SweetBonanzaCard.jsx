import React, { useState, useEffect, useRef } from 'react';
import styled, { css, keyframes } from 'styled-components';
import Header from '../header/Header';
import { fetchWithAuth } from '../auth_components/fetchWrapper';
import * as jwtDecodeModule from 'jwt-decode';

// Images
import backgroundImage from '../../images/backgroundCasinoNew.jpg';
import bananaImage from '../../images/sweetbonanza-items/banana.png';
import grapesImage from '../../images/sweetbonanza-items/grapes.png';
import watermelonImage from '../../images/sweetbonanza-items/watermelon.png';
import plumImage from '../../images/sweetbonanza-items/plum.png';
import appleImage from '../../images/sweetbonanza-items/apple.png';
import bluecandyImage from '../../images/sweetbonanza-items/bluecandy.png';
import greencandyImage from '../../images/sweetbonanza-items/greencandy.png';
import purpleImage from '../../images/sweetbonanza-items/purplecandy.png';
import scatter from '../../images/sweetbonanza-items/scatter.webp';
import redCandy from '../../images/sweetbonanza-items/redcandy.png';
import bomb100xImage from '../../images/sweetbonanza-items/bomb100x.png';
import bomb25xImage from '../../images/sweetbonanza-items/bomb25x.png';
import bomb50xImage from '../../images/sweetbonanza-items/bomb50x.png';
import bomb2xImage from '../../images/sweetbonanza-items/bomb2x.png';
import bomb3xImage from '../../images/sweetbonanza-items/bomb3x.png';
import bomb5xImage from '../../images/sweetbonanza-items/bomb5x.png';
import bomb10xImage from '../../images/sweetbonanza-items/bomb10x.png';
import sweetbonanzabackground from '../../images/sweetbonanza-items/sweetbonanza-background.png';
import backgroundMusic from '../../images/sweetbonanza-items/music/sweet-music.mp3';




// Стили для Wrapper
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
  border: 2px solid #444;
  box-shadow: 0 6px 24px rgba(0, 0, 0, 0.7);
  overflow: hidden;
`;

// Анимация для падения символов
const fallAnimation = keyframes`
  0% {
    transform: translateY(-300%) rotate(0deg) scale(0.6);
    opacity: 0;
    box-shadow: 0px 0px 0px rgba(0, 0, 0, 0.2);
  }
  40% {
    transform: translateY(50%) rotate(90deg) scale(1.1);
    opacity: 0.5;
    box-shadow: 0px 10px 15px rgba(0, 0, 0, 0.3);
  }
  70% {
    transform: translateY(10%) rotate(180deg) scale(1.05);
    opacity: 0.7;
    box-shadow: 0px 15px 20px rgba(0, 0, 0, 0.4);
  }
  85% {
    transform: translateY(0) rotate(270deg) scale(1);
    opacity: 0.9;
    box-shadow: 0px 20px 25px rgba(0, 0, 0, 0.5);
  }
  100% {
    transform: translateY(0) rotate(360deg) scale(1);
    opacity: 1;
    box-shadow: 0px 25px 30px rgba(0, 0, 0, 0.6);
  }
`;

// Анимация для перемещения GameBoard вниз
const boardAnimation = keyframes`
  0% { transform: translateY(-200%); }
  50% { transform: translateY(15%); }
  70% { transform: translateY(-5%); }
  90% { transform: translateY(5%); }
  100% { transform: translateY(0); }
`;

// Стили для GameBoard
const GameBoard = styled.div`
  display: grid;
  grid-template-columns: repeat(6, 1fr);
  grid-gap: 15px;
  background: rgba(255, 255, 255, 0.2);
  padding: 10px;
  border-radius: 15px;
  margin-top: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.5);
  border: 2px solid #555;
  position: relative;
  overflow: hidden;
  animation: ${props => props.isSpinning ? css`${boardAnimation} 1.2s cubic-bezier(0.52, 0.04, 0.37, 1) both` : 'none'};
`;

const bounceAnimation = keyframes`
  0% { transform: translateY(0); }
  50% { transform: translateY(-10px); }
  100% { transform: translateY(0); }
`;

// Стили для символов
const Symbol = styled.div`
  width: 70px;
  height: 70px;
  display: flex;
  justify-content: center;
  align-items: center;
  background: transparent;
  border-radius: 12px;
  font-size: 28px;
  color: #fff;
  animation: ${bounceAnimation} 1s infinite;
  animation: ${props => props.isSpinning ? css`${fallAnimation} 1s cubic-bezier(0.52, 0.04, 0.37, 1) both` : 'none'};
  animation-delay: ${props => props.delay || '0s'};
`;

// Стили для кнопок
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

// Стили для текста баланса
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

// Стили для заголовка
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

// Стили для внутреннего контейнера
const InnerWrapper = styled.div`
  border: 2px solid #444;
  border-radius: 15px;
  padding: 20px;
  background-image: url(${sweetbonanzabackground});
  background-size: cover;
  background-position: center;
  backdrop-filter: blur(6px);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.5);
`;

// Стили для инпута ставки
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

const MusicButton = styled.button`
  margin-top: 20px;
  padding: 12px 24px;
  background: #007bff;
  color: #fff;
  border: 2px solid #444;
  border-radius: 8px;
  cursor: pointer;
  font-size: 20px;
  font-weight: bold;
  transition: background 0.3s ease, transform 0.3s ease, box-shadow 0.3s ease;

  &:hover {
    background: #0056b3;
    transform: scale(1.05);
    box-shadow: 0 6px 12px rgba(0, 0, 0, 0.4);
  }
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
  { id: 11, name: 'Bomb2X', image: bomb2xImage },
  { id: 12, name: 'Bomb3X', image: bomb3xImage },
  { id: 13, name: 'Bomb5X', image: bomb5xImage },
  { id: 14, name: 'Bomb10X', image: bomb10xImage },
  { id: 15, name: 'Bomb25X', image: bomb25xImage },
  { id: 16, name: 'Bomb50X', image: bomb50xImage },
  { id: 17, name: 'Bomb100X', image: bomb100xImage }
];



const SweetBonanzaCard = () => {
  let defaultMatrix = [[10, 7, 6, 4, 1, 6], [4, 6, 5, 4, 6, 4], [3, 1, 3, 3, 4, 6], [6, 3, 4, 5, 5, 6], [7, 4, 6, 6, 4, 6]]
  const [gameBoard, setGameBoard] = useState(defaultMatrix);
  const [username, setUsername] = useState(null);
  const [balance, setBalance] = useState(null);
  const [isSpinning, setIsSpinning] = useState(false);
  const [bet, setBet] = useState(10);
  const [isMusicPlaying, setIsMusicPlaying] = useState(true); // Состояние для управления музыкой
  const audioRef = useRef(null); // Ссылка на аудио элемент

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      const decodedToken = jwtDecodeModule.jwtDecode(token);
      const username = decodedToken.username;
      setUsername(username);
      fetchUserBalance(username);
    }

    // Воспроизведение музыки при монтировании компонента
    if (audioRef.current) {
      audioRef.current.loop = true; // Зацикливаем музыку
      audioRef.current.volume = 0.5; // Устанавливаем громкость
      if (isMusicPlaying) {
        audioRef.current.play();
      } else {
        audioRef.current.pause();
      }
    }
  }, [isMusicPlaying]);

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
    setGameBoard(defaultMatrix);

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

        let data = await response.json();
        
        // Getting a string from json representing a string of numbers
        let sRow1 = data.sRow1;
        let sRow2 = data.sRow2;
        let sRow3 = data.sRow3;
        let sRow4 = data.sRow4;
        let sRow5 = data.sRow5;

        // ["2", "3", "4", ... "5"]
        let strRow1Array = sRow1.split(", ");
        let strRow2Array = sRow2.split(", ");
        let strRow3Array = sRow3.split(", ");
        let strRow4Array = sRow4.split(", ");
        let strRow5Array = sRow5.split(", ");

        let numRow1Array = strRow1Array.map(Number);
        let numRow2Array = strRow2Array.map(Number);
        let numRow3Array = strRow3Array.map(Number);
        let numRow4Array = strRow4Array.map(Number);
        let numRow5Array = strRow5Array.map(Number);

        let matrixArray = [
            numRow1Array,
            numRow2Array,
            numRow3Array,
            numRow4Array,
            numRow5Array
        ];


        setGameBoard(matrixArray);
        setBalance(data.balance);
      } catch (error) {
        console.error('Error fetching spin data:', error);
      } finally {
        setIsSpinning(false);
      }
    }, 500);
  };

  const toggleMusic = () => {
    setIsMusicPlaying(!isMusicPlaying);
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
          {isSpinning ? 'Spinning...' : 'Крутить'}
        </Button>
        <MusicButton onClick={toggleMusic}>
          {isMusicPlaying ? 'Отключить музыку' : 'Включить музыку'}
        </MusicButton>
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
      <audio ref={audioRef} src={backgroundMusic} />
    </Wrapper>
  );
};



export default SweetBonanzaCard;
